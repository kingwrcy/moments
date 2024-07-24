<template>
  <div class="header relative mb-14">
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
    <div class="absolute shadow -right-10 top-0 rounded p-1 flex flex-col gap-2 bg-white rounded p-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#FDE047"
           stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
           class="lucide lucide-moon-star-icon cursor-pointer">
        <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
        <path d="M20 3v4"></path>
        <path d="M22 5h-4"></path>
      </svg>
      <NuxtLink to="/user/login" v-if="!global.userinfo.token" title="登录">
        <UIcon name="i-carbon-login" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <UIcon @click="logout" name="i-carbon-logout" v-else class="text-[#9fc84a] w-5 h-5 cursor-pointer" title="登出"/>
      <NuxtLink to="/new" v-if="global.userinfo.token " title="发言">
        <UIcon name="i-carbon-edit" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/user/calendar" v-if="global.userinfo.token" title="日历">
        <UIcon name="i-carbon-calendar" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/" v-if="$route.path !== '/'" title="去首页">
        <UIcon name="i-carbon-arrow-left" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/sys/settings" v-if="global.userinfo.id === 1" title="系统设置">
        <UIcon name="i-carbon-settings" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
      <NuxtLink to="/user/settings" v-if="global.userinfo.token" title="用户设置">
        <UIcon name="i-carbon-user-profile" class="text-[#9fc84a] w-5 h-5 cursor-pointer"/>
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {useGlobalState} from "~/store";

const global = useGlobalState()
const props = defineProps<{ user: UserVO }>()

const logout = async () => {
  global.value.userinfo = {}
  await navigateTo('/')
}
</script>

<style scoped>

</style>