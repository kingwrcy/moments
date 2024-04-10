<template>
  <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md">
    <div class="lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-slate-600">
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

const { data: res } = await useFetch('/api/user/settings/get')
useState<User>('userinfo', () => (res.value?.data as any as User))

useHead({
  link: [
    {
      rel: 'shortcut icon',
      type: 'image/png',
      href: res.value?.data?.favicon || '/favicon.png',
    },
  ],
})
</script>