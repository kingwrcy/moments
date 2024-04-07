<template>
  <div class="p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-2">
      <Label for="coverUrl" class="font-bold">顶部图片</Label>
      <Input type="file" id="coverUrl" autocomplete="off" @change="(e: Event) => { uploadImgs(e, 'coverUrl') }" />
      <NuxtImg class="w-full h-[250px]" v-if="state.avatarUrl" :src="state.coverUrl" alt="" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="avatarUrl" class="font-bold">头像</Label>
      <Input type="file" id="avatarUrl" @change="(e: Event) => { uploadImgs(e, 'avatarUrl') }" />
      <NuxtImg :src="state.avatarUrl" alt="avatar" class="w-[70px] h-[70px] rounded-xl" v-if="state.avatarUrl" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="nickname" class="font-bold">昵称</Label>
      <Input type="text" id="nickname" placeholder="头像左边的作者名字" autocomplete="off" v-model="state.nickname" />
    </div>


    <div class="flex flex-col gap-2">
      <Label for="slogan" class="font-bold">心情状态</Label>
      <Input type="text" id="slogan" placeholder="头像下方文字,最好别超过15个汉字" autocomplete="off" v-model="state.slogan" />
    </div>

    <div class="flex flex-col gap-2">
      <Label for="password" class="font-bold">密码</Label>
      <Input type="password" id="password" placeholder="不修改密码不要填写" autocomplete="off"  v-model="state.password"/>
    </div>

    <div class="flex flex-col gap-2 ">
      <Button @click="saveSettings">保存</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner'
import {settingsUpdateEvent} from '~/lib/event'
const token = useCookie('token')


useHead({
  title: '设置-极简朋友圈',
})

const state = reactive({
  coverUrl: '',
  avatarUrl: '',
  nickname: '',
  slogan: '',
  password:'',
})

const { data: res } = await useFetch('/api/user/settings/get')

state.coverUrl = res.value?.data?.coverUrl || 'https://images.unsplash.com/photo-1711299253442-de19d4dacaae?q=80&w=3500&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
state.avatarUrl = res.value?.data?.avatarUrl || 'https://images.kingwrcy.cn/memo/20240386211829TtcOUOMaXyITlTkxhSjp'
state.nickname = res.value?.data?.nickname || 'Jerry'
state.slogan = res.value?.data?.slogan || '星垂平野阔，月涌大江流。'


const uploadImgs = async (event: Event, id: string) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }
  const formData = new FormData()
  formData.append('file', file)
  const res = await $fetch('/api/files/upload', {
    method: 'POST',
    body: formData
  })
  if (res.success) {
    (event.target as HTMLInputElement).value = ''
    if (id === 'coverUrl') {
      state.coverUrl = res.filename
    } else if (id === 'avatarUrl') {
      state.avatarUrl = res.filename
    }
  } else {
    toast.warning(res.message || '上传失败')
  }
}

const saveSettings = async () => {
  const { success } = await $fetch('/api/user/settings/save', {
    method: 'POST',
    body: JSON.stringify({
      coverUrl: state.coverUrl,
      avatarUrl: state.avatarUrl,
      nickname: state.nickname,
      slogan: state.slogan,
      password:state.password
    })
  })
  if (success) {    
    if(state.password){
      token.value = ''
      toast.success('密码修改成功,请重新登录')
      navigateTo('/login')
    }else{
      toast.success('保存成功')
    }
    state.password = ''
    settingsUpdateEvent.emit()
  }
}



</script>

<style scoped></style>