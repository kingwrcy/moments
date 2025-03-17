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
        <span v-else-if="$route.path.indexOf('/tags/') >= 0">话题专栏</span>
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
        title="登出"
        @click="logout"
      >
        <UIcon name="i-carbon-logout" class="w-5 h-5 cursor-pointer" />
      </NuxtLink>
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
      <UIcon
        v-if="$route.path == '/' && sysConfig.friendLinks"
        name="i-carbon-friendship"
        title="友情链接"
        class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        @click="showfriendLinks = true"
      />
      <NuxtLink v-if="!global.userinfo.token" to="/user/login" title="登录">
        <UIcon
          name="i-carbon-login"
          class="text-[#9fc84a] w-5 h-5 cursor-pointer"
        />
      </NuxtLink>
    </div>

    <img class="header-img w-full" :src="props.user.coverUrl" alt="" />
    <UIcon
        v-if="$route.path == '/' && sysConfig.friendLinks"
        name="i-carbon-friendship"
        title="友情链接"
        class="sm:hidden absolute right-3 top-3 text-[#9fc84a] dark:text-white w-5 h-5 cursor-pointer sm:w-[567px]"
        @click="showfriendLinks = true"
      />
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

    <template>
      <UModal v-model="showfriendLinks" :ui="{ container: 'sm:items-start fixed top-5 left-0 right-3 flex justify-center items-start' }">
        <div
          class="absolute top-2 right-2 cursor-pointer"
          @click="showfriendLinks = false"
        >
          <UIcon name="i-carbon-close" class="text-[#9fc84a] w-5 h-5" />
        </div>
        <div class="flex flex-col gap-2 p-4 text-gray-500 dark:text-white">
          <h3 class="flex items-center text font-bold mb-2">联系人</h3>
          <div class="grid grid-cols-2 sm:grid-cols-3 gap-4">
            <div
              v-for="link in friendLinkList"
              :key="link.url"
              class="flex items-center"
            >
              <a
                :href="link.url"
                target="_blank"
                class="flex items-center gap-2"
              >
                <img
                  :src="link.icon"
                  alt="Friend Avatar"
                  class="w-6 h-6 rounded-full"
                />
                <span>{{ link.name }}</span>
              </a>
            </div>
          </div>
          <div class="flex justify-center item-center mt-2 text-sm text-gray-400">共 {{ friendLinkList.length }} 个朋友</div>
        </div>
      </UModal>
    </template>
  </div>
</template>
<script setup lang="ts">
import { toast } from "vue-sonner";
import type { SysConfigVO, UserVO } from "~/types";
import { useGlobalState } from "~/store";

const global = useGlobalState();
const sysConfig = useState<SysConfigVO>("sysConfig");

const props = defineProps<{ user: UserVO }>();
const mode = useColorMode();
const { y } = useWindowScroll();

const showfriendLinks = ref(false);

const logout = async () => {
  global.value.userinfo = {};
  await navigateTo("/");
};

const friendLinkList = computed(() => {
  if (!sysConfig.value.friendLinks) {
    return [];
  }
  const lines = sysConfig.value.friendLinks.split("\n");
  return lines.map((line) => {
    const [name, url, icon] = line.split("|");
    return {
      name,
      url,
      icon,
    };
  });
});

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
</script>

<style scoped></style>
