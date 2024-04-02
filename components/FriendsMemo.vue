<template>

  <div class="flex flex-row gap-4 text-sm border-x-0 pt-2">
    <img src="https://images.kingwrcy.cn/memo/20240386211829TtcOUOMaXyITlTkxhSjp" class="w-9 h-9 rounded">
    <div class="flex flex-col gap-.5 flex-1">
      <div class="text-[#576b95] cursor-default mb-1">{{ props.memo.user.nickname }}</div>
      <div class="text-sm friend-md" ref="el" v-html="props.memo.content.replaceAll(/\n/g, '<br/>')"> </div>
      <div class="text-[#576b95] cursor-pointer" v-if="hh > 96 && !showAll" @click="showMore">全文</div>
      <div class="text-[#576b95] cursor-pointer " v-if="showAll" @click="showLess">收起</div>
      <div class="relative flex flex-row justify-between">
        <div class="flex-1 text-gray text-xs">{{
        dayjs(props.memo.createdAt).locale('zh-cn').fromNow().replaceAll(/\s+/g,
          '') }}</div>
        <div @click="showToolbar = !showToolbar"
          class="mb-2 px-2 py-1 bg-[#f7f7f7] hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center">
          <img src="~/assets/img/dian.svg" class="w-3 h-3" />
        </div>
        <div class="text-xs absolute top-[-8px] right-[30px] bg-[#4c4c4c] rounded text-white p-2" v-if="showToolbar"
          ref="toolbarRef">
          <div class="flex flex-row gap-4">
            <div class="flex flex-row gap-4 cursor-pointer">
              <img src="~/assets/img/zan.svg" />
              <div>赞</div>
            </div>
            <div class="flex flex-row gap-4 cursor-pointer"
              @click="showCommentInput = !showCommentInput; showUserCommentArray = []; showToolbar = false">
              <img src="~/assets/img/comment.svg" />
              <div>评论</div>
            </div>
          </div>
        </div>
      </div>
      <div class="rounded bottom-shadow bg-[#f7f7f7] dark:bg-[#262626] flex flex-col gap-1  ">
        <template v-if="props.memo._count.comments > 0">
          <div class="p-4 flex flex-col gap-1">
            <div class="flex flex-col gap-2 text-sm" v-for="(comment, index) in props.memo.comments" :key="index">
              <div class="flex flex-row">
                <div class="text-[#576b95] text-nowrap">{{ comment.username ?? '匿名' }}</div>
                <div class="mx-1">:</div>
                <div class="break-all cursor-pointer" @click="toggleUserComment(index)">{{ comment.content }}</div>
              </div>
              <FriendsCommentInput :memoId="props.memo.id" :commentId="comment.id" :reply="comment.username"
                v-if="showUserCommentArray[index]" />
            </div>
          </div>
        </template>

        <FriendsCommentInput :memoId="props.memo.id" v-if="showCommentInput" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Memo } from '@/lib/types';
import { useElementSize, onClickOutside, watchOnce } from '@vueuse/core';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import 'dayjs/locale/zh-cn';
dayjs.extend(relativeTime)
const props = withDefaults(
  defineProps<{
    memo: Memo
  }>(), {}
)


const showAll = ref(false)
const showToolbar = ref(false)
const showCommentInput = ref(false)
const toolbarRef = ref(null)
const showUserCommentArray = ref<Array<boolean>>([])
const el = ref<any>(null)
let hh = ref(0)
const { height } = useElementSize(el)

const toggleUserComment = (index: number) => {
  const current = showUserCommentArray.value[index]
  showUserCommentArray.value = []
  showUserCommentArray.value[index] = !current
}

onClickOutside(toolbarRef, () => showToolbar.value = false)

const showMore = () => {
  showAll.value = true
  el.value.classList.remove('line-clamp-4')
}
const showLess = () => {
  showAll.value = false
  el.value.classList.add('line-clamp-4')
}



watchOnce(height, () => {
  hh.value = height.value
  if (height.value > 96) {
    el.value.classList.add('line-clamp-4')
  }
})

</script>