<template>
  <USlideover v-model="open" side="left" class="w-1/3">
    <div class="flex flex-col gap-4 p-4 text-sm text-gray-500 dark:text-white">

      <div class="flex items-center gap-2">
        <svg @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="20" height="20" v-if="mode==='light'"
             viewBox="0 0 24 24" fill="none"
             stroke="#FDE047"
             stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
             class="lucide lucide-moon-star-icon cursor-pointer">
          <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
          <path d="M20 3v4"></path>
          <path d="M22 5h-4"></path>
        </svg>

        <svg @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" v-else
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
      </div>
      <div @click="navigate('/')" v-if="$route.path !== '/'" title="去首页" class="flex items-center gap-1">
        <UIcon name="i-carbon-arrow-left" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>去首页</span>
      </div>
      <div @click="navigate('/new')" v-if="global.userinfo.token " title="发言" class="flex gap-1">
        <UIcon name="i-carbon-edit" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>发言</span>
      </div>
      <div @click="navigate('/user/calendar')" v-if="global.userinfo.token" title="日历" class="flex gap-1">
        <UIcon name="i-carbon-calendar" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>日历</span>
      </div>

      <div @click="navigate('/sys/settings')" v-if="global.userinfo.id === 1" title="系统设置" class="flex gap-1">
        <UIcon name="i-carbon-settings" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>系统设置</span>
      </div>
      <div @click="navigate('/user/settings')" v-if="global.userinfo.token" title="用户设置" class="flex gap-1">
        <UIcon name="i-carbon-user-profile" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>用户设置</span>
      </div>
      <div @click="navigate('/user/login')" v-if="!global.userinfo.token" title="登录" class="flex gap-1">
        <UIcon name="i-carbon-login" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
        <span>登录</span>
      </div>
    </div>
  </USlideover>
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