<template>
  <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md dark:text-[#C0BEBF]">
    <div class="lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-[#181818]">
      <HeaderImg />
      <slot />
      <Footer />
    </div>

  </div>
  <Toaster position="top-center" rich-colors />
</template>

<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner';
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')
await callOnce(async () => {
  const { data: res } = await useAsyncData('userinfo', async () => await $fetch('/api/user/settings/get'))
  userinfo.value = res.value?.data as any as User
})

useHead({
  link: [
    {
      rel: 'shortcut icon',
      type: 'image/png',
      href: userinfo.value?.favicon || '/favicon.png',
    },
  ],
})
</script>