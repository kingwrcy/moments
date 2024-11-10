<template>
  <UModal v-model="open" :ui="{container: 'sm:items-end'}">
    <div @click="navigate('/new')" v-if="global.userinfo.token " title="发表" class="flex flex-col items-center p-4 pt-8 text-gray-500 dark:text-white">
      <UIcon name="i-mage-edit" class="text-[#9fc84a] w-7 h-7 cursor-pointer"/>
      <span>发表</span>
    </div>
    <div class="flex items-center justify-between gap-4 p-8 pt-2 text-gray-500 dark:text-white">
      <!-- <div @click="navigate('/')" v-if="$route.path !== '/'" title="主页" class="flex flex-col items-center">
        <UIcon name="i-carbon-home" class="text-[#9fc84a] w-6 h-6 cursor-pointer"/>
        <span>主页</span>
      </div> -->
      <div class="flex flex-col items-center gap-1">
        <svg @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="22" height="22" v-if="mode==='light'"
             viewBox="0 0 24 24" fill="none"
             stroke="#FDE047"
             stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
             class="lucide lucide-moon-star-icon cursor-pointer">
          <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
          <path d="M20 3v4"></path>
          <path d="M22 5h-4"></path>
        </svg>

        <svg @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" v-else
             fill="none"
             stroke="#FDE047" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
             class="lucide lucide-sun-icon cursor-pointer">
          <circle cx="12" cy="12" r="4"></circle>
          <path d="M12 2v2"></path>
          <path d="M12 20v2"></path>
          <path d="m4.93 4.93 1.41 1.41"></path>
          <path d="m17.66 17.66 1.41 1.41"></path>
          <path d="M2 12h2"></path>
          <path d="M20 12h2"></path>
          <path d="m6.34 17.66-1.41 1.41"></path>
          <path d="m19.07 4.93-1.41 1.41"></path>
        </svg>
        <span @click="toggleMode">{{mode==='light' ? '暗色' : '亮色'}}</span>
      </div>
      <div @click="navigate('/user/calendar')" v-if="$route.path !== '/user/calendar' && global.userinfo.token" title="日历检索" class="flex flex-col items-center">
        <UIcon name="i-jam-search-folder" class="text-[#9fc84a] w-6 h-6 cursor-pointer"/>
        <span>检索</span>
      </div>
      <div @click="navigate('/sys/settings')" v-if="$route.path !== '/sys/settings' && global.userinfo.id === 1" title="系统设置" class="flex flex-col items-center">
        <UIcon name="i-carbon-settings" class="text-[#9fc84a] w-6 h-6 cursor-pointer"/>
        <span>系统</span>
      </div>
      <div @click="navigate('/user/settings')" v-if="$route.path !== '/user/settings' && global.userinfo.token" title="用户中心" class="flex flex-col items-center">
        <UIcon name="i-carbon-user-avatar" class="text-[#9fc84a] w-6 h-6 cursor-pointer"/>
        <span>用户</span>
      </div>
    </div>
  </UModal>
</template>

<script setup lang="ts">

import {useGlobalState} from "~/store";
import {useColorMode} from '@vueuse/core'

const global = useGlobalState()
const mode = useColorMode()
const open = useState<boolean>('sidebarOpen',()=>false)
const toggleMode = () => {
  if (mode.value === 'dark') {
    mode.value = 'light'
  } else {
    mode.value = 'dark'
  }
}

const navigate = async (url :string)=>{
  open.value = false
  await navigateTo(url)
}
</script>

<style scoped>

</style>