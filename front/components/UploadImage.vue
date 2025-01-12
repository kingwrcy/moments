<template>
  <UPopover :popper="{ arrow: true }" mode="click">
    <svg class="cursor-pointer w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M19 14a3 3 0 1 0-3-3a3 3 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1"/><path fill="currentColor" d="M26 4H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2m0 22H6v-6l5-5l5.59 5.59a2 2 0 0 0 2.82 0L21 19l5 5Zm0-4.83l-3.59-3.59a2 2 0 0 0-2.82 0L18 19.17l-5.59-5.59a2 2 0 0 0-2.82 0L6 17.17V6h20Z"/></svg>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">本地上传</div>
        <UInput type="file" size="sm" icon="i-heroicons-folder" accept="image/*" @change="upload" multiple/>
        <UTextarea :rows="5" placeholder="或者输入在线图片地址,逗号分隔,最多9张" v-model="imgs"/>

        <p v-if="filename" class="text-xs text-gray-400">正在上传({{ current }}/{{ total }})</p>
        <p v-if="filename" class="text-xs text-gray-400">{{ filename }}</p>
        <UProgress :value="progress" v-if="progress > 0" indicator/>

        <UButtonGroup class="w-fit">
          <UButton @click="close()">确定</UButton>
          <UButton color="white" @click="clear(close)">清空并关闭</UButton>
        </UButtonGroup>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import {useUpload} from "~/utils";
import {toast} from "vue-sonner";

const imgs = defineModel<string>('imgs')
const progress = ref(0)
const filename = ref('')
const total = ref(0)
const current = ref(0)
const upload = async (files: FileList) => {
  const containsOtherFile = [...files].some(file => !file.type.startsWith('image/'))
  if (containsOtherFile) {
    toast.error("只能上传图片");
    return
  }

  const result = await useUpload(files, (totalSize: number, index: number, name: string, p: number) => {
    progress.value = Math.round(p * 100)
    filename.value = name
    total.value = totalSize
    current.value = index
  })
  if (result && result.length) {
    toast.success("上传成功")
    imgs.value = [imgs.value, ...result].filter(Boolean).join(',')
  }
}

const clear = (close: Function) => {
  imgs.value = ''
  close()
}
</script>

<style scoped>

</style>