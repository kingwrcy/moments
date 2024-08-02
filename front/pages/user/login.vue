<template>
  <Header v-if="currentUser" v-bind:user="currentUser"/>

 <div class="pb-20">
   <UCard :ui="{base:'max-w-[400px] mx-auto mt-20'}">

     <p class="text-center text-2xl font-sans"> 登录</p>
     <UForm class="space-y-4" size="sm" :state="state">
       <UFormGroup label="用户名" name="email">
         <UInput v-model="state.username"/>
       </UFormGroup>

       <UFormGroup label="密码" name="password">
         <UInput type="password" v-model="state.password"/>
       </UFormGroup>
       <UButtonGroup size="sm">
         <UButton @click="doLogin" :disabled="pending" :loading="pending">登录</UButton>
         <UButton color="gray" v-if="sysConfig.enableRegister" variant="solid" to="/user/reg">去注册</UButton>
       </UButtonGroup>

     </UForm>

   </UCard>
 </div>
</template>

<script setup lang="ts">
import type {LoginResp, SysConfigVO, UserVO} from "~/types";
import {useGlobalState} from "~/store";
import {toast} from "vue-sonner";
const sysConfig = useState<SysConfigVO>('sysConfig')
const currentUser = useState<UserVO>('userinfo')
const global = useGlobalState()
const state = reactive({
  username: "",
  password: ""
})
const pending = ref(false)
const doLogin = async () => {
  pending.value = true
  let success =false
  try{
    global.value.userinfo = await useMyFetch<LoginResp>('/user/login', state)
    toast.success("登录成功,跳转到首页...")
    success = true
  }finally {
    pending.value = false
  }
  if (success){
    location.href='/'
  }

}
</script>

<style scoped>

</style>