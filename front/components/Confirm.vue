<template>
  <UPopover v-model:open="open">
      <slot  @click="open = true"/>
    <template #panel>
     <div class="px-4 py-2">
       <p class="text-gray-600 my-2">你确定要删除吗?</p>
       <div class="flex gap-2 items-center">
         <UButton @click="ok()">确定</UButton>
         <UButton color="white" @click="cancel()">取消</UButton>
       </div>
     </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
const open = ref(false)
const emit = defineEmits(['ok','cancel'])
const ok = () => {
  open.value = false
  emit('ok')
}
const cancel = () => {
  emit('cancel')
  open.value = false
}
</script>

<style scoped>

</style>