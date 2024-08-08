<template>
  <Header v-if="currentUser" v-bind:user="currentUser"/>

  <div class="pb-20">
    <UCard :ui="{base:'w-4/5 mx-auto mt-20'}">

      <p class="text-center text-2xl font-sans">注册用户</p>
      <UForm class="space-y-4" size="sm" :state="state">
        <UFormGroup label="用户名" name="email">
          <UInput v-model="state.username"/>
        </UFormGroup>

        <UFormGroup label="密码" name="password">
          <UInput type="password" v-model="state.password"/>
        </UFormGroup>
        <UFormGroup label="重复密码" name="repeatPassword">
          <UInput type="password" v-model="state.repeatPassword"/>
        </UFormGroup>
        <UButtonGroup size="sm">
          <UButton @click="doReg" :disabled="pending" :loading="pending">注册</UButton>
          <UButton color="gray" variant="solid" to="/user/login">去登录</UButton>
        </UButtonGroup>
      </UForm>
    </UCard>
  </div>
</template>

<script setup lang="ts">
import type {UserVO} from "~/types";
import {toast} from "vue-sonner";

const state = reactive({
  username: "",
  password: "",
  repeatPassword: ""
})
const pending = ref(false)
const currentUser = useState<UserVO>('userinfo')
const doReg = async () => {
  if (state.username.length < 3) {
    toast.warning("用户名最少3个字符")
    return
  }
  let success = false
  pending.value = true
  try {
    await useMyFetch('/user/reg', state)
    toast.success("注册成功,快去登录吧!")
    success = true
  } finally {
    pending.value = false
  }
  if (success) {
    await navigateTo('/user/login')
  }
}
</script>

<style scoped>

</style>