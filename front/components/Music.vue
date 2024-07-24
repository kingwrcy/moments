<template>
  <UPopover :ui="{base:'w-[300px]'}">
    <UIcon name="i-carbon-music" class="cursor-pointer"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">嵌入网易云音乐</div>
        <UInput type="text" size="sm"  />
        <UInput  placeholder="或者输入在线音频地址"/>

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