<template>
  <iframe v-if="url" class="w-full h-[250px] rounded" :src="youtubeUrl"
          title="YouTube video player" frameborder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
</template>

<script setup lang="ts">
const props = defineProps<{ url: string }>()
const youtubeUrlRegs = [/v=([^&#]+)/, /youtu\.be\/(.*)\?/]
const youtubeUrl = computed(() => {
  for (let i = 0; i < youtubeUrlRegs.length; i++) {
    const match = props.url.match(youtubeUrlRegs[i])
    if (match && match.length > 1) {
      const id = match[1]
      return `https://www.youtube.com/embed/${id}?autoplay=0&frameborder="0"`
    }
  }
  return props.url
})


</script>

<style scoped>

</style>