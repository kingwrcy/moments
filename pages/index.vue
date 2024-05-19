<template>
  <div class="mt-12">

    <MemoInput v-if="token" @memo-added="loadData(1, 'add')"
      @memo-updated="loadData(Math.ceil((memoUpdateIndex + 1) / state.size), 'edit')" />

    <div class="content flex flex-col divide-y divide-[#C0BEBF]/10">
      <div v-if="(state.memoList).length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">空空如也，赶紧去发表Moments吧!</div>
        <Button @click="navigateTo('/login')">登录</Button>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in state.memoList" :index="index" :key="index"
        :show-more="true" @memo-update="loadData(Math.ceil((index + 1) / state.size), 'edit')" />
    </div>
    <div class="cursor-pointer text-center text-sm opacity-70  my-4" @click="loadData(state.page + 1, 'more')"
      v-if="state.hasNext">- 加载更多 -
    </div>
  </div>
</template>

<script setup lang="ts">
import type { User, Memo, PublicConfig } from '~/lib/types';
const token = useCookie('token')
const userinfo = useState<User>('userinfo')
const publicConfig = useState<PublicConfig>('publicConfig')
const route = useRoute()
useHead({
  title: userinfo.value.title || '极简朋友圈',
})

const state = reactive({
  memoList: Array<Memo>(),
  page: 1,
  size: publicConfig.value.pageSize,
  hasNext: false
})

const memoUpdateIndex = useState<number>('memoUpdateIndex', () => -1)

const { data } = await useFetch('/api/memo/list', {
  key: 'memoList',
  method: 'POST',
  body: JSON.stringify({
    page: state.page,
    size: state.size,
    tag: route.query.tag
  })
})

state.memoList = data.value?.data as any as Memo[]
state.hasNext = data.value?.hasNext || false

watch(() => route.query, async () => {
  await loadData(1, 'refresh')
})

const loadData = async (page: number, type: 'add' | 'edit' | 'more' | 'refresh') => {

  page = page <= 0 ? 1 : page;
  const { data, hasNext } = await $fetch('/api/memo/list', {
    key: 'memoList',
    method: 'POST',
    body: JSON.stringify({
      page, size: state.size, tag: route.query.tag
    })
  })
  const list = data as any as Memo[]
  if (type === 'add') {
    //删除首页
    state.memoList.splice(0, Math.min(state.size, data.length))
    //添加到首页
    state.memoList.unshift(...list)
  } else if (type === 'edit') {
    //删除当前页,添加到当前页
    const currentPageStart = (page - 1) * state.size
    state.memoList.splice(currentPageStart, Math.min(state.size, state.memoList.length - currentPageStart), ...list)
  } else if (type === 'more') {
    //添加到最后页
    state.memoList.push(...list)
    state.page = page
    state.hasNext = hasNext
  } else {
    //刷新
    state.memoList = list
    state.page = 1
    state.hasNext = hasNext
  }
}

</script>

<style scoped></style>