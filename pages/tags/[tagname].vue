<template>
  <div>
    <div class="content flex flex-col divide-y divide-[#C0BEBF]/10 gap-2">
      <div v-if="state.memoList.length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">关于该话题什么也没有,赶紧去登录发表Moments，内容里添加上 #{{ tagname }} 吧!</div>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in state.memoList" :key="index" :show-more="true" />

      <div class="cursor-pointer text-center text-sm opacity-70  my-4" @click="loadData(state.page + 1)"
        v-if="state.hasNext">- 加载更多 -
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import FriendsMemo from "~/components/FriendsMemo.vue";
import { type Memo, type PublicConfig } from '~/lib/types';
const token = useCookie('token')
const route = useRoute();
const tagname = route.params.tagname;
const publicConfig = useState<PublicConfig>('publicConfig')
const state = reactive({
  memoList: Array<Memo>(),
  page: 1,
  size: publicConfig.value.pageSize,
  hasNext: false
})

const { data, } = await useFetch('/api/memo/list', {
  method: "POST",
  body: {
    tagname: tagname,
    page: state.page,
  }
})

state.memoList = data.value?.data as any as Memo[]
state.hasNext = data.value?.hasNext || false

const loadData = async (page: number) => {
  page = page <= 0 ? 1 : page;
  const { data, hasNext } = await $fetch('/api/memo/list', {
    key: 'memoList',
    method: 'POST',
    body: JSON.stringify({
      page, size: state.size,tagname
    })
  })
  state.memoList.push(...data as any as Memo[])  
  state.page = page
  state.hasNext = hasNext
}



</script>
