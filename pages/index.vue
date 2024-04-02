<template>
  <div>
    <HeaderImg />

    <MemoInput  v-if="token" @memo-added="refresh"/>

    <div class="flex flex-col divide-y divide-gray-100/50 p-4 gap-2">
      <FriendsMemo :memo="memo" v-for="(memo,index) in data?.data as Memo[]" :key="index" />
    </div>

  </div>
</template>

<script setup lang="ts">
import type { Memo } from '~/lib/types';
const token = useCookie('token')

const { data, pending, error, refresh } = await useFetch('/api/memo/list', {
  method:'POST',
  body: JSON.stringify({
    page: 1,    
  })
})
</script>

<style scoped></style>