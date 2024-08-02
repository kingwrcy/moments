<template>
  <div class="header relative mb-14" v-if="$route.path!=='/new' && $route.path.indexOf('/edit/') < 0">

    <div
        class="dark:bg-neutral-800 hidden sm:flex sm:absolute sm:-right-10 sm:rounded sm:p-2 sm:flex-col sm:w-fit justify-end shadow w-full flex-row  top-0  p-1 flex  gap-2 bg-white ">
      <NuxtLink to="/" v-if="$route.path !== '/'" title="去首页">
        <UIcon name="i-carbon-arrow-left" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
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


      <NuxtLink to="/new" v-if="global.userinfo.token " title="发言">
        <UIcon name="i-carbon-edit" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/user/calendar" v-if="global.userinfo.token" title="日历">
        <UIcon name="i-carbon-calendar" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>

      <NuxtLink to="/sys/settings" v-if="global.userinfo.id === 1" title="系统设置">
        <UIcon name="i-carbon-settings" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/user/settings" v-if="global.userinfo.token" title="用户设置">
        <UIcon name="i-carbon-user-profile" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/user/login" v-if="!global.userinfo.token" title="登录">
        <UIcon name="i-carbon-login" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
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