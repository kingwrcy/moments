<template>
  <div>
    <span v-if="props.comment.author == props.memoUserId" class="text-[#576b95] text-nowrap">
      {{ props.comment.username }}
      <UBadge color="gray" variant="solid" size="xs">作者</UBadge>
    </span>
    <span v-else class="text-[#576b95] text-nowrap">{{ props.comment.username }}</span>
   <template v-if="props.comment.replyTo">
     <span class="mx-1">回复</span>
     <span>{{props.comment.replyTo}}</span>
   </template>
    <span class="mx-0.5">:</span>
    <span class="inline break-all cursor-pointer" @click="toggle">
    <span>{{ props.comment.content }}</span>
    <UIcon v-if="(global.userinfo.id === props.memoUserId || global.userinfo.id === 1)" name="i-carbon-trash-can"
           @click="removeComment" class="ml-4 text-red-400"/>
    </span>
  </div>
  <CommentBox :memo-id="props.memoId" :reply-to="props.comment.username" :comment-id="props.comment.id"/>
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
</script>

<style scoped>

</style>