<template>
  <Header v-if="memos.length>0" v-bind:user="memos[0].user"/>

  <div class="flex justify-end gap-2 sm:hidden px-4 my-2">
    <UButton @click="navigateTo('/')" icon="i-carbon-arrow-left" size="xs" color="gray" variant="solid">返回</UButton>
  </div>

  <div class="flex flex-col divide-y divide-[#C0BEBF]/20 ">
    <Memo v-bind:memo="m" v-for="m in memos" :key="m.id"/>
  </div>
  <div ref="loadMoreEle" class="text-xs text-center text-gray-500 py-2" @click="loadMore" v-if="hasNext">
    点击加载更多
  </div>
  <div class="text-xs text-center text-gray-500 py-2" @click="loadMore" v-else>
    已经到底啦
  </div>
</template>

<script setup lang="ts">
import type {MemoVO} from "~/types";
import Memo from "~/components/Memo.vue";
import {memoChangedEvent, memoReloadEvent} from "~/event";
import {useElementVisibility} from "@vueuse/core";
const loadMoreEle = ref(null)
const targetIsVisible = useElementVisibility(loadMoreEle)
watch(targetIsVisible, async (visible) => {
  if (visible) {
    await loadMore()
  }
})
const hasNext = ref(false)
const route = useRoute()
const userId = route.params.id as any as string
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
    total: number,
    hasNext: boolean
  }>('/memo/list', {
    ...state, userId:parseInt(userId),
  })
  memos.value = res.list
  hasNext.value = res.hasNext
}

const loadMore = async () => {
  state.page = state.page + 1
  const res = await useMyFetch<{
    list: Array<MemoVO>,
    total: number,
    hasNext: boolean
  }>('/memo/list', state)
  memos.value = [...memos.value, ...res.list]
  hasNext.value = res.hasNext
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