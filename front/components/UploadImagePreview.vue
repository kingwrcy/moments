<template>

  <MyFancyBox :style="gridStyle" class="grid gap-2" v-if="images.length>0">
    <img class="cursor-zoom-in rounded" :class="imgs.length === 1 ? 'full-cover-image-single' : 'full-cover-image-mult'"
         :src="img" alt="" :key="index" v-for="(img,index) in images">
  </MyFancyBox>

</template>

<script setup lang="ts">
const props = defineProps<{ imgs: string }>()
const images = computed(() => {
  if (!props.imgs || props.imgs === ',') {
    return []
  }
  return props.imgs.split(",")
})

const gridCols = computed(() => {
  const len = images.value.length
  if (len === 1) {
    return 1
  }
  return 3
})

const gridStyle = computed(() => {
  let style = 'align-items: start;'; // 确保内容顶部对齐
  switch (images.value.length) {
    case 1:
      style += 'grid-template-columns: 1fr;max-width:60%;';
      break;
    case 2:
      style += 'grid-template-columns: 1fr 1fr; aspect-ratio: 2 / 1;';
      break;
    case 3:
      style += 'grid-template-columns: 1fr 1fr 1fr; aspect-ratio: 3 / 1;';
      break;
    case 4:
      style += 'grid-template-columns: 1fr 1fr; aspect-ratio: 1;';
      break;
    default:
      style += 'grid-template-columns: 1fr 1fr 1fr; aspect-ratio: 3 / 1;';
  }
  return style;
});
</script>


<style scoped>
.full-cover-image-mult {
  object-fit: cover;
  object-position: center;
  width: 100%;
  aspect-ratio: 1 / 1;
  border: transparent 1px solid;
}

.full-cover-image-single {
  object-fit: cover;
  object-position: center;
  max-height: 300px;
  height: auto;
  width: auto;
  border: transparent 1px solid;
}
</style>