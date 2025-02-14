<template>
  <div class="w-full md:w-[567px] mx-auto h-full shadow-2xl">
    <slot/>
    <Footer/>
  </div>

  <div title="到顶部" v-if="y>200 && $route.path === '/'"
       class="hidden sm:block bottom-[20%] sm:right-[20%] md:right-[10%] lg:right-[15%] xl:right-[20%] 2xl:right-[28%] fixed  flex items-center justify-center">
    <UIcon name="i-lets-icons-expand-top-stop" class="w-10 h-10 text-gray-500 cursor-pointer" @click="y=0"></UIcon>
  </div>


  <div class="sm:hidden relative">
    <div class="right-0 bottom-10 fixed flex items-center justify-end"
         v-if="global.userinfo.token">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <UIcon name="i-lets-icons-expand-top-stop" class="w-6 h-6 text-[#9fc84a] cursor-pointer"></UIcon>
        </div>
        <NuxtLink to="/new" v-if="$route.path === '/'" class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-camera" class="w-6 h-6 text-[#9fc84a]"></UIcon>
        </NuxtLink>
        <div class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl" @click="open = true">
          <UIcon name="i-icon-park-solid-more-four" class="w-6 h-6 text-[#9fc84a] cursor-pointer"></UIcon>
        </div>
      </div>
    </div>

    <div class="right-0 bottom-10 fixed flex items-center justify-end"
         v-if="!global.userinfo.token && $route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <UIcon name="i-lets-icons-expand-top-stop" class="w-6 h-6 text-[#9fc84a] cursor-pointer"></UIcon>
        </div>
        <NuxtLink to="/user/login"
                  class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-login" class="w-6 h-6 text-[#9fc84a]"></UIcon>
        </NuxtLink>
      </div>
    </div>

    <MobileNav :open="open"/>
  </div>
</template>

<script lang="ts" setup>
import type {SysConfigVO, UserVO} from "~/types";
import {useGlobalState} from "~/store";

const global = useGlobalState()
const open = useState<boolean>('sidebarOpen', () => false)
const currentUser = useState<UserVO>('userinfo')
const sysConfig = useState<SysConfigVO>('sysConfig')
const currentProfile = await useMyFetch<UserVO>("/user/profile")
const sysConfigVO = await useMyFetch<SysConfigVO>("/sysConfig/get")
if (currentProfile) {
  currentUser.value = currentProfile
  sysConfig.value = sysConfigVO
}
const {y} = useWindowScroll()
useHead({
  title: sysConfigVO.title,
  link: [
    {
      rel: 'shortcut icon',
      type: 'image/png',
      href: sysConfigVO.favicon || '/favicon.png',
    },
    {
      rel: 'apple-touch-icon-precomposed',
      href: sysConfigVO.favicon || '/favicon.png',
    },
    {
      rel: 'alternate',
      type: 'application/rss+xml',
      title: '我的 RSS 订阅',
      href: sysConfigVO.rss || `/rss`,
    },
  ],
  style: [
    {
      innerHTML: sysConfigVO.css || '',
    }
  ],
  script: [
    {
      type: 'text/javascript',
      innerHTML: sysConfigVO.js || '',
    },

  ]
})

if (sysConfigVO.enableGoogleRecaptcha) {
  useHead({
    script: [
      {
        type: 'text/javascript',
        src: `https://recaptcha.net/recaptcha/api.js?render=${sysConfigVO.googleSiteKey}`,
      },
    ],
  })
}
</script>
