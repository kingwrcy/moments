<template>
  <UModal v-model="open" :ui="{container: 'sm:items-end'}">
    <div @click="navigate('/new')" v-if="global.userinfo.token " title="发表" class="flex flex-col items-center p-4 pt-8 text-gray-500 dark:text-white">
      <svg class="text-[#9fc84a] w-7 h-7 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M29 26H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h6.46l1.71-2.55A1 1 0 0 1 12 4h8a1 1 0 0 1 .83.45L22.54 7H29a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1M4 24h24V9h-6a1 1 0 0 1-.83-.45L19.46 6h-6.92l-1.71 2.55A1 1 0 0 1 10 9H4Z"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a6 6 0 0 1-6 6m0-10a4 4 0 1 0 4 4a4 4 0 0 0-4-4"/></svg>
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
        <svg class="text-[#9fc84a] w-6 h-6 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="-2 -2 24 24"><path fill="currentColor" d="M17 16a1 1 0 0 1 0-2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H9.415l-.471-1.334A1 1 0 0 0 8 2H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1a1 1 0 0 1 0 2a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h5c1.306 0 2.417.835 2.83 2H17a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3m-5.01-.096a5.002 5.002 0 0 1-6.293-7.707a5 5 0 0 1 7.707 6.293l2.9 2.9a1 1 0 0 1-1.415 1.413l-2.9-2.899zm-.636-2.05A3 3 0 1 0 7.11 9.61a3 3 0 0 0 4.243 4.243z"/></svg>
        <span>检索</span>
      </div>
      <div @click="navigate('/sys/settings')" v-if="$route.path !== '/sys/settings' && global.userinfo.id === 1" title="系统设置" class="flex flex-col items-center">
        <svg class="text-[#9fc84a] w-6 h-6 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M27 16.76v-1.53l1.92-1.68A2 2 0 0 0 29.3 11l-2.36-4a2 2 0 0 0-1.73-1a2 2 0 0 0-.64.1l-2.43.82a11 11 0 0 0-1.31-.75l-.51-2.52a2 2 0 0 0-2-1.61h-4.68a2 2 0 0 0-2 1.61l-.51 2.52a11.5 11.5 0 0 0-1.32.75l-2.38-.86A2 2 0 0 0 6.79 6a2 2 0 0 0-1.73 1L2.7 11a2 2 0 0 0 .41 2.51L5 15.24v1.53l-1.89 1.68A2 2 0 0 0 2.7 21l2.36 4a2 2 0 0 0 1.73 1a2 2 0 0 0 .64-.1l2.43-.82a11 11 0 0 0 1.31.75l.51 2.52a2 2 0 0 0 2 1.61h4.72a2 2 0 0 0 2-1.61l.51-2.52a11.5 11.5 0 0 0 1.32-.75l2.42.82a2 2 0 0 0 .64.1a2 2 0 0 0 1.73-1l2.28-4a2 2 0 0 0-.41-2.51ZM25.21 24l-3.43-1.16a8.9 8.9 0 0 1-2.71 1.57L18.36 28h-4.72l-.71-3.55a9.4 9.4 0 0 1-2.7-1.57L6.79 24l-2.36-4l2.72-2.4a8.9 8.9 0 0 1 0-3.13L4.43 12l2.36-4l3.43 1.16a8.9 8.9 0 0 1 2.71-1.57L13.64 4h4.72l.71 3.55a9.4 9.4 0 0 1 2.7 1.57L25.21 8l2.36 4l-2.72 2.4a8.9 8.9 0 0 1 0 3.13L27.57 20Z"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a5.94 5.94 0 0 1-6 6m0-10a3.91 3.91 0 0 0-4 4a3.91 3.91 0 0 0 4 4a3.91 3.91 0 0 0 4-4a3.91 3.91 0 0 0-4-4"/></svg>
        <span>系统</span>
      </div>
      <div @click="navigate('/user/settings')" v-if="$route.path !== '/user/settings' && global.userinfo.token" title="用户中心" class="flex flex-col items-center">
        <svg class="text-[#9fc84a] w-6 h-6 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M16 8a5 5 0 1 0 5 5a5 5 0 0 0-5-5m0 8a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3"/><path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14.016 14.016 0 0 0 16 2m-6 24.377V25a3.003 3.003 0 0 1 3-3h6a3.003 3.003 0 0 1 3 3v1.377a11.9 11.9 0 0 1-12 0m13.993-1.451A5 5 0 0 0 19 20h-6a5 5 0 0 0-4.992 4.926a12 12 0 1 1 15.985 0"/></svg>
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