<template>
  <Header v-if="memos.length > 0" v-bind:user="memos[0].user" />

  <div v-if="sysConfig.enableNewMemo" class="flex flex-col">
    <div v-for="(memo, index) in pinnedMemos" :key="index">
      <MemoUser v-bind:memo="memo" />
    </div>
    <div v-for="(memo, index) in nonPinnedMemoList" :key="index">
      <div v-if="memo.displayYear" class="pl-4 py-4">
        <span class="text-xl">{{ memo.displayYear }}</span>
        <span class="text-sm">年</span>
      </div>
      <MemoUser v-bind:memo="memo" />
    </div>
  </div>
  <div v-else class="flex flex-col divide-y divide-[#C0BEBF]/20">
    <Memo v-bind:memo="memo" v-for="memo in memos" :key="memo.id" />
  </div>
  <div
    v-if="hasNext"
    ref="loadMoreEle"
    class="text-xs text-center text-gray-500 py-2 cursor-pointer"
    @click="loadMore"
  >
    点击加载更多
  </div>
  <div class="text-xs text-center text-gray-500 py-2" @click="loadMore" v-else>
    已经到底啦
  </div>
</template>

<script setup lang="ts">
import type { MemoVO, SysConfigVO } from "~/types";
import Memo from "~/components/Memo.vue";
import MemoUser from "~/components/MemoUser.vue";
import { memoChangedEvent, memoReloadEvent } from "~/event";
import { useElementVisibility } from "@vueuse/core";
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
  state.page = 1;
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
  const currentPage = state.page;
  state.page = currentPage + 1;
  const res = await useMyFetch<{
    list: Array<MemoVO>;
    total: number;
    hasNext: boolean;
  }>("/memo/list", {
    ...state,
    userId: parseInt(userId),
  });
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

const pinnedMemos = computed(() => memos.value.filter((memo) => memo.pinned));
const nonPinnedMemos = computed(() =>
  memos.value.filter((memo) => !memo.pinned)
);

const nonPinnedMemoList = computed(() => {
  if (!nonPinnedMemos.value.length) return [];
  let lastYear: string | null = null;
  let lastDate: string | null = null;
  return nonPinnedMemos.value.map((memo) => {
    const currentYear = dayjs(memo.createdAt).locale("zh-cn").format("YYYY");
    const currentDate = dayjs(memo.createdAt).locale("zh-cn").format("YYYY-MM-DD"); // 新增：当前日期
    let returns = memo;

    if (currentYear !== lastYear) {
      lastYear = currentYear;
      returns = Object.assign({}, returns, { displayYear: currentYear });
    } else {
      returns = Object.assign({}, returns, { displayYear: null });
    }

    // 处理日期显示（仅当日期变化时标记显示）
    if (currentDate !== lastDate) {
      lastDate = currentDate;
      returns = Object.assign({}, returns, { displayDate: currentDate });
    } else {
      returns = Object.assign({}, returns, { displayDate: null });
    }

    return returns;
  });
});
</script>

<style scoped></style>