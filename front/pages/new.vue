<template>
  <div>
    <Header v-if="currentUser" v-bind:user="currentUser"/>
    <MemoEdit/>

  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import MemoEdit from "~/components/MemoEdit.vue";

const currentUser = useState<UserVO>('userinfo')

</script>

<style scoped>

</style>