<template>
  <div class="flex flex-col gap-2 mt-2" v-if="currentCommentBox === `${props.memoId}#${props.commentId}`">
    <div class="relative">
      <UTextarea :rows="3" autofocus :placeholder="`回复给${replyTo}`" v-model="state.content"/>
      <div class="animate-bounce absolute right-2 bottom-1 cursor-pointer text-xl select-none" @click="toggleEmoji">😊</div>
    </div>
    <Emoji v-if="emojiShow" @selected="emojiSelected"/>
    <div class="flex gap-1">
      <template v-if="!currentUser">
        <UInput placeholder="姓名" v-model="state.username"/>
        <UInput placeholder="网站" v-model="state.website"/>
      </template>
      <UButton color="white" @click="doComment">评论</UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {toast} from "vue-sonner";
import {memoChangedEvent} from "~/event";
import Emoji from "~/components/Emoji.vue";

const emojiShow = ref(false)
const currentCommentBox = useState('currentCommentBox')
const props = defineProps<{
  commentId: number
  memoId: number
  replyTo: string
}>()
const state = reactive({
  content: "",
  memoId: props.memoId,
  replyTo: props.replyTo,
  username: "",
  website: "",
})
const currentUser = useState<UserVO>('userinfo')
const doComment = async () => {
  await useMyFetch('/comment/add', state)
  toast.success("评论成功!")
  currentCommentBox.value = ''
  state.username = ''
  state.content = ''
  state.website = ''
  memoChangedEvent.emit(props.memoId)
}
const toggleEmoji = ()=>{
  emojiShow.value = !emojiShow.value
}
const emojiSelected = (emoji:string)=>{
  state.content = state.content+emoji
}
</script>

<style scoped>

</style>