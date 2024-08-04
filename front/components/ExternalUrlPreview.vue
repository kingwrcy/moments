<template>
  <div v-if="props.url"
       class="flex flex-row gap-2 my-2 bg-[#f7f7f7] dark:bg-[#212121] items-center dark:border-gray-700/50 p-2 border rounded"
  >
    <img class="w-8 h-8" :src="props.favicon" alt=""><a
      :href="props.url" target="_blank" class="text-[#576b95]">{{ props.title }}</a>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  url: string
  title: string
  favicon: string
}>()
</script>

<style scoped>

</style>