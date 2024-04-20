<template>
  <div class="footer p-4 flex justify-center items-center text-xs ">
    <div class="flex flex-col gap-1 items-center">
      <NuxtLink to="https://github.com/kingwrcy/moments">
        <img src="https://img.shields.io/github/stars/kingwrcy/moments" class="h-6 inline-block mr-1" alt="github" />
      </NuxtLink>
      <div class="my-2 text-gray-500" v-if="userinfo.beianNo">{{ userinfo.beianNo }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')
</script>

<style scoped>

</style>