<template>
      <div class="douban-card-block">
        <a class="douban-card" :href="book?.url" target="_blank">
          <div  class="douban-card-bgimg" :style="`background-image: url('${book?.image}')`">
          </div>
          <div class="douban-card-left">
            <div class="douban-card-img" :style="`background-image: url('${book?.image}')`">
            </div>
          </div>
          <div class="douban-card-right" style="line-height: 1.7;">
            <div class="douban-card-item"><span>书名: </span><strong>{{ book?.title }}</strong></div>
            <div class="douban-card-item"><span>作者: </span><span>{{ book?.author }}</span></div>
            <div class="douban-card-item"><span>出版年份: </span><span>{{ book?.pubDate }}</span></div>
            <div class="douban-card-item"><span>评分: </span><span>{{ book?.rating }}</span></div>
          </div>
        </a>
      </div>
</template>

<script setup lang="ts">
import type { DoubanBook } from '~/lib/types';
const props = withDefaults(
  defineProps<{
    book: DoubanBook | undefined,
  }>(), {}
)
</script>

<style scoped>
.douban-card-block {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  max-height: 400px;
}

.douban-card {
  display: flex;
  margin: 10px;
  padding: 15px;
  border-radius: 15px;
  position: relative;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  color: antiquewhite;
  text-decoration: none;
}

.douban-card:hover {
  text-decoration: none;
}

.douban-card-bgimg {
  position: absolute;
  width: 115%;
  height: 115%;
  filter: blur(15px) brightness(0.6);
  background-size: 100%;
  background-position: center;
  background-repeat: no-repeat;
}

.douban-card-img {
  position: relative;
  height: 130px;
  width: 80px;
  background-size: 100%;
  background-position: center;
  background-repeat: no-repeat;
}

.douban-card-left:hover .douban-card-img {
  filter: blur(5px) brightness(0.6);
  transform: perspective(800px) rotateX(180deg);
}

.douban-card-left .douban-card-img {
  transition: all 500ms ease;
}

.douban-card-left {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.douban-card-left .douban-card-status {
  height: 130px;
  width: 80px;
  text-align: center;
  font-weight: bold;
  position: absolute;
  left: 0;
  top: 30%;
  transform: rotateX(180deg);
  backface-visibility: hidden;
  transition: all 500ms ease;
}

.douban-card-left:hover .douban-card-status {
  transform: perspective(800px) rotateX(0deg);
}

.douban-card-right {
  position: relative;
  display: flex;
  flex-direction: column;
  margin-left: 12px;
  font-size: 16px;
  font-family: "Courier New", Courier, monospace;
  line-height: 1.3;
  color: antiquewhite;
}

.douban-card-item {
  margin-top: 4px;
}
</style>