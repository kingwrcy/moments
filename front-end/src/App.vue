<script setup lang="ts">
import {  RouterView } from 'vue-router'
import {Toaster} from 'vue-sonner'
</script>

<template>
    <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md dark:text-[#C0BEBF]">
      <div class="lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-[#181818] rounded">
        <RouterView />
      </div>
    </div>
    <Toaster position="top-center" rich-colors />
</template>

<style>
html,body,#app{
  width:100%;
  height:100%
}
</style>
