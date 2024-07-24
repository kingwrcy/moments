<template>
  <UPopover>
    <UIcon name="i-carbon-link"/>
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
            <UButton color="white" variant="solid" @click="clear(close)">清空</UButton>
          </UButtonGroup>
        </div>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import {toast} from "vue-sonner";

const props = defineProps<{
  externalTitle: string,
  externalUrl: string,
  externalFavicon: string,
}>()

const emit = defineEmits(['confirm'])
const pending = ref(false)
const url = ref(props.externalUrl)
const title = ref(props.externalTitle)
const favicon = ref(props.externalFavicon)
export type ExternalUrlDTO = {
  url: string
  title: string
  favicon: string
}

const clear = (close: Function) => {
  favicon.value = ''
  url.value = ''
  title.value = ''
  emit('confirm', {favicon: '', url: '', title: ''})
  close()
}

const confirmExternalUrl = async (close: Function) => {
  if (url.value && title.value && favicon.value) {
    close()
    emit('confirm', {
      url: url.value,
      title: title.value,
      favicon: favicon.value
    })
    return
  }
  toast.error("请完整填写相关内容")
}
const getFavicon = async () => {
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