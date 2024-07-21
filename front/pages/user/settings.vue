<template>
  <Header :user="currentUser"/>
  <div class="space-y-4  flex flex-col p-4 my-4">
    <UFormGroup label="头像" name="avatarUrl" :ui="{label:{base:'font-bold'}}">
      <UInput type="file" size="sm" icon="i-heroicons-folder" @change="uploadAvatarUrl"/>
      <div class="text-gray-500 text-sm my-2">或者输入在线地址</div>
      <UInput v-model="state.avatarUrl" class="mb-2"/>
      <UAvatar :src="state.avatarUrl" size="lg"/>
    </UFormGroup>
    <UFormGroup label="顶部图片" name="coverUrl" :ui="{label:{base:'font-bold'}}">
      <UInput type="file" size="sm" icon="i-heroicons-folder" @change="uploadCoverUrl"/>
      <div class="text-gray-500 text-sm my-2">或者输入在线地址</div>
      <UInput v-model="state.coverUrl" class="mb-2"/>
      <img :src="state.coverUrl" class="w-full rounded object-cover" alt="" />
    </UFormGroup>
    <UFormGroup label="登录名" name="username" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.username" disabled />
    </UFormGroup>
    <UFormGroup label="昵称" name="nickname" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.nickname"/>
    </UFormGroup>
    <UFormGroup label="心情状态" name="slogan" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.slogan"/>
    </UFormGroup>

    <UButton class="justify-center" @click="save">保存</UButton>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {toast} from "vue-sonner";
import {useUpload} from "~/utils";

const currentUser = useState<UserVO>('userinfo')
const state = reactive({
  username: "",
  nickname: "",
  slogan: "",
  avatarUrl: "",
  coverUrl: "",
  css: "",
  js: "",
})

const reload = async () => {
  const res = await useMyFetch<UserVO>('/user/profile')
  if (res) {
    Object.assign(state, res)
    currentUser.value = res
  }
}

const save = async () => {
  await useMyFetch('/user/saveProfile', state)
  toast.success("保存成功")
  await reload()
}

const uploadAvatarUrl = async (files: FileList) => {
  const result = await useUpload(files)
  toast.success("上传成功")
  if (result) {
    state.avatarUrl = result[0]
  }
}

const uploadCoverUrl = async (files: FileList) => {
  const result = await useUpload(files)
  toast.success("上传成功")
  if (result) {
    state.coverUrl = result[0]
  }
}

onMounted(async () => {
  Object.assign(state,currentUser.value)
})

</script>

<style scoped>

</style>