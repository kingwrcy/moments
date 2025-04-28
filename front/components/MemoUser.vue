<template>
  <div
    class="flex flex-row text-sm w-full"
    :class="{ 'bg-slate-100 dark:bg-neutral-800': props.memo.pinned }"
  >
    <div class="flex flex-col w-24 pt-2">
      <template v-if="!isPinned">
        <div class="flex justify-center">
          <span class="text-xl font-bold">{{
            $dayjs(item.createdAt).format("DD")
          }}</span>
          <span class="flex items-end text-xs"
            >{{ $dayjs(item.createdAt).format("MM") }}月</span
          >
        </div>
        <div
          class="flex justify-center text-[#576b95] font-medium dark:text-white text-xs mt-2 select-none"
        >
          {{ location }}
        </div>
      </template>
      <div v-else class="flex justify-center items-center">
        <span class="text-lg">置顶</span>
      </div>
    </div>
    <div class="flex w-full flex-col pr-4 py-2">
      <NuxtLink class="flex" :to="`/memo/${item.id}`">
        <div
          class="sm:w-24 sm:h-24 w-20 h-20 relative overflow-hidden"
          v-if="imageCount > 0"
        >
          <div
            :class="getImageGridClass(imageCount)"
            class="h-full w-full gap-0.1 grid"
          >
            <div
              v-for="(img, index) in images.slice(0, Math.min(imageCount, 9))"
              :key="index"
              :class="getGridClass(index, imageCount)"
              class="border border-white dark:border-neutral-800 relative"
            >
              <img
                :src="img"
                class="absolute inset-0 w-full h-full object-cover"
              />
            </div>
          </div>
        </div>
        <div class="flex-1 flex flex-col justify-between">
          <div
            class="markdown-content bg-neutral-100 dark:bg-neutral-800 p-2 !leading-7 line-clamp-3"
            v-if="imageCount === 0"
            v-html="content"
          ></div>
          <div
            class="markdown-content ml-1 !leading-5 line-clamp-3"
            v-if="imageCount > 0"
            v-html="content"
          ></div>
          <div v-if="imageCount > 0" class="text-sm text-gray-500 ml-1">
            有{{ imageCount }}图
          </div>
        </div>
      </NuxtLink>
      <div class="flex flex-col gap-2">
        <external-url-preview
          v-if="item.externalFavicon && item.externalTitle && item.externalUrl"
          :favicon="item.externalFavicon"
          :title="item.externalTitle"
          :url="item.externalUrl"
          class="pt-2"
        />
        <music-preview
          v-if="extJSON.music && extJSON.music.id"
          v-bind="extJSON.music"
          class="pt-2"
        />
        <douban-book-preview
          v-if="extJSON.doubanBook && extJSON.doubanBook.title"
          :book="extJSON.doubanBook"
          class="pt-2"
        />
        <douban-movie-preview
          v-if="extJSON.doubanMovie && extJSON.doubanMovie.title"
          :movie="extJSON.doubanMovie"
          class="pt-2"
        />
        <video-preview-iframe
          v-if="
            extJSON.video &&
            ['bilibili', 'youtube'].includes(extJSON.video.type) &&
            extJSON.video.value
          "
          :url="extJSON.video.value"
          class="pt-2"
        />
        <video-preview
          v-if="
            extJSON.video &&
            extJSON.video.type === 'online' &&
            extJSON.video.value
          "
          :url="extJSON.video.value"
          class="pt-2"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ExtDTO, MemoVO, SysConfigVO } from "~/types";
import { md } from "~/utils";
import { ref, computed } from "vue";
import { toast } from "vue-sonner";
import { useRouter } from "vue-router";

const router = useRouter();
const sysConfig = useState<SysConfigVO>("sysConfig");
const props = defineProps<{
  memo: MemoVO;
}>();

const item = computed(() => props.memo);
const isPinned = computed(() => item.value.pinned);
const location = computed(() =>
  (item.value.location || "").replaceAll(" ", " · ")
);

const extJSON = computed(() => {
  try {
    return JSON.parse(item.value.ext || "{}") as ExtDTO;
  } catch (error) {
    console.error("解析 ext 字段时出错:", error);
    return {} as ExtDTO;
  }
});

const content = computed(() => {
  if (item.value.content && item.value.content.length > 0) {
    try {
      return md.render(item.value.content);
    } catch (error) {
      console.error("内容渲染错误，请重新编辑:", error);
      toast.error("内容渲染错误，请重新编辑");
      return "内容渲染错误，请重新编辑";
    }
  }
  return "";
});

const imageCount = computed(() => {
  const imgs = item.value.imgs || "";
  return imgs.split(",").filter(Boolean).length;
});

const images = computed(() => {
  const imgs = item.value.imgs || "";
  return imgs.split(",").filter(Boolean);
});

const getGridClass = (index: number, count: number) => {
  switch (count) {
    case 3:
      if (index === 0) return "col-span-1 row-span-2";
      return "";
    case 5:
      if (index === 0) return "col-span-2 row-span-2";
      if (index === 1) return "col-span-1 row-span-2";
      return "";
    case 6:
      if (index === 0) return "col-span-2 row-span-2";
      return "";
    case 7:
      if (index === 0 || index === 1) return "col-span-1 row-span-2";
      return "";
    case 8:
      if (index === 0) return "col-span-1 row-span-2";
      return "";
    default:
      return "";
  }
};

const getImageGridClass = (count: number) => {
  switch (count) {
    case 1:
      return "";
    case 2:
      return "grid grid-cols-2";
    case 3:
      return "grid grid-cols-2";
    case 4:
      return "grid grid-cols-2 grid-rows-2";
    case 5:
    case 6:
    case 7:
    case 8:
    case 9:
      return "grid grid-cols-3 grid-rows-3";
    default:
      return "";
  }
};
</script>

<style scoped></style>
