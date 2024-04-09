<template>
  <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md">
    <div class="lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-slate-600">
      <HeaderImg />
      <slot />
      <Footer />
    </div>
    <div class="fixed right-10 bottom-10">
      <LogIn  :size="30" class="cursor-pointer my-4"  color="#9FC84A"
        @click="navigateTo('/login')" />
      <Sun color="#FDE047" :size="30" class="cursor-pointer" v-if="colorMode.value === 'dark'"
        @click="colorMode.value = 'light'" />
      <MoonStar color="#FDE047" :size="30" class="cursor-pointer" v-else @click="colorMode.value = 'dark'" />
    </div>
  </div>
  <Toaster position="top-center" rich-colors />
</template>

<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner'
import { Sun, MoonStar, LogIn } from 'lucide-vue-next'
const colorMode = useColorMode()
</script>