<template>
  <UCard :ui="{base:'w-[400px] mx-auto mt-20'}">

    <p class="text-center text-2xl font-sans">注册用户</p>
    <UForm  class="space-y-4" size="sm">
      <UFormGroup label="用户名" name="email">
        <UInput />
      </UFormGroup>

      <UFormGroup label="密码" name="password">
        <UInput type="current-password" />
      </UFormGroup>
      <UFormGroup label="重复密码" name="password">
        <UInput type="current-password" />
      </UFormGroup>
      <UButtonGroup size="sm">
        <UButton>注册</UButton>
        <UButton color="gray" variant="solid" to="/user/login">去登录</UButton>
      </UButtonGroup>

    </UForm>

  </UCard>
</template>

<script setup lang="ts">

</script>

<style scoped>

</style>