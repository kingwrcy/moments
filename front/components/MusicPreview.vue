<template>
  <meting-js v-if="id && server && type && api" :server="server" :type="type" :id="id" :api="api"/>
</template>

<script setup lang="ts">
import type {MetingJSDTO} from "~/types";

defineProps<MetingJSDTO>()
</script>

<style scoped>

</style>