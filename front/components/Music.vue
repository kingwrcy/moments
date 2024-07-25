<template>
  <UPopover :ui="{base:'w-[350px] min-h-[350px]'}">
    <UIcon name="i-carbon-music" class="cursor-pointer"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2">
        <div class="text-xs text-gray-400">嵌入在线音乐</div>
        <UFormGroup label="选择平台" :ui="{label:{base:'font-bold'}}">
          <template #hint>
            <div class="text-xs text-gray-400">
              <ULink class="underline" to="https://github.com/metowolf/MetingJS">MetingJS文档</ULink>
            </div>
          </template>
          <USelectMenu v-model="server" :options="servers" value-attribute="value" option-attribute="label"
                       placeholder="选择平台"></USelectMenu>
        </UFormGroup>

        <UFormGroup label="选择类型" :ui="{label:{base:'font-bold'}}">
          <USelectMenu v-model="type" :options="types" value-attribute="value" option-attribute="label"
                       placeholder="选择平台"></USelectMenu>
        </UFormGroup>

        <UFormGroup label="ID" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="id" placeholder="输入歌曲ID/播放列表ID/专辑ID/搜索关键字"/>
        </UFormGroup>
        <UFormGroup label="API接口地址" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="api" />
        </UFormGroup>
<!--        <MusicPreview :id="id" :server="server" :type="type"/>-->

        <UButtonGroup class="w-fit">
          <UButton @click="close">确定</UButton>
          <UButton color="white" @click="reset(close)">清空</UButton>
        </UButtonGroup>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import type {MetingMusicServer, MetingMusicType} from "@/types"

const id = defineModel<string>('id')
const server = defineModel<MetingMusicServer>('server')
const type = defineModel<MetingMusicType>('type')
const api = defineModel<string>('api')
const reset = (close: Function) => {
  id.value = ""
  server.value = undefined
  type.value = undefined
  api.value = "https://api.i-meto.com/meting/api?server=:server&type=:type&id=:id&r=:r"
  close()
}
const servers = ref([{
  value: "netease",
  label: "网易云音乐",
}, {
  value: "tencent",
  label: "QQ音乐",
}, {
  value: "kugou",
  label: "酷狗音乐",
}, {
  value: "xiami",
  label: "虾米音乐",
}, {
  value: "baidu",
  label: "百度音乐",
},])

const types = ref([{
  value: "song",
  label: "歌曲",
}, {
  value: "playlist",
  label: "播放列表",
}, {
  value: "album",
  label: "专辑",
}, {
  value: "search",
  label: "搜索",
}, {
  value: "artist",
  label: "艺术家",
},])
</script>

<style scoped>

</style>