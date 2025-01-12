<template>
  <div>
    <span v-if="props.comment.author == props.memoUserId" class="text-[#576b95] text-nowrap">
      {{ props.comment.username }}
      <UBadge color="gray" variant="solid" size="xs">作者</UBadge>
    </span>
    <span v-else class="text-[#576b95] text-nowrap">{{ props.comment.username }}</span>
   <template v-if="props.comment.replyTo">
     <span class="mx-1">回复</span>
     <span  class="text-[#576b95] text-nowrap">{{props.comment.replyTo}}</span>
   </template>
    <span class="mx-0.5">:</span>
    <span class="inline break-all cursor-pointer" @click="toggle">{{ props.comment.content }}</span>
    <span class="text-xs text-gray-400 ml-2 hidden sm:inline-block">{{$dayjs(props.comment.createdAt).fromNow()}}</span>
    <span class="text-xs text-gray-400 ml-2 inline-flex" v-if="(global.userinfo.id === props.memoUserId || global.userinfo.id === 1)">
      <Confirm @ok="removeComment">
        <svg class="w-3 h-3 cursor-pointer text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M12 12h2v12h-2zm6 0h2v12h-2z"/><path fill="currentColor" d="M4 6v2h2v20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8h2V6zm4 22V8h16v20zm4-26h8v2h-8z"/></svg>
      </Confirm>
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