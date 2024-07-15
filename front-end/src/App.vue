<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <n-message-provider>
    <div class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md dark:text-[#C0BEBF]">
      <div class="lg:w-[567px] mx-auto shadow-2xl bg-white dark:bg-[#181818] rounded">
        <RouterView />
      </div>
    </div>
  </n-message-provider>
</template>

<style>
html,body,#app{
  width:100%;
  height:100%
}
</style>
