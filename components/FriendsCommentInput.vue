<template>
  <div class="p-2 rounded text-sm ">
    <div class="relative">
      <Textarea autocomplete="new-text" rows="3" v-model="content" class="dark:bg-slate-500 border-separate" :placeholder="placeholder" </Textarea>
      <div class="absolute right-2 bottom-1 cursor-pointer text-xl" @click="toggleShowEmoji" ref="showEmojiRef">😊</div>
    </div>
    <Emoji v-if="showEmoji" class="mt-2" @emoji-selected="emojiSelected"/>
    <div class="flex flex-row items-center justify-end mt-2 gap-2 ">
      <Input placeholder="昵称,必填" type="text"  v-model="info.username" class="dark:bg-slate-500  py-0.5"></Input>
      <Input placeholder="主页,可空" type="text" v-model="info.website" class="dark:bg-slate-500  py-0.5"> </Input>
      <Input placeholder="邮箱,可空" type="text" v-model="info.email" class="dark:bg-slate-500  py-0.5"></Input>
      <Button size="sm" @click="saveComment">发表评论</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner';
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import {useAnimate, useStorage} from '@vueuse/core'

const content = ref('')
const placeholder = ref('发表评论')
const emit = defineEmits(['commentAdded'])
const showEmoji = ref(false)
const keyframes = { transform: 'rotate(360deg)' }
const props = defineProps<{ memoId: number, reply?: string }>()
const showEmojiRef = ref<HTMLElement>()
const token =  useCookie('token')
const info = useStorage('anonymous', {
  email:'',
  website:'',
  username:''
})

const toggleShowEmoji = ()=>{
  showEmoji.value = !showEmoji.value
  useAnimate(showEmojiRef.value, keyframes, { duration: 1000, easing: 'ease-in-out' })
}
const emojiSelected = (emoji: string) => {
  content.value += emoji
  showEmoji.value = false
}

const saveComment = async () => {

  if (!content.value) {
    toast.warning('先填写评论')
    return
  }
  if (!info.value.username) {
    toast.warning('用户名必填')
    return
  }
  const res = await $fetch('/api/comment/save', {
    method: 'POST',
    body: JSON.stringify({
      content: content.value,
      memoId: props.memoId,      
      replyTo: props.reply,
      author:token.value != '',
      email:info.value.email,
      website:info.value.website,
      username:info.value.username
    })
  })

  if (res.success) {
    toast.success('评论成功')
    content.value=''
    emit('commentAdded')
  } else {
    toast.warning('评论失败')
  }
}

onMounted(async () => {

  if (props.reply) {
    placeholder.value = "回复给@" + props.reply
  }
})
</script>

<style scoped></style>
