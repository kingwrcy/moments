<template>
  <UPopover :ui="{base:'w-[300px]'}" :popper="{ arrow: true }" mode="click">
    <UIcon name="i-carbon-video-player" class="cursor-pointer w-6 h-6"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">嵌入b站视频</div>
        <UInput type="text" size="sm" v-model="bilibiliUrl">
          <template #leading>
            <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="24" height="24" viewBox="0 0 48 48">
              <path fill="#1e88e5"
                    d="M36.5,12h-7.086l3.793-3.793c0.391-0.391,0.391-1.023,0-1.414s-1.023-0.391-1.414,0L26.586,12 h-5.172l-5.207-5.207c-0.391-0.391-1.023-0.391-1.414,0s-0.391,1.023,0,1.414L18.586,12H12.5C9.467,12,7,14.467,7,17.5v15 c0,3.033,2.467,5.5,5.5,5.5h2c0,0.829,0.671,1.5,1.5,1.5s1.5-0.671,1.5-1.5h14c0,0.829,0.671,1.5,1.5,1.5s1.5-0.671,1.5-1.5h2 c3.033,0,5.5-2.467,5.5-5.5v-15C42,14.467,39.533,12,36.5,12z M39,32.5c0,1.378-1.122,2.5-2.5,2.5h-24c-1.378,0-2.5-1.122-2.5-2.5 v-15c0-1.378,1.122-2.5,2.5-2.5h24c1.378,0,2.5,1.122,2.5,2.5V32.5z"></path>
              <rect width="2.75" height="7.075" x="30.625" y="18.463" fill="#1e88e5"
                    transform="rotate(-71.567 32.001 22)"></rect>
              <rect width="7.075" height="2.75" x="14.463" y="20.625" fill="#1e88e5"
                    transform="rotate(-18.432 17.998 21.997)"></rect>
              <path fill="#1e88e5"
                    d="M28.033,27.526c-0.189,0.593-0.644,0.896-1.326,0.896c-0.076-0.013-0.139-0.013-0.24-0.025 c-0.013,0-0.05-0.013-0.063,0c-0.341-0.05-0.745-0.177-1.061-0.467c-0.366-0.265-0.808-0.745-0.947-1.477 c0,0-0.29,1.174-0.896,1.49c-0.076,0.05-0.164,0.114-0.253,0.164l-0.038,0.025c-0.303,0.164-0.682,0.265-1.086,0.278 c-0.568-0.051-0.947-0.328-1.136-0.821l-0.063-0.164l-1.427,0.656l0.05,0.139c0.467,1.124,1.465,1.768,2.74,1.768 c0.922,0,1.667-0.303,2.209-0.909c0.556,0.606,1.288,0.909,2.209,0.909c1.856,0,2.55-1.288,2.765-1.843l0.051-0.126l-1.427-0.657 L28.033,27.526z"></path>
            </svg>
          </template>
        </UInput>
        <div class="text-xs text-gray-400">嵌入Youtube站视频</div>
        <UInput type="text" size="sm" icon="i-carbon-video-player" v-model="youtubeUrl">
          <template #leading>
            <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="24" height="24" viewBox="0 0 48 48">
              <path fill="#FF3D00"
                    d="M43.2,33.9c-0.4,2.1-2.1,3.7-4.2,4c-3.3,0.5-8.8,1.1-15,1.1c-6.1,0-11.6-0.6-15-1.1c-2.1-0.3-3.8-1.9-4.2-4C4.4,31.6,4,28.2,4,24c0-4.2,0.4-7.6,0.8-9.9c0.4-2.1,2.1-3.7,4.2-4C12.3,9.6,17.8,9,24,9c6.2,0,11.6,0.6,15,1.1c2.1,0.3,3.8,1.9,4.2,4c0.4,2.3,0.9,5.7,0.9,9.9C44,28.2,43.6,31.6,43.2,33.9z"></path>
              <path fill="#FFF" d="M20 31L20 17 32 24z"></path>
            </svg>
          </template>
        </UInput>
        <div class="text-xs text-gray-400">嵌入在线站视频</div>
        <UInput type="text" size="sm" v-model="onlineUrl">
          <template #leading>
            <UIcon name="i-carbon-video-player" class="w-6 h-6"/>
          </template>
        </UInput>

        <div class="text-xs text-gray-400">上传本地视频</div>
        <UInput type="file" size="sm" icon="i-heroicons-folder" @change="handleUploadVideo"/>

        <p v-if="filename" class="text-xs text-gray-400">正在上传({{ current }}/{{ total }})</p>
        <p v-if="filename" class="text-xs text-gray-400">{{ filename }}</p>
        <UProgress :value="progress" v-if="progress > 0" indicator/>

        <UButtonGroup>
          <UButton @click="confirm(close)">确定</UButton>
          <UButton color="white" @click="reset(close)">清空</UButton>
        </UButtonGroup>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import type {Video, VideoType} from "~/types";
import {toast} from "vue-sonner";
import {useUpload} from "~/utils";

const props = withDefaults(defineProps<Video>(), {
  type: "youtube",
  value: ""
})
const emit = defineEmits(['confirm'])
const videoType = ref<VideoType>(props.type)
const youtubeUrl = ref('')
const bilibiliUrl = ref('')
const onlineUrl = ref('')
const youtubeUrlRegs = [/v=([^&#]+)/, /youtu\.be\/(.*)\?/]
const progress = ref(0)
const filename = ref('')
const total = ref(0)
const current = ref(0)

watch(props, () => {
  videoType.value = props.type
  if (props.type === 'youtube') {
    youtubeUrl.value = props.value
  } else if (props.type === 'bilibili') {
    bilibiliUrl.value = props.value
  } else if (props.type === 'online') {
    onlineUrl.value = props.value
  }
})

const handleUploadVideo = async (files: FileList) => {
  for (let i = 0; i < files.length; i++) {
    if (files[i].type.indexOf("video") < 0){
      toast.error("只能上传视频文件");
      return
    }
  }
  const result = await useUpload(files, (totalSize: number, index: number, name: string, p: number) => {
    progress.value = Math.round(p * 100)
    filename.value = name
    total.value = totalSize
    current.value = index
  }) as string[]
  toast.success("上传成功")
  if (result) {
    onlineUrl.value = result[0]
  }
}
const confirm = (close: Function) => {
  const match = bilibiliUrl.value.match(/src=['"]([^'"]+)['"]/)
  if (bilibiliUrl.value.trim()) {
    videoType.value = 'bilibili'
    if (match && match.length > 1) {
      emit('confirm', {
        type: videoType.value,
        value: bilibiliUrl.value
      })
      close()
    } else {
      toast.warning("无效的B站视频地址")
      return
    }
    return
  }

  if (youtubeUrl.value.trim()) {
    videoType.value = 'youtube'
    let success = false
    for (let i = 0; i < youtubeUrlRegs.length; i++) {
      const match = youtubeUrl.value.match(youtubeUrlRegs[i])
      if (match && match.length > 1) {
        success = true
        break
      }
    }
    if (success) {
      emit('confirm', {
        type: videoType.value,
        value: youtubeUrl.value
      })
      close()
    } else {
      toast.warning("无效的Youtube视频地址")
    }
    return
  }
  if (onlineUrl.value.trim()) {
    videoType.value = 'online'
    emit('confirm', {
      type: videoType.value,
      value: onlineUrl.value.trim()
    })
    close()
    return
  }
}

const reset = (close: Function) => {
  videoType.value = 'youtube'
  youtubeUrl.value = ''
  bilibiliUrl.value = ''
  onlineUrl.value = ''
  emit('confirm', {
    type: videoType.value,
    value: ""
  })
}


</script>

<style scoped>

</style>