<template>
  <Header :user="currentUser"/>
  <div class="space-y-4  flex flex-col p-4 my-4">
    <div class="flex justify-end text-xs text-gray-400">
      <div>版本号:v0.3</div>
    </div>
    <UFormGroup label="管理员账号" name="adminUserName" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.adminUserName"/>
    </UFormGroup>
    <UFormGroup label="网站标题" name="title" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.title"/>
    </UFormGroup>
    <UFormGroup label="Favicon" name="favicon"
                :ui="{wrapper:'border rounded p-2 border-gray-300',label:{base:'font-bold'}}">
      <UInput type="file" size="sm" icon="i-heroicons-folder" @change="uploadFavicon"/>
      <div class="text-gray-500 text-sm my-2">或者输入在线地址</div>
      <UInput v-model="state.favicon" class="mb-2"/>
      <UAvatar :src="state.favicon"/>
    </UFormGroup>
    <UFormGroup label="备案号" name="beiAnNo" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.beiAnNo" placeholder="没有可以不填写"/>
    </UFormGroup>
    <UFormGroup label="自定义CSS" name="css" :ui="{label:{base:'font-bold'}}">
      <UTextarea v-model="state.css" :rows="5"/>
    </UFormGroup>
    <UFormGroup label="自定义JS" name="js" :ui="{label:{base:'font-bold'}}">
      <UTextarea v-model="state.js" :rows="5"/>
    </UFormGroup>
    <div class="border rounded border-gray-300 p-2 space-y-4  flex flex-col">
      <UFormGroup label="是否启用S3存储" name="s3" :ui="{label:{base:'font-bold'}}">
        <UToggle v-model="state.enableS3"/>
      </UFormGroup>
      <template v-if="state.enableS3">
        <UFormGroup label="域名" name="domain" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.domain"/>
        </UFormGroup>
        <UFormGroup label="桶名" name="bucket" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.bucket"/>
        </UFormGroup>
        <UFormGroup label="地区" name="region" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.region"/>
        </UFormGroup>
        <UFormGroup label="AccessKey" name="accessKey" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.accessKey"/>
        </UFormGroup>
        <UFormGroup label="SecretKey" name="secretKey" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.secretKey"/>
        </UFormGroup>
        <UFormGroup label="S3 API接口地址" name="endpoint" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.endpoint"/>
        </UFormGroup>
        <UFormGroup label="图片后缀" name="thumbnailSuffix" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.s3.thumbnailSuffix"/>
        </UFormGroup>
      </template>
    </div>

    <UButton class="justify-center" @click="save">保存</UButton>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {toast} from "vue-sonner";
import {useUpload} from "~/utils";

const currentUser = useState<UserVO>('userinfo')
const state = reactive({
  adminUserName: "admin",
  title: "极简朋友圈",
  favicon: "/favicon.ico",
  beiAnNo: "",
  css: "",
  js: "",
  enableS3: false,
  s3: {
    domain: "",
    bucket: "",
    region: "",
    accessKey: "",
    secretKey: "",
    endpoint: "",
    thumbnailSuffix: ""
  }
})

const reload = async () => {
  const res = await useMyFetch('/sysConfig/getFull')
  if (res) {
    Object.assign(state, res)
  }
}

const save = async () => {
  await useMyFetch('/sysConfig/save', state)
  toast.success("保存成功")
  location.reload()
}

const uploadFavicon = async (files: FileList) => {
  console.log(files)
  const result = await useUpload(files)
  toast.success("上传成功")
  if (result) {
    state.favicon = result[0]
  }
}

onMounted(async () => {
  await reload()
})

</script>

<style scoped>

</style>