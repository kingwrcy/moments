<template>
  <div class="w-full md:w-[567px] mx-auto h-full shadow-2xl ">
    <slot/>
    <Footer/>
  </div>
  <div class="sm:hidden relative">
    <div class="left-0 bottom-10 w-full fixed  flex items-center justify-end" v-if="$route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <NuxtLink to="/new" class="mr-4 rounded-full bg-slate-50 w-14 h-14 flex items-center justify-center shadow-xl">
          <UIcon name="i-carbon-edit" class="w-8 h-8 text-[#9fc84a]"></UIcon>
        </NuxtLink>
      </div>
    </div>

    <div class="right-2 top-2 w-full fixed  flex items-center justify-end" v-if="$route.path === '/'">
      <div class="flex flex-col items-center gap-2">
        <div class="flex rounded bg-slate-50 p-1 items-center justify-center shadow-xl" @click="open = true">
          <UIcon name="i-carbon-settings" class="w-4 h-4 text-[#9fc84a]" ></UIcon>
        </div>
      </div>
    </div>
    <MobileNav :open="open"/>
  </div>
</template>

<script lang="ts" setup>
import type {SysConfigVO, UserVO} from "~/types";
const open = useState<boolean>('sidebarOpen',()=>false)
const currentUser = useState<UserVO>('userinfo')
const sysConfig = useState<SysConfigVO>('sysConfig')
const currentProfile = await useMyFetch<UserVO>("/user/profile")
const sysConfigVO = await useMyFetch<SysConfigVO>("/sysConfig/get")
if (currentProfile) {
  currentUser.value = currentProfile
  sysConfig.value = sysConfigVO
}
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
