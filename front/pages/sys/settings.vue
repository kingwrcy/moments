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
    <UFormGroup label="是否启用评论" name="enableComment" :ui="{label:{base:'font-bold'}}">
      <UToggle v-model="state.enableComment"/>
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
    <UFormGroup label="评论最大字数" name="maxCommentLength" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.maxCommentLength"/>
    </UFormGroup>
    <UFormGroup label="发言最大高度(单位px)" name="memoMaxHeight" :ui="{label:{base:'font-bold'}}">
      <UInput v-model="state.memoMaxHeight"/>
    </UFormGroup>
    <UFormGroup label="评论排序方式(按日期)" name="commentOrder" :ui="{label:{base:'font-bold'}}">
      <USelectMenu v-model="state.commentOrder"
                   :options="[{label:'倒序,越晚发布越靠前',value:'desc'},{label:'正序,越早发布越靠前',value:'asc'}]"
                   value-attribute="value" option-attribute="label"></USelectMenu>
    </UFormGroup>
    <UFormGroup label="日期格式" name="timeFormat" :ui="{label:{base:'font-bold'}}">
      <USelectMenu v-model="state.timeFormat"
                   :options="[{label:'几分钟前',value:'timeAgo'},{label:'2024-07-24 09:56:55',value:'time'}]"
                   value-attribute="value" option-attribute="label"></USelectMenu>
    </UFormGroup>
    <div class="border rounded border-gray-300 p-2 space-y-4  flex flex-col">
      <UFormGroup label="是否启用Google Recaptcha" name="enableGoogleRecaptcha" :ui="{label:{base:'font-bold'}}">
        <UToggle v-model="state.enableGoogleRecaptcha"/>
      </UFormGroup>
      <template v-if="state.enableGoogleRecaptcha">
        <UFormGroup label="SiteKey" name="googleSiteKey" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.googleSiteKey"/>
        </UFormGroup>
        <UFormGroup label="SecretKey" name="googleSecretKey" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="state.googleSecretKey"/>
        </UFormGroup>
      </template>
    </div>
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
  enableGoogleRecaptcha: false,
  googleSiteKey:"",
  googleSecretKey:"",
  enableComment: true,
  maxCommentLength: 120,
  memoMaxHeight: 300,
  commentOrder: 'desc',
  timeFormat: 'timeAgo',
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