<template>
  <div class="p-2 rounded text-sm ">
    <div class="relative">
      <Textarea @keyup.ctrl.enter="saveComment" ref="textareaRef" autocomplete="new-text" rows="3" v-model="content"
        class="dark:bg-slate-500 border-separate" :placeholder="placeholder" </Textarea>
      <div class="absolute right-2 bottom-1 cursor-pointer text-xl" @click="toggleShowEmoji" ref="showEmojiRef">😊</div>
    </div>
    <Emoji v-if="showEmoji" class="mt-2" @emoji-selected="emojiSelected"/>
    <div class="flex flex-row items-center justify-end mt-2 gap-2">
      <Input placeholder="昵称,必填" type="text"  v-model="info.username" class=" dark:bg-slate-500 text-xs sm:text-sm  py-0.5"></Input>
      <Input placeholder="主页,可空" type="text" v-model="info.website" class="dark:bg-slate-500 text-xs sm:text-sm  py-0.5"> </Input>
      <Input placeholder="邮箱,可空" type="text" v-model="info.email" class="hidden sm:block dark:bg-slate-500 text-xs sm:text-sm py-0.5"></Input>
      <Button size="sm" @click="saveComment" :disabled="pending">发表评论</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner';
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import { useAnimate, useStorage } from '@vueuse/core'
import { insertTextAtCursor } from '~/lib/utils';
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')
const config = useRuntimeConfig()
const textareaRef = ref()
const token = useCookie('token')
const content = ref('')
const placeholder = ref('发表评论')
const emit = defineEmits(['commentAdded'])
const showEmoji = ref(false)
const keyframes = { transform: 'rotate(360deg)' }
const props = defineProps<{ memoId: number, reply?: string }>()
const showEmojiRef = ref<HTMLElement>()
const info = useStorage('anonymous', {
  email: '',
  website: '',
  username: ''
})



onMounted(() => {
  textareaRef.value?.getRef().focus()
  if (token.value && userinfo.value && userinfo.value.nickname) {
    info.value.username = userinfo.value.nickname
  }
})

const toggleShowEmoji = () => {
  showEmoji.value = !showEmoji.value
  useAnimate(showEmojiRef.value, keyframes, { duration: 1000, easing: 'ease-in-out' })
}
const emojiSelected = (emoji: string) => {
  const target = textareaRef.value?.getRef() as HTMLTextAreaElement
  insertTextAtCursor(emoji, target)
  content.value = target.value!
  // showEmoji.value = false
}

const pending = ref(false)

const saveComment = async () => {

  if (!content.value) {
    toast.warning('先填写评论')
    return
  }
  if (content.value.length > parseInt(config.public.momentsCommentMaxLength)) {
    toast.warning('评论超长')
    return
  }
  if (!info.value.username) {
    toast.warning('用户名必填')
    return
  }
  if (info.value.username.length > 10) {
    toast.warning('用户名')
    return
  }
  if (info.value.website.length > 100) {
    toast.warning('网站地址超长')
    return
  }

  if (config.public.googleRecaptchaSiteKey) {
    //@ts-ignore
    grecaptcha.ready(function () {
      //@ts-ignore
      grecaptcha.execute(config.public.googleRecaptchaSiteKey, { action: 'submit' }).then(async function (token) {
        submitComment(token)
      });
    });
  } else {
    submitComment()
  }

}

const submitComment = async (token?: string) => {
  pending.value = true
  const res = await $fetch('/api/comment/save', {
    method: 'POST',
    body: JSON.stringify({
      content: content.value,
      memoId: props.memoId,
      replyTo: props.reply,
      author: false,
      email: info.value.email,
      website: info.value.website,
      username: info.value.username,
      token
    })
  })

  if (res.success) {
    toast.success('评论成功')
    content.value = ''
    emit('commentAdded')
  } else {
    toast.warning(res.message || '评论失败')
  }
  pending.value = false
}

onMounted(async () => {

  if (props.reply) {
    placeholder.value = "回复给@" + props.reply
  }
})
</script>

<style scoped></style>
