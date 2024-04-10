<template>
  <div class="p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-2">
      <Label for="title" class="font-bold">网站标题</Label>
      <Input type="text" id="title" placeholder="网站标题" autocomplete="off" v-model="state.title" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="favicon" class="font-bold">Favicon</Label>
      <Input type="file" id="favicon" autocomplete="off" @change="(e: Event) => { uploadImgs(e, 'favicon') }" />
      <img class="max-w-[50px] max-h-[50px]" v-if="state.favicon" :src="state.favicon" alt="" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="coverUrl" class="font-bold">顶部图片</Label>
      <Input type="file" id="coverUrl" autocomplete="off" @change="(e: Event) => { uploadImgs(e, 'coverUrl') }" />
      <img class="w-full h-[250px]" v-if="state.avatarUrl" :src="state.coverUrl" alt="" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="avatarUrl" class="font-bold">头像</Label>
      <Input type="file" id="avatarUrl" @change="(e: Event) => { uploadImgs(e, 'avatarUrl') }" />
      <img :src="state.avatarUrl" alt="avatar" class="w-[70px] h-[70px] rounded-xl" v-if="state.avatarUrl" />
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
      <Input type="password" id="password" placeholder="不修改密码不要填写" autocomplete="off" v-model="state.password" />
    </div>

    <div class="flex flex-col gap-2">
      <Label for="enableS3" class="font-bold">启用S3存储</Label>
      <Switch id="enableS3" v-model:checked="state.enableS3" />
    </div>

    <template v-if="state.enableS3">
      <div class="flex flex-col gap-2">
        <Label for="domain" class="font-bold">域名</Label>
        <Input type="text" id="domain" placeholder="S3 CDN域名" autocomplete="off" v-model="state.domain" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="bucket" class="font-bold">桶名</Label>
        <Input type="text" id="bucket" placeholder="bucket" autocomplete="off" v-model="state.bucket" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="region" class="font-bold">地区</Label>
        <Input type="text" id="region" placeholder="" autocomplete="off" v-model="state.region" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="accessKey" class="font-bold">accessKey</Label>
        <Input type="text" id="accessKey" placeholder="" autocomplete="off" v-model="state.accessKey" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="secretKey" class="font-bold">secretKey</Label>
        <Input type="text" id="secretKey" placeholder="" autocomplete="off" v-model="state.secretKey" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="endpoint" class="font-bold">S3接口地址</Label>
        <Input type="text" id="endpoint" placeholder="" autocomplete="off" v-model="state.endpoint" />
      </div>

      <div class="flex flex-col gap-2">
        <Label for="thumbnailSuffix" class="font-bold">后缀</Label>
        <Input type="text" id="thumbnailSuffix" placeholder="" autocomplete="off" v-model="state.thumbnailSuffix" />
      </div>


    </template>

    <div class="flex flex-col gap-2 ">
      <Button @click="saveSettings">保存</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner'
import { settingsUpdateEvent } from '~/lib/event'
const token = useCookie('token')
import { useStorage } from "@vueuse/core";
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')

useHead({
  title: '设置-'+(userinfo.value.title || '极简朋友圈'),
})

const enableS3 = useStorage("enableS3", false);


const state = reactive({
  coverUrl: '',
  avatarUrl: '',
  nickname: '',
  slogan: '',
  password: '',
  enableS3: false,
  domain: '',
  bucket: '',
  region: '',
  accessKey: '',
  secretKey: '',
  endpoint: '',
  thumbnailSuffix: '',
  title: '',
  favicon: "",
})

const { data: res } = await useFetch<{ data: typeof state }>('/api/user/settings/full')

state.title = res.value?.data?.title || '极简朋友圈'
state.favicon = res.value?.data?.favicon || '/upload/favicon.png'
state.coverUrl = res.value?.data?.coverUrl || 'https://images.unsplash.com/photo-1711299253442-de19d4dacaae?q=80&w=3500&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
state.avatarUrl = res.value?.data?.avatarUrl || 'https://images.kingwrcy.cn/memo/20240386211829TtcOUOMaXyITlTkxhSjp'
state.nickname = res.value?.data?.nickname || 'Jerry'
state.slogan = res.value?.data?.slogan || '星垂平野阔，月涌大江流。'
state.enableS3 = res.value?.data?.enableS3 || false
state.domain = res.value?.data?.domain || ''
state.bucket = res.value?.data?.bucket || ''
state.region = res.value?.data?.region || ''
state.accessKey = res.value?.data?.accessKey || ''
state.secretKey = res.value?.data?.secretKey || ''
state.endpoint = res.value?.data?.endpoint || ''
state.thumbnailSuffix = res.value?.data?.thumbnailSuffix || ''
enableS3.value = state.enableS3


const uploadImgs = async (event: Event, id: string) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) {
    return
  }

  await useUpload(file, async (res) => {
    if (res.success) {
      (event.target as HTMLInputElement).value = ''
      if (id === 'coverUrl') {
        state.coverUrl = res.filename
      } else if (id === 'avatarUrl') {
        state.avatarUrl = res.filename
      } else if (id === 'favicon') {
        state.favicon = res.filename
      }
    } else {
      toast.warning(res.message || '上传失败')
    }
  })
}

const saveSettings = async () => {
  const { success } = await $fetch('/api/user/settings/save', {
    method: 'POST',
    body: JSON.stringify(state)
  })
  if (success) {
    enableS3.value = state.enableS3
    if (state.password) {
      token.value = ''
      toast.success('密码修改成功,请重新登录')
      navigateTo('/login')
    } else {
      toast.success('保存成功')
      location.reload()
    }
    state.password = ''
    settingsUpdateEvent.emit()
  }
}
</script>

<style scoped></style>