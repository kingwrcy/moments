<template>
  <Header :user="currentUser" v-if="!isAdminMode"/>

  <div class="space-y-4 flex flex-col p-4 my-4" :class="{ 'pt-0': isAdminMode }">
    <UFormGroup label="登录名" name="username" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.username" disabled />
    </UFormGroup>
    <UFormGroup label="昵称" name="nickname" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.nickname"/>
    </UFormGroup>
    <UFormGroup label="心情状态" name="slogan" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.slogan"/>
    </UFormGroup>
    <UFormGroup class="dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-neutral-700 overflow-hidden">
      <div class="flex items-center justify-between p-2 hover:bg-gray-50 dark:hover:bg-neutral-700/50 transition-colors cursor-pointer" @click="showAvatar = !showAvatar">
        <span class="text-sm font-medium text-gray-900 dark:text-white">头像</span>
        <div class="flex items-center space-x-2">
          <UAvatar :src="state.avatarUrl" size="sm"/>
          <UIcon 
            :name="showAvatar ? 'i-carbon-chevron-down' : 'i-carbon-chevron-right'"
            class="w-5 h-5 text-gray-400 transition-transform duration-200"
            :class="{'rotate-180': showAvatar}"
          />
        </div>
      </div>
      <div v-show="showAvatar" class="px-4 pb-4 space-y-3 border-t border-gray-100 dark:border-neutral-700">
        <div class="space-y-3 pt-3">
          <UInput type="file" size="sm" icon="i-heroicons-photo" @change="uploadAvatarUrl" accept="image/*"/>
          <UInput v-model="state.avatarUrl" placeholder="或输入头像地址" size="sm"/>
        </div>
      </div>
    </UFormGroup>
    <UFormGroup class="dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-neutral-700 overflow-hidden">
      <div class="flex items-center justify-between p-2 hover:bg-gray-50 dark:hover:bg-neutral-700/50 transition-colors cursor-pointer" @click="showCover = !showCover">
        <span class="text-sm font-medium text-gray-900 dark:text-white">顶部图片</span>
        <div class="flex items-center space-x-2">
          <img v-if="state.coverUrl" :src="state.coverUrl" class="w-8 h-6 rounded object-cover" alt=""/>
          <UIcon 
            :name="showCover ? 'i-carbon-chevron-down' : 'i-carbon-chevron-right'"
            class="w-5 h-5 text-gray-400 transition-transform duration-200"
            :class="{'rotate-180': showCover}"
          />
        </div>
      </div>
      <div v-show="showCover" class="px-4 pb-4 space-y-3 border-t border-gray-100 dark:border-neutral-700">
        <div class="space-y-3 pt-3">
          <UInput type="file" size="sm" icon="i-heroicons-photo" @change="uploadCoverUrl" accept="image/*"/>
          <UInput v-model="state.coverUrl" placeholder="或输入图片地址" size="sm"/>
        </div>
        <div v-if="state.coverUrl" class="rounded overflow-hidden">
          <img :src="state.coverUrl" class="w-full h-full object-cover" alt=""/>
        </div>
      </div>
    </UFormGroup>
    <UFormGroup label="密码" name="password" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.password" type="password" placeholder="留空则不修改密码"/>
    </UFormGroup>
    <UFormGroup label="邮箱" name="email" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.email" type="email" placeholder="若管理员启用了邮件通知，将在收到评论时发送邮件通知"/>
    </UFormGroup>
    <UButton class="justify-center" @click="save">保存</UButton>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {toast} from "vue-sonner";
import {useUpload} from "~/utils";
import {useGlobalState} from "~/store";

const props = defineProps<{
  targetUser?: UserVO,
  isAdminMode?: boolean,
  onSave?: () => void
}>()

const global = useGlobalState()
const currentUser = useState<UserVO>('userinfo')
const state = reactive({
  password: "",
  username: "",
  nickname: "",
  slogan: "",
  avatarUrl: "",
  coverUrl: "",
  email: "",
  css: "",
  js: "",
})

const showAvatar = ref(false)
const showCover = ref(false)
const logout = async () => {
  global.value.userinfo = {}
  await navigateTo('/')
}
const reload = async () => {
  const res = await useMyFetch<UserVO>('/user/profile')
  if (res) {
    Object.assign(state, res)
    currentUser.value = res
  }
}

const save = async () => {
  try {
    if (state.password && state.password.length < 6) {
      toast.warning("密码长度至少6位")
      return
    }

    if (props.isAdminMode) {
      // 管理员模式：更新其他用户信息
      await useMyFetch('/user/update', {
        id: props.targetUser!.id,
        username: state.username,
        nickname: state.nickname,
        email: state.email,
        slogan: state.slogan,
        avatarUrl: state.avatarUrl,
        coverUrl: state.coverUrl,
        ...(state.password && { password: state.password })
      })
      toast.success("用户更新成功")
      if (props.onSave) {
        props.onSave()
      }
    } else {
      // 普通用户模式：更新自己的信息
  await useMyFetch('/user/saveProfile', state)
  toast.success("保存成功")
  location.reload()
}
  } catch (error) {
    toast.error(props.isAdminMode ? "用户更新失败" : "保存失败")
  }
}

const uploadAvatarUrl = async (files: FileList) => {
  for (let i = 0; i < files.length; i++) {
    if (files[i].type.indexOf("image") < 0){
      toast.error("只能上传图片");
      return
    }
  }
  const result = await useUpload(files)
  if (result.length) {
    toast.success("上传成功")
    state.avatarUrl = result[0]
  }
}

const uploadCoverUrl = async (files: FileList) => {
  for (let i = 0; i < files.length; i++) {
    if (files[i].type.indexOf("image") < 0){
      toast.error("只能上传图片");
      return
    }
  }
  const result = await useUpload(files)
  if (result.length) {
    toast.success("上传成功")
    state.coverUrl = result[0]
  }
}

onMounted(async () => {
  if (props.isAdminMode && props.targetUser) {
    Object.assign(state, props.targetUser)
    state.password = ""
  } else {
  Object.assign(state,currentUser.value)
  }
})

</script>

<style scoped>

</style>
