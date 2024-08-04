<template>
  <UPopover :popper="{ arrow: true }" mode="click">
    <UIcon name="i-carbon-image" class="cursor-pointer w-10 h-10 sm:w-6 sm:h-6"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">本地上传</div>
        <UInput type="file" size="sm" icon="i-heroicons-folder" @change="upload" multiple/>
        <UTextarea :rows="5" placeholder="或者输入在线图片地址,逗号分隔,最多9张" v-model="imgs"/>

        <p v-if="filename" class="text-xs text-gray-400">正在上传({{current}}/{{total}})</p>
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
  const result = await useUpload(files, (totalSize: number, index: number, name: string, p: number) => {
    progress.value = Math.round(p * 100)
    filename.value = name
    total.value = totalSize
    current.value = index
  }) as string[]
  toast.success("上传成功")
  if (result) {
    imgs.value = (imgs.value ? ',' : '') + result.join(",")
  }
}

const clear = (close: Function) => {
  imgs.value = ''
  close()
}
</script>

<style scoped>

</style>