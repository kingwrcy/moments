<template>
  <div class="p-4 pb-2 border-b bg-white">
    <div class="flex flex-row my-2 ">
    <div class="flex flex-1 gap-2 ">
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

      <Popover :open="music163Open">
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
              <NuxtLink to="https://www.baidu.com" class="text-gray-500 underline">如何获取?</NuxtLink>
            </div>
            <Input class="my-2" placeholder="请输入网易云音乐代码" v-model="music163Url" />
            <Button size="sm" @click="importMusic">提交</Button>
          </div>
        </PopoverContent>
      </Popover>


      <Popover :open="bilibiliOpen">
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
          <div class="">
            <div class=" text-xs my-2 flex justify-between"><span>嵌入B站视频</span>
              <NuxtLink to="https://www.baidu.com" class="text-gray-500 underline">如何获取?</NuxtLink>
            </div>
            <Input class="my-2" placeholder="请输入B站视频代码" v-model="bilibiliUrl" />
            <Button size="sm" @click="importBiliBili">提交</Button>
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



      <Textarea autocomplete="new-text" v-model="content" rows="2" placeholder="今天发点什么呢?" class="bg-white"></Textarea>
      <div class="absolute right-2 bottom-1 cursor-pointer text-xl" @click="toggleShowEmoji" ref="showEmojiRef">😊</div>
    </div>

    <Emoji v-if="showEmoji" class="mt-2" @emoji-selected="emojiSelected" />

    <iframe class="rounded" frameborder="no" border="0" marginwidth="0" marginheight="0" width=330 height=86
      :src="music163IfrUrl" v-if="music163IfrUrl"></iframe>

    <iframe class="w-full h-[250px] my-2" v-if="bilibiliIfrUrl" :src="bilibiliIfrUrl" scrolling="no" border="0"
      frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>

    <div class="grid grid-cols-3 my-2 gap-2" v-if="imgs && imgs.length > 0">
      <div v-for="(img, index) in imgs" :key="index" class="relative">
        <NuxtImg format="avif,webp" :src="`${img}`" class="rounded" />
        <Trash2 color="#379d1b" :size="15" class="absolute top-1 right-1 cursor-pointer"
          @click="imgs.splice(index, 1)" />
      </div>
    </div>
    <div class="flex flex-row justify-end mt-2 items-center gap-2">


      <Button @click="submitMemo">提交</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'
import { memoUpdateEvent } from '@/lib/event'
import type { Memo } from '~/lib/types';
import { useAnimate } from '@vueuse/core';
import { Image, Music4, Settings, Trash2, LogOut, Youtube } from 'lucide-vue-next'

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from '@/components/ui/tooltip'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
const showEmojiRef = ref<HTMLElement>()
const keyframes = { transform: 'rotate(360deg)' }
const showEmoji = ref(false)
const emit = defineEmits(['memoAdded'])
const toggleShowEmoji = () => {
  showEmoji.value = !showEmoji.value
  useAnimate(showEmojiRef.value, keyframes, { duration: 1000, easing: 'ease-in-out' })
}
const content = ref('')
const id = ref(-1)
const music163Url = ref('')
const music163IfrUrl = ref('')
const music163Open = ref(false)

const bilibiliUrl = ref('')
const bilibiliIfrUrl = ref('')
const bilibiliOpen = ref(false)

// const music163Ifr = ref<HTMLIFrameElement>()
const imgs = ref<string[]>([])
const submitMemo = async () => {
  const res = await $fetch('/api/memo/save', {
    method: 'POST',
    body: JSON.stringify({
      id: id.value,
      content: content.value,
      imgUrls: imgs.value,
      music163Url: music163IfrUrl.value,
      bilibiliUrl: bilibiliIfrUrl.value
    })
  })
  if (res.success) {
    toast.success('提交成功')
    content.value = ''
    id.value = -1
    imgs.value = []
    music163IfrUrl.value = ''
    music163Url.value = ''
    bilibiliIfrUrl.value = ''
    bilibiliUrl.value = ''
    emit('memoAdded')
  } else {
    toast.warning('提交失败')
  }
}

const token = useCookie('token')
const logout = () => {
  token.value = ''
  navigateTo('/', { replace: true })
}

const uploadImgs = async (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }
  const formData = new FormData()
  formData.append('file', file)
  const res = await $fetch('/api/files/upload', {
    method: 'POST',
    body: formData
  })
  if (res.success) {
    (event.target as HTMLInputElement).value = ''
    imgs.value = [...imgs.value, res.filename]
  } else {
    toast.warning(res.message || '上传失败')
  }
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


const importBiliBili = () => {
  if (bilibiliUrl.value === '') {
    toast.warning('请输入B站视频代码')
    return
  }
  const match = bilibiliUrl.value.match(/src="(.*?)"/)
  if (match && match.length > 1) {
    const url = match[1]
    bilibiliIfrUrl.value = url + '&autoplay=0&high_quality=1&as_wide=1'
    bilibiliOpen.value = false
  }
}

const emojiSelected = (emoji: string) => {
  content.value += emoji
  showEmoji.value = false
}
memoUpdateEvent.on((event: Memo) => {
  content.value = event.content
  id.value = event.id
  if (event.imgs) {
    imgs.value = event.imgs?.split(',')
  }
})
</script>

<style scoped></style>