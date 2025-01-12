<template>
  <UPopover :popper="{ arrow: true }" mode="click">
    <svg class="cursor-pointer w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M29.25 6.76a6 6 0 0 0-8.5 0l1.42 1.42a4 4 0 1 1 5.67 5.67l-8 8a4 4 0 1 1-5.67-5.66l1.41-1.42l-1.41-1.42l-1.42 1.42a6 6 0 0 0 0 8.5A6 6 0 0 0 17 25a6 6 0 0 0 4.27-1.76l8-8a6 6 0 0 0-.02-8.48"/><path fill="currentColor" d="M4.19 24.82a4 4 0 0 1 0-5.67l8-8a4 4 0 0 1 5.67 0A3.94 3.94 0 0 1 19 14a4 4 0 0 1-1.17 2.85L15.71 19l1.42 1.42l2.12-2.12a6 6 0 0 0-8.51-8.51l-8 8a6 6 0 0 0 0 8.51A6 6 0 0 0 7 28a6.07 6.07 0 0 0 4.28-1.76l-1.42-1.42a4 4 0 0 1-5.67 0"/></svg>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <UInput v-model="url" placeholder="请输入分享的链接"/>
        <UButtonGroup>
          <UInput v-model="title" placeholder="请输入分享的标题"/>
          <UButton color="white" variant="solid" @click="getFavicon" :disabled="pending"
                   :loading="pending">自动获取标题
          </UButton>
        </UButtonGroup>
        <div class="flex gap-2 ">
          <UInput v-model="favicon" class="flex-1" placeholder="请输入分享的favicon地址"/>
          <UAvatar :src="favicon" size="xs"/>
        </div>

        <div class="w-fit ">
          <UButtonGroup>
            <UButton @click="confirmExternalUrl(close)">确定</UButton>
            <UButton color="white" variant="solid" @click="clear(close)">清空并关闭</UButton>
          </UButtonGroup>
        </div>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import {toast} from "vue-sonner";

const url = defineModel<string>("url");
const favicon = defineModel<string>("favicon");
const title = defineModel<string>("title");
const pending = ref(false)
const clear = (close: Function) => {
  favicon.value = ''
  url.value = ''
  title.value = ''
  close()
}

const confirmExternalUrl = async (close: Function) => {
  if (url.value && title.value && favicon.value) {
    close()

    return
  }
  toast.error("请完整填写相关内容")
}
const getFavicon = async () => {
  if (!url.value) {
    toast.error("请先填写地址")
    return
  }
  pending.value = true
  try {
    const res = await useMyFetch<{
      favicon: string,
      title: string
    }>('/memo/getFaviconAndTitle?url=' + encodeURIComponent(url.value))
    title.value = res.title
    favicon.value = res.favicon
  } finally {
    pending.value = false
  }
}
</script>

<style scoped>

</style>