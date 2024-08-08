<template>
  <div class="w-full md:w-[567px] mx-auto h-full shadow-2xl">
    <slot/>
    <Footer/>
  </div>

  <div title="到顶部" v-if="y>200"
       class="fixed bottom-[20%] sm:right-[20%] md:right-[10%] lg:right-[15%] xl:right-[20%] 2xl:right-[28%] fixed  flex items-center justify-center">
    <UIcon name="i-carbon-up-to-top" class="w-12 h-12 text-gray-500 cursor-pointer" @click="y=0"></UIcon>
  </div>


  <div class="sm:hidden relative">
    <div class="left-0 bottom-10 w-full fixed  flex items-center justify-end"
         v-if="global.userinfo.token && $route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="mr-4 rounded-full bg-slate-50 w-14 h-14 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-up-to-top" class="w-8 h-8 text-[#9fc84a]"></UIcon>
        </div>
        <NuxtLink to="/new" class="mr-4 rounded-full bg-slate-50 w-14 h-14 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-edit" class="w-8 h-8 text-[#9fc84a]"></UIcon>
        </NuxtLink>
      </div>
    </div>

    <div class="left-0 bottom-10 w-full fixed  flex items-center justify-end"
         v-if="!global.userinfo.token && $route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="mr-4 rounded-full bg-slate-50 w-14 h-14 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-up-to-top" class="w-8 h-8 text-[#9fc84a]"></UIcon>
        </div>
        <NuxtLink to="/user/login"
                  class="mr-4 rounded-full bg-slate-50 w-14 h-14 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-login" class="w-8 h-8 text-[#9fc84a]"></UIcon>
        </NuxtLink>
      </div>
    </div>


    <div class="right-2 top-2 w-full fixed  flex items-center justify-end" v-if="global.userinfo.token">
      <div class="flex flex-col items-center gap-2">
        <div class="flex rounded bg-slate-50 p-1 gap-2 items-center justify-center shadow-xl" @click="open = true">
          <UIcon name="i-carbon-settings" class="w-4 h-4 text-[#9fc84a]"></UIcon>
        </div>
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
    }
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
