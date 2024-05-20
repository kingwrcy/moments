<template>
  <div>
    <div class="pt-4">
      <FriendsMemo :memo="data?.data as any as Memo" v-if="data?.data" :show-more="false" @memo-update="refresh" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Memo } from '~/lib/types';
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')

useHead({
  title: userinfo.value.title || '极简朋友圈',
})
const route = useRoute()
const id = route.params.id

const {data,refresh} = await useFetch('/api/memo/detail', {
  key:`memoDetail-${id}`,
  method: 'POST',
  body: JSON.stringify({ id: parseInt(id as string) })
})


</script>

<style scoped></style>