<template>
  <Header v-if="memo && memo.user" v-bind:user="memo.user"/>
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