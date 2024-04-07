<template>
  <div class="relative mb-8">
    <div class="absolute top-0 left-0 w-full bg-white/80" v-if="showBack()">
      <img src="~/assets/img/back.svg" class="w-10 h-10 px-1 py-2 cursor-pointer" alt="返回" title="返回"
        @click="navigateTo('/')">
    </div>
    <NuxtImg class="w-full h-[250px]"
      :src="res?.data?.coverUrl!"
      alt=""/>
    <div class="absolute right-2 bottom-[-40px]">
      <div class="flex flex-col">
        <div class="flex flex-row items-center gap-4">
          <div class="text-lg font-bold text-white">Jerry</div>
          <NuxtImg :src="res?.data?.avatarUrl!"
            class="w-[70px] h-[70px] rounded-xl" />
        </div>
        <div class="text-gray truncate w-full text-end text-xs mt-2">{{res?.data?.slogan}}</div>
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">
import {settingsUpdateEvent} from '~/lib/event'

const route = useRoute()
const showBack = ()=>{
  return route.path.startsWith('/detail') || route.path.startsWith('/settings')
}
const { data: res ,refresh} = await useFetch('/api/user/settings/get')

settingsUpdateEvent.on(async ()=>{
  await refresh()
})

</script>

<style scoped></style>