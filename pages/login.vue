<template>
  <div class="p-2 sm:p-4 flex justify-center min-h-[500px w-full]">
    <div class="p-8 rounded shadow-md max-w-sm w-full">
      <div class="mb-4">
        <Label for="username" class="block text-gray-700 mb-2">用户名</Label>
        <Input v-model="state.username" autocomplete="off" type="text" id="username" />
      </div>
      <div class="mb-6">
        <Label for="password" class="block text-gray-700 mb-2">密码</Label>
        <Input v-model="state.password" autocomplete="off" type="password" id="password" />
      </div>
      <div class="flex flex-row gap-2">
        <Button @click="login" type="button">登录</Button>
        <Button variant="ghost" @click="navigateTo('/')" type="button">返回首页</Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import type { User } from '~/lib/types';

const userinfo = useState<User>('userinfo')

useHead({
  title: '登录-'+(userinfo.value.title || '极简朋友圈'),
})


const state = reactive({
  username: '',
  password: ''
})

const login = async () => {
  const res = await $fetch('/api/user/login', {
    method: 'POST',
    body: JSON.stringify(state)
  })

  if (res.success) {
    toast.success('登录成功',)
    await navigateTo('/')
  } else {
    toast.warning(res.message || '登录失败')
  }
}
</script>

<style scoped></style>