<template>
  <div class="header relative mb-14">
    <img class="header-img w-full" :src="props.user.coverUrl"
         alt="">
    <div class="absolute left-2 top-2 rounded p-1 flex flex-row gap-2">
      <div title="登录">
        <UIcon name="i-carbon-user" class="text-[#9fc84a] w-5 h-5"/>
      </div>
    </div>
    <div class="absolute right-2 bottom-[-40px]">
      <div class="userinfo flex flex-col">
        <div class="flex flex-row items-center gap-4 justify-end">
          <div class="username text-lg font-bold text-white">{{ props.user.nickname }}</div>
          <img :src="props.user.avatarUrl"
               class="avatar w-[70px] h-[70px] rounded-xl"></div>
        <div class="slogon text-gray truncate w-full text-end text-xs mt-2">{{ props.user.slogan }}</div>
      </div>
    </div>
    <div class="absolute right-2 top-2 rounded p-1 flex flex-col gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#FDE047"
           stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
           class="lucide lucide-moon-star-icon cursor-pointer">
        <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
        <path d="M20 3v4"></path>
        <path d="M22 5h-4"></path>
      </svg>
      <NuxtLink to="/new" v-if="$route.path === '/'">
        <UIcon name="i-carbon-edit" class="text-[#9fc84a] w-5 h-5"  />
      </NuxtLink>
      <NuxtLink to="/" v-else><UIcon name="i-carbon-arrow-left" class="text-[#9fc84a] w-5 h-5" /></NuxtLink>
      <NuxtLink to="/sys/settings"  v-if="currentUser && currentUser.id === 1"><UIcon name="i-carbon-settings" class="text-[#9fc84a] w-5 h-5" /></NuxtLink>
      <NuxtLink to="/user/settings"  v-if="currentUser"><UIcon name="i-carbon-user-profile" class="text-[#9fc84a] w-5 h-5" /></NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
const currentUser = useState<UserVO>('userinfo')

const props = defineProps<{ user: UserVO }>()
</script>

<style scoped>

</style>