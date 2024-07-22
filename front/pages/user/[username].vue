<template>
  <Header v-if="memos.length>0" v-bind:user="memos[0].user"/>
  <div class="flex flex-col divide-y divide-[#C0BEBF]/20 ">
    <Memo v-bind:memo="m" v-for="m in memos" :key="m.id"/>
  </div>
</template>

<script setup lang="ts">
import type {MemoVO} from "~/types";
import Memo from "~/components/Memo.vue";
import {memoChangedEvent, memoReloadEvent} from "~/event";

const route = useRoute()
const username = route.params.username as any as string
const state = reactive({
  page: 1,
  size: 10,
})

const memos = ref<Array<MemoVO>>([])
onMounted(async () => {
  await reload()
})

const reload = async () => {
  const res = await useMyFetch<{
    list: Array<MemoVO>,
    total: number
  }>('/memo/list', {
    ...state, username,
  })
  memos.value = res.list
}

memoReloadEvent.on(async () => {
  await reload()
})

memoChangedEvent.on(async (id: number) => {
  const res = await useMyFetch<MemoVO>('/memo/get?latest=1&id=' + id)
  const index = memos.value.findIndex(r => r.id === id)
  if (index >= 0) {
    memos.value[index] = res
  }
})
</script>

<style scoped>

</style>