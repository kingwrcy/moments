<template>
  <div>
    <div class="p-2 sm:p-4">
      <FriendsMemo :memo="memo" v-if="memo" :show-more="false" />
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

const memo = ref<Memo>()
const {data} = await useFetch('/api/memo/detail', {
  key:`memoDetail-${id}`,
  method: 'POST',
  body: JSON.stringify({ id: parseInt(id as string) })
})

if (data.value?.success) {
  memo.value = data.value?.data! as any as Memo
}
</script>

<style scoped></style>