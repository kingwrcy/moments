<template>
  <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-600">
    <div class="lg:w-[567px] mx-auto bg-white shadow-2xl dark:bg-slate-900" >
    <HeaderImg />
    <slot />
    <Footer />
  </div>
  <div class="fixed right-10 bottom-10">
    <Sun color="#FDE047" :size="30" class="cursor-pointer" v-if="colorMode.value === 'dark'" @click="colorMode.value = 'light'"/>
    <MoonStar color="#FDE047" :size="30" class="cursor-pointer" v-else  @click="colorMode.value = 'dark'"/>
  </div>
  </div>
  <Toaster position="top-center" rich-colors/>
</template>

<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner'
import {Sun,MoonStar} from 'lucide-vue-next'
const colorMode = useColorMode()
console.log(colorMode.value)
</script>