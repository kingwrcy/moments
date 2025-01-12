<template>
  <div class="header relative mb-14" v-if="$route.path === `/memo/${item.id}`">
    <div :class="{ 'bg-[#4c4c4c]/80 z-10': y > 100 }" class="flex fixed justify-between items-center p-4 w-full md:w-[567px] text-white top-0">
      <NuxtLink class="flex items-center" title="返回主页">
        <svg @click="navigateTo('/')" class="w-5 h-5 cursor-pointer mr-4"xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M10 16L20 6l1.4 1.4l-8.6 8.6l8.6 8.6L20 26z"/></svg>
        <span>详情</span>
      </NuxtLink>
      <svg v-if="global.userinfo.id === 1 || global.userinfo.id === item.userId" class="w-5 h-5 cursor-pointer" @click="moreToolbar = true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><circle cx="8" cy="16" r="2" fill="currentColor"/><circle cx="16" cy="16" r="2" fill="currentColor"/><circle cx="24" cy="16" r="2" fill="currentColor"/></svg>
    </div>
  </div>
  <div>
    <div class="relative flex gap-4 text-sm dark:bg-neutral-800 p-4"
         :class="[item.pinned ? 'bg-slate-100 dark:bg-neutral-700' : '']">
      <div class="avatar ">
        <NuxtLink :to="`/memo/${item.id}`">
          <UAvatar
              :src="item.user.avatarUrl"
              alt="Avatar"
          />
        </NuxtLink>
      </div>
      <div class="flex flex-col gap-1  flex-1">
        <div class="username text-[#576b95] mb-1 dark:text-white  flex justify-between">
          <NuxtLink class="cursor-pointer" :to="`/user/${item.user.id}`">{{ item.user.nickname }}</NuxtLink>
          <div class="flex">
            <svg v-if="item.pinned" class="w-3 h-3 text-red-500 ml-2 dark:text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M28.59 13.31L30 11.9L20 2l-1.31 1.42l1.18 1.18l-11.49 9.72l-1.72-1.71L5.25 14l5.66 5.68L2 28.58L3.41 30l8.91-8.91L18 26.75l1.39-1.42l-1.71-1.71l9.72-11.49ZM16.26 22.2L9.8 15.74L21.29 6L26 10.71Z"/></svg>
            <svg v-if="item.showType === 0" class="w-3 h-3 text-red-500 ml-2 dark:text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M24 14h-2V8a6 6 0 0 0-12 0v6H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2M12 8a4 4 0 0 1 8 0v6h-8Zm12 20H8V16h16Z"/></svg>
          </div>
        </div>
        <div class="mb-2">
          <div :style="getMemoMaxHeightStyle()" class="overflow-hidden">
            <div class="markdown-content " ref="contentRef"
                 v-html="content"></div>
          </div>
          <div class="text-[#576b95] text-sm my-1 cursor-pointer"
               @click="doShowMore" v-if="showMore">{{ getMemoMaxHeightStyle() === '' ? '收起' : '全文' }}
          </div>
          <div class="flex gap-2 mt-2" v-if="tags.length > 0">

            <span v-for="(tag,index) in tags" :key="`tag-${index}`">
              <NuxtLink :to="`/tags/${item.user.username}/${tag}`">
                <UBadge size="xs" color="gray" variant="solid">{{ tag }}</UBadge>
              </NuxtLink>
            </span>
          </div>
        </div>

        <div class="flex flex-col gap-2">
          <external-url-preview :favicon="item.externalFavicon" :title="item.externalTitle" :url="item.externalUrl"
                                v-if="item.externalFavicon&&item.externalTitle&&item.externalUrl"/>
          <upload-image-preview :imgs="item.imgs||''" :memo-id="item.id"/>

          <music-preview v-if="extJSON.music && extJSON.music.id" v-bind="extJSON.music"/>
          <douban-book-preview v-if="extJSON.doubanBook && extJSON.doubanBook.title" :book="extJSON.doubanBook"/>
          <douban-movie-preview v-if="extJSON.doubanMovie && extJSON.doubanMovie.title" :movie="extJSON.doubanMovie"/>
          <youtube-preview v-if="extJSON.video && extJSON.video.type === 'youtube' && extJSON.video.value"
                           :url="extJSON.video.value"/>
          <bilibili-preview v-if="extJSON.video && extJSON.video.type === 'bilibili' && extJSON.video.value"
                            :url="extJSON.video.value"/>
          <video-preview v-if="extJSON.video && extJSON.video.type === 'online' && extJSON.video.value"
                         :url="extJSON.video.value"/>
        </div>

        <div class="text-[#576b95] font-medium dark:text-white text-xs mt-2 mb-1 select-none flex items-center gap-0.5"
             v-if="location">
          <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M16 18a5 5 0 1 1 5-5a5.006 5.006 0 0 1-5 5m0-8a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3"/><path fill="currentColor" d="m16 30l-8.436-9.949a35 35 0 0 1-.348-.451A10.9 10.9 0 0 1 5 13a11 11 0 0 1 22 0a10.9 10.9 0 0 1-2.215 6.597l-.001.003s-.3.394-.345.447ZM8.813 18.395s.233.308.286.374L16 26.908l6.91-8.15c.044-.055.278-.365.279-.366A8.9 8.9 0 0 0 25 13a9 9 0 1 0-18 0a8.9 8.9 0 0 0 1.813 5.395"/></svg>
          <span>{{ location }}</span>
        </div>

        <div class="flex justify-between items-center relative">
          <div class="flex text-xs text-[#9DA4B0]">{{
              sysConfig.timeFormat === 'timeAgo' ? $dayjs(item.createdAt).fromNow() : $dayjs(item.createdAt).format("YYYY-MM-DD HH:mm:ss")
            }}
          </div>
          <div @click="showToolbar=true"
               class="toolbar-icon px-2 py-1 bg-[#f7f7f7] dark:bg-slate-700 hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center"
          >
            <img
                src="data:image/svg+xml,%3csvg%20t='1709204592505'%20class='icon'%20viewBox='0%200%201024%201024'%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'%20p-id='16237'%20width='16'%20height='16'%3e%3cpath%20d='M229.2%20512m-140%200a140%20140%200%201%200%20280%200%20140%20140%200%201%200-280%200Z'%20p-id='16238'%20fill='%238a8a8a'%3e%3c/path%3e%3cpath%20d='M794.8%20512m-140%200a140%20140%200%201%200%20280%200%20140%20140%200%201%200-280%200Z'%20p-id='16239'%20fill='%238a8a8a'%3e%3c/path%3e%3c/svg%3e"
                class="w-3 h-3"/>
          </div>

          <div ref="toolbarRef" v-if="showToolbar"
               class="absolute top-[-8px] right-[32px] bg-[#4c4c4c] rounded text-white p-2">
            <div class="flex flex-row gap-2">
              <div class="flex flex-row gap-1 cursor-pointer items-center px-4" @click="likeMemo(item.id)">
                <svg class="w-4 h-4 cursor-pointer" :class="[liked ? 'text-red-400' : '']" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M22.45 6a5.47 5.47 0 0 1 3.91 1.64a5.7 5.7 0 0 1 0 8L16 26.13L5.64 15.64a5.7 5.7 0 0 1 0-8a5.48 5.48 0 0 1 7.82 0l2.54 2.6l2.53-2.58A5.44 5.44 0 0 1 22.45 6m0-2a7.47 7.47 0 0 0-5.34 2.24L16 7.36l-1.11-1.12a7.49 7.49 0 0 0-10.68 0a7.72 7.72 0 0 0 0 10.82L16 29l11.79-11.94a7.72 7.72 0 0 0 0-10.82A7.5 7.5 0 0 0 22.45 4"/></svg>
                <div>赞</div>
              </div>
              <template v-if="sysConfig.enableComment">
                <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>
                <div class="flex flex-row gap-1 cursor-pointer items-center px-4" @click="doComment">
                  <svg class="w-4 h-4 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="currentColor" fill-rule="evenodd" d="M10.503 17.8H20.5a.3.3 0 0 0 .3-.3v-12a.3.3 0 0 0-.3-.3h-17a.3.3 0 0 0-.3.3v12a.3.3 0 0 0 .3.3h4.7v2.303zM11 19l-2.293 2.293A1 1 0 0 1 7 20.586V19H3.5A1.5 1.5 0 0 1 2 17.5v-12A1.5 1.5 0 0 1 3.5 4h17A1.5 1.5 0 0 1 22 5.5v12a1.5 1.5 0 0 1-1.5 1.5z"/></svg>
                  <div>评论</div>
                </div>
              </template>
            </div>
          </div>
          <template>
            <UModal v-model="moreToolbar" :ui="{container: 'sm:items-end'}">
              <div class="flex items-center justify-center gap-8 p-4 text-gray-500 dark:text-white">
                <template v-if="global.userinfo.id === 1">
                  <div class="flex flex-col gap-1 cursor-pointer items-center" @click="setPinned(item.id)">
                    <svg class="text-[#9fc84a] w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M28.59 13.31L30 11.9L20 2l-1.31 1.42l1.18 1.18l-11.49 9.72l-1.72-1.71L5.25 14l5.66 5.68L2 28.58L3.41 30l8.91-8.91L18 26.75l1.39-1.42l-1.71-1.71l9.72-11.49ZM16.26 22.2L9.8 15.74L21.29 6L26 10.71Z"/></svg>
                    <div>{{ item.pinned ? '取消' : '' }}置顶</div>
                  </div>
                </template>
                <template v-if="global&&global.userinfo.id === item.userId">
                  <div class="flex flex-col gap-1 cursor-pointer items-center" @click="go2Edit(item.id)">
                    <svg class="text-[#9fc84a] w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M2 26h28v2H2zM25.4 9c.8-.8.8-2 0-2.8l-3.6-3.6c-.8-.8-2-.8-2.8 0l-15 15V24h6.4zm-5-5L24 7.6l-3 3L17.4 7zM6 22v-3.6l10-10l3.6 3.6l-10 10z"/></svg>
                    <div>编辑</div>
                  </div>
                </template>
                <template v-if="(global.userinfo.id === 1 || global.userinfo.id === item.userId) ">
                <Confirm @ok="removeMemo(item.id)" @cancel="moreToolbar = false">
                  <div class="flex flex-col gap-1 cursor-pointer items-center">
                    <svg class="text-[#9fc84a] w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M12 12h2v12h-2zm6 0h2v12h-2z"/><path fill="currentColor" d="M4 6v2h2v20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8h2V6zm4 22V8h16v20zm4-26h8v2h-8z"/></svg>
                  <div>删除</div>
                  </div>
                </Confirm>
              </template>
              </div>
            </UModal>
          </template>
        </div>

        <div class="rounded bottom-shadow bg-[#f7f7f7] dark:bg-[#202020] flex flex-col gap-1"
        >
          <div class="flex flex-row py-2 px-4 gap-2 items-center text-sm" v-if="item.favCount>0">
            <svg class="w-4 h-4 text-red-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M22.45 6a5.47 5.47 0 0 1 3.91 1.64a5.7 5.7 0 0 1 0 8L16 26.13L5.64 15.64a5.7 5.7 0 0 1 0-8a5.48 5.48 0 0 1 7.82 0l2.54 2.6l2.53-2.58A5.44 5.44 0 0 1 22.45 6m0-2a7.47 7.47 0 0 0-5.34 2.24L16 7.36l-1.11-1.12a7.49 7.49 0 0 0-10.68 0a7.72 7.72 0 0 0 0 10.82L16 29l11.79-11.94a7.72 7.72 0 0 0 0-10.82A7.5 7.5 0 0 0 22.45 4"/></svg>
            <div class="text-[#576b95]"><span class="mx-1">{{ item.favCount }}</span>位访客</div>
          </div>
          <div class="flex flex-col gap-1" v-if="sysConfig.enableComment">
            <CommentBox :comment-id="0" :memo-id="item.id"/>
            <div class="space-y-1" :class="[item.comments && item.comments.length>0 ? 'py-2' : '']">
              <div class="px-4 relative flex-col text-sm " v-for="c in item.comments" :key="c.id"
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
import type {ExtDTO, MemoVO, SysConfigVO} from "~/types";
import {toast} from "vue-sonner";
import {memoChangedEvent, memoReloadEvent} from "~/event";
import Comment from "~/components/Comment.vue";
import {useGlobalState} from "~/store";
import {md} from "~/utils"

const showMore = ref(false)
const showMoreClicked = ref(false)
const isDetailPage = computed(() => {
  return route.path.startsWith("/memo/")
})
const contentRef = ref<HTMLDivElement | null>(null)
const sysConfig = useState<SysConfigVO>('sysConfig')
const route = useRoute()
const {y} = useWindowScroll()

const getMemoMaxHeightStyle = () => {
  if (isDetailPage.value || showMoreClicked.value) {
    return ""
  }
  if (sysConfig.value.memoMaxHeight) {
    return `max-height:${sysConfig.value.memoMaxHeight}px`
  }
  return ""
}
const currentCommentBox = useState('currentCommentBox')
const props = defineProps<{
  memo: MemoVO,
}>()
const extJSON = computed(() => {
  return JSON.parse(props.memo.ext || "{}") as ExtDTO
})
const item = computed(() => {
  return props.memo
})

const global = useGlobalState()

const moreToolbar = ref(false)

const showToolbar = ref(false)
const toolbarRef = ref(null)
const liked = ref(false)
onClickOutside(toolbarRef, () => showToolbar.value = false)

const location = computed(() => {
  return (item.value.location || "").replaceAll(" ", " · ")
})

const tags = computed(() => {
  const tagsStr = item.value.tags
  if (!tagsStr) {
    return []
  }
  const len = tagsStr.length
  if (tagsStr[len - 1] === ',') {
    return tagsStr.substring(0, len - 1).split(",")
  }
  return tagsStr.split(",")
})

const doComment = () => {
  const value = item.value.id + '#0'
  if (currentCommentBox.value === value) {
    currentCommentBox.value = ''
  } else {
    currentCommentBox.value = value
  }
  showToolbar.value = false
}

const doShowMore = () => {
  showMoreClicked.value = !showMoreClicked.value
}

const go2Edit = async (id: number) => {
  await navigateTo('/edit/' + id)
}

const removeMemo = async (id: number) => {
  await useMyFetch('/memo/remove?id=' + id)
  toast.success("删除成功!")
  if (isDetailPage.value) {
    await navigateTo('/')
  } else {
    memoReloadEvent.emit()
  }
  moreToolbar.value = false
}
const setPinned = async (id: number) => {
  await useMyFetch('/memo/setPinned?id=' + id)
  toast.success("操作成功!")
  if (isDetailPage.value) {
    await navigateTo('/')
  } else {
    memoReloadEvent.emit()
  }
  moreToolbar.value = false
}

const doLike = async (id: number, token: string = '') => {
  const likes = JSON.parse(localStorage.getItem('likeMemos') || '[]') as Array<number>
  await useMyFetch(`/memo/like?id=${id}&token=${token}`)
  toast.success("点赞成功!")
  likes.push(id)
  localStorage.setItem('likeMemos', JSON.stringify(likes))
  memoChangedEvent.emit(id)
  liked.value = true
}


const likeMemo = async (id: number) => {
  showToolbar.value = false
  const likes = JSON.parse(localStorage.getItem('likeMemos') || '[]') as Array<number>
  if (likes.includes(id)) {
    toast.warning("您已经点赞过了!")
    return
  }

  if (sysConfig.value.enableGoogleRecaptcha) {
    grecaptcha.ready(() => {
      grecaptcha.execute(sysConfig.value.googleSiteKey, {action: 'newComment'}).then(async (token) => {
        await doLike(id, token)
      })
    })
  } else {
    await doLike(id)
  }


}


onMounted(() => {
  const likes = JSON.parse(localStorage.getItem('likeMemos') || '[]') as Array<number>
  liked.value = likes.findIndex(r => r === item.value.id) >= 0
  if (!isDetailPage.value) {
    setTimeout(() => {
      const {height} = useElementSize(contentRef.value);
      if (height.value > sysConfig.value.memoMaxHeight) {
        showMore.value = true
      }
    }, 20)
  }
})


const content = computed(() => {
  if (item.value.content && item.value.content.length > 0) {
    try{
      return md.render(item.value.content)
    }catch (e) {
      console.log('内容渲染错误,请重新编辑',e)
      return "内容渲染错误,请重新编辑"
    }
  }
  return ""
})

</script>

<style lang="scss" scoped>

</style>