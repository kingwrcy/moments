<template>
  <div class="header relative mb-8 ">
    <img class="header-img w-full  " :src="getImgUrl(res?.data?.coverUrl!)" alt="" />
	  
    <div class="absolute left-2 top-2 rounded p-1 flex flex-row gap-2">
      <div title="返回" v-if="showBack()" @click="navigateTo('/')">
        <ArrowLeft color="#9FC84A" :size="20" class="cursor-pointer" />
      </div>
      <div title="登录" v-if="!token">
        <User :size="20" class="cursor-pointer" color="#9FC84A" @click="navigateTo('/login')" />
      </div>
    </div>

    <div class="absolute right-2 bottom-[-40px]">
      <div class="userinfo flex flex-col">
        <div class="flex flex-row items-center gap-4 justify-end">
          <div class="username text-lg font-bold text-white">{{ res?.data?.nickname }}</div>
          <img :src="getImgUrl(res?.data?.avatarUrl!)" class="avatar w-[70px] h-[70px] rounded-xl" />
        </div>
        <div class="slogon text-gray truncate w-full text-end text-xs mt-2">{{ res?.data?.slogan }}</div>
      </div>
    </div>

    <div class="absolute right-2 top-2 rounded p-1 flex flex-row gap-2">
      <div title="亮色" v-if="colorMode === 'dark'" @click="colorMode = 'light'">
        <Sun color="#FDE047" :size="20" class="cursor-pointer" />
      </div>
      <div title="暗色" v-else>
        <MoonStar color="#FDE047" :size="20" class="cursor-pointer" @click="colorMode = 'dark'" />
      </div>
    </div>

  </div>

</template>

<script setup lang="ts">
import { settingsUpdateEvent } from '~/lib/event'
import { getImgUrl } from '~/lib/utils';
import { Sun, MoonStar, User, ArrowLeft } from 'lucide-vue-next'
import { useColorMode } from '@vueuse/core';

const colorMode = useColorMode()
const token = useCookie('token')
const route = useRoute()
const showBack = () => {
  return route.path.startsWith('/detail') || route.path.startsWith('/settings') || route.path.startsWith('/tags') || route.query.tag
}
const { data: res, refresh } = await useAsyncData('userinfo', async () => await $fetch('/api/user/settings/get'))

settingsUpdateEvent.on(async () => {
  await refresh()
})

</script>

<style scoped></style>