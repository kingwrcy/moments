<template>
  <Header v-if="memos.length > 0" v-bind:user="memos[0].user" />

  <div class="flex flex-col">
    <div v-for="(memo, index) in pinnedMemos" :key="index">
      <userMemo v-bind:memo="memo" />
    </div>
    <div v-for="(memo, index) in nonPinnedMemoList" :key="index">
      <div v-if="memo.displayYear" class="pl-4 py-4">
        <span class="text-xl">{{ memo.displayYear }}</span>
        <span class="text-sm">年</span>
      </div>
      <userMemo v-bind:memo="memo" />
    </div>
  </div>
  <div
    ref="loadMoreEle"
    class="text-xs text-center text-gray-500 py-2 cursor-pointer"
    @click="loadMore"
    v-if="hasNext"
  >
    点击加载更多
  </div>
  <div class="text-xs text-center text-gray-500 py-2" @click="loadMore" v-else>
    已经到底啦
  </div>
</template>

<script setup lang="ts">
import type { MemoVO, SysConfigVO } from "~/types";
import userMemo from "~/components/userMemo.vue";
import { memoChangedEvent, memoReloadEvent } from "~/event";
import { useElementVisibility } from "@vueuse/core";
import { computed, onMounted, reactive, ref, watch } from "vue";
import dayjs from "dayjs";

const loadMoreEle = ref(null);
const targetIsVisible = useElementVisibility(loadMoreEle);
const sysConfig = useState<SysConfigVO>("sysConfig");

watch(targetIsVisible, async (visible) => {
  if (visible && sysConfig.value?.enableAutoLoadNextPage) {
    await loadMore();
  }
});

const hasNext = ref(false);
const route = useRoute();
const userId = route.params.id as any as string;
const state = reactive({
  page: 1,
  size: 10,
});

const memos = ref<Array<MemoVO>>([]);

onMounted(async () => {
  await reload();
});

const reload = async () => {
  const res = await useMyFetch<{
    list: Array<MemoVO>;
    total: number;
    hasNext: boolean;
  }>("/memo/list", {
    ...state,
    userId: parseInt(userId),
  });
  memos.value = res.list;
  hasNext.value = res.hasNext;
};

const loadMore = async () => {
  state.page = state.page + 1;
  const res = await useMyFetch<{
    list: Array<MemoVO>;
    total: number;
    hasNext: boolean;
  }>("/memo/list", state);
  memos.value = [...memos.value, ...res.list];
  hasNext.value = res.hasNext;
};

memoReloadEvent.on(async () => {
  await reload();
});

memoChangedEvent.on(async (id: number) => {
  const res = await useMyFetch<MemoVO>("/memo/get?latest=1&id=" + id);
  const index = memos.value.findIndex((r) => r.id === id);
  if (index >= 0) {
    memos.value[index] = res;
  }
});

// 分离置顶和非置顶memo
const pinnedMemos = computed(() => memos.value.filter((memo) => memo.pinned));
const nonPinnedMemos = computed(() =>
  memos.value.filter((memo) => !memo.pinned)
);

const nonPinnedMemoList = computed(() => {
  if (!nonPinnedMemos.value.length) return [];
  let lastYear: string | null = null;
  return nonPinnedMemos.value.map((memo) => {
    const currentYear = dayjs(memo.createdAt).locale("zh-cn").format("YYYY");
    let returns = memo;
    if (currentYear !== lastYear) {
      lastYear = currentYear;
      returns = { ...returns, displayYear: currentYear };
    } else {
      returns = { ...returns, displayYear: null };
    }
    return returns;
  });
});
</script>

<style scoped></style>
