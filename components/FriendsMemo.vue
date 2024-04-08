<template>

  <div class="flex flex-row gap-4 text-sm border-x-0 pt-2 bg-white">
    <NuxtImg :src="props.memo.user.avatarUrl" class="w-9 h-9 rounded"/>
    <div class="flex flex-col gap-.5 flex-1">
      <div class="text-[#576b95] cursor-default mb-1">{{ props.memo.user.nickname }}</div>
      <div class="text-sm friend-md" ref="el" v-html="props.memo.content.replaceAll(/\n/g, '<br/>')"> </div>

      <iframe class="rounded" frameborder="no" border="0" marginwidth="0" marginheight="0" width=330 height=86
      :src="props.memo.music163Url" v-if="props.memo.music163Url"></iframe>

      <iframe class="w-full h-[250px] my-2" v-if="props.memo.bilibiliUrl"
        :src="props.memo.bilibiliUrl"
        scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>


      <div v-if="props.memo.imgs">
        <FancyBox class="grid grid-cols-3 my-1 gap-2" :options="{
          Carousel: {
            infinite: false,
          },
        }">
            <NuxtImg loading="lazy" placeholder format="avif,webp" :src="`${img}`" class="cursor-pointer rounded"
              v-for="(img, index) in props.memo.imgs?.split(',')" :key="index" />
        </FancyBox>
      </div>
      <div class="text-[#576b95] cursor-pointer" v-if="hh > 96 && !showAll" @click="showMore">全文</div>
      <div class="text-[#576b95] cursor-pointer " v-if="showAll" @click="showLess">收起</div>
      <div class="relative flex flex-row justify-between select-none my-1">
        <div class="flex-1 text-gray text-xs text-[#9DA4B0] ">{{
          dayjs(props.memo.createdAt).locale('zh-cn').fromNow().replaceAll(/\s+/g,
            '') }}</div>
        <div @click="showToolbar = !showToolbar"
          class="mb-2 px-2 py-1 bg-[#f7f7f7] dark:bg-slate-700 hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center">
          <img src="~/assets/img/dian.svg" class="w-3 h-3" />
        </div>
        <div class="text-xs absolute top-[-8px] right-[30px] bg-[#4c4c4c] rounded text-white p-2" v-if="showToolbar"
          ref="toolbarRef">
          <div class="flex flex-row gap-4">
            <div class="flex flex-row gap-2 cursor-pointer items-center" v-if="token" @click="editMemo">
              <FilePenLine :size=14 />
              <div>编辑</div>
            </div>
            <AlertDialog>
              <AlertDialogTrigger asChild>
                <div class="flex flex-row gap-2 cursor-pointer items-center" v-if="token">
                  <Trash2 :size=14 />
                  <div>删除</div>
                </div>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>确定删除吗?</AlertDialogTitle>
                  <AlertDialogDescription>
                    无法恢复,你确定删除吗?
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>取消</AlertDialogCancel>
                  <AlertDialogAction @click="removeMemo">确定</AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>

            <div class="flex flex-row gap-2 cursor-pointer items-center" @click="like">
              <Heart :size=14 v-if="likeList.findIndex((id) => id === props.memo.id) < 0" />
              <HeartCrack :size=14 v-else />
              <div>{{ likeList.findIndex((id) => id === props.memo.id) >= 0 ? '取消' : '赞' }}</div>
            </div>
            <div class="flex flex-row gap-2 cursor-pointer items-center"
              @click="showCommentInput = !showCommentInput; showUserCommentArray = []; showToolbar = false">
              <MessageSquareMore :size=14 />
              <div>评论</div>
            </div>
          </div>
        </div>
      </div>
      <div class="rounded bottom-shadow bg-[#f7f7f7] dark:bg-[#262626] flex flex-col gap-1  ">
        <div class="flex flex-row py-2 px-4 gap-2 items-center text-sm" v-if="props.memo.favCount > 0">
          <Heart :size=14 />
          <div class="text-[#576b95]"><span class="mx-1">{{ props.memo.favCount }}</span>位访客赞过</div>
        </div>
        <template v-if="props.memo.comments.length > 0">
          <div class="px-4 py-2 flex flex-col gap-1">
            <div class="flex flex-col gap-2 text-sm" v-for="(comment, index) in props.memo.comments" :key="index">
              <div class="">
                <span class="text-[#576b95] text-nowrap">{{ comment.username ?? '匿名' }}
                  <b v-if="comment.author" class="border text-xs border-[#C64A4A] rounded px-0.5 text-[#C64A4A]">作者</b></span>
                <span v-if="comment.replyTo" class="text-nowrap mx-1">回复<span class="text-[#576b95] ml-1">{{
                  comment.replyTo }}</span> </span>
                <span class="mr-0.5">:</span>
                <span :title="`点击回复${comment.username}`" class="break-all cursor-pointer"
                  @click="toggleUserComment(index)">{{ comment.content }}</span>
              </div>
              <FriendsCommentInput @commentAdded="refreshComment" :memoId="props.memo.id" :commentId="comment.id"
                :reply="comment.username" v-if="showUserCommentArray[index]" />
            </div>
            <div v-if="props.memo._count.comments > 5 && props.showMore" class="text-[#576b95] cursor-pointer"
              @click="navigateTo(`/detail/${props.memo.id}`)">查看更多...</div>
          </div>
        </template>

        <FriendsCommentInput :memoId="props.memo.id" @commentAdded="refreshComment" v-if="showCommentInput" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Memo } from '@/lib/types';
import { useElementSize, onClickOutside, watchOnce, useStorage } from '@vueuse/core';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import 'dayjs/locale/zh-cn';
import { Heart, HeartCrack, MessageSquareMore, Trash2, FilePenLine } from 'lucide-vue-next'
import { memoUpdateEvent } from '@/lib/event'
import { toast } from 'vue-sonner';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'
const token = useCookie('token')



dayjs.extend(relativeTime)
const props = withDefaults(
  defineProps<{
    memo: Memo,
    showMore: boolean
  }>(), {}
)

const emit = defineEmits(['memo-update'])

const showAll = ref(false)
const showToolbar = ref(false)
const showCommentInput = ref(false)
const toolbarRef = ref(null)
const showUserCommentArray = ref<Array<boolean>>([])
const el = ref<any>(null)
let hh = ref(0)
const { height } = useElementSize(el)
const likeList = useStorage<Array<number>>('likeList', [])

const like = async () => {
  showToolbar.value = false
  const contain = likeList.value.find((id) => id === props.memo.id)
  const res = await $fetch('/api/memo/like', {
    method: 'POST',
    body: JSON.stringify({
      memoId: props.memo.id,
      like: !contain
    })
  })
  if (res.success) {
    if (contain) {
      likeList.value = likeList.value.filter((id) => id !== props.memo.id)
    } else {
      likeList.value.push(props.memo.id)
    }
    emit('memo-update')
  }
}

const removeMemo = async () => {
  showToolbar.value = false
  const res = await $fetch('/api/memo/remove', {
    method: 'POST',
    body: JSON.stringify({
      memoId: props.memo.id
    })
  })
  if (res.success) {
    toast.success('删除成功')
    emit('memo-update')
  }
}

const editMemo = async () => {
  showToolbar.value = false
  memoUpdateEvent.emit(props.memo)
}


const refreshComment = async () => {
  emit('memo-update', props.memo)
  showUserCommentArray.value = []
  showCommentInput.value = false
}

const toggleUserComment = (index: number) => {
  const current = showUserCommentArray.value[index]
  showUserCommentArray.value = []
  showUserCommentArray.value[index] = !current
  showCommentInput.value = false
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