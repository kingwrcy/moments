<template>
  <div
    v-if="$route.path !== '/new' && $route.path.indexOf('/edit/') < 0"
    class="header relative mb-14"
  >
    <div
      v-if="$route.path !== '/' && $route.path.indexOf('/memo/') < 0"
      :class="{ 'bg-[#4c4c4c]/80 z-10': y > 100 }"
      class="flex fixed justify-between items-center p-4 w-full md:w-[567px] text-white top-0"
    >
      <NuxtLink class="flex items-center" title="返回主页">
        <UIcon
          @click="navigateTo('/')"
          name="i-carbon-chevron-left"
          class="w-5 h-5 cursor-pointer mr-4"
        />
        <span v-if="$route.path === '/user/calendar'">日历检索</span>
        <span v-else-if="$route.path === '/sys/settings'">系统设置</span>
        <span v-else-if="$route.path === '/user/settings'">用户中心</span>
        <span v-else-if="$route.path.indexOf('/tags/') >= 0">
          {{ getTagFromRoute() || "话题专栏" }}
        </span>
        <span v-else-if="$route.path === '/links'">友情链接</span>
        <span v-else>
          <span v-if="!global.userinfo.token && $route.path === '/user/login'">
            登录
          </span>
          <span
            v-else-if="!global.userinfo.token && $route.path === '/user/reg'"
          >
            注册
          </span>
          <span v-else>{{ props.user.nickname }} 的空间</span>
        </span>
      </NuxtLink>
      <NuxtLink
        v-if="$route.path === '/user/settings' && global.userinfo.token"
        class="hidden sm:flex"
        title="登出"
        @click="logout"
      >
        <UIcon name="i-carbon-logout" class="w-5 h-5 cursor-pointer" />
      </NuxtLink>
      <span
        v-if="$route.path == '/links' && global.userinfo.id === 1"
        class="flex"
      >
        <UIcon
          name="i-carbon-add"
          class="w-6 h-6 cursor-pointer"
          @click="$emit('add-links')"
        />
      </span>
    </div>

    <div
      class="dark:bg-neutral-800 hidden sm:flex sm:absolute sm:-right-10 sm:rounded sm:p-2 sm:flex-col sm:w-fit justify-end shadow w-full flex-row top-0 p-1 flex gap-2 bg-white"
    >
      <svg
        v-if="mode.value === 'light'"
        class="lucide lucide-moon-star-icon cursor-pointer"
        @click="toggleMode"
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="#FDE047"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9"></path>
        <path d="M20 3v4"></path>
        <path d="M22 5h-4"></path>
      </svg>

      <svg
        v-else
        class="lucide lucide-sun-icon cursor-pointer"
        @click="toggleMode"
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="#FDE047"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <circle cx="12" cy="12" r="4"></circle>
        <path d="M12 2v2"></path>
        <path d="M12 20v2"></path>
        <path d="m4.93 4.93 1.41 1.41"></path>
        <path d="m17.66 17.66 1.41 1.41"></path>
        <path d="M2 12h2"></path>
        <path d="M20 12h2"></path>
        <path d="m6.34 17.66-1.41 1.41"></path>
        <path d="m19.07 4.93-1.41 1.41"></path>
      </svg>

      <NuxtLink v-if="global.userinfo.token" to="/new" title="发表">
        <UIcon
          name="i-carbon-camera"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
      <NuxtLink
        v-if="$route.path !== '/user/calendar' && global.userinfo.token"
        to="/user/calendar"
        title="日历检索"
      >
        <UIcon
          name="i-jam-search-folder"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
      <NuxtLink v-if="$route.path == '/'" to="/links" title="友情链接">
        <UIcon
          name="i-carbon-friendship"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
      <NuxtLink
        v-if="$route.path !== '/sys/settings' && global.userinfo.id === 1"
        to="/sys/settings"
        title="系统设置"
      >
        <UIcon
          name="i-carbon-settings"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
      <NuxtLink
        v-if="$route.path !== '/user/settings' && global.userinfo.token"
        to="/user/settings"
        title="用户中心"
      >
        <UIcon
          name="i-carbon-user-avatar"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
      <NuxtLink v-if="!global.userinfo.token" to="/user/login" title="登录">
        <UIcon
          name="i-carbon-login"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
    </div>

    <img class="header-img w-full" :src="props.user.coverUrl" alt="" />
    <div class="absolute right-2 bottom-[-40px]">
      <div class="userinfo flex flex-col">
        <div class="flex flex-row items-center gap-4 justify-end">
          <div class="username text-lg font-bold text-white">
            {{ props.user.nickname }}
          </div>
          <img
            :src="props.user.avatarUrl"
            class="avatar w-[70px] h-[70px] rounded-xl"
          />
        </div>
        <div class="slogon text-gray truncate w-full text-end text-xs mt-2">
          {{ props.user.slogan }}
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { toast } from "vue-sonner";
import type { UserVO } from "~/types";
import { useGlobalState } from "~/store";

const global = useGlobalState();
const route = useRoute();

const props = defineProps<{ user: UserVO }>();
const mode = useColorMode();
const { y } = useWindowScroll();

const logout = async () => {
  global.value.userinfo = {};
  await navigateTo("/");
};

const toggleMode = () => {
  if (mode.preference === "system") {
    mode.preference = "dark";
  } else if (mode.preference === "dark") {
    mode.preference = "light";
  } else {
    mode.preference = "system";
    toast.success("显示模式将跟随系统设置");
  }
};

const getTagFromRoute = () => {
  if (route.path.indexOf("/tags/") >= 0) {
    const parts = route.path.split("/");
    const tag = parts[parts.length - 1];
    try {
      return decodeURIComponent(tag);
    } catch (error) {
      console.error("标签名称解码出错:", error);
      return tag;
    }
  }
  return null;
};
</script>

<style scoped></style>
