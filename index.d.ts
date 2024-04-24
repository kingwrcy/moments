declare module 'nuxt/schema' {
  interface RuntimeConfig {
  }
  interface PublicRuntimeConfig {
    momentsShowComment:string
    momentsCommentEnable:string
    momentsCommentOrderBy:string
    momentsCommentMaxLength:string
    momentsToolbarEnableDouban:string
    momentsToolbarEnableMusic163:string
    momentsToolbarEnableVideo:string
    momentsMaxLine:string
  }
}
// It is always important to ensure you import/export something when augmenting a type
export {}