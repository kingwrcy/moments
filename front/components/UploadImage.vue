<template>
  <UPopover>
    <UIcon name="i-carbon-image" class="cursor-pointer"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">本地上传</div>
        <UInput type="file" size="sm" icon="i-heroicons-folder" @change="upload"/>
        <input name="localUpload" id="localUpload" type="file" class="hidden">
        <UTextarea :rows="5" placeholder="或者输入在线图片地址,逗号分隔,最多9张" v-model="imgs"/>

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
const upload = async (files: FileList) => {
  const result = await useUpload(files)
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