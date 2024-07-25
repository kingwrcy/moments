<template>
  <div class="px-4 space-y-2">
    <div class="flex gap-2 text-lg text-gray-600">
      <ExternalUrl v-model:favicon="state.externalFavicon" v-model:title="state.externalTitle"
                   v-model:url="state.externalUrl"/>

      <upload-image v-model:imgs="state.imgs"/>
      <music v-model:id="state.music.id" v-model:type="state.music.type" v-model:server="state.music.server"/>
      <upload-video/>
      <douban-edit :data="state.douban"/>

    </div>

    <div class="w-full" @contextmenu.prevent="onContextMenu" v-if="tags">
      <UTextarea ref="contentRef" v-model="state.content" :rows="8" autoresize padded autofocus/>

      <UContextMenu v-model="isOpen" :virtual-element="virtualElement">
        <div class="px-2 py-1 flex flex-col gap-2 text-xs">
          <div class="mb-2 text-gray-300">点击标签插入</div>
          <div v-for="(tag,index) in tags" :key="index" class="cursor-pointer">
            <UBadge size="xs" color="gray" variant="solid" @click="clickTag(tag)">{{ tag }}</UBadge>
          </div>
        </div>
      </UContextMenu>
    </div>

    <div class="flex flex-col gap-2">
      <external-url-preview :favicon="state.externalFavicon" :title="state.externalTitle" :url="state.externalUrl"/>
      <upload-image-preview :imgs="state.imgs"/>
      <music-preview v-if="state.music.id && state.music.type && state.music.server" :id="state.music.id"
                     :type="state.music.type" :server="state.music.server" :api="state.music.api"/>
      <douban-book-preview :book="state.douban.data"
                           v-if="state.douban && state.douban.data.title && state.douban.type === 'book'"/>
      <douban-movie-preview :book="state.douban.data"
                            v-if="state.douban && state.douban.data.title && state.douban.type === 'movie'"/>
    </div>
    <div class="flex justify-between items-center">
      <div class="flex flex-row gap-1 items-center text-[#576b95] text-sm cursor-pointer">
        <UPopover>
          <div class="flex items-center gap-1">
            <UIcon name="i-carbon-location"/>
            <span>{{ state.location ? locationLabel : '自定义位置' }}</span>
          </div>
          <template #panel="{close}">
            <div class="p-4">
              <UButtonGroup>
                <UInput v-model="state.location" placeholder="自定义位置,空格分隔"/>
                <UButton @click="close" color="white" variant="solid">关闭</UButton>
              </UButtonGroup>
            </div>
          </template>
        </UPopover>
      </div>

      <UButtonGroup>
        <UButton color="white" variant="solid" @click="navigateTo('/')">返回首页</UButton>
        <UButton @click="saveMemo">发表</UButton>
        <UButton color="white" @click="reset">清空</UButton>
      </UButtonGroup>
    </div>
  </div>


</template>

<script setup lang="ts">
import {useMouse, useWindowScroll} from '@vueuse/core'
import type {DoubanBook, DoubanMovie, ExtDTO, MemoVO, MetingMusicServer, MetingMusicType} from "~/types";
import {toast} from "vue-sonner";
import UploadImage from "~/components/UploadImage.vue";


const contentRef = ref(null)
const props = defineProps<{ id?: number }>()
const defaultState = {
  id: props.id || 0,
  content: "",
  ext: "",
  pinned: false,
  showType: 1,
  location: "",
  externalFavicon: "",
  externalTitle: "",
  externalUrl: "",
  music163Url: "",
  bilibiliUrl: "",
  imgs: "",
  music: {
    id: '',
    api: '',
    server: 'netease' as MetingMusicServer,
    type: 'song' as MetingMusicType
  },
  douban: {
    type: 'book',
    data: {} as DoubanBook | DoubanMovie
  }
}

const state = reactive({
  ...defaultState
})
const tags = ref<string[]>([])
const reset = () => {
  Object.assign(state, defaultState)

}

const locationLabel = computed(() => {
  return state.location.split(" ").join(" · ")
})


const {x, y} = useMouse()
const {y: windowY} = useWindowScroll()
const isOpen = ref(false)
const virtualElement = ref({getBoundingClientRect: () => ({})})

function onContextMenu() {
  if (tags.value.length <= 0) {
    return
  }
  const top = unref(y) - unref(windowY)
  const left = unref(x)

  virtualElement.value.getBoundingClientRect = () => ({
    width: 0,
    height: 0,
    top,
    left
  })

  isOpen.value = true
}

const loadTags = async () => {
  const res = await useMyFetch<{
    tags: string[]
  }>("/tag/list")
  tags.value = res.tags
}

const clickTag = (tag: string) => {
  state.content = '#' + tag + ' ' + state.content
  isOpen.value = false;
  //@ts-ignore
  (contentRef.value?.textarea as HTMLTextAreaElement).focus()
}
onMounted(async () => {
  if (state.id > 0) {
    const res = await useMyFetch<MemoVO>('/memo/get?id=' + state.id)
    Object.assign(state, res)
    const ext = JSON.parse(res.ext) as ExtDTO
    Object.assign(state.music, ext.music)
    if (res.tags) {
      const tagsArray = res.tags.split(",")
      const len = tagsArray.length
      let tagStr = ""
      for (let i = 0; i < len; i++) {
        if (tagsArray[i]) {
          tagStr += ('#' + tagsArray[i] + ',')
        }
      }
      if (tagStr) {
        tagStr = tagStr.substring(0, tagStr.length - 1)
      }
      state.content = tagStr + '\n' + res.content
    }
  }
  await loadTags()
})

// const keydown=(event:KeyboardEvent)=>{
//   if(event.key === '#'){
//     tagPopoverOpen.value = true
//   }
// }

const saveMemo = async () => {


  const res = await useMyFetch('/memo/save', {
    id: state.id,
    content: state.content,
    ext: {
      music: state.music,
    },
    pinned: state.pinned,
    showType: state.showType,
    externalFavicon: state.externalFavicon,
    externalTitle: state.externalTitle,
    externalUrl: state.externalUrl,
    imgs: state.imgs.split(','),
    location: state.location,
  })
  toast.success("保存成功!")
  await navigateTo('/')
}

</script>

<style scoped>

</style>