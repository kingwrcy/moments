<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import { useRouter } from 'vue-router'

const router = useRouter()

const message = useMessage()
const formRef = ref<FormInst | null>(null)
const formValue = reactive({
  username: '',
  password: '',
})
const rules = {
  username: {
    required: true,
    message: '请输入姓名',
    trigger: 'blur',
  },
  password: {
    required: true,
    message: '请输入密码',
    trigger: 'blur',
  },
}

async function doLogin(e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate((errors) => {
    if (!errors) {
      message.success('Valid')
    }
    else {
      message.error('必填信息没有填写')
    }
  })
}
</script>

<template>
  <n-card title="登录">
    <div class="p-8 rounded shadow max-w-xs mx-auto">
      <n-form
        ref="formRef"
        :label-width="80"
        :model="formValue"
        :rules="rules"
        size="small"
        label-placement="top"
      >
        <n-form-item label="姓名" path="username">
          <n-input v-model:value="formValue.username" placeholder="输入姓名" />
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="formValue.password"
            type="password" show-password-on="mousedown" placeholder="输入密码"
          />
        </n-form-item>
        <n-form-item>
          <div class="flex justify-center w-full gap-4">
            <n-button type="primary" size="medium" @click="doLogin">
              登录
            </n-button>
            <n-button type="warning" size="medium" @click="router.push('reg')">
              去注册
            </n-button>
          </div>
        </n-form-item>
      </n-form>
    </div>
  </n-card>
</template>

<style scoped></style>
