<template>
  <div
    class="flex flex-row text-sm w-full"
    :class="{ 'bg-slate-100 dark:bg-neutral-800': props.memo.pinned }"
  >
    <div class="flex flex-col w-24 pt-2">
      <template v-if="!isPinned">
        <div class="flex justify-center">
          <span class="text-xl font-bold">{{ formattedDate.day }}</span>
          <span class="flex items-end text-xs"
            >{{ formattedDate.month }}月</span
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
        <div class="sm:w-24 sm:h-24 w-20 h-20" v-if="imageCount > 0">
          <div
            v-if="imageCount === 1"
            class="h-full w-full border border-white dark:border-neutral-800"
          >
            <img :src="images[0]" class="h-full w-full object-cover" />
          </div>
          <div v-if="imageCount === 2" class="h-full w-full flex gap-0.1">
            <div class="w-1/2 border border-white dark:border-neutral-800">
              <img :src="images[0]" class="h-full w-full object-cover" />
            </div>
            <div class="w-1/2 border border-white dark:border-neutral-800">
              <img :src="images[1]" class="h-full w-full object-cover" />
            </div>
          </div>
          <div v-if="imageCount === 3" class="h-full w-full flex gap-0.1">
            <div class="w-1/2 border border-white dark:border-neutral-800">
              <img :src="images[0]" class="h-full w-full object-cover" />
            </div>
            <div class="w-1/2 flex flex-col gap-0.1">
              <div class="h-1/2 border border-white dark:border-neutral-800">
                <img :src="images[1]" class="h-full w-full object-cover" />
              </div>
              <div class="h-1/2 border border-white dark:border-neutral-800">
                <img :src="images[2]" class="h-full w-full object-cover" />
              </div>
            </div>
          </div>

          <div
            v-if="imageCount === 4"
            class="h-full w-full grid grid-cols-2 grid-rows-2 gap-0.1"
          >
            <div
              class="border border-white dark:border-neutral-800"
              v-for="(img, index) in images.slice(0, 4)"
              :key="index"
            >
              <img :src="img" class="h-full w-full object-cover" />
            </div>
          </div>

          <div
            v-if="imageCount >= 5 && imageCount <= 9"
            class="h-full w-full grid grid-cols-3 grid-rows-3 gap-0.1"
          >
            <div
              v-if="imageCount >= 1"
              :class="getGridClass(0, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[0]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 2"
              :class="getGridClass(1, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[1]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 3"
              :class="getGridClass(2, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[2]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 4"
              :class="getGridClass(3, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[3]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 5"
              :class="getGridClass(4, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[4]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 6"
              :class="getGridClass(5, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[5]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 7"
              :class="getGridClass(6, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[6]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount >= 8"
              :class="getGridClass(7, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[7]" class="h-full w-full object-cover" />
            </div>
            <div
              v-if="imageCount === 9"
              :class="getGridClass(8, imageCount)"
              class="border border-white dark:border-neutral-800"
            >
              <img :src="images[8]" class="h-full w-full object-cover" />
            </div>
          </div>
        </div>
        <div class="flex-1 flex flex-col justify-between">
          <div
            ref="contentRef"
            class="markdown-content bg-neutral-100 dark:bg-neutral-800 p-2"
            v-if="imageCount === 0"
            v-html="content"
          ></div>
          <div
            ref="contentRef"
            class="markdown-content ml-1 !leading-5"
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
          v-if="hasExternalUrl"
          :favicon="item.externalFavicon"
          :title="item.externalTitle"
          :url="item.externalUrl"
          class="pt-2"
        />
        <music-preview v-if="hasMusic" v-bind="extJSON.music" class="pt-2" />
        <douban-book-preview
          v-if="hasDoubanBook"
          :book="extJSON.doubanBook"
          class="pt-2"
        />
        <douban-movie-preview
          v-if="hasDoubanMovie"
          :movie="extJSON.doubanMovie"
          class="pt-2"
        />
        <video-preview-iframe
          v-if="hasVideoIframe"
          :url="extJSON.video.value"
          class="pt-2"
        />
        <video-preview
          v-if="hasVideo"
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
import { useGlobalState } from "~/store";
import { ref, computed } from "vue";
import { toast } from "vue-sonner";

const router = useRouter();
const sysConfig = useState<SysConfigVO>("sysConfig");
const props = defineProps<{
  memo: MemoVO;
}>();

const item = computed(() => props.memo);
const isPinned = computed(() => item.value.pinned);
const formattedDate = computed(() => {
  const date = new Date(item.value.createdAt);
  return {
    day: date.getDate().toString().padStart(2, "0"),
    month: (date.getMonth() + 1).toString().padStart(2, "0"),
  };
});
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

const hasExternalUrl = computed(
  () =>
    item.value.externalFavicon &&
    item.value.externalTitle &&
    item.value.externalUrl
);
const hasMusic = computed(() => extJSON.value.music && extJSON.value.music.id);
const hasDoubanBook = computed(
  () => extJSON.value.doubanBook && extJSON.value.doubanBook.title
);
const hasDoubanMovie = computed(
  () => extJSON.value.doubanMovie && extJSON.value.doubanMovie.title
);
const hasVideoIframe = computed(
  () =>
    extJSON.value.video &&
    ["bilibili", "youtube"].includes(extJSON.value.video.type) &&
    extJSON.value.video.value
);
const hasVideo = computed(
  () =>
    extJSON.value.video &&
    extJSON.value.video.type === "online" &&
    extJSON.value.video.value
);

const getGridClass = (index: number, count: number) => {
  switch (count) {
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
const contentRef = ref<HTMLElement | null>(null);
onMounted(() => {
  if (contentRef.value) {
    const lineHeight = parseFloat(
      getComputedStyle(contentRef.value).lineHeight
    );
    const maxHeight = lineHeight * 3;
    if (contentRef.value.offsetHeight > maxHeight) {
      let text = props.memo.content;
      while (contentRef.value.offsetHeight > maxHeight && text.length > 0) {
        text = text.slice(0, -1);
        contentRef.value.textContent = text + "...";
      }
    }
  }
});
</script>

<style scoped>
.upload-image-preview img {
  width: 100%;
  height: auto;
}
</style>
