<template>
  <div class="p-2 rounded text-sm">
    <Textarea rows="3" class="border bg-white" :placeholder="placeholder"</Textarea>
    <div class="flex flex-row items-center justify-end mt-2 gap-2 ">
      <Input placeholder="昵称"  v-model:value="username" class="bg-white"></Input>
      <Input placeholder="主页"  v-model:value="link" class="bg-white"> </Input>
      <Input placeholder="邮箱"  v-model:value="email" class="bg-white"></Input>
      <Button size="sm" @click="saveComment">发表评论</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner';
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'

const content = ref('')
const username = ref('')
const email = ref('')
const link = ref('')
const placeholder = ref('发表评论')

const props = defineProps<{ memoId: number, reply?: string }>()

const saveComment = async () => {

  if (!content.value) {
    toast.warning('先填写评论')
    return
  }
  const res = await $fetch('/api/comment/save', {
    method: 'POST',
    body: JSON.stringify({
      content: content.value,
      memoId: props.memoId,
      username: username.value || '匿名',
      link: link.value,
      email: email.value,
    })
  })

  if (res.success) {
    toast.success('评论成功')
  } else {
    toast.warning('评论失败')
  }
}

onMounted(async () => {

  if (props.reply) {
    placeholder.value = "回复给@" + props.reply
  }
})
</script>

<style scoped></style>
