<template>
  <UModal
    v-model="open"
    :ui="{
      container:
        'fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center',
    }"
  >
    <div
      v-if="global.userinfo.token"
      class="flex flex-col items-center p-4 pt-8 text-gray-500 dark:text-white"
      @click="navigate('/new')"
      title="发表"
    >
      <span
        class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
      >
        <UIcon name="i-carbon-camera" class="w-7 h-7 cursor-pointer" />
      </span>
      <span class="text-sm mt-1">发表</span>
    </div>
    <div
      class="flex items-center justify-center gap-3 p-4 text-gray-500 dark:text-white min-h-[120px]"
    >
      <div class="flex flex-col items-center gap-1">
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <svg
            v-if="mode.value === 'light'"
            class="lucide lucide-moon-star-icon cursor-pointer"
            @click="toggleMode"
            xmlns="http://www.w3.org/2000/svg"
            width="22"
            height="22"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
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
            width="22"
            height="22"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
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
        </span>
        <span @click="toggleMode" class="text-sm mt-1">
          {{
            mode.preference === "system"
              ? "自动"
              : mode.preference === "light"
              ? "亮色"
              : "暗色"
          }}
        </span>
      </div>
      <div
        v-if="$route.path !== '/user/calendar' && global.userinfo.token"
        class="flex flex-col items-center"
        @click="navigate('/user/calendar')"
        title="日历检索"
      >
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <UIcon name="i-jam-search-folder" class="w-6 h-6 cursor-pointer" />
        </span>
        <span class="text-sm mt-1">检索</span>
      </div>
      <div
        v-if="$route.path == '/'"
        class="flex flex-col items-center"
        @click="navigate('/links')"
        title="友情链接"
      >
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <UIcon name="i-carbon-friendship" class="w-6 h-6 cursor-pointer" />
        </span>
        <span class="text-sm mt-1">友链</span>
      </div>
      <div
        v-if="$route.path !== '/sys/settings' && global.userinfo.id === 1"
        class="flex flex-col items-center"
        @click="navigate('/sys/settings')"
        title="系统设置"
      >
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <UIcon name="i-carbon-settings" class="w-6 h-6 cursor-pointer" />
        </span>
        <span class="text-sm mt-1">系统</span>
      </div>
      <div
        v-if="$route.path !== '/user/settings' && global.userinfo.token"
        class="flex flex-col items-center"
        @click="navigate('/user/settings')"
        title="用户中心"
      >
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <UIcon name="i-carbon-user-avatar" class="w-6 h-6 cursor-pointer" />
        </span>
        <span class="text-sm mt-1">用户</span>
      </div>
      <div
        v-if="$route.path === '/user/settings' && global.userinfo.token"
        class="flex flex-col items-center"
        title="登出"
        @click="logout"
      >
        <span
          class="flex items-center bg-gray-200/75 dark:bg-gray-800/75 p-3 rounded-full"
        >
          <UIcon name="i-carbon-logout" class="w-6 h-6 cursor-pointer" />
        </span>
        <span class="text-sm mt-1">退出</span>
      </div>
    </div>
  </UModal>
</template>

<script setup lang="ts">
import { toast } from "vue-sonner";
import { useGlobalState } from "~/store";
import type { SysConfigVO } from "~/types";

const sysConfig = useState<SysConfigVO>("sysConfig");
const global = useGlobalState();
const mode = useColorMode();
const open = useState<boolean>("sidebarOpen", () => false);

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

const navigate = async (url: string) => {
  open.value = false;
  await navigateTo(url);
};

const logout = async () => {
  global.value.userinfo = {};
  await navigateTo("/");
};
</script>

<style scoped></style>
