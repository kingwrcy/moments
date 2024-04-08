<template>
  <div>

    <MemoInput v-if="token" @memo-added="refresh" />

    <div class="flex flex-col divide-y divide-gray-100/50 p-4 gap-2">
      <div v-if="(data?.data as Memo[]).length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">什么也没有,赶紧去登录发表Moments吧!</div>
        <Button @click="navigateTo('/login')">去登录</Button>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in data?.data as Memo[]" :key="index" :show-more="true"
        @memo-update="refresh" />
    </div>

  </div>
</template>

<script setup lang="ts">
import type { Memo } from '~/lib/types';
const token = useCookie('token')

useHead({
  title: '极简朋友圈',
})


const { data, pending, error, refresh } = await useFetch('/api/memo/list', {
  method: 'POST',
  body: JSON.stringify({
    page: 1,
  })
})
</script>

<style scoped></style>