<template>
  <div class="w-full md:w-[567px] mx-auto h-full shadow-2xl">
    <slot/>
    <Footer/>
  </div>

  <div title="到顶部" v-if="y>200 && $route.path === '/'"
       class="hidden sm:block bottom-[20%] sm:right-[20%] md:right-[10%] lg:right-[15%] xl:right-[20%] 2xl:right-[28%] fixed  flex items-center justify-center">
    <svg class="w-10 h-10 text-gray-500 cursor-pointer" @click="y=0" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-width="2" d="m18 17l-6-6l-6 6M7 6h10"/></svg>
  </div>


  <div class="sm:hidden relative">
    <div class="right-0 bottom-10 fixed flex items-center justify-end"
         v-if="global.userinfo.token">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <svg class="w-6 h-6 text-[#9fc84a] cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-width="2" d="m18 17l-6-6l-6 6M7 6h10"/></svg>
        </div>
        <NuxtLink to="/new" v-if="$route.path === '/'" class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <svg class="text-[#9fc84a] w-6 h-6 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M29 26H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h6.46l1.71-2.55A1 1 0 0 1 12 4h8a1 1 0 0 1 .83.45L22.54 7H29a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1M4 24h24V9h-6a1 1 0 0 1-.83-.45L19.46 6h-6.92l-1.71 2.55A1 1 0 0 1 10 9H4Z"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a6 6 0 0 1-6 6m0-10a4 4 0 1 0 4 4a4 4 0 0 0-4-4"/></svg>
        </NuxtLink>
        <div class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl" @click="open = true">
          <svg class="w-6 h-6 text-[#9fc84a] cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48"><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 11l7.071-7.071L31.142 11l-7.07 7.071zm13 13l7.071-7.071L44.142 24l-7.07 7.071zM4 24l7.071-7.071L18.142 24l-7.07 7.071zm13 13l7.071-7.071L31.142 37l-7.07 7.071z"/></svg>
        </div>
      </div>
    </div>

    <div class="right-0 bottom-10 fixed flex items-center justify-end"
         v-if="!global.userinfo.token && $route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <div v-if="y>300" @click="y=0"
             class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <svg class="w-6 h-6 text-[#9fc84a] cursor-pointer" @click="y=0" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-width="2" d="m18 17l-6-6l-6 6M7 6h10"/></svg>
        </div>
        <NuxtLink to="/user/login"
                  class="dark:bg-gray-900/85 mr-4 rounded-full bg-slate-50 w-10 h-10 flex items-center justify-center shadow-xl">
          <svg class="text-[#9fc84a] w-6 h-6 cursor-pointer" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path fill="currentColor" d="M26 30H14a2 2 0 0 1-2-2v-3h2v3h12V4H14v3h-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2"/><path fill="currentColor" d="M14.59 20.59L18.17 17H4v-2h14.17l-3.58-3.59L16 10l6 6l-6 6z"/></svg>
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
