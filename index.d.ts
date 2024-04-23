declare module 'nuxt/schema' {
  interface RuntimeConfig {
  }
  interface PublicRuntimeConfig {
    showComment:string
    commentEnable:string
    commentOrderBy:string
    commentMaxLength:string
    toolbarEnableDouban:string
    toolbarEnableMusic163:string
    toolbarEnableVideo:string
  }
}
// It is always important to ensure you import/export something when augmenting a type
export {}