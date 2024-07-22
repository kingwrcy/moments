<template>
  <div :class="[item.pinned ? 'bg-slate-100' : '']">
    <div class="flex gap-2 text-sm dark:bg-neutral-800 px-4">
      <div class="avatar p-2">
        <NuxtLink :to="`/memo/${item.id}`">
          <UAvatar
              :src="item.user.avatarUrl"
              alt="Avatar"
          />
        </NuxtLink>
      </div>
      <div class="flex flex-col gap-1 p-2 w-full">
        <div class="username text-[#576b95] mb-1 dark:text-white cursor-pointer">
          <NuxtLink :to="`/user/${item.user.username}`">{{ item.user.nickname }}</NuxtLink>
        </div>
        <div class="content mb-2" v-html="md.render(item.content)">
        </div>

        <div v-if="item.externalTitle"
             class="flex flex-row gap-2 my-2 bg-[#f7f7f7] dark:bg-[#212121] items-center p-2 border rounded"
        >
          <img class="w-8 h-8" :src="item.externalFavicon" alt=""><a
            :href="item.externalUrl" target="_blank" class="text-[#576b95]">{{ item.externalTitle }}</a>
        </div>


        <div class="grid " :class="`grid-cols-${gridCols}`" v-if="images.length>0">
          <MyFancyBox v-for="(img,index) in images" :key="index">
            <img class="cursor-zoom-in rounded object-cover" :src="img" ></img>
          </MyFancyBox>
        </div>


        <div class="text-[#576b95] font-medium dark:text-white text-xs mt-2 mb-1 select-none">{{ location }}</div>

        <div class="flex justify-between items-center relative">
          <div class="text-xs text-[#9DA4B0]">{{ $dayjs(item.createdAt).fromNow() }}</div>
          <div @click="showToolbar=true"
               class="toolbar-icon px-2 py-1 bg-[#f7f7f7] dark:bg-slate-700 hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center"
          >
            <img
                src="data:image/svg+xml,%3csvg%20t='1709204592505'%20class='icon'%20viewBox='0%200%201024%201024'%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'%20p-id='16237'%20width='16'%20height='16'%3e%3cpath%20d='M229.2%20512m-140%200a140%20140%200%201%200%20280%200%20140%20140%200%201%200-280%200Z'%20p-id='16238'%20fill='%238a8a8a'%3e%3c/path%3e%3cpath%20d='M794.8%20512m-140%200a140%20140%200%201%200%20280%200%20140%20140%200%201%200-280%200Z'%20p-id='16239'%20fill='%238a8a8a'%3e%3c/path%3e%3c/svg%3e"
                class="w-3 h-3"/>
          </div>

          <div ref="toolbarRef" v-if="showToolbar"
               class="absolute top-[-8px] right-[32px] bg-[#4c4c4c] rounded text-white p-2 px-4">
            <div class="flex flex-row gap-2">
              <template v-if="global.userinfo.id === 1">
                <div class="flex flex-row gap-1 cursor-pointer items-center" @click="setPinned(item.id)">
                  <UIcon name="i-carbon-pin"/>
                  <div>{{ item.pinned ? '取消' : '' }}置顶</div>
                </div>
                <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>
              </template>
              <div class="flex flex-row gap-1 cursor-pointer items-center" @click="likeMemo(item.id)">
                <UIcon name="i-carbon-favorite" :class="[liked ? 'text-red-400' : '']"/>
                <div>赞</div>
              </div>
              <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>
              <div class="flex flex-row gap-1 cursor-pointer items-center" @click="doComment">
                <UIcon name="i-carbon-chat"/>
                <div>评论</div>
              </div>
              <template v-if="global&&global.userinfo.id === item.userId">
                <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>
                <div class="flex flex-row gap-1 cursor-pointer items-center" @click="go2Edit(item.id)">
                  <UIcon name="i-carbon-edit"/>
                  <div>编辑</div>
                </div>
              </template>
              <template v-if="(global.userinfo.id === 1 || global.userinfo.id === item.userId) ">
                <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>
                <Confirm @ok="removeMemo(item.id)" @cancel="showToolbar = false">
                  <div class="flex flex-row gap-1 cursor-pointer items-center">
                    <UIcon name="i-carbon-trash-can"/>
                    <div>删除</div>
                  </div>
                </Confirm>
              </template>
            </div>
          </div>
        </div>


        <div class="rounded bottom-shadow bg-[#f7f7f7] dark:bg-[#202020] flex flex-col gap-1"
        >
          <div class="flex flex-row py-2 px-4 gap-2 items-center text-sm" v-if="item.favCount>0">
            <UIcon name="i-carbon-favorite" class="text-red-500"/>
            <div class="text-[#576b95]"><span class="mx-1">{{ item.favCount }}</span>位访客</div>
          </div>
          <div class="flex flex-col gap-1">
            <CommentBox :comment-id="0" :memo-id="item.id"/>
            <div :class="[item.comments && item.comments.length>0 ? 'py-2' : '']">
              <div class="px-4 relative flex-col text-sm" v-for="c in item.comments" :key="c.id"
                   v-if="item.comments && item.comments.length>0">
                <Comment :comment="c" :memo-id="item.id" :memo-user-id="item.user.id"/>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">
import type {MemoVO} from "~/types";
import {toast} from "vue-sonner";
import {memoChangedEvent, memoReloadEvent} from "~/event";
import Comment from "~/components/Comment.vue";
import {useGlobalState} from "~/store";
import markdownit from 'markdown-it'

const md = markdownit({
  html: true,
  linkify: true,
  typographer: true,
  breaks: true,
})
const currentCommentBox = useState('currentCommentBox')
const props = defineProps<{
  memo: MemoVO
}>()
const item = computed(() => {
  return props.memo
})
const global = useGlobalState()
const route = useRoute()
const isDetailPage = computed(() => {
  return route.path.startsWith("/memo/")
})
const showToolbar = ref(false)
const toolbarRef = ref(null)
const liked = ref(false)
onClickOutside(toolbarRef, () => showToolbar.value = false)

const location = computed(() => {
  return (item.value.location || "").replaceAll(" ", " · ")
})

const doComment = () => {
  const value = item.value.id + '#0'
  console.log('value is ', value, '===>', currentCommentBox.value)
  if (currentCommentBox.value === value) {
    currentCommentBox.value = ''
  } else {
    currentCommentBox.value = value
  }
  showToolbar.value = false
}

const go2Edit = async (id: number) => {
  await navigateTo('/edit/' + id)
}

const removeMemo = async (id: number) => {
  await useMyFetch('/memo/remove?id=' + id)
  toast.success("删除成功!")
  if (isDetailPage) {
    await navigateTo('/')
  } else {
    memoReloadEvent.emit()
  }
}
const setPinned = async (id: number) => {
  await useMyFetch('/memo/setPinned?id=' + id)
  toast.success("操作成功!")
  showToolbar.value = false
  memoReloadEvent.emit()
}
const likeMemo = async (id: number) => {
  showToolbar.value = false
  const likes = JSON.parse(localStorage.getItem('likeMemos') || '[]') as Array<number>
  if (likes.includes(id)) {
    toast.warning("您已经点赞过了!")
    return
  }
  await useMyFetch('/memo/like?id=' + id)
  toast.success("点赞成功!")
  likes.push(id)
  localStorage.setItem('likeMemos', JSON.stringify(likes))
  memoChangedEvent.emit(id)
  liked.value = true
}

onMounted(() => {
  const likes = JSON.parse(localStorage.getItem('likeMemos') || '[]') as Array<number>
  liked.value = likes.findIndex(r => r === item.value.id) >= 0

})

const images = computed(() => {
  if (!item.value.imgs) {
    return []
  }
  return item.value.imgs.split(",")
})

const gridCols = computed(() => {
  const len = images.value.length
  if (len === 1) {
    return 1
  }
  return 3
})

// const content = computed(() => {
//   if (!item.value.content) {
//     return ""
//   }
//   return item.value.content.replaceAll("\n", "<br/>")
// })
</script>

<style lang="scss" scoped>

</style>