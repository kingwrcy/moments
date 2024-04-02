<template>
  <div class="lg:w-[567px] mx-auto bg-white shadow-2xl">
    <slot />
  </div>
  <Toaster position="top-center" rich-colors/>
</template>

<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner'
</script>