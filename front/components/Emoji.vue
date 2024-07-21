<template>
  <div @click="selectEmoji" class="flex flex-wrap gap-0.5 text-lg rounded border p-2 shadow-lg select-none mt-2 *:cursor-pointer">
    <div>🤔</div>
    <div>😐</div>
    <div>😑</div>
    <div>😶</div>
    <div>🙄</div>
    <div>😣</div>
    <div>😏</div>
    <div>😥</div>
    <div>😮</div>
    <div>🤐</div>
    <div>😪</div>
    <div>😯</div>
    <div>😫</div>
    <div>😴</div>
    <div>😌</div>
    <div>🤓</div>
    <div>😛</div>
    <div>😜</div>
    <div>😝</div>
    <div>🤤</div>
    <div>😒</div>
    <div>😓</div>
    <div>😔</div>
    <div>😕</div>
    <div>🙃</div>
    <div>🤑</div>
    <div>😲</div>
    <div>🤗</div>
    <div>🙂</div>
    <div>😚</div>
    <div>😙</div>
    <div>😗</div>
    <div>😘</div>
    <div>😍</div>
    <div>😎</div>
    <div>😋</div>
    <div>😊</div>
    <div>😉</div>
    <div>😆</div>
    <div>😅</div>
    <div>😄</div>
    <div>😃</div>
    <div>🤣</div>
    <div>😂</div>
    <div>😁</div>
    <div>😀</div>
    <div>😷</div>
    <div>🤒</div>
    <div>🤕</div>
    <div>🤢</div>
    <div>🤧</div>
    <div>😇</div>
    <div>🤠</div>
    <div>🤡</div>
    <div>🤥</div>
    <div>😈</div>
    <div>👿</div>
    <div>👹</div>
    <div>👺</div>
    <div>💀</div>
    <div>☠️</div>
    <div>👻</div>
    <div>👽</div>
    <div>💩</div>
  </div>
</template>

<script setup lang="ts">
const emit = defineEmits(['selected'])
const selectEmoji = (event:Event)=>{
  const value = (event.target as HTMLDivElement).innerHTML
  emit('selected',value)
}
</script>

<style scoped>

</style>