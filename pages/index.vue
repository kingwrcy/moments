<template>
  <div>

    <MemoInput v-if="token" @memo-added="refresh" />

    <div class="content flex flex-col divide-y divide-gray-100/50 gap-2">
      <div v-if="(data?.data as any as Memo[]).length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">什么也没有,赶紧去登录发表Moments吧!</div>
        <Button @click="navigateTo('/login')">去登录</Button>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in data?.data as any as Memo[]" :key="index" :show-more="true"
        @memo-update="refresh" />

    </div>
    <div class="cursor-pointer text-center text-sm opacity-70  my-4" @click="loadMore()" v-if="state.hasNext">点击加载更多...
    </div>
  </div>
</template>

<script setup lang="ts">
import { type User, type Memo } from '~/lib/types';
const token = useCookie('token')
const userinfo = useState<User>('userinfo')

useHead({
  title: userinfo.value.title || '极简朋友圈',
})


const state = reactive({
  memoList: Array<Memo>(),
  page: 1,
  hasNext: false
})


const { data, refresh } = await useFetch('/api/memo/list', {
  key: 'memoList',
  method: 'POST',
  body: JSON.stringify({
    page: state.page,
  })
})
state.memoList = data.value?.data as any as Memo[]
state.hasNext = data.value?.hasNext || false

const loadMore = async () => {
  const { data, hasNext } = await $fetch('/api/memo/list', {
    key: 'memoList',
    method: 'POST',
    body: JSON.stringify({
      page: ++state.page,
    })
  })
  state.memoList.push(...data as any as Memo[])
  state.hasNext = hasNext
}

</script>

<style scoped></style>