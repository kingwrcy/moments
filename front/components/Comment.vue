<template>
  <div>
    <span v-if="props.comment.author == props.memoUserId" class="text-[#576b95] text-nowrap">
      {{ props.comment.username }}
    </span>
    <span v-else class="text-[#576b95] text-nowrap">
      <a v-if="props.comment.author" :href="`/user/${props.comment.author}`" target="_blank">
        {{ props.comment.username }}
      </a>
      <a v-else-if="props.comment.website" :href="formatWebsite(props.comment.website)" target="_blank">
        {{ props.comment.username }}
      </a>
      <span v-else>{{ props.comment.username }}</span>
    </span>
   <template v-if="props.comment.replyTo">
     <span class="mx-1">回复</span>
     <span  class="text-[#576b95] text-nowrap">{{props.comment.replyTo}}</span>
   </template>
    <span class="mx-0.5">:</span>
    <span class="inline break-all cursor-pointer" @click="toggle">{{ props.comment.content }}</span>
    <span class="text-xs text-gray-400 ml-2 hidden sm:inline-block">{{$dayjs(props.comment.createdAt).fromNow()}}</span>
    <span class="text-xs text-gray-400 ml-2 inline-flex" v-if="(global.userinfo.id === props.memoUserId || global.userinfo.id === 1)">
      <Confirm @ok="removeComment">
        <UIcon name="i-carbon-trash-can" class="cursor-pointer text-red-400"/>
      </Confirm>
    </span>

  </div>
  <CommentBox :memo-id="props.memoId" :reply-to="props.comment.username" :comment-id="props.comment.id" :reply-email="props.comment.email"/>
</template>

<script setup lang="ts">
import type {CommentVO, UserVO} from "~/types";
import CommentBox from "~/components/CommentBox.vue";
import {toast} from "vue-sonner";
import {memoChangedEvent} from "~/event";
import {useGlobalState} from "~/store";

const global = useGlobalState()
const currentCommentBox = useState('currentCommentBox')
const toggle = () => {
  const value = props.memoId + '#' + props.comment.id
  if (currentCommentBox.value === value) {
    currentCommentBox.value = ''
    return
  }
  currentCommentBox.value = value
}
const props = defineProps<{
  comment: CommentVO
  memoId: number
  memoUserId: number
}>()
const removeComment = async () => {
  await useMyFetch('/comment/remove?id=' + props.comment.id)
  toast.success("删除成功!")
  memoChangedEvent.emit(props.memoId)
}
const formatWebsite = (website: string) => {
  if (/^https?:\/\//i.test(website)) {
    return website;
  }
  return `http://${website}`
}
</script>

<style scoped>

</style>
