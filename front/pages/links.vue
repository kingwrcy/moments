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
          name="linksName"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.linksName"
            class="mb-2"
            placeholder="(*此项必填)"
          />
        </UFormGroup>
        <UFormGroup
          label="图标"
          name="linksIcon"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.linksIcon"
            class="mb-2"
            placeholder="(*此项必填)"
          />
        </UFormGroup>
        <UFormGroup
          label="网址"
          name="linksUrl"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.linksUrl"
            class="mb-2"
            placeholder="(*此项必填，须以 http(s):// 开头)"
          />
        </UFormGroup>
        <UFormGroup
          label="描述"
          name="linksDesc"
          :ui="{ label: { base: 'font-bold' } }"
        >
          <UInput
            v-model="state.linksDesc"
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
        v-for="links in linksList"
        :key="links.id"
        class="bg-neutral-100 dark:bg-neutral-700 rounded-lg shadow-md overflow-hidden transition-transform hover:scale-105 duration-300 relative"
        @mouseenter="showDeleteIcon(links.id)"
        @mouseleave="hideDeleteIcon(links.id)"
      >
        <a :href="links.linksUrl" target="_blank" class="block p-4">
          <div class="flex items-center gap-2 mb-2">
            <img
              :src="links.linksIcon"
              alt="Friend Avatar"
              class="w-8 h-8 rounded-full"
            />
            <span class="text font-semibold">{{ links.linksName }}</span>
          </div>
          <p class="text-gray-600 dark:text-gray-300 text-sm">
            {{ links.linksDesc || "暂无描述" }}
          </p>
        </a>
        <div
          v-if="showDelete[links.id] && global.userinfo.id === 1"
          class="absolute top-0 right-0 px-1 bg-white dark:bg-gray-900 m-2 rounded hover:text-red-500 cursor-pointer"
          @click="showConfirmModal(links.id)"
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
import type { Links, UserVO } from "~/types";
import { toast } from "vue-sonner";
import { useGlobalState } from "~/store";

const global = useGlobalState();
const state = reactive({
  linksName: "",
  linksIcon: "",
  linksUrl: "",
  linksDesc: "",
});

const currentUser = useState<UserVO>("userinfo");
const linksList = ref<Links[]>([]);
const showAddModal = ref(false);
const showDelete = ref<{ [key: number]: boolean }>({});
const showConfirm = ref(false);
const selectedLinkId = ref<number>(0);

const addLinks = async () => {
  if (state.linksName.length === 0) {
    toast.warning("名称不能为空");
    return;
  }
  if (state.linksIcon.length === 0) {
    toast.warning("图标地址不能为空");
    return;
  }
  if (state.linksUrl.length === 0) {
    toast.warning("网址不能为空");
    return;
  }
  if (!/^https?:\/\//.test(state.linksUrl)) {
    toast.warning("必须以 http 或 https 开头");
    return;
  }
  if (state.linksDesc.length === 0) {
    state.linksDesc = "暂无描述";
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
    linksList.value = response as Links[];
    linksList.value.forEach((links) => {
      showDelete.value[links.id] = false;
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
