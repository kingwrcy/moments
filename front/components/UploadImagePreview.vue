<template>
  <div ref="el" v-if="($route.path.startsWith('/new') || $route.path.startsWith('/edit')) && images.length>0"
       :style="gridStyle" class="grid gap-2">
    <div :key="img" v-for="img in images" class="relative">
      <img :src="img" alt="" class="cursor-move rounded relative"
           :class="images.length === 1 ? 'full-cover-image-single' : 'full-cover-image-mult'">
      <div class="absolute right-0 top-0 px-1 bg-white m-2 rounded hover:text-red-500 cursor-pointer"
           @click="removeImage(img)">
        <UIcon name="i-carbon-trash-can" class=""/>
      </div>
    </div>
  </div>


  <MyFancyBox v-else :style="gridStyle" class="grid gap-2" v-if="images.length>0">
    <img class="cursor-zoom-in rounded"
         :class="images.length === 1 ? 'full-cover-image-single' : 'full-cover-image-mult'"
         :src="img" alt="" :key="z" v-for="(img,z) in images">
  </MyFancyBox>

</template>

<script setup lang="ts">
import {useSortable} from '@vueuse/integrations/useSortable'

const route = useRoute()
const el = ref(null)
const props = defineProps<{ imgs: string }>()
const emit = defineEmits(['removeImage', 'dragImage'])
const images = ref<string[]>((!props.imgs || props.imgs === ',') ? [] : props.imgs.split(","))
watch(props, () => {
  console.log('changed')
  if (!props.imgs || props.imgs === ',') {
    images.value = []
  } else {
    images.value = props.imgs.split(",")
  }
})

watch(images, () => {
  emit('dragImage', images.value)
})
// const images = computed(() => {
//   if (!props.imgs || props.imgs === ',') {
//     return []
//   }
//   return props.imgs.split(",")
// })

const removeImage = async (img: string) => {
  // await useMyFetch('/memo/removeImage',{
  //   img
  // })
  emit('removeImage', img)
}

onMounted(() => {
  if (route.path.startsWith('/new') || route.path.startsWith('/edit')) {
    setTimeout(() => {
      console.log('1234')
      useSortable(el, images)
    }, 500)
  }
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
  max-height: 300px;
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