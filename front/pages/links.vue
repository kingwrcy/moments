<template>
  <Header v-bind:user="currentUser" @add-links="showAddModal = true" />
  <UModal
    v-model="showAddModal"
    :ui="{
      container:
        'fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center',
    }"
  >
    <div class="p-4">
      <p class="text-center text-lg font-bold mb-2">添加友情链接</p>
      <UForm class="space-y-4" size="sm" :state="state">
        <UFormGroup
          label="名称"
          name="name"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput v-model="state.name" class="mb-2" placeholder="(*此项必填)" />
        </UFormGroup>
        <UFormGroup
          label="图标"
          name="icon"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput v-model="state.icon" class="mb-2" placeholder="(*此项必填)" />
        </UFormGroup>
        <UFormGroup
          label="网址"
          name="url"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.url"
            class="mb-2"
            placeholder="(*此项必填，须以 http(s):// 开头)"
          />
        </UFormGroup>
        <UFormGroup
          label="描述"
          name="desc"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.desc"
            class="mb-2"
            placeholder="(此项非必填)"
          />
        </UFormGroup>
        <div class="flex justify-end gap-2 mt-4">
          <UButton color="white" @click="showAddModal = false">取消</UButton>
          <UButton @click="addLinks">确认添加</UButton>
        </div>
      </UForm>
    </div>
  </UModal>
  <div class="bg-white dark:bg-neutral-800">
    <div class="grid sm:grid-cols-2 grid-cols gap-4 p-4">
      <div
        v-for="link in linksList"
        :key="link.id"
        class="bg-neutral-100 dark:bg-neutral-700 rounded-lg shadow-md overflow-hidden transition-transform hover:scale-105 duration-300 relative"
        @mouseenter="showDeleteIcon(link.id)"
        @mouseleave="hideDeleteIcon(link.id)"
      >
        <a :href="link.url" target="_blank" class="block p-4">
          <div class="flex items-center gap-2 mb-2">
            <img
              :src="link.icon"
              alt="Friend Avatar"
              class="w-8 h-8 rounded-full"
            />
            <span class="text font-semibold">{{ link.name }}</span>
          </div>
          <p class="text-gray-600 dark:text-gray-300 text-sm">
            {{ link.desc || "暂无描述" }}
          </p>
        </a>
        <div
          v-if="showDelete[link.id] && global.userinfo.id === 1"
          class="absolute top-0 right-0 px-1 bg-white dark:bg-gray-900 m-2 rounded hover:text-red-500 cursor-pointer"
          @click="showConfirmModal(link.id)"
        >
          <UIcon name="i-carbon-trash-can" />
        </div>
      </div>
    </div>
    <div
      class="flex justify-center items-center text-sm text-gray-400 pt-4 pb-10"
    >
      <span v-if="linksList && linksList.length > 0"
        >共 {{ linksList.length }} 个朋友</span
      >
      <span v-else class="text-gray-600 dark:text-gray-300">
        <span class="text font-semibold">空空如也</span>
        <UButton
          v-if="global.userinfo.id === 1"
          class="ml-2"
          @click="showAddModal = true"
          >点击添加</UButton
        >
      </span>
    </div>
  </div>
  <UModal
    v-model="showConfirm"
    :ui="{
      container:
        'fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center',
    }"
  >
    <div class="p-4 bg-white dark:bg-neutral-800 rounded-lg shadow-md">
      <p class="text-center text-lg font-bold mb-2">确认删除</p>
      <p class="text-gray-600 mb-4">你确定要删除这个友情链接吗？</p>
      <div class="flex justify-end gap-2 mt-4">
        <UButton color="white" @click="hideConfirmModal">取消</UButton>
        <UButton @click="deleteLinks(selectedLinkId)">确认删除</UButton>
      </div>
    </div>
  </UModal>
</template>

<script setup lang="ts">
import type { Link, UserVO } from "~/types";
import { toast } from "vue-sonner";
import { useGlobalState } from "~/store";

const global = useGlobalState();
const state = reactive({
  name: "",
  icon: "",
  url: "",
  desc: "",
});

const currentUser = useState<UserVO>("userinfo");
const linksList = ref<Link[]>([]);
const showAddModal = ref(false);
const showDelete = ref<{ [key: number]: boolean }>({});
const showConfirm = ref(false);
const selectedLinkId = ref<number>(0);

const addLinks = async () => {
  if (state.name.length === 0) {
    toast.warning("名称不能为空");
    return;
  }
  if (state.icon.length === 0) {
    toast.warning("图标地址不能为空");
    return;
  }
  if (state.url.length === 0) {
    toast.warning("网址不能为空");
    return;
  }
  if (!/^https?:\/\//.test(state.url)) {
    toast.warning("必须以 http 或 https 开头");
    return;
  }
  if (state.desc.length === 0) {
    state.desc = "暂无描述";
  }

  try {
    const response = await useMyFetch("/links/add", state);
    toast.success("友情链接添加成功");
    await getLinksList();
    showAddModal.value = false;
  } catch (error) {
    toast.error("友情链接添加失败，请稍后重试");
  }
};

const getLinksList = async () => {
  try {
    const response = await useMyFetch("/links/list");
    linksList.value = response as Link[];
    linksList.value.forEach((link) => {
      showDelete.value[link.id] = false;
    });
  } catch (error) {
    linksList.value = [];
  }
};

const showDeleteIcon = (id: number) => {
  showDelete.value[id] = true;
};

const hideDeleteIcon = (id: number) => {
  showDelete.value[id] = false;
};

const showConfirmModal = (id: number) => {
  selectedLinkId.value = id;
  showConfirm.value = true;
};

const hideConfirmModal = () => {
  showConfirm.value = false;
};

const deleteLinks = async (id: number) => {
  try {
    await useMyFetch(`/links/delete?id=${id}`);
    toast.success("友情链接删除成功");
    await getLinksList();
    showConfirm.value = false;
  } catch (error) {
    toast.error("友情链接删除失败，请稍后重试");
  }
};

onMounted(() => {
  getLinksList();
});
</script>

<style scoped></style>
