<template>
  <Header v-if="memo && memo.user" v-bind:user="memo.user"/>

  <div class="flex justify-end gap-2 sm:hidden px-4 my-2">
    <UButton @click="navigateTo('/')" icon="i-carbon-arrow-left" size="xs" color="gray" variant="solid">返回</UButton>
  </div>

  <Memo v-if="memo" v-bind:memo="memo"/>

</template>

<script setup lang="ts">
import type {MemoVO} from "~/types";
import {memoChangedEvent} from "~/event";

const route = useRoute()
const id = route.params.id as any as number
const memo = ref<MemoVO>()
const reload = async () => {
  memo.value = await useMyFetch<MemoVO>('/memo/get?id=' + id)
}

memoChangedEvent.on(async () => {
  await reload()
})
onMounted(async () => {
  await reload()
})


</script>

<style scoped>

</style>