<template>
  <div class="px-4 space-y-2">
    <div class="flex justify-between items-center pt-4 text-gray-600">
      <NuxtLink class="flex items-center" title="返回主页">
        <svg @click="navigateTo('/')" class="w-5 h-5 cursor-pointer mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M10 16L20 6l1.4 1.4l-8.6 8.6l8.6 8.6L20 26z"/></svg>
        <span v-if="$route.path==='/new'">新增内容</span>
        <span v-else>修改内容</span>
      </NuxtLink>
      <UButton @click="saveMemo">发表</UButton>
    </div>
    <div class="flex gap-2 text-lg text-gray-600 pt-4 ">
      <ExternalUrl v-model:favicon="state.externalFavicon" v-model:title="state.externalTitle"
                   v-model:url="state.externalUrl"/>

      <upload-image v-model:imgs="state.imgs"/>
      <music v-bind="state.music" @confirm="updateMusic"/>
      <upload-video @confirm="handleVideo" v-bind="state.video"/>
      <douban-edit v-model:type="doubanType" v-model:data="doubanData"/>
      <svg @click="reset" class="w-6 h-6 cursor-pointer" title="清空" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="m29.438 16.572l-7.985-7.986a2 2 0 0 0-2.829 0l-5.358 5.358L9 2H7L2 16h2l.999-3h6l.803 2.408l-7.216 7.216a2 2 0 0 0 0 2.829L9.132 30h9.59l10.716-10.717a1.917 1.917 0 0 0 0-2.712M5.665 11l2.331-7l2.336 7Zm12.23 17H9.96L6 24.038l6.312-6.311l7.928 7.927Zm3.76-3.76l-7.928-7.927L20.039 10l7.927 7.927Z"/></svg>
    </div>

    <div class="w-full" @contextmenu.prevent="onContextMenu">
      <div class="relative">
        <UTextarea ref="contentRef" v-model="state.content" :rows="8" autoresize padded autofocus/>
        <div class="animate-bounce absolute right-2 bottom-1 cursor-pointer text-xl select-none" @click="toggleEmoji">
          😊
        </div>
      </div>

      <Emoji v-if="emojiShow" @selected="emojiSelected" @close="emojiShow=false"/>

      <USelectMenu v-model="selectedLabel" :options="existTags" show-create-option-when="always"
                   multiple searchable creatable placeholder="选择标签" class="my-2" >
        <template #label>
          <span v-if="selectedLabel.length" class="truncate">{{ selectedLabel.join(',') }}</span>
          <span v-else>选择标签</span>
        </template>
      </USelectMenu>
      

      <UContextMenu v-model="isOpen" :virtual-element="virtualElement">
        <div class="px-2 py-1 flex flex-col gap-2 text-xs">
          <div class="mb-2 text-gray-300">点击标签插入</div>
          <div v-for="(tag,index) in existTags" :key="index" class="cursor-pointer">
            <UBadge size="xs" color="gray" variant="solid" @click="clickTag(tag)">{{ tag }}</UBadge>
          </div>
        </div>
      </UContextMenu>
    </div>

    <div class="flex justify-between items-center">
      <div class="flex flex-row gap-1 items-center text-[#576b95] text-sm cursor-pointer">
        <UPopover :popper="{ arrow: true }" mode="click">
          <div class="flex items-center gap-1">
            <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M16 18a5 5 0 1 1 5-5a5.006 5.006 0 0 1-5 5m0-8a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3"/><path fill="currentColor" d="m16 30l-8.436-9.949a35 35 0 0 1-.348-.451A10.9 10.9 0 0 1 5 13a11 11 0 0 1 22 0a10.9 10.9 0 0 1-2.215 6.597l-.001.003s-.3.394-.345.447ZM8.813 18.395s.233.308.286.374L16 26.908l6.91-8.15c.044-.055.278-.365.279-.366A8.9 8.9 0 0 0 25 13a9 9 0 1 0-18 0a8.9 8.9 0 0 0 1.813 5.395"/></svg>
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

      <div class="flex gap-1 text-gray-500 gap-4">
        <div class="flex gap-1 items-center">
          <span>{{ state.showType ? '公开' : '私密' }}</span>
          <UToggle v-model="state.showType"/>
        </div>
      </div>
    </div>
    
    <div class="flex flex-col gap-2">
      <external-url-preview :favicon="state.externalFavicon" :title="state.externalTitle" :url="state.externalUrl"/>
      <upload-image-preview :imgs="state.imgs" @remove-image="handleRemoveImage" @drag-image="handleDragImage"/>
      <music-preview v-if="state.music && state.music.id && state.music.type && state.music.server"
                     v-bind="state.music"/>
      <douban-book-preview :book="doubanData" v-if="doubanType === 'book' && doubanData&& doubanData.title"/>
      <douban-movie-preview :movie="doubanData" v-if="doubanType === 'movie' && doubanData&& doubanData.title"/>
      <youtube-preview v-if="state.video.type === 'youtube' && state.video.value" :url="state.video.value"/>
      <bilibili-preview v-if="state.video.type === 'bilibili' && state.video.value" :url="state.video.value"/>
      <video-preview v-if="state.video.type === 'online' && state.video.value" :url="state.video.value"/>
    </div>
  </div>


</template>

<script setup lang="ts">
import {useMouse, useWindowScroll} from '@vueuse/core'
import type {
  DoubanBook,
  DoubanMovie,
  ExtDTO,
  MemoVO,
  MetingMusicServer,
  MetingMusicType,
  MusicDTO,
  Video,
  VideoType
} from "~/types";
import {toast} from "vue-sonner";
import UploadImage from "~/components/UploadImage.vue";
import Emoji from "~/components/Emoji.vue";

const doubanType = ref<'book' | 'movie'>('book')
const doubanData = ref<DoubanBook | DoubanMovie>({})
const contentRef = ref(null)
const props = defineProps<{ id?: number }>()
const defaultState = {
  id: props.id || 0,
  content: "",
  ext: "",
  pinned: false,
  showType: true,
  location: "",
  externalFavicon: "",
  externalTitle: "",
  externalUrl: "",
  imgs: "",
  music: {
    id: '',
    api: 'https://api.i-meto.com/meting/api?server=:server&type=:type&id=:id&r=:r',
    server: 'netease' as MetingMusicServer,
    type: 'song' as MetingMusicType
  },
  video: {
    type: 'youtube' as VideoType,
    value: ""
  },
  doubanBook: {} as DoubanBook,
  doubanMovie: {} as DoubanMovie,
  tags: Array<string>(),
}
const selectedTags = ref<Array<string>>([])
const selectedLabel = computed({
  get:()=>selectedTags.value,
  set:(labels:Array<string>)=>{
    const tempLabels = Array<string>()
    labels.map(label=>{
      // @ts-ignore
      if(typeof  label !== 'string'){
        // @ts-ignore
        label = label.label
      }
      tempLabels.push(label)
      if(!existTags.value.includes(label)){
        existTags.value.push(label)
      }
    })
    selectedTags.value = [...tempLabels]
    console.log('selectedTags',selectedTags.value)
  }
})
const state = reactive({
  ...defaultState
})
const existTags = ref<string[]>([])
const reset = () => {
  Object.assign(state, defaultState)
}

const locationLabel = computed(() => {
  return state.location.split(" ").join(" · ")
})

const handleDragImage = (imgs: string[]) => {
  state.imgs = imgs.join(",")
}

const updateMusic = (music: MusicDTO) => {
  state.music.id = ""
  setTimeout(() => {
    Object.assign(state.music, music)
  }, 500)
}

const handleVideo = (video: Video) => {
  state.video = video
}

const {x, y} = useMouse()
const {y: windowY} = useWindowScroll()
const isOpen = ref(false)
const virtualElement = ref({getBoundingClientRect: () => ({})})
const handleRemoveImage = (img: string) => {
  const imgs = state.imgs.split(",")
  const index = imgs.findIndex(r => r === img)
  if (index < 0) {
    return
  }
  imgs.splice(index, 1)
  state.imgs = imgs.join(",")
}

function onContextMenu() {
  if (existTags.value.length <= 0) {
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
  existTags.value = res.tags || []
}

const emojiShow = ref(false)

const toggleEmoji = () => {
  emojiShow.value = !emojiShow.value
}
const emojiSelected = (emoji: string) => {
  state.content = state.content + emoji
}

const clickTag = (tag: string) => {
  isOpen.value = false;
  if (!selectedLabel.value.includes(tag)){
    if (selectedLabel.value) {
      selectedLabel.value = [...selectedLabel.value , tag]
    } else {
      selectedLabel.value = [tag]
    }
  }

  //@ts-ignore
  (contentRef.value?.textarea as HTMLTextAreaElement).focus()
}
onMounted(async () => {
  if (state.id > 0) {
    const res = await useMyFetch<MemoVO>('/memo/get?id=' + state.id)
    Object.assign(state, res)
    state.showType = res.showType === 1
    const ext = JSON.parse(res.ext) as ExtDTO
    Object.assign(state.music, ext.music)
    Object.assign(state.video, ext.video)
    doubanType.value = ext.doubanBook && ext.doubanBook.title ? 'book' : 'movie'
    doubanData.value = doubanType.value === 'book' ? ext.doubanBook : ext.doubanMovie
    selectedLabel.value = res.tags ? res.tags.substring(0,res.tags.length-1).split(',') : []
  }
  await loadTags()
})

// const keydown=(event:KeyboardEvent)=>{
//   if(event.key === '#'){
//     tagPopoverOpen.value = true
//   }
// }

const saveMemo = async () => {

  const doubanKey = doubanType.value === 'book' ? 'doubanBook' : 'doubanMovie'
  await useMyFetch('/memo/save', {
    id: state.id,
    content: state.content,
    ext: {
      music: state.music.id ? state.music : {},
      [doubanKey]: doubanData.value,
      video: state.video.value ? state.video : {},
    },
    pinned: state.pinned,
    showType: state.showType ? 1 : 0,
    externalFavicon: state.externalUrl ? state.externalFavicon : "",
    externalTitle: state.externalTitle,
    externalUrl: state.externalUrl,
    imgs: state.imgs.split(','),
    location: state.location,
    tags: selectedLabel.value
  })
  toast.success("保存成功!")
  await navigateTo('/')
}

</script>

<style scoped>

</style>