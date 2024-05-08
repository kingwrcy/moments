<template>
  <div class="wrapper w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md dark:text-[#C0BEBF]">
    <ScrollArea class="h-full" type="hover">
      <div class="main lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-[#181818]">
        <HeaderImg />
        <slot />
        <Footer />
      </div>
      <ScrollBar orientation="vertical" />
    </ScrollArea>
  </div>

  <Toaster position="top-center" rich-colors />

</template>

<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner';
import type { Memo, User } from '~/lib/types';
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
import { memoUpdateEvent } from '~/lib/event';


const { data } = await useFetch('/api/user/validateToken', {
  method: 'POST'
})
if (!data.value?.success) {
  const token = useCookie('token')
  token.value = undefined
}


const userinfo = useState<User>('userinfo')

await callOnce(async () => {
  const { data: res } = await useAsyncData('userinfo', async () => await $fetch('/api/user/settings/get'))
  userinfo.value = res.value?.data as any as User
})

const config = useRuntimeConfig()

memoUpdateEvent.on((event: Memo & { index?: number }) => {
  const target = document.querySelector('div[data-radix-scroll-area-viewport]')
  if (target) {
    target.scrollTop = 0
  }
})


useHead({
  link: [
    {
      rel: 'shortcut icon',
      type: 'image/png',
      href: userinfo.value?.favicon || '/favicon.png',
    },
  ],
  style: [
    {
      innerHTML: userinfo.value?.css || '',
    }
  ],
  script: [
    {
      type: 'text/javascript',
      innerHTML: userinfo.value?.js || '',
    }
  ]
})

if (config.public.googleRecaptchaSiteKey) {
  useHead({
    script: [
      {
        type: 'text/javascript',
        src: 'https://recaptcha.net/recaptcha/api.js?render=' + config.public.googleRecaptchaSiteKey
      }
    ]
  })
}
</script>
