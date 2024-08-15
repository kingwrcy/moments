<template>
  <UPopover :popper="{ arrow: true }" mode="click">
    <UIcon name="i-carbon-link" class="w-6 h-6"/>
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