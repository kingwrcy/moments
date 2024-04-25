<template>
  <div class="flex flex-col gap-4 p-2 sm:p-4">
    <div class="flex flex-col gap-2">
      <Label for="version" class="text-right text-gray-400 text-xs my-2" v-if="versionData">版本号:v{{ versionData.version
        }}</Label>

    </div>
    <div class="flex flex-col gap-2">
      <Label for="username" class="font-bold">管理员账号</Label>
      <Input type="text" id="username" placeholder="管理员账号" autocomplete="off" v-model="state.username" />
    </div>
    <div class="flex flex-col gap-2">
      <Label for="title" class="font-bold">网站标题</Label>
      <Input type="text" id="title" placeholder="网站标题" autocomplete="off" v-model="state.title" />
    </div>
    <div class="flex flex-col gap-2 border rounded p-2">
      <Label for="avatarUrl" class="font-bold">头像</Label>
      <Input type="file" id="avatarUrl" @change="(e: Event) => { uploadImgs(e, 'avatarUrl') }" />
      <Label for="avatarUrl-input" class="font-medium">或者输入在线地址</Label>
      <Input type="text" id="avatarUrl-input" placeholder="或者填入在线地址" autocomplete="off" v-model="state.avatarUrl" />
      <img :src="state.avatarUrl" alt="avatar" class="w-[70px] h-[70px] rounded-xl" v-if="state.avatarUrl" />
    </div>
    <div class="flex flex-col gap-2 border rounded p-2">
      <Label for="favicon" class="font-bold">Favicon</Label>
      <Input type="file" id="favicon" autocomplete="off" @change="(e: Event) => { uploadImgs(e, 'favicon') }" />
      <Label for="favicon-input" class="font-medium">或者输入在线地址</Label>
      <Input type="text" id="favicon-input" placeholder="或者填入在线地址" autocomplete="off" v-model="state.favicon" />
      <img class="max-w-[50px] max-h-[50px]" v-if="state.favicon" :src="state.favicon" alt="" />
    </div>
    <div class="flex flex-col gap-2 border rounded p-2">
      <Label for="coverUrl" class="font-bold">顶部图片</Label>
      <Input type="file" id="coverUrl" autocomplete="off" @change="(e: Event) => { uploadImgs(e, 'coverUrl') }" />
      <Label for="coverUrl-input" class="font-medium">或者输入在线地址</Label>
      <Input type="text" id="coverUrl-input" placeholder="或者填入在线地址" autocomplete="off" v-model="state.coverUrl" />
      <img class="w-full " v-if="state.avatarUrl" :src="state.coverUrl" alt="" />
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
      <Label for="css" class="font-bold">自定义CSS</Label>
      <Textarea id="css" v-model="state.css" rows="3"></Textarea>
    </div>

    <div class="flex flex-col gap-2">
      <Label for="js" class="font-bold">自定义JS</Label>
      <Textarea id="js" v-model="state.js" rows="3"></Textarea>
    </div>

    <div class="flex flex-col gap-2">
      <Label for="beianNo" class="font-bold">备案号</Label>
      <Input type="text" id="beianNo" placeholder="备案号,没有可不填写" autocomplete="off" v-model="state.beianNo" />
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
    <Collapsible >
      <CollapsibleTrigger>
        <div class="cursor-pointer font-bold text-sm flex justify-between items-center"><div>基础设置</div> <ChevronsUpDown :size=16 /></div>
      </CollapsibleTrigger>
      <CollapsibleContent>
        <table class="w-full border my-2 text-xs" >
          <thead>
            <tr class="*:border *:p-2 *:bg-gray-300">
              <td>配置项(去掉了前缀`NUXT_PUBLIC`)</td>
              <td>内容</td>
            </tr>
          </thead>
          <tbody>
            <tr class="*:border *:p-2">
              <td>MOMENTS_COMMENT_ENABLE(是否可以评论)</td>
              <td>{{ $config.public.momentsCommentEnable }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_SHOW_COMMENT(是否展示评论)</td>
              <td>{{ $config.public.momentsShowComment }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_COMMENT_MAX_LENGTH(评论最大字数)</td>
              <td>{{ $config.public.momentsCommentMaxLength }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_COMMENT_ORDER_BY(评论排序)</td>
              <td>{{ $config.public.momentsCommentOrderBy }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_TOOLBAR_ENABLE_DOUBAN(是否显示豆瓣按钮)</td>
              <td>{{ $config.public.momentsToolbarEnableDouban }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_TOOLBAR_ENABLE_MUSIC163(是否显示音乐按钮)</td>
              <td>{{ $config.public.momentsToolbarEnableMusic163 }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_TOOLBAR_ENABLE_VIDEO(是否显示视频按钮)</td>
              <td>{{ $config.public.momentsToolbarEnableVideo }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>MOMENTS_MAX_LINE(发言最大行数,超过折叠)</td>
              <td>{{ $config.public.momentsMaxLine }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>GOOGLE_RECAPTCHA_SITE_KEY(RECAPTCHA网站key)</td>
              <td>{{ $config.public.googleRecaptchaSiteKey }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>SITE_URL(本站访问地址站点地址)</td>
              <td>{{ $config.public.siteUrl }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>NOTIFY_BY_EMAIL_ENABLE(是否开启评论邮件通知)</td>
              <td>{{ $config.public.notifyByEmailEnable }}</td>
            </tr>
            <tr class="*:border *:p-2">
              <td>ALIYUN_TEXT_JUDGE_ENABLE(是否开启评论内容阿里云鉴权)</td>
              <td>{{ $config.public.aliyunTextJudgeEnable }}</td>
            </tr>
          </tbody>
        </table>
      </CollapsibleContent>
      
    </Collapsible>

    
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
import {ChevronsUpDown} from 'lucide-vue-next'

const { data: versionData } = await useAsyncData('version', async () => $fetch('/api/version'))

const userinfo = useState<User>('userinfo')

useHead({
  title: '设置-' + (userinfo.value.title || '极简朋友圈'),
})

const enableS3 = useStorage("enableS3", false);


const state = reactive({
  username: '',
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
  css: "",
  js: "",
  beianNo: ""
})

const { data: res } = await useFetch<{ data: typeof state }>('/api/user/settings/full', { key: 'settings' })
const data = res.value?.data
state.username = data?.username || 'admin'
state.title = data?.title || '极简朋友圈'
state.favicon = data?.favicon || '/avatar.webp'
state.coverUrl = data?.coverUrl || ''
state.avatarUrl = data?.avatarUrl || '/cover.webp'
state.nickname = data?.nickname || 'Jerry'
state.slogan = data?.slogan || '星垂平野阔，月涌大江流。'
state.enableS3 = data?.enableS3 || false
state.domain = data?.domain || ''
state.bucket = data?.bucket || ''
state.region = data?.region || ''
state.accessKey = data?.accessKey || ''
state.secretKey = data?.secretKey || ''
state.endpoint = data?.endpoint || ''
state.thumbnailSuffix = data?.thumbnailSuffix || ''
state.css = data?.css || ''
state.js = data?.js || ''
state.beianNo = data?.beianNo || ''
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