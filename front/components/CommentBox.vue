<template>
  <div class="px-4 py-2 flex flex-col gap-2 mt-2" v-if="currentCommentBox === pid">
    <div class="relative">
      <UTextarea :rows="3" autofocus :placeholder="replyTo ? `回复给${replyTo}` : ''" v-model="state.content"/>
      <div class="animate-bounce absolute right-2 bottom-1 cursor-pointer text-xl select-none" @click="toggleEmoji">😊
      </div>
    </div>
    <Emoji v-if="emojiShow" @selected="emojiSelected"/>
    <div class="flex gap-1">
      <template v-if="!global.userinfo.token">
        <UInput placeholder="姓名" v-model="state.username"/>
        <UInput placeholder="网站" v-model="state.website"/>
      </template>
      <UButton color="white" @click="doComment">发布评论</UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import {toast} from "vue-sonner";
import {memoChangedEvent} from "~/event";
import Emoji from "~/components/Emoji.vue";
import {useGlobalState} from "~/store";
import {useStorage} from '@vueuse/core'
const props = defineProps<{
  commentId: number
  memoId: number
  replyTo?: string
}>()
const pid = computed(()=>{
  return `${props.memoId}#${props.commentId}`
})
const global = useGlobalState()
const localCommentUserinfo = useStorage('localCommentUserinfo', {
  username: "",
  website: ""
})
const emojiShow = ref(false)
const currentCommentBox = useState('currentCommentBox')

const state = reactive({
  content: "",
  memoId: props.memoId,
  replyTo: props.replyTo,
  username: localCommentUserinfo.value.username,
  website: localCommentUserinfo.value.website,
})
const doComment = async () => {
  if (!global.value.userinfo.token) {
    localCommentUserinfo.value = {
      username: state.username,
      website: state.website,
    }
  }
  await useMyFetch('/comment/add', state)
  toast.success("评论成功!")
  currentCommentBox.value = ''
  state.username = ''
  state.content = ''
  state.website = ''
  memoChangedEvent.emit(props.memoId)
}
const toggleEmoji = () => {
  emojiShow.value = !emojiShow.value
}
const emojiSelected = (emoji: string) => {
  state.content = state.content + emoji
}
</script>

<style scoped>

</style>