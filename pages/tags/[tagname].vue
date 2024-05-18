<template>
  <div>



    <div class="content flex flex-col divide-y divide-[#C0BEBF]/10 gap-2">
      <div v-if="state.memoList.length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">关于该话题什么也没有,赶紧去登录发表Moments，内容里添加上 #{{ tagname }} 吧!</div>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in state.memoList" :key="index" :show-more="true" />
    </div>


  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import FriendsMemo from "~/components/FriendsMemo.vue";
import { type Memo } from '~/lib/types';
const token = useCookie('token')
const route = useRoute();
const tagname = route.params.tagname;
const state = reactive({
  memoList: Array<Memo>(),
  page: 1,
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

</script>

<style scoped></style>