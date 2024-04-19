<template>
  <div class="flex flex-wrap gap-0.5 text-lg rounded border p-2 shadow-lg select-none">
    <div v-for="(item,index) in emoji" :key="index" class="cursor-pointer" @click="emit('emojiSelected',item)">{{ item }}</div>
  </div>
</template>

<script setup lang="ts">
const emoji = ['🤔', '😐', '😑', '😶', '🙄', '😣', '😏', '😥', '😮', '🤐', '😪', '😯', '😫', '😴', '😌', '🤓', '😛', '😜', '😝', '🤤', '😒', '😓', '😔', '😕', '🙃', '🤑', '😲', '🤗', '🙂', '😚', '😙', '😗', '😘', '😍', '😎', '😋', '😊', '😉', '😆', '😅', '😄', '😃', '🤣', '😂', '😁', '😀', '😷', '🤒', '🤕', '🤢', '🤧', '😇', '🤠', '🤡', '🤥', '😈', '👿', '👹', '👺', '💀', '☠️', '👻', '👽', '💩']


const emit = defineEmits(['emojiSelected'])
</script>

<style scoped></style>