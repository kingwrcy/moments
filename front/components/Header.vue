<template>
  <div class="header relative mb-14" v-if="$route.path!=='/new' && $route.path.indexOf('/edit/') < 0">
    <div v-if="$route.path!=='/' && $route.path.indexOf('/memo/') < 0" :class="{ 'bg-[#4c4c4c]/80 z-10': y > 100 }" class="flex fixed justify-between items-center p-4 w-full md:w-[567px] text-white top-0">
      <NuxtLink class="flex items-center" title="返回主页">
        <svg @click="navigateTo('/')" class="w-5 h-5 cursor-pointer mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M10 16L20 6l1.4 1.4l-8.6 8.6l8.6 8.6L20 26z"/></svg>
        <span v-if="$route.path==='/user/calendar'">日历检索</span>
        <span v-else-if="$route.path==='/sys/settings'">系统设置</span>
        <span v-else-if="$route.path==='/user/settings'">用户中心</span>
        <span v-else-if="$route.path.indexOf('/tags/') >= 0">话题专栏</span>
        <span v-else>
          <span v-if="!global.userinfo.token && $route.path==='/user/login'">登录</span>
          <span v-else-if="!global.userinfo.token && $route.path==='/user/reg'">注册</span>
          <span v-else>{{ props.user.nickname }} 的空间</span>
        </span>
      </NuxtLink>
      <NuxtLink @click="logout" v-if="$route.path === '/user/settings' && global.userinfo.token" title="登出">
        <svg class="w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M6 30h12a2 2 0 0 0 2-2v-3h-2v3H6V4h12v3h2V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2"/><path fill="currentColor" d="M20.586 20.586L24.172 17H10v-2h14.172l-3.586-3.586L22 10l6 6l-6 6z"/></svg>
      </NuxtLink>
    </div>

    <div
        class="dark:bg-neutral-800 hidden sm:flex sm:absolute sm:-right-10 sm:rounded sm:p-2 sm:flex-col sm:w-fit justify-end shadow w-full flex-row top-0  p-1 flex  gap-2 bg-white ">
      <svg v-if="mode==='light'" @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="20" height="20"
           viewBox="0 0 24 24" fill="none"
           stroke="#FDE047"
           stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
           class="lucide lucide-moon-star-icon cursor-pointer">
        <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
        <path d="M20 3v4"></path>
        <path d="M22 5h-4"></path>
      </svg>

      <svg v-else @click="toggleMode" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
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


      <NuxtLink to="/new" v-if="global.userinfo.token " title="发表">
        <svg class="text-[#9fc84a] w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M29 26H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h6.46l1.71-2.55A1 1 0 0 1 12 4h8a1 1 0 0 1 .83.45L22.54 7H29a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1M4 24h24V9h-6a1 1 0 0 1-.83-.45L19.46 6h-6.92l-1.71 2.55A1 1 0 0 1 10 9H4Z"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a6 6 0 0 1-6 6m0-10a4 4 0 1 0 4 4a4 4 0 0 0-4-4"/></svg>
      </NuxtLink>
      <NuxtLink to="/user/calendar" v-if="$route.path !== '/user/calendar' && global.userinfo.token" title="日历检索">
        <svg class="text-[#9fc84a] w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="-2 -2 24 24"><path fill="currentColor" d="M17 16a1 1 0 0 1 0-2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H9.415l-.471-1.334A1 1 0 0 0 8 2H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1a1 1 0 0 1 0 2a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h5c1.306 0 2.417.835 2.83 2H17a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3m-5.01-.096a5.002 5.002 0 0 1-6.293-7.707a5 5 0 0 1 7.707 6.293l2.9 2.9a1 1 0 0 1-1.415 1.413l-2.9-2.899zm-.636-2.05A3 3 0 1 0 7.11 9.61a3 3 0 0 0 4.243 4.243z"/></svg>
      </NuxtLink>

      <NuxtLink to="/sys/settings" v-if="$route.path !== '/sys/settings' && global.userinfo.id === 1" title="系统设置">
        <svg class="text-[#9fc84a] w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M27 16.76v-1.53l1.92-1.68A2 2 0 0 0 29.3 11l-2.36-4a2 2 0 0 0-1.73-1a2 2 0 0 0-.64.1l-2.43.82a11 11 0 0 0-1.31-.75l-.51-2.52a2 2 0 0 0-2-1.61h-4.68a2 2 0 0 0-2 1.61l-.51 2.52a11.5 11.5 0 0 0-1.32.75l-2.38-.86A2 2 0 0 0 6.79 6a2 2 0 0 0-1.73 1L2.7 11a2 2 0 0 0 .41 2.51L5 15.24v1.53l-1.89 1.68A2 2 0 0 0 2.7 21l2.36 4a2 2 0 0 0 1.73 1a2 2 0 0 0 .64-.1l2.43-.82a11 11 0 0 0 1.31.75l.51 2.52a2 2 0 0 0 2 1.61h4.72a2 2 0 0 0 2-1.61l.51-2.52a11.5 11.5 0 0 0 1.32-.75l2.42.82a2 2 0 0 0 .64.1a2 2 0 0 0 1.73-1l2.28-4a2 2 0 0 0-.41-2.51ZM25.21 24l-3.43-1.16a8.9 8.9 0 0 1-2.71 1.57L18.36 28h-4.72l-.71-3.55a9.4 9.4 0 0 1-2.7-1.57L6.79 24l-2.36-4l2.72-2.4a8.9 8.9 0 0 1 0-3.13L4.43 12l2.36-4l3.43 1.16a8.9 8.9 0 0 1 2.71-1.57L13.64 4h4.72l.71 3.55a9.4 9.4 0 0 1 2.7 1.57L25.21 8l2.36 4l-2.72 2.4a8.9 8.9 0 0 1 0 3.13L27.57 20Z"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a5.94 5.94 0 0 1-6 6m0-10a3.91 3.91 0 0 0-4 4a3.91 3.91 0 0 0 4 4a3.91 3.91 0 0 0 4-4a3.91 3.91 0 0 0-4-4"/></svg>
      </NuxtLink>
      <NuxtLink to="/user/settings" v-if="$route.path !== '/user/settings' && global.userinfo.token" title="用户中心">
        <svg class="text-[#9fc84a] w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M16 8a5 5 0 1 0 5 5a5 5 0 0 0-5-5m0 8a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3"/><path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14.016 14.016 0 0 0 16 2m-6 24.377V25a3.003 3.003 0 0 1 3-3h6a3.003 3.003 0 0 1 3 3v1.377a11.9 11.9 0 0 1-12 0m13.993-1.451A5 5 0 0 0 19 20h-6a5 5 0 0 0-4.992 4.926a12 12 0 1 1 15.985 0"/></svg>
      </NuxtLink>
      <NuxtLink to="/user/login" v-if="!global.userinfo.token" title="登录">
        <svg class="text-[#9fc84a] w-5 h-5 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M26 30H14a2 2 0 0 1-2-2v-3h2v3h12V4H14v3h-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2"/><path fill="currentColor" d="M14.59 20.59L18.17 17H4v-2h14.17l-3.58-3.59L16 10l6 6l-6 6z"/></svg>
      </NuxtLink>
    </div>

    <img class="header-img w-full" :src="props.user.coverUrl"
         alt="">
    <div class="absolute right-2 bottom-[-40px]">
      <div class="userinfo flex flex-col">
        <div class="flex flex-row items-center gap-4 justify-end">
          <div class="username text-lg font-bold text-white">{{ props.user.nickname }}</div>
          <img :src="props.user.avatarUrl"
               class="avatar w-[70px] h-[70px] rounded-xl"></div>
        <div class="slogon text-gray truncate w-full text-end text-xs mt-2">{{ props.user.slogan }}</div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type {UserVO} from "~/types";
import {useGlobalState} from "~/store";
import {useColorMode} from '@vueuse/core'

const global = useGlobalState()
const props = defineProps<{ user: UserVO }>()
const mode = useColorMode()
const {y} = useWindowScroll()

const logout = async () => {
  global.value.userinfo = {}
  await navigateTo('/')
}

const toggleMode = () => {
  if (mode.value === 'dark') {
    mode.value = 'light'
  } else {
    mode.value = 'dark'
  }
}
</script>

<style scoped>

</style>