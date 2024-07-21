<template>
  <div class="p-4 space-y-2">
    <div class="flex gap-2 text-lg *:cursor-pointer text-gray-600">
      <UIcon name="i-carbon-link"/>
      <UIcon name="i-carbon-image"/>
      <UIcon name="i-carbon-music"/>
      <UIcon name="i-carbon-video-player"/>
      <svg class="focus:outline-0 cursor-pointer w-[18px] h-[18px]" xmlns="http://www.w3.org/2000/svg"
           viewBox="0 0 24 24" fill="currentColor" data-state="closed" >
        <path
            d="M15.2735 15H5V7H19V15H17.3764L16.0767 19H21V21H3V19H7.6123L6.8 16.5L8.70211 15.882L9.71522 19H13.9738L15.2735 15ZM3.5 3H20.5V5H3.5V3ZM7 9V13H17V9H7Z"
        ></path>
      </svg>
    </div>
    <UTextarea v-model="state.content" :rows="8" autoresize padded/>
    <div class="flex justify-between items-center">
      <div class="flex flex-row gap-1 items-center text-[#576b95] text-sm cursor-pointer">
        <UIcon name="i-carbon-location"/>
        自定义位置
      </div>

      <UButtonGroup>
        <UButton @click="saveMemo">发表</UButton>
        <UButton color="white">清空</UButton>
      </UButtonGroup>
    </div>
  </div>
</template>

<script setup lang="ts">

import type {MemoVO, UserVO} from "~/types";
import {toast} from "vue-sonner";

const props = defineProps<{id?:number}>()
const state = reactive({
  id:props.id || 0,
  content:"",
  ext:"",
  pinned:0,
  showType:1,
  externalFavicon:"",
  externalTitle:"",
  externalUrl:"",
  music163Url:"",
  bilibiliUrl:"",
  imgs:""
})
const currentUser = useState<UserVO>('userinfo')
onMounted(async () => {
  if(state.id >0){
    const res = await useMyFetch<MemoVO>('/memo/get?id='+state.id)
    Object.assign(state,res)

  }
})

const saveMemo = async()=>{
  const res = await useMyFetch('/memo/save',{
    id:state.id,
    content:state.content,
    ext: {

    },
    pinned:state.pinned,
    showType:state.showType,
    externalFavicon:state.externalFavicon,
    externalTitle:state.externalTitle,
    externalUrl:state.externalUrl,
    music163Url:state.music163Url,
    bilibiliUrl:state.bilibiliUrl,
    imgs:state.imgs.split(','),
  })
  toast.success("保存成功!")
  await navigateTo('/')
}

</script>

<style scoped>

</style>