<template>
  <UPopover :ui="{base:'w-[300px]'}">
    <UIcon name="i-carbon-video-player" class="cursor-pointer"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">嵌入b站视频</div>
        <UInput type="text" size="sm" icon="i-carbon-video-player" />
        <div class="text-xs text-gray-400">嵌入Youtube站视频</div>
        <UInput type="text" size="sm" icon="i-carbon-video-player" />
        <div class="text-xs text-gray-400">嵌入在线站视频</div>
        <UInput type="text" size="sm" icon="i-carbon-video-player" />
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