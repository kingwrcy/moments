<template>
  <iframe v-if="url" class="w-full h-[250px] rounded" :src="bilibiliUrl"
          title="YouTube video player" frameborder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
</template>

<script setup lang="ts">
const props = defineProps<{ url: string }>()
const bilibiliUrl = computed(() => {
  const match = props.url.match(/src=['"]([^'"]+)['"]/)
  if (match && match.length > 1) {
    const url = match[1]
    return url + '&autoplay=0&high_quality=1&as_wide=1'
  }
  return ""
})
</script>

<style scoped>

</style>