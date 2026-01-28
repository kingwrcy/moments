<template>
  <div class="px-4 space-y-2">
    <div class="flex justify-between items-center pt-4 text-gray-600">
      <NuxtLink class="flex items-center" title="返回主页">
        <UIcon @click="navigateTo('/')" name="i-carbon-chevron-left" class="w-5 h-5 cursor-pointer mr-4"/>
        <span v-if="$route.path==='/new'">新增内容</span>
        <span v-else>修改内容</span>
      </NuxtLink>
      <UButton @click="saveMemo">发表</UButton>
    </div>
    
    <!-- 工具栏区域 -->
    <div class="flex gap-2 text-lg text-gray-600 pt-4 flex-wrap">
      <!-- 外部链接 -->
      <ExternalUrl v-model:favicon="state.externalFavicon" v-model:title="state.externalTitle"
                   v-model:url="state.externalUrl"/>

      <!-- 图片上传 -->
      <upload-image v-model:imgs="state.imgs"/>
      
      <!-- 音乐 -->
      <music v-bind="state.music" @confirm="updateMusic"/>
      
      <!-- 视频 -->
      <upload-video @confirm="handleVideo" v-bind="state.video"/>
      
      <!-- 豆瓣 -->
      <douban-edit v-model:type="doubanType" v-model:data="doubanData"/>

      <!-- [新增] Steam 游戏 -->
      <UPopover :popper="{ arrow: true }" mode="click">
        <UIcon name="i-icon-park-solid-game-handle" class="w-6 h-6 cursor-pointer hover:text-blue-500" title="Steam游戏" />
        <template #panel="{ close }">
          <div class="p-4 flex flex-col gap-2 w-[300px]">
            <div class="text-sm font-bold">添加 Steam 游戏</div>
            <UInput v-model="steamInput" placeholder="输入 Steam 游戏 ID 或 商店链接" icon="i-carbon-game-console" />
            <div class="text-xs text-gray-400">例如: 2357570 或 完整URL</div>
            <UButtonGroup class="flex justify-end">
              <UButton color="white" @click="close">取消</UButton>
              <UButton @click="() => { fetchSteam(steamInput); close(); }">获取</UButton>
            </UButtonGroup>
          </div>
        </template>
      </UPopover>

      <!-- [新增] TMDB 影视 -->
      <UPopover :popper="{ arrow: true }" mode="click">
        <UIcon name="i-icon-park-solid-movie" class="w-6 h-6 cursor-pointer hover:text-yellow-500" title="TMDB影视" />
        <template #panel="{ close }">
          <div class="p-4 flex flex-col gap-2 w-[300px]">
            <div class="text-sm font-bold">添加 TMDB 影视</div>
            <div class="flex gap-2">
               <USelectMenu v-model="tmdbType" :options="[{label:'电影', value:'movie'}, {label:'剧集', value:'tv'}]" value-attribute="value" option-attribute="label" class="w-24" />
               <UInput v-model="tmdbInput" class="flex-1" placeholder="输入 TMDB ID" />
            </div>
            <div class="text-xs text-gray-400">例如: 157336 (星际穿越)</div>
            <UButtonGroup class="flex justify-end">
              <UButton color="white" @click="close">取消</UButton>
              <UButton @click="() => { fetchTmdb(tmdbInput, tmdbType); close(); }">获取</UButton>
            </UButtonGroup>
          </div>
        </template>
      </UPopover>

      <!-- 自定义时间 -->
      <UPopover :popper="{ arrow: true }" mode="click">
        <UIcon name="i-carbon-calendar" class="w-6 h-6" title="自定义时间"/>
        <template #panel="{close}">
          <DatePicker
            v-model="state.createdAt"
            mode="datetime"
            is24hr
            :time-accuracy="2"
            :rules="{ seconds: 0 }"
            @close="close"
          />
        </template>
      </UPopover>
      
      <!-- 清空按钮 -->
      <UIcon name="i-carbon-text-clear-format" @click="reset" class="w-6 h-6 cursor-pointer text-red-400" title="清空"></UIcon>
    </div>

    <!-- 文本编辑区域 -->
    <div class="w-full" @contextmenu.prevent="onContextMenu">
      <div class="relative">
        <UTextarea ref="contentRef" v-model="state.content" :rows="8" autoresize padded autofocus placeholder="此刻的想法..."/>
        <UIcon class="text-[#9fc84a] w-6 h-6 animate-bounce absolute right-2 bottom-1 cursor-pointer select-none" name="i-carbon-face-satisfied" @click="toggleEmoji"/>
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

    <!-- 底部选项 -->
    <div class="flex justify-between items-center">
      <div class="flex flex-row gap-1 items-center text-[#576b95] text-sm cursor-pointer">
        <UPopover :popper="{ arrow: true }" mode="click">
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

      <div class="flex gap-1 text-gray-500 items-center">
          <span>{{ state.showType ? '公开' : '私密' }}</span>
          <UToggle v-model="state.showType"/>
        </div>
      </div>

    <!-- 预览区域 -->
    <div class="flex flex-col gap-2 pb-10">
      <external-url-preview :favicon="state.externalFavicon" :title="state.externalTitle" :url="state.externalUrl"/>
      
      <upload-image-preview :imgs="state.imgs" @remove-image="handleRemoveImage" @drag-image="handleDragImage"/>
      
      <music-preview v-if="state.music && state.music.id && state.music.type && state.music.server"
                     v-bind="state.music"/>
      
      <douban-book-preview :book="doubanData" v-if="doubanType === 'book' && doubanData&& doubanData.title"/>
      <douban-movie-preview :movie="doubanData" v-if="doubanType === 'movie' && doubanData&& doubanData.title"/>
      
      <!-- [新增] Steam 预览 -->
      <steam-game-preview v-if="state.steamGame && state.steamGame.name" :game="state.steamGame" />

      <!-- [新增] TMDB 预览 -->
      <tmdb-preview v-if="state.tmdbItem && state.tmdbItem.title" :item="state.tmdbItem" />

      <video-preview-iframe v-if="['bilibili', 'youtube'].includes(state.video.type) && state.video.value" :url="state.video.value"/>
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
  VideoType,
  SteamGame, // 确保 types/index.ts 已定义
  TmdbItem   // 确保 types/index.ts 已定义
} from "~/types";
import {toast} from "vue-sonner";
import UploadImage from "~/components/UploadImage.vue";
import Emoji from "~/components/Emoji.vue";
import dayjs from "dayjs";

// [新增] 输入框状态
const steamInput = ref('')
const tmdbInput = ref('')
const tmdbType = ref<'movie'|'tv'>('movie')

const doubanType = ref<'book' | 'movie'>('book')
const doubanData = ref<DoubanBook | DoubanMovie>({})
const contentRef = ref(null)
const props = defineProps<{ id?: number }>()

const defaultState = {
  id: props.id || 0,
  createdAt: '' as string,
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
  // [新增] 默认状态
  steamGame: {} as SteamGame,
  tmdbItem: {} as TmdbItem,
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
  }
})

const state = reactive({
  ...defaultState
})

const existTags = ref<string[]>([])

// [修改] 重置时包含新字段
const reset = () => {
  // 简单重置，注意对象引用的问题，这里手动清空一下深层对象
  Object.assign(state, {
    ...defaultState,
    music: { ...defaultState.music },
    video: { ...defaultState.video },
    doubanBook: {},
    doubanMovie: {},
    steamGame: {},
    tmdbItem: {},
    tags: []
  })
  doubanData.value = {}
  selectedLabel.value = []
  steamInput.value = ''
  tmdbInput.value = ''
}

const locationLabel = computed(() => {
  return state.location.split(" ").join(" · ")
})

const handleDragImage = (imgs: string[]) => {
  state.imgs = imgs.filter(Boolean).join(",")
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

// [新增] 获取 Steam 信息
const fetchSteam = async (urlOrId: string) => {
    if (!urlOrId) return toast.error("请输入内容");
    const idMatch = urlOrId.match(/\/app\/(\d+)/) || [null, urlOrId];
    const id = idMatch[1];
    
    try {
        const res = await useMyFetch<SteamGame>('/memo/getSteamGameInfo?id=' + id);
        if(res) {
            state.steamGame = res;
            toast.success("获取 Steam 游戏成功: " + res.name);
        }
    } catch (e: any) {
        toast.error("获取失败: " + e.message);
    }
}

// [新增] 获取 TMDB 信息
const fetchTmdb = async (id: string, type: 'movie' | 'tv' = 'movie') => {
    if (!id) return toast.error("请输入ID");
    try {
        const res = await useMyFetch<TmdbItem>(`/memo/getTmdbInfo?id=${id}&type=${type}`);
        if(res) {
            state.tmdbItem = res;
            toast.success("获取 TMDB 信息成功: " + res.title);
        }
    } catch (e: any) {
        toast.error("获取失败: " + e.message);
    }
}

const {x, y} = useMouse()
const {y: windowY} = useWindowScroll()
const isOpen = ref(false)
const virtualElement = ref({getBoundingClientRect: () => ({})})

const handleRemoveImage = (index: number) => {
  const arr = state.imgs.split(",").filter(Boolean)
  arr.splice(index, 1)
  state.imgs = arr.join(",")
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
    
    // 回显数据
    Object.assign(state.music, ext.music)
    Object.assign(state.video, ext.video)
    if(ext.steamGame) state.steamGame = ext.steamGame
    if(ext.tmdbItem) state.tmdbItem = ext.tmdbItem

    doubanType.value = ext.doubanBook && ext.doubanBook.title ? 'book' : 'movie'
    doubanData.value = doubanType.value === 'book' ? ext.doubanBook : ext.doubanMovie
    selectedLabel.value = res.tags ? res.tags.substring(0,res.tags.length-1).split(',') : []
    state.createdAt = dayjs(res.createdAt).format()
  }
  await loadTags()
})

const saveMemo = async () => {
  const doubanKey = doubanType.value === 'book' ? 'doubanBook' : 'doubanMovie'
  await useMyFetch('/memo/save', {
    id: state.id,
    content: state.content,
    ext: {
      music: state.music.id ? state.music : {},
      [doubanKey]: doubanData.value,
      video: state.video.value ? state.video : {},
      // [新增] 保存 Steam 和 TMDB 数据
      steamGame: state.steamGame && state.steamGame.name ? state.steamGame : {},
      tmdbItem: state.tmdbItem && state.tmdbItem.title ? state.tmdbItem : {},
    },
    pinned: state.pinned,
    showType: state.showType ? 1 : 0,
    externalFavicon: state.externalUrl ? state.externalFavicon : "",
    externalTitle: state.externalTitle,
    externalUrl: state.externalUrl,
    imgs: state.imgs.split(",").filter(Boolean),
    location: state.location,
    tags: selectedLabel.value,
    createdAt: state.createdAt || dayjs().format(),
  })
  toast.success("保存成功!")
  await navigateTo('/')
}

</script>

<style scoped>

</style>