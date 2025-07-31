<template>
  <Header :user="currentUser" />
  <div class="flex flex-col p-4 my-4 dark:bg-neutral-800">
    <div class="flex justify-between items-center mb-4">
      <div class="text-lg font-bold">
        <UButtonGroup size="sm">
          <UButton
            :color="viewMode === 'card' ? 'primary' : 'gray'"
            variant="ghost"
            @click="viewMode = 'card'"
            icon="i-heroicons-squares-2x2"
            :title="'卡片视图'"
          />
          <UButton
            :color="viewMode === 'table' ? 'primary' : 'gray'"
            variant="ghost"
            @click="viewMode = 'table'"
            icon="i-heroicons-table-cells"
            :title="'列表视图'"
          />
          <UButton
            :color="sortOrder === 'asc' ? 'primary' : 'gray'"
            variant="ghost"
            @click="toggleSort"
            :icon="
              sortOrder === 'asc'
                ? 'i-heroicons-arrow-up'
                : 'i-heroicons-arrow-down'
            "
            :title="sortOrder === 'asc' ? '升序排序' : '降序排序'"
          />
        </UButtonGroup>
      </div>
      <div class="flex items-center space-x-4">
        <UInput
          v-model="keyword"
          placeholder="用户名或昵称..."
          icon="i-heroicons-magnifying-glass"
          size="sm"
          @keyup.enter="handleSearch"
          class="w-48"
        />
      </div>
    </div>

    <!-- 卡片视图 -->
    <div
      v-if="viewMode === 'card'"
      class="grid grid-cols-2 sm:grid-cols-3 gap-4"
    >
      <UCard
        v-for="user in users"
        :key="user.id"
        class="hover:shadow-lg transition-shadow duration-200"
      >
        <div class="flex flex-col items-center space-y-3">
          <UAvatar
            :src="user.avatarUrl"
            size="xl"
            class="ring-2 ring-gray-200 dark:ring-gray-700"
          />
          <div class="text-center">
            <h3 class="font-semibold text-lg">
              {{ user.nickname || user.username }}
            </h3>
            <p class="text-sm text-gray-500">@{{ user.username }}</p>
          </div>
          <div class="text-xs text-gray-500">
            注册于 {{ $dayjs(user.createdAt).format("YYYY-MM-DD") }}
          </div>
          <div class="flex justify-center items-center space-x-2">
            <UButton
              color="blue"
              variant="ghost"
              size="xs"
              icon="i-heroicons-pencil"
              @click="openUserSettings(user)"
              :title="'编辑用户'"
            >
              编辑
            </UButton>
            <UButton
              color="red"
              variant="ghost"
              size="xs"
              icon="i-heroicons-trash"
              @click="confirmDelete(user)"
              :disabled="user.id === currentUser?.id"
              :title="
                user.id === currentUser?.id ? '不能删除自己的账户' : '删除用户'
              "
            >
              删除
            </UButton>
          </div>
        </div>
      </UCard>
    </div>
    <!-- 列表视图 -->
    <div v-else-if="viewMode === 'table'" class="overflow-x-auto">
      <UTable
        :columns="columns"
        :rows="users"
        :loading="loading"
        class="w-full"
      >
        <template #id-data="{ row }">
          <span class="text-sm">{{ row.id }}</span>
        </template>

        <template #avatarUrl-data="{ row }">
          <UAvatar :src="row.avatarUrl" size="sm" />
        </template>

        <template #username-data="{ row }">
          <div class="space-y-2">
            <div class="font-medium">{{ row.username }}</div>
          </div>
        </template>

        <template #nickname-data="{ row }">
          <div class="text-sm text-gray-500">{{ row.nickname }}</div>
        </template>

        <template #email-data="{ row }">
          <span class="text-sm">{{ row.email || "-" }}</span>
        </template>

        <template #createdAt-data="{ row }">
          <span class="text-sm">
            {{ $dayjs(row.createdAt).format("YYYY-MM-DD HH:mm") }}
          </span>
        </template>

        <template #actions-data="{ row }">
          <div class="flex items-center space-x-2">
            <UButton
              color="blue"
              variant="ghost"
              size="xs"
              icon="i-heroicons-pencil"
              @click="openUserSettings(row)"
            >
              编辑
            </UButton>
            <UButton
              color="red"
              variant="ghost"
              size="xs"
              icon="i-heroicons-trash"
              @click="confirmDelete(row)"
              :disabled="row.id === currentUser?.id"
              :title="
                row.id === currentUser?.id ? '不能删除自己的账户' : '删除用户'
              "
            >
              删除
            </UButton>
          </div>
        </template>
      </UTable>
    </div>
    <div class="flex justify-center text-md text-gray-500 py-4">
      已加载 {{ users.length }} 个用户
    </div>
    <!-- 加载更多 -->
    <div
      ref="loadMoreEle"
      class="text-xs text-center text-gray-500 py-4 cursor-pointer"
      @click="loadMore"
      v-if="hasNext"
    >
      点击加载更多
    </div>
    <div class="text-xs text-center text-gray-500 py-4" v-else>已经到底啦</div>

    <!-- 用户设置对话框 -->
    <UModal
      v-model="showUserSettingsModal"
      :ui="{
        container: 'flex justify-center items-center backdrop-blur',
      }"
    >
      <UCard>
        <template #header>
          <h3 class="text-lg font-semibold">
            编辑用户 - {{ settingsTargetUser?.username }}
          </h3>
        </template>

        <UserSettings
          v-if="settingsTargetUser"
          :target-user="settingsTargetUser"
          :is-admin-mode="true"
          :on-save="handleUserSettingsSave"
        />
      </UCard>
    </UModal>

    <!-- 删除确认对话框 -->
    <UModal
      v-model="showDeleteModal"
      :ui="{
        container: 'flex justify-center items-center backdrop-blur',
      }"
    >
      <UCard>
        <template #header>
          <h3 class="text-lg font-semibold">确认删除</h3>
        </template>

        <p>
          确定要删除用户 "{{ deleteTarget?.username }}" 吗？此操作不可恢复。
        </p>
        <p
          v-if="deleteTarget?.id === currentUser?.id"
          class="text-red-500 font-medium mt-2"
        >
          ⚠️ 警告：您正在尝试删除自己的账户，这是不允许的操作！
        </p>

        <div class="flex justify-end space-x-2 mt-4">
          <UButton color="gray" @click="showDeleteModal = false">取消</UButton>
          <UButton color="red" :loading="deleting" @click="doDelete"
            >删除</UButton
          >
        </div>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { UserVO, SysConfigVO } from "~/types";
import { toast } from "vue-sonner";
import UserSettings from "~/pages/user/settings.vue";
import { useElementVisibility } from "@vueuse/core";

const currentUser = useState<UserVO>("userinfo");
const sysConfig = useState<SysConfigVO>("sysConfig");
// 表格配置
const columns = [
  { key: "id", label: "ID" },
  { key: "avatarUrl", label: "头像", width: "60px" },
  { key: "username", label: "登录名" },
  { key: "nickname", label: "昵称" },
  { key: "email", label: "邮箱" },
  { key: "createdAt", label: "注册时间" },
  { key: "actions", label: "操作" },
];

// 列表数据
const users = ref<UserVO[]>([]);
const loading = ref(false);
const hasNext = ref(false);
const viewMode = ref<"table" | "card">("card");

// 分页状态
const state = reactive({
  page: 1,
  size: 6,
  sort: "desc",
  keyword: "",
});

// 搜索关键词
const keyword = ref("");

// 排序状态
const sortOrder = ref<"asc" | "desc">("desc");

const loadMoreEle = ref(null);
const targetIsVisible = useElementVisibility(loadMoreEle);

// 定义响应数据接口
interface UserListResponse {
  list: UserVO[];
  hasNext: boolean;
  keyword?: string;
}

// 用户设置相关
const showUserSettingsModal = ref(false);
const settingsTargetUser = ref<UserVO | null>(null);

// 删除相关
const showDeleteModal = ref(false);
const deleteTarget = ref<UserVO | null>(null);
const deleting = ref(false);

// 加载用户列表
const loadUsers = async () => {
  state.page = 1;
  state.sort = sortOrder.value;
  state.keyword = keyword.value;
  const res = await useMyFetch<UserListResponse>("/user/list", state);
  if (res) {
    users.value = res.list || [];
    hasNext.value = res.hasNext || false;
  }
};

// 搜索处理
const handleSearch = () => {
  loadUsers();
};

const loadMore = async () => {
  state.page = state.page + 1;
  state.sort = sortOrder.value;
  state.keyword = keyword.value;
  const res = await useMyFetch<UserListResponse>("/user/list", state);
  if (res) {
    users.value = [...users.value, ...res.list];
    hasNext.value = res.hasNext || false;
  }
};

watch(targetIsVisible, async (visible) => {
  if (visible && sysConfig.value.enableAutoLoadNextPage) {
    await loadMore();
  }
});

// 确认删除
const confirmDelete = (user: UserVO) => {
  deleteTarget.value = user;
  showDeleteModal.value = true;
};

// 执行删除
const doDelete = async () => {
  if (!deleteTarget.value) return;

  // 防止删除自己的账户
  if (deleteTarget.value.id === currentUser?.value?.id) {
    toast.error("不能删除自己的账户");
    showDeleteModal.value = false;
    return;
  }

  deleting.value = true;
  try {
    await useMyFetch(`/user/${deleteTarget.value.id}`, null);
    toast.success("删除成功");
    showDeleteModal.value = false;
    await loadUsers();
  } catch (error) {
    toast.error("删除失败");
  } finally {
    deleting.value = false;
  }
};

// 打开用户设置
const openUserSettings = (user: UserVO) => {
  settingsTargetUser.value = user;
  showUserSettingsModal.value = true;
};

// 处理用户设置保存完成
const handleUserSettingsSave = () => {
  showUserSettingsModal.value = false;
  state.page = 1; // 重置分页
  loadUsers(); // 重新加载用户列表
};

// 切换排序
const toggleSort = () => {
  sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
  loadUsers();
};

onMounted(() => {
  loadUsers();
});
</script>
