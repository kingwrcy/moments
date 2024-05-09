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
            <div class="flex flex-col gap-2">
              <div class="text-xs my-2 flex justify-between"><span>插入链接</span>
              </div>
              <Input class="my-2" placeholder="请输入链接地址" v-model="externalUrl" />
              <template v-if="externalFetchError">
                <Input class="my-2" placeholder="请输入链接标题" v-model="externalTitle" />
                <Input class="my-2" placeholder="请输入链接图标,选填" v-model="externalFavicon" />
              </template>
              <div class="text-sm my-1" v-if="externalPending">获取信息中...</div>
              <Button size="sm" @click="addLink">提交</Button>
              <Button size="sm" class="ml-2" variant="secondary" @click="clearExternalUrl()">清空并关闭</Button>
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

          <input type="file" id="imgUpload" class="hidden" name="file" @change="uploadImgs">
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
            <div class="">
              <div class=" text-xs my-2 flex justify-between"><span>嵌入网易云音乐</span>
                <NuxtLink to="https://jerry.mblog.club/simple-moments-import-music-and-video"
                  class="text-gray-500 underline">
                  如何获取?</NuxtLink>
              </div>
              <Input class="my-2" placeholder="请输入网易云音乐代码" v-model="music163Url" />
              <Button size="sm" class="mr-2" @click="importMusic">提交</Button>
              <Button size="sm" variant="ghost"
                @click="music163IfrUrl = ''; music163Url = ''; music163Open = false;">清空</Button>
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
            <div class="flex flex-col gap-2">
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
              <Button size="sm" @click="importVideo">提交</Button>
            </div>
          </PopoverContent>
        </Popover>

        <Popover :open="doubanOpen" v-if="privateConfig.enableDouban">
          <PopoverTrigger as="div">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger as-child>
                  <img src="https://www.douban.com/favicon.ico" @click="doubanOpen = true" t="1713855207689"
                    class="focus:outline-0 cursor-pointer w-[18px] h-[18px] " />
                </TooltipTrigger>
                <TooltipContent>
                  <p>引入豆瓣读书和豆瓣电影</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>

          </PopoverTrigger>
          <PopoverContent as-child @interact-outside="doubanOpen = false">
            <div class="">
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
                <Loader2 class="w-4 h-4 mr-2 animate-spin" v-if="doubanSubmitting" />提交
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
              <a href="/settings">
                <Settings :stroke-width="1.5" class="cursor-pointer w-[20px] h-[20px]" />
              </a>
            </TooltipTrigger>
            <TooltipContent>
              <p>进入设置</p>
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

    <iframe class="w-full h-[250px] my-2" v-if="bilibiliIfrUrl" :src="bilibiliIfrUrl" scrolling="no" border="0"
      frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>

    <iframe class="w-full h-[250px] my-2" v-if="youtubeIfrUrl" :src="youtubeIfrUrl" scrolling="no" border="0"
      frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>

    <div class="relative"  v-if="videoIfrUrl">
      <video class="w-full h-[250px] my-2" :src="videoIfrUrl" controls></video>
      <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer"
        @click="videoIfrUrl = ''; videoUrl = ''" />
    </div>

    <div class="relative" v-if="localVideoUrl && !localVideoUploading">
      <video class="w-full h-[250px] my-2" :src="localVideoUrl" controls></video>
      <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer"
        @click="localVideoUploading = false; localVideoUrl = ''" />
    </div>
    <div v-if="localVideoUploading" class="text-sm my-2">视频上传中,请耐心等待上传完成!</div>



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

    <div class="grid grid-cols-3 my-2 gap-2" v-if="imgs && imgs.length > 0">
      <div v-for="(img, index) in imgs" :key="index" class="relative">
        <img :src="getImgUrl(img)" class="rounded object-cover h-full aspect-square max-h-[200px]" />
        <Trash2 color="rgb(234 88 12)" :size="15" class="absolute top-1 right-1 cursor-pointer"
          @click="removePreviewImg(index)" />
      </div>
    </div>
    <div class="flex flex-row justify-between mt-2 items-center gap-2 ">
      <div class="text-sm flex flex-row gap-1 flex-1 items-center">
        <Popover>
          <PopoverTrigger>
            <div class="text-[#576b95] text-sm cursor-pointer">{{ fmtLocation }}</div>
          </PopoverTrigger>
          <PopoverContent class="w-80">
            <div class="flex flex-row gap-2 text-sm">
              <Input v-model="location" placeholder="空格分隔,火星都行!" />
              <Button variant="outline" @click="location = ''">清空</Button>
            </div>
          </PopoverContent>
        </Popover>
      </div>
      <Button @click="submitMemo">提交</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getImgUrl, insertTextAtCursor } from '~/lib/utils';
import { Textarea } from '@/components/ui/textarea'
import { Button, buttonVariants } from '@/components/ui/button'
import { toast } from 'vue-sonner'
import { memoUpdateEvent } from '@/lib/event'
import type { DoubanBook, DoubanMovie, Memo, MemoExt, PrivateConfig } from '~/lib/types';
import { useAnimate } from '@vueuse/core';
import { Image, Music4, Settings, Trash2, LogOut, Link, Youtube, CircleX, Check, Loader2 } from 'lucide-vue-next'

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
  return '自定义位置?'
})
const content = ref('')
const id = ref(-1)
const music163Url = ref('')
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
const submitMemo = async () => {
  if (content.value === '' && imgs.value.length === 0
    && music163IfrUrl.value === '' && bilibiliIfrUrl.value === ''
    && videoIfrUrl.value === '' && youtubeIfrUrl.value === ''
    && externalUrl.value === '' && !doubanBook.value
    && !doubanMovie.value && localVideoUrl.value === '') {
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
        localVideoUrl: localVideoUrl.value
      }
    })
  })
  if (res.success) {
    toast.success('提交成功')
    content.value = ''

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
    emit(id.value > 0 ? 'memoUpdated' : 'memoAdded')
    id.value = -1
  } else {
    toast.warning('提交失败')
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
  await useUpload(items[0], async (res) => {
    if (res.success) {
      imgs.value = [...imgs.value, res.filename]
    } else {
      toast.warning(res.message || '上传失败')
    }
  })
}

const uploadImgs = async (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }
  localVideoUploading.value = true
  await useUpload(file, async (res) => {
    if (res.success) {
      (event.target as HTMLInputElement).value = ''
      imgs.value = [...imgs.value, res.filename]
    } else {
      toast.warning(res.message || '上传失败')
    }
  })
  localVideoUploading.value = false
}

const importMusic = () => {
  if (music163Url.value === '') {
    toast.warning('请输入网易云音乐代码')
    return
  }
  const match = music163Url.value.match(/src="(.*)&auto.*"/)
  if (match && match.length > 1) {
    const url = match[1]
    music163IfrUrl.value = url + '&auto=0&height=66'
    music163Open.value = false
  }
}


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
    const match = youtubeUrl.value.match(/v=([^&#]+)/)
    if (match && match.length > 1) {
      const id = match[1]
      youtubeIfrUrl.value = `https://www.youtube.com/embed/${id}?autoplay=0&frameborder="0"`
    }
  } else if (videoUrl.value) {
    // https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4
    videoIfrUrl.value = videoUrl.value
  }
  bilibiliOpen.value = false
}

const validVideoTypes = ['video/mp4', 'video/webm', 'video/ogg']
const uploadLocalVideo = async (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }
  if (!validVideoTypes.includes(file.type)) {
    toast.warning('不支持的视频类型,只支持mp4/webm/ogg格式.')
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
})
</script>

<style scoped></style>