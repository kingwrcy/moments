<template>
  <div >
    <div class="p-4">
      <FriendsMemo :memo="memo" v-if="memo" :show-more="false"/>
    </div>
    
    
  </div>
</template>

<script setup lang="ts">
import type { Memo } from '~/lib/types';

  const route = useRoute()
  const id = route.params.id

  const memo = ref<Memo>()
  const res = await $fetch('/api/memo/detail',{
    method: 'POST',
    body: JSON.stringify({id:parseInt(id as string)})
  })

  if(res.success){
    memo.value = res.data! as Memo
  }
</script>

<style scoped>

</style>