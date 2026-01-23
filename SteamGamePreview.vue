<template>
  <div class="steam-card-block my-2" v-if="game && game.name">
    <a class="steam-card bg-[#1b2838] hover:bg-[#223042] flex p-3 rounded-xl border border-gray-700/50 relative overflow-hidden text-decoration-none transition-all duration-300 group" :href="game.url" target="_blank">
      
      <!-- 左侧：横版封面 -->
      <div class="flex-shrink-0 w-[140px] h-[78px] self-center relative rounded overflow-hidden shadow-md group-hover:shadow-blue-900/20 transition-shadow">
        <img 
          :src="game.image" 
          class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" 
          referrerpolicy="no-referrer" 
        />
        <!-- 价格小标签，浮在图片左下角 -->
        <div class="absolute bottom-0 right-0 bg-black/70 text-[#66c0f4] px-1.5 py-0.5 text-[10px] font-bold backdrop-blur-sm">
          {{ game.price || 'Free' }}
        </div>
      </div>

      <!-- 右侧：信息 -->
      <div class="flex flex-col justify-center flex-1 ml-3 min-w-0 overflow-hidden">
        <div class="flex justify-between items-center mb-1">
          <div class="text-[15px] font-bold text-[#c7d5e0] group-hover:text-white truncate">
            {{ game.name }}
          </div>
          <!-- Steam Logo -->
          <UIcon name="i-icon-park-solid-game-handle" class="w-4 h-4 text-gray-600 group-hover:text-[#66c0f4] transition-colors" />
        </div>
        
        <div class="text-xs text-gray-500 mb-1.5 flex items-center gap-2">
          <span>发行: {{ game.releaseDate }}</span>
        </div>

        <div class="text-xs text-[#8f98a0] line-clamp-2 leading-relaxed">
          {{ game.description }}
        </div>
      </div>
    </a>
  </div>
</template>

<script setup lang="ts">
import type { SteamGame } from '~/types';

defineProps<{
  game: SteamGame | undefined,
}>()
</script>

<style scoped>
.steam-card-block {
  width: 100%;
}
.steam-card {
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
}
.steam-card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.2);
  border-color: rgba(102, 192, 244, 0.3); /* Steam Blue Border on Hover */
}
</style>