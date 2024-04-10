<template>
  <div class="relative mb-8 ">
    <img class="w-full  max-h-[300px]" :src="getImgUrl(res?.data?.coverUrl!)" alt="" />
    <div class="absolute right-2 bottom-[-40px]">
      <div class="flex flex-col ">
        <div class="flex flex-row items-center gap-4 justify-end">
          <div class="text-lg font-bold text-white">{{ res?.data?.nickname }}</div>
          <img :src="getImgUrl(res?.data?.avatarUrl!)" class="w-[70px] h-[70px] rounded-xl" />
        </div>
        <div class="text-gray truncate w-full text-end text-xs mt-2">{{ res?.data?.slogan }}</div>
      </div>
    </div>

    <div class="absolute right-2 top-2 bg-slate-100 rounded p-1 flex flex-row gap-2">
      <div title="返回" v-if="showBack()" @click="navigateTo('/')">
        <ArrowLeft color="#9FC84A" :size="20" class="cursor-pointer" />
      </div>
      <div title="登录" v-if="!token">
        <LogIn :size="20" class="cursor-pointer" color="#9FC84A" @click="navigateTo('/login')" />
      </div>

      <div title="亮色" v-if="colorMode.value === 'dark'" @click="colorMode.value = 'light'">
        <Sun color="#FDE047" :size="20" class="cursor-pointer" />
      </div>
      <div title="暗色" v-else>
        <MoonStar color="#FDE047" :size="20" class="cursor-pointer" @click="colorMode.value = 'dark'" />
      </div>
    </div>

  </div>

</template>

<script setup lang="ts">
import { settingsUpdateEvent } from '~/lib/event'
import { getImgUrl } from '~/lib/utils';
import { Sun, MoonStar, LogIn, ArrowLeft } from 'lucide-vue-next'
const colorMode = useColorMode()
const token = useCookie('token')
const route = useRoute()
const showBack = () => {
  return route.path.startsWith('/detail') || route.path.startsWith('/settings')
}
const { data: res, refresh } = await useFetch('/api/user/settings/get')

settingsUpdateEvent.on(async () => {
  await refresh()
})

</script>

<style scoped></style>