<template>
  <div class="footer p-4 flex justify-center items-center text-xs ">
    <div class="flex flex-col gap-1 items-center">
      <NuxtLink to="https://github.com/kingwrcy/moments" target="_blank">
        <img src="https://img.shields.io/github/stars/kingwrcy/moments" class="h-6 inline-block mr-1" alt="github" />
      </NuxtLink>
      <a class="my-2 text-gray-500" v-if="sysConfig.beiAnNo" href="https://beian.miit.gov.cn/" target="_blank">{{ sysConfig.beiAnNo }}</a>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {SysConfigVO} from "~/types";

const sysConfig = useState<SysConfigVO>('sysConfig')

</script>

<style scoped>

</style>