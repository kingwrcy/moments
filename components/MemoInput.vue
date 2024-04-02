<template>
  <div class="p-4 border-b pb-2">
    <div class="flex flex-row gap-2 my-2">
      <img src="~/assets/img/uploadImg.svg" alt="上传图片,最多9张" title="上传图片,最多9张" class="cursor-pointer w-[20px] h-[20px]">
    </div>
    <Textarea v-model="context" rows="5"></Textarea>
    <div class="flex flex-row justify-end mt-2">
      <Button @click="submitMemo">提交</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'


const emit = defineEmits(['memoAdded'])

const context = ref('')
const submitMemo = async () => {
  const res = await $fetch('/api/memo/save', {
    method: 'POST',
    body: JSON.stringify({
      content: context.value
    })
  })
  if (res.success) {
    toast.success('提交成功')
    context.value = ''
    emit('memoAdded')
  } else {
    toast.warning('提交失败')
  }
}
</script>

<style scoped></style>