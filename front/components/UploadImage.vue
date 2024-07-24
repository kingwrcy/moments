<template>
  <UPopover >
    <UIcon name="i-carbon-image" class="cursor-pointer"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">本地上传</div>
        <UInput type="file" size="sm" icon="i-heroicons-folder" />
        <input name="localUpload" id="localUpload" type="file" class="hidden">
        <UTextarea :rows="5" placeholder="或者输入在线图片地址,一行一个,最多9行"/>

        <UButtonGroup>
          <UButton>确定</UButton>
          <UButton color="white">清空</UButton>
        </UButtonGroup>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">

</script>

<style scoped>

</style>