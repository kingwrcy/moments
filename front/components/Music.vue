<template>
  <UPopover :ui="{base:'w-[350px] min-h-[350px]'}" :popper="{ arrow: true }" mode="click">
    <UIcon name="i-carbon-music" class="cursor-pointer w-6 h-6"/>
    <template #panel="{close}">
      <div class="p-4 flex flex-col gap-2 max-h-[400px] overflow-auto">
        <div class="text-xs text-gray-400">嵌入在线音乐</div>
        <UFormGroup label="选择平台" :ui="{label:{base:'font-bold'}}">
          <template #hint>
            <div class="text-xs text-gray-400">
              <ULink class="underline" target="_blank" to="https://github.com/metowolf/MetingJS">MetingJS文档</ULink>
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
          <UInput v-model="id" placeholder="输入歌曲ID/播放列表ID/专辑ID"/>
        </UFormGroup>
        <UFormGroup label="API接口地址" :ui="{label:{base:'font-bold'}}">
          <UInput v-model="api"/>
        </UFormGroup>
        <MusicPreview v-if="previewing" :id="id" :server="server" :type="type" :api="api"/>

        <div class="flex gap-2">
          <UButton color="indigo" variant="solid" @click="preview(close)" :disabled="previewLoading"
                   :loading="previewLoading">预览
          </UButton>
          <UButton @click="confirm(close)">确定</UButton>
          <UButton color="white" @click="reset(close)">清空</UButton>
        </div>
      </div>
    </template>
  </UPopover>
</template>

<script setup lang="ts">
import type {MetingMusicServer, MetingMusicType, MusicDTO} from "@/types"
import {toast} from "vue-sonner";

const props = withDefaults(defineProps<MusicDTO>(), {
  id: "",
  server: "netease" as MetingMusicServer,
  type: "song" as MetingMusicType,
  api: "https://api.i-meto.com/meting/api?server=:server&type=:type&id=:id&r=:r"
})

const id = ref<string>(props.id)
const server = ref<MetingMusicServer>(props.server)
const type = ref<MetingMusicType>(props.type)
const api = ref<string>(props.api)
const emit = defineEmits(['confirm'])


watch(props, () => {
  id.value = props.id
  server.value = props.server
  type.value = props.type
  api.value = props.api
})

const previewing = ref(false)
const previewLoading = ref(false)
const preview = (close: Function) => {
  if (!server.value || !api.value || !id.value || !type.value) {
    toast.error("请完整填写所需信息")
    return
  }
  previewing.value = false
  previewLoading.value = true
  setTimeout(() => {
    previewing.value = true
    previewLoading.value = false
  }, 500)
}
const confirm = (close: Function) => {
  emit('confirm', {
    id: id.value,
    server: server.value,
    type: type.value,
    api: api.value
  })
  close()
}
const reset = (close: Function) => {
  previewing.value = false
  id.value = ""
  server.value = "netease" as MetingMusicServer
  type.value = "song" as MetingMusicType
  api.value = "https://api.i-meto.com/meting/api?server=:server&type=:type&id=:id&r=:r"
  emit('confirm', {
    id: id.value,
    server: server.value,
    type: type.value,
    api: api.value
  })
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