<template>
  <div class="mt-12">

    <MemoInput v-if="token" @memo-added="loadData(1, 'add')"
      @memo-updated="loadData(Math.ceil((memoUpdateIndex + 1) / state.size), 'edit')" />

    <div class="content flex flex-col divide-y divide-[#C0BEBF]/10 gap-2">
      <div v-if="(data?.data as any as Memo[]).length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">什么也没有,赶紧去登录发表Moments吧!</div>
        <Button @click="navigateTo('/login')">去登录</Button>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in data?.data as any as Memo[]" :index="index" :key="index"
        :show-more="true" @memo-update="loadData(Math.ceil((index + 1) / state.size), 'edit')" />
    </div>
    <div class="cursor-pointer text-center text-sm opacity-70  my-4" @click="loadData(state.page + 1, 'more')"
      v-if="state.hasNext">点击加载更多...
    </div>
  </div>
</template>

<script setup lang="ts">
import type { User, Memo, PublicConfig } from '~/lib/types';
const token = useCookie('token')
const userinfo = useState<User>('userinfo')
const publicConfig = useState<PublicConfig>('publicConfig')
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
    size: state.size
  })
})
state.memoList = data.value?.data as any as Memo[]
state.hasNext = data.value?.hasNext || false



const loadData = async (page: number, type: 'add' | 'edit' | 'more') => {
  page = page <= 0 ? 1 : page;
  const { data, hasNext } = await $fetch('/api/memo/list', {
    key: 'memoList',
    method: 'POST',
    body: JSON.stringify({
      page, size: state.size
    })
  })
  if (type === 'add') {
    //删除首页
    state.memoList.splice(0, Math.min(state.size, data.length))
    //添加到首页
    state.memoList.unshift(...data as any as Memo[])
  } else if (type === 'edit') {
    //删除当前页,添加到当前页
    const currentPageStart = (page - 1) * state.size
    state.memoList.splice(currentPageStart, Math.min(state.size, state.memoList.length - currentPageStart), ...data as any as Memo[])
  } else {
    //添加到最后页
    state.memoList.push(...data as any as Memo[])
  }
  state.page = page
  state.hasNext = hasNext
}

</script>

<style scoped></style>