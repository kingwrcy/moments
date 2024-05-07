<template>
  <div>

    <MemoInput v-if="token" @memo-added="firstLoad"/>

    <div class="content flex flex-col divide-y divide-[#C0BEBF]/10 gap-2">
      <div v-if="state.memoList.length === 0 && !token" class="text-center">
        <div class="my-2 text-sm">关于该话题什么也没有,赶紧去登录发表Moments，内容里添加上 #{{ tagname }} 吧!</div>
      </div>
      <FriendsMemo :memo="memo" v-for="(memo, index) in state.memoList" :key="index" :show-more="true"
                   @memo-update="firstLoad" />
    </div>

    <div id="get-more" ref="getMore" class="cursor-pointer text-center text-sm opacity-70 my-4" @click="loadMore()" v-if="state.hasNext" >
      加载中...
    </div>
    <div class="cursor-pointer text-center text-sm opacity-70 my-4">
      ———— 没有更多啦～ ————
    </div>
  </div>
</template>

<script setup lang="ts">
import { type User, type Memo } from '~/lib/types';
import { onMounted, onUnmounted, watch, ref } from 'vue';
import jsonp from 'jsonp';
import {toast} from "vue-sonner";
import HeaderImg from "~/components/HeaderImg.vue";
import MemoInput from "~/components/MemoInput.vue";
import FriendsMemo from "~/components/FriendsMemo.vue";

const getMore = ref(null);
const token = useCookie('token')
const route = useRoute();
const tagname = route.params.tagname;

let observer: IntersectionObserver | null = null;


onMounted(async () => {
  await firstLoad();
  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      loadMore();
    }
  }, {
    // 上边框距离屏幕底部一定距离时触发
    rootMargin: '500px',
  });

  if (getMore.value) {
    observer.observe(getMore.value);
  }

  // 当组件卸载时，停止观察
  onUnmounted(() => {
    if (getMore.value) {
      observer.unobserve(getMore.value);
    }
  });
  // 监听 getMore 引用的变化，并重新设置观察者
  watch(getMore, () => {
    setupObserver();
  }, {
    immediate: true // 立即触发，确保初始 setup
  });
});

const setupObserver = () => {
  // 清除现有的观察者（如果有）
  if (observer && getMore.value) {
    observer.unobserve(getMore.value);
  }

  // 创建新的观察者实例
  observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      loadMore();
    }
  }, {
    // 上边框距离屏幕底部一定距离时触发
    rootMargin: '500px',
  });

  // 设置新的观察目标
  if (getMore.value) {
    observer.observe(getMore.value);
  }
};


const state = reactive({
  memoList: Array<Memo>(),
  page: 1,
  hasNext: false
})

const firstLoad = async () => {
  state.page = 1
  toast.promise($fetch('/api/memo/list', {
        key: 'memoList',
        method: 'POST',
        body: JSON.stringify({
          tagname: tagname,
          page: state.page,
        })
      }), {
        loading: '加载中...',
        success: (data) => {
          if (data.success) {
            state.memoList = data.data as any as Memo[]
            state.hasNext = data.hasNext || false
            return '加载成功';
          } else {
            return '加载失败: ' + data.message;
          }
        },
        error: (error) => {
          if (error.response && error.response.status === 429) {
            return '请求过于频繁，请稍后再试';
          } else {
            return `加载失败: ${error.message || '未知错误'}`;
          }
        },
        finally() {
          loadLock = false; // 确保加载锁被重置
        },
      }
  );
}

let loadLock = false;

const loadMore = async () => {
  if(loadLock) return;
  loadLock = true;

  toast.promise(
      $fetch('/api/memo/list', {
        key: 'memoList',
        method: 'POST',
        body: JSON.stringify({
          tagname: tagname,
          page: state.page + 1
        })
      }), {
        loading: '加载中...',
        success: (data) => {
          if (data.success) {
            state.page += 1; // 成功后增加页码
            if (Array.isArray(data.data)) { // 确保数据是数组
              state.memoList.push(...data.data);
            }
            state.hasNext = data.hasNext;
            return '加载成功';
          } else {
            return '加载失败: ' + data.message;
          }
        },
        error: (error) => {
          if (error.response && error.response.status === 429) {
            return '请求过于频繁，请稍后再试';
          } else {
            return `加载失败: ${error.message || '未知错误'}`;
          }
        },
        finally() {
          loadLock = false; // 确保加载锁被重置
        },
      }
  );
}

</script>

<style scoped></style>