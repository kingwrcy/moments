<template>
  <div class="w-full md:w-[567px] mx-auto h-full shadow-2xl ">
    <slot/>
    <Footer/>
  </div>
</template>

<script lang="ts" setup>
import type {SysConfigVO, UserVO} from "~/types";
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
    }
  ]
})



</script>
