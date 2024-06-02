<template>
  <div class="p-2 sm:p-4 pb-2 border-b dark:border-[#C0BEBF]/10">
    <div class="flex flex-row my-2 ">
      <div class="flex flex-1 gap-2 items-center">
        <Popover :open="linkOpen">
          <PopoverTrigger as="div">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger as-child>
                  <Link :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" @click="linkOpen = true" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>插入链接</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </PopoverTrigger>
          <PopoverContent as-child @interact-outside="linkOpen = false">
            <div class="flex flex-col gap-2" @keyup.enter="addLink()">
              <div class="text-xs my-2 flex justify-between"><span>插入链接</span>
              </div>
              <Input class="my-2" placeholder="请输入地址" v-model="externalUrl" />
              <template v-if="externalFetchError">
                <Input class="my-2" placeholder="请输入标题" v-model="externalTitle" />
                <Input class="my-2" placeholder="请输入图标(选填)" v-model="externalFavicon" />
              </template>
              <div class="text-sm my-1" v-if="externalPending">获取信息中...</div>
              <Button size="sm" @click="addLink">确定</Button>
              <Button size="sm" variant="secondary" @click="clearExternalUrl()">清空</Button>
            </div>
          </PopoverContent>
        </Popover>


        <Label for="imgUpload">
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger as-child>
                <Image :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" />
              </TooltipTrigger>
              <TooltipContent>
                <p>上传本地图片</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>

          <input type="file" id="imgUpload" class="hidden" name="file" @change="uploadImgs" multiple>
        </Label>

        <Popover :open="music163Open" v-if="privateConfig.enableMusic163">
          <PopoverTrigger as="div">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger as-child>
                  <Music4 :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" @click="music163Open = true" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>嵌入网易云音乐</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>

          </PopoverTrigger>
          <PopoverContent as-child @interact-outside="music163Open = false">
            <div @keyup.enter="importMusic()">
              <div class="flex flex-col space-y">
                <div class=" text-xs my-2 flex justify-between"><span>嵌入网易云音乐</span>
                  <NuxtLink to="https://jerry.mblog.club/simple-moments-import-music-and-video"
                    class="text-gray-500 underline">
                    如何获取?</NuxtLink>
                </div>
                <Input class="my-2" placeholder="请输入网易云音乐代码" v-model="music163Url" />

                <Input class="my-2" placeholder="请输入在线音频地址" v-model="audioUrl" />

              </div>
              <Button size="sm" class="mr-2" @click="importMusic">确定</Button>
              <Button size="sm" variant="ghost"
                @click="music163IfrUrl = ''; music163Url = ''; audioUrl = ''; music163Open = false;">清空</Button>
            </div>
          </PopoverContent>
        </Popover>


        <Popover :open="bilibiliOpen" v-if="privateConfig.enableVideo">
          <PopoverTrigger as="div">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger as-child>
                  <Youtube :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" @click="bilibiliOpen = true" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>嵌入B站视频</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>

          </PopoverTrigger>
          <PopoverContent as-child @interact-outside="bilibiliOpen = false">
            <div class="flex flex-col gap-2" @keyup.enter="importVideo()">
              <div class="flex flex-col gap-2 mb-2">
                <Label for="localVideo" :class="buttonVariants({ variant: 'outline' })">上传本地视频</Label>
                <Input class="my-2 hidden" type="file" id="localVideo" name="localVideo" @change="uploadLocalVideo" />
              </div>
              <div class="flex flex-col gap-2">
                <div class=" text-xs  flex justify-between"><span>嵌入B站视频</span>
                  <NuxtLink to="https://jerry.mblog.club/simple-moments-import-music-and-video"
                    class="text-gray-500 underline">
                    如何获取?</NuxtLink>
                </div>
                <Input class="my-2" placeholder="请输入B站视频代码" v-model="bilibiliUrl" />
              </div>

              <div class="flex flex-col gap-2">
                <div class=" text-xs  flex justify-between"><span>嵌入Youtube视频</span></div>
                <Input class="my-2" placeholder="请输入Youtube视频链接" v-model="youtubeUrl" />
              </div>


              <div class="flex flex-col gap-2">
                <div class=" text-xs  flex justify-between"><span>嵌入在线视频</span>

                </div>
                <Input class="my-2" placeholder="请输入在线视频地址" v-model="videoUrl" />
              </div>
              <Button size="sm" @click="importVideo">确定</Button>
            </div>
          </PopoverContent>
        </Popover>

        <Popover :open="doubanOpen" v-if="privateConfig.enableDouban">
          <PopoverTrigger as="div">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger as-child>
                  <svg @click="doubanOpen = true" class="focus:outline-0 cursor-pointer w-[18px] h-[18px]"
                    xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                    <path
                      d="M15.2735 15H5V7H19V15H17.3764L16.0767 19H21V21H3V19H7.6123L6.8 16.5L8.70211 15.882L9.71522 19H13.9738L15.2735 15ZM3.5 3H20.5V5H3.5V3ZM7 9V13H17V9H7Z">
                    </path>
                  </svg>
                </TooltipTrigger>
                <TooltipContent>
                  <p>引入豆瓣读书和豆瓣电影</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>

          </PopoverTrigger>
          <PopoverContent as-child @interact-outside="doubanOpen = false">
            <div @keyup.enter="importDouban()">
              <RadioGroup :default-value="douban.type" class="flex flex-row gap-2 text-sm" v-model="douban.type">
                <div class="flex items-center space-x-2">
                  <RadioGroupItem id="book" value="book" />
                  <Label for="book">豆瓣读书</Label>
                </div>
                <div class="flex items-center space-x-2">
                  <RadioGroupItem id="movie" value="movie" />
                  <Label for="movie">豆瓣电影</Label>
                </div>
              </RadioGroup>
              <Input class="my-2" placeholder="请输入豆瓣读书/电影的ID" v-model="douban.id" />
              <Button size="sm" @click="importDouban" :disabled="doubanSubmitting">
                <Loader2 class="w-4 h-4 mr-2 animate-spin" v-if="doubanSubmitting" />确定
              </Button>
              <span class="text-xs ml-2 text-gray-400" v-if="doubanSubmitting">请耐心等待下载豆瓣图片并上传!</span>
            </div>
          </PopoverContent>
        </Popover>

      </div>


      <div class="flex mx-auto gap-2">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger as-child>
              <NuxtLink to="/settings">
                <Settings :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" />
              </NuxtLink>
            </TooltipTrigger>
            <TooltipContent>
              <p>设置</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger as-child>
              <LogOut :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" @click="logout" />
            </TooltipTrigger>
            <TooltipContent>
              <p>退出</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </div>


    </div>
    <div class="relative">
      <Textarea ref="textareaRef" @paste="pasteImg" autocomplete="new-text" v-model="content" rows="4"
        @keyup.ctrl.enter="submitMemo()" placeholder="今天发点什么呢?" class=" dark:text-[#C0BEBF]"></Textarea>
      <div class="absolute right-2 bottom-1 cursor-pointer text-xl" @click="toggleShowEmoji" ref="showEmojiRef">😊</div>
    </div>

    <Emoji v-if="showEmoji" class="mt-2" @emoji-selected="emojiSelected" />

    <iframe class="rounded" frameborder="no" border="0" marginwidth="0" marginheight="0" width=330 height=86
      :src="music163IfrUrl" v-if="music163IfrUrl"></iframe>

    <audio class="w-full my-2" :src="audioUrl" controls v-if="audioUrl && !music163Open"></audio>

    <iframe class="w-full h-[250px] my-2" v-if="bilibiliIfrUrl" :src="bilibiliIfrUrl" scrolling="no" border="0"
      frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>

    <iframe class="w-full h-[250px] my-2" v-if="youtubeIfrUrl" :src="youtubeIfrUrl" scrolling="no" border="0"
      frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>

    <div class="relative" v-if="videoIfrUrl">
      <video class="w-2/3 my-2" :src="videoIfrUrl" controls></video>
      <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer rounded"
        @click="videoIfrUrl = ''; videoUrl = ''" />
    </div>

    <div class="relative" v-if="localVideoUrl && !localVideoUploading">
      <video class="w-2/3 my-2" :src="localVideoUrl" controls></video>
      <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer rounded"
        @click="localVideoUploading = false; localVideoUrl = ''" />
    </div>
    <div v-if="localVideoUploading" class="text-sm my-2">上传中,请耐心等待上传完成!</div>



    <DoubanBook :book="doubanBook" v-if="doubanBook" />
    <DoubanMovie :movie="doubanMovie" v-if="doubanMovie" />

    <div class="flex flex-row gap-2 my-2 bg-[#f7f7f7] dark:bg-[#212121] items-center p-2 border rounded"
      v-if="externalFavicon && externalTitle">
      <div class="flex-1 flex flex-row gap-2 items-center"><img class="w-8 h-8" :src="externalFavicon" alt="">
        <div class="text-sm text-[#576b95] cursor-pointer" v-if="!externalTitleEditing" title="点击编辑标题"
          @click="externalTitleEditing = true">{{ externalTitle }}</div>
        <Input placeholder="请输入链接标题" v-model="externalTitle" v-if="externalTitleEditing" />
      </div>
      <Check class="w-5 h-5 mr-2 cursor-pointer" color="green" v-if="externalTitleEditing"
        @click="externalTitleEditing = false" />
      <CircleX class="w-5 h-5 cursor-pointer" color="red" @click="clearExternalUrl" />
    </div>

    <div class="editImgGrid grid grid-cols-3 my-2 gap-2" v-if="imgs && imgs.length > 0">
      <div v-for="(img, index) in imgs" :key="index" class="relative" draggable="true"
        @dragstart="event => dragStart(event, index)" @dragover="dragOver" @drop="event => drop(event, index)">
        <img :src="getImgUrl(img)" class="rounded object-cover h-full aspect-square max-h-[200px] cursor-move" />
        <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer"
          @click="removePreviewImg(index)" />
      </div>
    </div>
    <div class="flex flex-row justify-between mt-2 items-center gap-2 ">
      <div class="text-sm flex flex-row gap-1 flex-1 items-center">
        <Popover>
          <PopoverTrigger>
            <div class="flex flex-row gap-1 items-center text-[#576b95] text-sm cursor-pointer">
              <MapPin :size=14 />{{ fmtLocation }}
            </div>
          </PopoverTrigger>
          <PopoverContent class="w-80">
            <div class="flex flex-row gap-2 text-sm">
              <Input v-model="location" placeholder="空格分隔" />
              <Button variant="outline" @click="updateLocation" v-if="publicConfig.tencentMapKey">自动获取</Button>
              <Button variant="outline" @click="location = ''">清空</Button>
            </div>
          </PopoverContent>
        </Popover>
      </div>
      <label class="text-sm" :class="[showType ? 'text-lime-600' : 'text-stone-400']">{{ showType ? '公开' : '私密' }}</label>
      <Switch id="showType" v-model:checked="showType"></Switch>
      <Button size="sm" @click="submitMemo">发布</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getImgUrl, insertTextAtCursor } from '~/lib/utils';
import { Textarea } from '@/components/ui/textarea'
import { Button, buttonVariants } from '@/components/ui/button'
import { toast } from 'vue-sonner'
import { memoUpdateEvent } from '@/lib/event'
import type { DoubanBook, DoubanMovie, Memo, MemoExt, PrivateConfig, PublicConfig } from '~/lib/types';
import { useAnimate } from '@vueuse/core';
import { Image, Music4, Settings, Trash2, LogOut, Link, Youtube, CircleX, Check, Loader2, MapPin } from 'lucide-vue-next'
import jsonp from "jsonp";

const publicConfig = useState<PublicConfig>('publicConfig')
const privateConfig = useState<PrivateConfig>('privateConfig')

const textareaRef = ref()
const showEmojiRef = ref<HTMLElement>()
const keyframes = { transform: 'rotate(360deg)' }
const showEmoji = ref(false)
const emit = defineEmits(['memoAdded', 'memoUpdated'])
const toggleShowEmoji = () => {
  showEmoji.value = !showEmoji.value
  useAnimate(showEmojiRef.value, keyframes, { duration: 1000, easing: 'ease-in-out' })
}
const location = ref('')
const fmtLocation = computed(() => {
  if (location.value) {
    return location.value.split(' ').join(' · ')
  }
  return '自定义位置'
})
const content = ref('')
const id = ref(-1)
const music163Url = ref('')
const audioUrl = ref('')
const music163IfrUrl = ref('')
const music163Open = ref(false)

const localVideoUrl = ref('')
const localVideoUploading = ref(false)
const youtubeUrl = ref('')
const youtubeIfrUrl = ref('')
const videoUrl = ref('')
const videoIfrUrl = ref('')
const bilibiliUrl = ref('')
const bilibiliIfrUrl = ref('')
const bilibiliOpen = ref(false)
const doubanOpen = ref(false)

const doubanSubmitting = ref(false)
const showType = ref(true)

const douban = reactive({
  id: '',
  type: 'book'
})
const doubanBook = ref<DoubanBook>()
const doubanMovie = ref<DoubanMovie>()

const linkOpen = ref(false)
const externalUrl = ref('')
const externalTitle = ref('')
const externalFavicon = ref('')
const externalPending = ref(false)
const externalFetchError = ref(false)
const externalTitleEditing = ref(false)

onMounted(() => {
  textareaRef.value?.getRef().focus()
})

const clearExternalUrl = () => {
  externalUrl.value = ''
  externalTitle.value = ''
  externalFavicon.value = ''
  linkOpen.value = false
  externalFetchError.value = false
}
const addLink = async () => {
  if (externalFetchError.value && externalTitle.value === '') {
    toast.warning('请填写标题和图标')
    return
  }
  if (externalFetchError.value && externalTitle.value !== '') {
    externalFetchError.value = false
    linkOpen.value = false
    externalPending.value = false
    externalFavicon.value = externalFavicon.value || '/favicon.png'
    return
  }
  externalPending.value = true
  externalFetchError.value = false
  const { data: res } = await useAsyncData('external_' + externalUrl.value, async () => {
    return await $fetch('/api/memo/readExternal', {
      method: 'POST',
      body: JSON.stringify({ url: externalUrl.value })
    })
  })
  if (res.value?.success) {
    externalTitle.value = res.value?.title || '无法获取标题'
    externalFavicon.value = res.value?.favicon || '/favicon.png'
    linkOpen.value = false
    externalPending.value = false
  } else {
    toast.warning(res.value?.message || '获取失败')
    externalPending.value = false
    externalFetchError.value = true
  }
}
const memoUpdateIndex = useState<number>('memoUpdateIndex', () => -1)

const imgs = ref<string[]>([])

const dragStart = (event: any, index: any) => {
  event.dataTransfer.setData('text/plain', index);
}

const dragOver = (event: any) => {
  event.preventDefault();
}

const drop = (event: any, dropIndex: any) => {
  event.preventDefault();
  const dragIndex = event.dataTransfer.getData('text/plain');
  const dragImg = imgs.value[dragIndex];
  imgs.value.splice(dragIndex, 1);  // 删除被拖拽的图片
  imgs.value.splice(dropIndex, 0, dragImg);  // 在放置位置插入被拖拽的图片
}

const submitMemo = async () => {
  if (content.value === '' && imgs.value.length === 0
    && music163IfrUrl.value === '' && bilibiliIfrUrl.value === ''
    && videoIfrUrl.value === '' && youtubeIfrUrl.value === ''
    && externalUrl.value === '' && !doubanBook.value
    && !doubanMovie.value && localVideoUrl.value === ''&& audioUrl.value === '') {
    toast.warning('请输入内容')
    return
  }
  const res = await $fetch('/api/memo/save', {
    method: 'POST',
    body: JSON.stringify({
      id: id.value,
      content: content.value,
      imgUrls: imgs.value,
      music163Url: music163IfrUrl.value,
      bilibiliUrl: bilibiliIfrUrl.value,
      location: location.value,
      externalFavicon: externalFavicon.value,
      externalTitle: externalTitle.value,
      externalUrl: externalUrl.value,
      ext: {
        doubanBook: doubanBook.value,
        doubanMovie: doubanMovie.value,
        youtubeUrl: youtubeIfrUrl.value,
        videoUrl: videoIfrUrl.value,
        localVideoUrl: localVideoUrl.value,
        audioUrl: audioUrl.value
      },
      showType: showType.value
    })
  })
  if (res.success) {
    toast.success('发布成功')
    content.value = ''

    showType.value = true
    imgs.value = []
    music163IfrUrl.value = ''
    music163Url.value = ''
    bilibiliIfrUrl.value = ''
    bilibiliUrl.value = ''
    videoIfrUrl.value = ''
    videoUrl.value = ''
    youtubeIfrUrl.value = ''
    youtubeUrl.value = ''
    location.value = ''
    externalFavicon.value = ''
    externalTitle.value = ''
    externalUrl.value = ''
    showEmoji.value = false
    doubanBook.value = undefined
    doubanMovie.value = undefined
    douban.id = ''
    douban.type = 'book'
    localVideoUrl.value = ''
    audioUrl.value = ''
    emit(id.value > 0 ? 'memoUpdated' : 'memoAdded')
    id.value = -1
  } else {
    toast.warning('发布失败')
  }
}

const removePreviewImg = async (index: number) => {
  await $fetch('/api/files/removePreviewImg', {
    method: 'POST',
    body: JSON.stringify({
      path: imgs.value[index]
    })
  })
  imgs.value.splice(index, 1)
}

const token = useCookie('token')
const logout = () => {
  token.value = ''
  navigateTo('/', { replace: true })
}

const fetchDoubanBook = async () => {
  return await $fetch('/api/memo/doubanBook', {
    method: 'POST',
    body: JSON.stringify({
      id: douban.id,
      type: douban.type
    })
  })
}

const fetchDoubanMovie = async () => {
  return await $fetch('/api/memo/doubanMovie', {
    method: 'POST',
    body: JSON.stringify({
      id: douban.id,
      type: douban.type
    })
  })
}

const importDouban = async () => {
  if (douban.id === '' || douban.type === '') {
    toast.warning('请输入豆瓣读书/电影的ID')
    return
  }
  doubanSubmitting.value = true
  if (douban.type === 'book') {
    const res = await fetchDoubanBook();
    if (res.success) {
      doubanBook.value = res
      doubanOpen.value = false
    }
  } else if (douban.type === 'movie') {
    const res = await fetchDoubanMovie();
    if (res.success) {
      doubanMovie.value = res
      doubanOpen.value = false
    }
  }
  doubanSubmitting.value = false
}


const pasteImg = async (event: ClipboardEvent) => {
  var items = event.clipboardData?.files
  if (!items || items.length === 0) {
    return;
  }
  for (var i = 0; i < items.length; i++) {
    await useUpload(items[i], async (res) => {
      if (res.success) {
        imgs.value = [...imgs.value, res.filename]
      } else {
        toast.warning(res.message || '上传失败')
      }
    })
  }
}

const uploadImgs = async (event: Event) => {
  const files = ((event.target as HTMLInputElement).files)
  if (!files) {
    return
  }
  localVideoUploading.value = true
  for (var i = 0; i < files.length; i++) {
    var file = files[i];
    useUpload(file, async (res) => {
      if (res.success) {
        (event.target as HTMLInputElement).value = ''
        imgs.value = [...imgs.value, res.filename]
      } else {
        toast.warning(res.message || '上传失败')
      }
    })
  }
  localVideoUploading.value = false
}

const importMusic = () => {
  if (music163Url.value === '' && audioUrl.value === '') {
    toast.warning('请输入网易云音乐代码或者在线音频的地址')
    return
  }
  if (music163Url.value) {
    const match = music163Url.value.match(/src="(.*)&auto.*"/)
    if (match && match.length > 1) {
      const url = match[1]
      music163IfrUrl.value = url + '&auto=0&height=66'     
    }
  }
  music163Open.value = false

}


const youtubeUrlRegs = [/v=([^&#]+)/, /youtu\.be\/(.*)\?/]
const importVideo = () => {
  if (bilibiliUrl.value === '' && videoUrl.value === '' && youtubeUrl.value === '') {
    toast.warning('请输入视频地址/代码')
    return
  }
  if (bilibiliUrl.value) {
    const match = bilibiliUrl.value.match(/src="(.*?)"/)
    if (match && match.length > 1) {
      const url = match[1]
      bilibiliIfrUrl.value = url + '&autoplay=0&high_quality=1&as_wide=1'
    }
  } else if (youtubeUrl.value) {
    for (let i = 0; i < youtubeUrlRegs.length; i++) {
      const match = youtubeUrl.value.match(youtubeUrlRegs[i])
      if (match && match.length > 1) {
        const id = match[1]
        youtubeIfrUrl.value = `https://www.youtube.com/embed/${id}?autoplay=0&frameborder="0"`
        break
      }
    }

  } else if (videoUrl.value) {
    // https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4
    videoIfrUrl.value = videoUrl.value
  }
  bilibiliOpen.value = false
}

const validVideoTypes = ['video/mp4', 'video/webm', 'video/ogg', 'video/quicktime']
const uploadLocalVideo = async (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }
  if (!validVideoTypes.includes(file.type)) {
    toast.warning('不支持的视频类型,只支持mp4/webm/ogg/mov格式.')
    return;
  }
  bilibiliOpen.value = false
  await useUpload(file, async (res) => {
    if (res.success) {
      (event.target as HTMLInputElement).value = ''
      localVideoUrl.value = res.filename
    } else {
      toast.warning(res.message || '上传失败')
    }
  })
}


const emojiSelected = (emoji: string) => {
  const target = textareaRef.value?.getRef() as HTMLTextAreaElement
  insertTextAtCursor(emoji, target)
  content.value = target.value!
  // showEmoji.value = false
}
memoUpdateEvent.on((event: Memo & { index?: number }) => {
  if (event.index) {
    memoUpdateIndex.value = event.index!
  }
  content.value = event.content
  id.value = event.id
  if (event.imgs) {
    imgs.value = event.imgs?.split(',')
  } else {
    imgs.value = []
  }

  location.value = event.location || ''
  externalFavicon.value = event.externalFavicon || ''
  externalTitle.value = event.externalTitle || ''
  externalUrl.value = event.externalUrl || ''
  const memoExt = JSON.parse(event.ext || '{}') as MemoExt
  doubanBook.value = memoExt.doubanBook
  doubanMovie.value = memoExt.doubanMovie
  youtubeIfrUrl.value = memoExt.youtubeUrl
  videoIfrUrl.value = memoExt.videoUrl
  localVideoUrl.value = memoExt.localVideoUrl
  music163IfrUrl.value = event.music163Url || ''
  music163Url.value = `<iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width=330 height=86 src="${event.music163Url}"></iframe>`
  textareaRef.value?.getRef().focus()
  showType.value = event.showType == 1
})

const getTmpLocation = async () => {
  return new Promise(async (resolve, reject) => {
    try {
      let tencentMapKey: string = '';
      if (publicConfig.value?.tencentMapKey !== undefined && publicConfig.value?.tencentMapKey !== '') {
        tencentMapKey = publicConfig.value?.tencentMapKey;
      } else {
        reject('当前站点未开启地图服务，请手动输入位置或者联系管理员开启地图服务');
      }
      const url = 'https://apis.map.qq.com/ws/location/v1/ip';
      const params = {
        key: tencentMapKey,
        output: 'jsonp'
      };
      const queryString = new URLSearchParams(params).toString();
      const jsonpUrl = `${url}?${queryString}`;

      jsonp(jsonpUrl, null, (err: any, data: any) => {
        if (err) {
          return '获取位置失败';
        } else {
          const ipLocation = data;
          if (ipLocation.status === 0) {
            let pos = ipLocation.result.ad_info.nation;
            if (ipLocation.result.ad_info.province !== undefined && ipLocation.result.ad_info.province !== '') {
              pos += ' ' + ipLocation.result.ad_info.province;
            }
            if (ipLocation.result.ad_info.city !== undefined && ipLocation.result.ad_info.city !== '' && ipLocation.result.ad_info.city !== ipLocation.result.ad_info.province) {
              pos += ' ' + ipLocation.result.ad_info.city;
            }
            if (ipLocation.result.ad_info.district !== undefined && ipLocation.result.ad_info.district !== '') {
              pos += ' ' + ipLocation.result.ad_info.district;
            }
            if (ipLocation.result.address_reference !== undefined && ipLocation.result.address_reference !== '') {
              if (ipLocation.result.address_reference.famous_area !== undefined && ipLocation.result.address_reference.famous_area !== '') {
                pos += ' ' + ipLocation.result.address_reference.famous_area.title;
              } else if (ipLocation.result.address_reference.business_area !== undefined && ipLocation.result.address_reference.business_area !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.town !== undefined && ipLocation.result.address_reference.town !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.landmark_l1 !== undefined && ipLocation.result.address_reference.landmark_l1 !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.landmark_l2 !== undefined && ipLocation.result.address_reference.landmark_l2 !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.street !== undefined && ipLocation.result.address_reference.street !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.street_number !== undefined && ipLocation.result.address_reference.street_number !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.crossroad !== undefined && ipLocation.result.address_reference.crossroad !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.water !== undefined && ipLocation.result.address_reference.water !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              } else if (ipLocation.result.address_reference.ocean !== undefined && ipLocation.result.address_reference.ocean !== '') {
                pos += ' ' + ipLocation.result.address_reference.town.title;
              }
            }
            resolve(pos);
          }
        }
        reject('获取位置失败');
      });
    } catch (error) {
      reject('获取位置异常');
      console.error(error);
    }
  });
}

async function updateLocation() {
  try {
    toast.promise(getTmpLocation(), {
      loading: '获取位置中...',
      success: (data: any) => {
        typeof data === "string" ? location.value = data : location.value = ''
        return '获取位置成功';
      },
      error: (error: any) => {
        location.value = '';
        return error;
      }
    });
  } catch (error) {
    console.error(error);
  }
}
</script>

<style scoped>
.editImgGrid img {
  pointer-events: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -o-user-select: none;
  user-select: none;
}
</style>