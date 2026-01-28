<template>
  <div class="tmdb-card-block my-2" v-if="item && item.title">
    <a class="tmdb-card bg-[#0d253f] hover:bg-[#133050] flex p-3 rounded-xl border border-gray-700/50 relative overflow-hidden text-decoration-none transition-all duration-300 group" 
       :href="item.url" 
       target="_blank">
      
      <!-- 背景氛围图 -->
      <div class="absolute right-0 top-0 w-2/3 h-full bg-cover bg-center opacity-[0.08] group-hover:opacity-[0.15] transition-opacity pointer-events-none mask-image-gradient" 
           :style="`background-image: url('${item.backdropPath || item.posterPath}')`">
      </div>
      
      <!-- 左侧：海报 -->
      <div class="flex-shrink-0 w-[72px] h-[108px] relative rounded overflow-hidden shadow-lg border border-white/5 z-10">
        <img class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" 
             :src="item.posterPath" 
             referrerpolicy="no-referrer" 
             :alt="item.title" />
      </div>
      
      <!-- 右侧：信息 -->
      <div class="flex flex-col flex-1 ml-3 min-w-0 z-10 justify-between py-0.5">
        <div>
          <!-- 第一行：标题 + 评分 -->
          <div class="flex justify-between items-start">
            <div class="text-[15px] font-bold text-[#e1e1e1] group-hover:text-white truncate">
              {{ item.title }}
            </div>
            <div v-if="item.voteAverage" 
                 class="flex items-center gap-0.5 text-xs font-bold text-yellow-500 bg-yellow-500/10 px-1.5 py-0.5 rounded ml-2 flex-shrink-0">
              <UIcon name="i-heroicons-star-solid" class="w-3 h-3" />
              {{ item.voteAverage }}
            </div>
          </div>
          
          <!-- 第二行：基本信息 + 导演/主创 -->
          <div class="text-xs text-gray-400 mt-1 flex flex-wrap items-center gap-x-2">
            <span v-if="item.releaseDate">
              {{ item.releaseDate.split('-')[0] }}
            </span>
            
            <span v-if="item.releaseDate" class="w-0.5 h-0.5 bg-gray-500 rounded-full"></span>
            
            <span>{{ item.type === 'tv' ? '剧集' : '电影' }}</span>
            
            <!-- 导演/主创信息 -->
            <template v-if="item.director">
              <span class="w-0.5 h-0.5 bg-gray-500 rounded-full"></span>
              <span class="truncate max-w-[140px]" :title="item.director">
                {{ item.type === 'tv' ? '主创:' : '导演:' }} {{ item.director }}
              </span>
            </template>
          </div>
          
          <!-- 第三行：主演 -->
          <div v-if="item.actors" 
               class="text-xs text-gray-400 mt-1 truncate opacity-90" 
               :title="item.actors">
            <span class="opacity-70">主演: </span>{{ item.actors }}
          </div>
        </div>
        
        <!-- 第四行：简介 -->
        <div class="text-xs text-gray-300 mt-1 line-clamp-2 leading-relaxed opacity-80">
          {{ item.overview }}
        </div>
      </div>
    </a>
  </div>
</template>

<script setup lang="ts">
import type { TmdbItem } from '~/types';

defineProps<{
  item: TmdbItem | undefined,
}>()
</script>

<style scoped>
.tmdb-card-block {
  width: 100%;
}

.tmdb-card {
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
}

.tmdb-card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.2);
  border-color: rgba(1, 180, 228, 0.3);
}

.mask-image-gradient {
  -webkit-mask-image: linear-gradient(to right, transparent, black);
  mask-image: linear-gradient(to right, transparent, black);
}
</style>