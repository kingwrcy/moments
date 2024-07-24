<template>
  <UPopover  :ui="{base:'w-[300px]'}">
    <svg class="focus:outline-0 cursor-pointer w-[18px] h-[18px]" xmlns="http://www.w3.org/2000/svg"
         viewBox="0 0 24 24" fill="currentColor" data-state="closed">
      <path
          d="M15.2735 15H5V7H19V15H17.3764L16.0767 19H21V21H3V19H7.6123L6.8 16.5L8.70211 15.882L9.71522 19H13.9738L15.2735 15ZM3.5 3H20.5V5H3.5V3ZM7 9V13H17V9H7Z"
      ></path>
    </svg>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <URadioGroup
            legend="选择类型"
            v-model="type"
            :options="[{ value: 'book', label: '豆瓣读书' }, { value: 'movie', label: '豆瓣电影' }]"
        />
        <UInput type="text" size="sm" placeholder="请输入豆瓣读书/豆瓣电影的ID" />


        <UButtonGroup>
          <UButton>确定</UButton>
          <UButton color="white">清空</UButton>
        </UButtonGroup>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
const type = ref('book')
</script>

<style scoped>

</style>