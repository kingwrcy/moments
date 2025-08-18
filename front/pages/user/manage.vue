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

    <!-- 加载骨架屏 -->
    <div v-if="loading && users.length === 0" class="space-y-4">
      <div
        v-if="viewMode === 'card'"
        class="grid grid-cols-2 sm:grid-cols-3 gap-4"
      >
        <USkeleton v-for="i in 6" :key="i" class="h-64" />
      </div>
      <div v-else class="space-y-2">
        <USkeleton class="h-12" />
        <USkeleton v-for="i in 5" :key="i" class="h-16" />
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
        class="hover:shadow-lg transition-shadow duration-200 relative group"
      >
        <div class="absolute top-2 right-2 z-10">
          <UTooltip text="编辑用户">
            <UButton
              color="primary"
              variant="ghost"
              size="xs"
              icon="i-heroicons-pencil-square"
              @click="openUserSettings(user)"
              class="opacity-100 group-hover:opacity-100 transition-opacity"
            />
          </UTooltip>
        </div>

        <div class="flex flex-col items-center space-y-3">
          <NuxtLink :to="'/user/' + user.id">
            <UAvatar
              :src="user.avatarUrl"
              size="xl"
              class="ring-2 ring-gray-200 dark:ring-gray-700"
            />
          </NuxtLink>
          <div class="text-center">
            <h3 class="font-semibold text-lg">
              {{ user.nickname || user.username }}
            </h3>
            <p class="text-sm text-gray-500">@{{ user.username }}</p>
          </div>
          <div class="text-xs text-gray-500">
            注册于 {{ $dayjs(user.createdAt).format("YYYY-MM-DD") }}
          </div>
          <div class="flex items-center space-x-1">
            <div
              class="w-2 h-2 rounded-full"
              :class="user.id === 1 ? 'bg-green-500' : 'bg-blue-500'"
            ></div>
            <span
              class="text-xs font-medium"
              :class="
                user.id === 1
                  ? 'text-green-700 dark:text-green-400'
                  : 'text-blue-700 dark:text-blue-400'
              "
            >
              {{ user.id === 1 ? "管理员" : "普通用户" }}
            </span>
          </div>
          <div
            class="absolute bottom-0 left-0 right-0 p-2 opacity-0 group-hover:opacity-100 transition-all duration-200 pointer-events-auto"
          >
            <UButton
              v-if="user.id !== 1"
              color="primary"
              variant="soft"
              size="sm"
              icon="i-heroicons-trash"
              @click="confirmDelete(user)"
              title="删除用户"
              class="w-full justify-center"
              block
            >
              删除
            </UButton>
          </div>
        </div>
      </UCard>
    </div>

    <!-- 列表视图 -->
    <div
      v-else-if="viewMode === 'table'"
      class="border rounded-lg overflow-hidden"
    >
      <!-- 移动端优化视图 -->
      <div class="sm:hidden">
        <div
          v-for="user in users"
          :key="user.id"
          class="p-4 border-b border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800/50 last:border-b-0"
        >
          <div class="flex items-center space-x-3">
            <NuxtLink :to="'/user/' + user.id">
              <UAvatar
                :src="user.avatarUrl"
                size="md"
                class="ring-2 ring-gray-200 dark:ring-gray-700 flex-shrink-0"
              />
            </NuxtLink>
            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between">
                <div>
                  <div
                    class="text-sm font-medium text-gray-900 dark:text-gray-100"
                  >
                    {{ user.nickname || user.username }}
                  </div>
                  <div class="text-sm text-gray-500 dark:text-gray-400">
                    @{{ user.username }}
                  </div>
                </div>
                <span
                  :class="[
                    'inline-flex items-center px-2 py-1 rounded text-xs font-medium',
                    user.id === 1
                      ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200'
                      : 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
                  ]"
                >
                  {{ user.id === 1 ? "管理员" : "普通用户" }}
                </span>
              </div>
              <div class="mt-2 flex items-center justify-between">
                <div class="text-xs text-gray-500">
                  注册于 {{ $dayjs(user.createdAt).format("YYYY-MM-DD") }}
                </div>
                <div class="flex space-x-1">
                  <UTooltip v-if="user.id !== 1" text="删除用户">
                    <UButton
                      color="red"
                      variant="ghost"
                      size="xs"
                      icon="i-heroicons-trash"
                      @click="confirmDelete(user)"
                      class="hover:bg-red-50 dark:hover:bg-red-900/20"
                    />
                  </UTooltip>
                  <UTooltip text="编辑用户">
                    <UButton
                      color="primary"
                      variant="ghost"
                      size="xs"
                      icon="i-heroicons-pencil-square"
                      @click="openUserSettings(user)"
                      class="hover:bg-primary-50 dark:hover:bg-primary-900/20"
                    />
                  </UTooltip>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 桌面端表格视图 -->
      <div class="hidden sm:block">
        <UTable
          :columns="columns"
          :rows="users"
          :loading="loading"
          class="w-full"
          :ui="tableUi"
        >
          <template #id-data="{ row }">
            <span
              class="text-sm font-mono font-medium text-gray-900 dark:text-gray-100"
              >{{ row.id }}</span
            >
          </template>

          <template #username-data="{ row }">
            <div class="flex items-center">
              <NuxtLink :to="'/user/' + row.id">
                <UAvatar
                  :src="row.avatarUrl"
                  size="sm"
                  class="ring-2 ring-gray-200 dark:ring-gray-700"
                />
              </NuxtLink>
              <div class="ml-3">
                <div
                  class="text-sm font-medium text-gray-900 dark:text-gray-100"
                >
                  {{ row.nickname || row.username }}
                </div>
                <div class="text-sm text-gray-500 dark:text-gray-400">
                  <span class="flex items-center">
                    @{{ row.username || "-" }}
                    <span
                      :class="[
                        'inline-flex items-center px-2 py-0.5 rounded text-xs font-medium',
                        row.id === 1
                          ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200'
                          : 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
                      ]"
                    >
                      {{ row.id === 1 ? "管理员" : "普通用户" }}
                    </span>
                  </span>
                </div>
              </div>
            </div>
          </template>

          <template #createdAt-data="{ row }">
            <div class="text-sm text-gray-900 dark:text-gray-100">
              <div>{{ $dayjs(row.createdAt).format("YYYY-MM-DD") }}</div>
              <div class="text-xs text-gray-500">
                {{ $dayjs(row.createdAt).format("HH:mm") }}
              </div>
            </div>
          </template>

          <template #actions-data="{ row }">
            <div class="flex items-center justify-end space-x-1">
              <UTooltip v-if="row.id !== 1" text="删除用户">
                <UButton
                  color="red"
                  variant="ghost"
                  size="xs"
                  icon="i-heroicons-trash"
                  @click="confirmDelete(row)"
                  class="hover:bg-red-50 dark:hover:bg-red-900/20"
                />
              </UTooltip>
              <UTooltip text="编辑用户">
                <UButton
                  color="primary"
                  variant="ghost"
                  size="xs"
                  icon="i-heroicons-pencil-square"
                  @click="openUserSettings(row)"
                  class="hover:bg-primary-50 dark:hover:bg-primary-900/20"
                />
              </UTooltip>
            </div>
          </template>
        </UTable>
      </div>
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
      <UCard
        :ui="{ body: { base: 'max-h-[60vh] sm:max-h-[90vh] overflow-y-auto' } }"
      >
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
// 表格配置 - 优化响应式设计
const columns = [
  { key: "id", label: "ID", class: "w-12 md:w-16 hidden sm:table-cell" },
  { key: "username", label: "用户信息", class: "min-w-[200px] flex-1" },
  { key: "createdAt", label: "创建时间", class: "w-28 hidden sm:table-cell" },
  { key: "actions", label: "操作", class: "w-24 text-right" },
];

// 表格UI配置
const tableUi = {
  base: "divide-y divide-gray-200 dark:divide-gray-700",
  thead: "bg-gray-50 dark:bg-gray-800/50",
  tbody: "divide-y divide-gray-200 dark:divide-gray-700",
  tr: {
    base: "hover:bg-gray-50 dark:hover:bg-gray-800/50 transition-colors duration-150",
    selected: "bg-gray-100 dark:bg-gray-700",
  },
  th: {
    base: "px-3 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider",
  },
  td: {
    base: "px-3 py-4 whitespace-nowrap",
  },
};

// 列表数据
const users = ref<UserVO[]>([]);
const loading = ref(false);
const hasNext = ref(false);
const viewMode = ref<"table" | "card">("card");

// 分页状态
const state = reactive({
  page: 1,
  size: 12,
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

  if (deleteTarget.value.id === 1) {
    toast.warning("不能删除管理员账号");
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
  state.page = 1;
  loadUsers();
};

// 切换排序
const toggleSort = () => {
  sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
  loadUsers();
};

onMounted(async () => {
  // 检查是否为管理员
  if (!currentUser.value || currentUser.value.id !== 1) {
    toast.warning("无权限访问用户管理页面");
    navigateTo("/");
    return;
  }
  loadUsers();
});
</script>
