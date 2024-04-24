declare module "nuxt/schema" {
  interface RuntimeConfig {
    googleRecaptchaSecretKey: string;
  }
  interface PublicRuntimeConfig {
    momentsShowComment: boolean;
    momentsCommentEnable: boolean;
    momentsCommentOrderBy: string;
    momentsCommentMaxLength: number;
    momentsToolbarEnableDouban: boolean;
    momentsToolbarEnableMusic163: boolean;
    momentsToolbarEnableVideo: boolean;
    momentsMaxLine: number;
  }
}
// It is always important to ensure you import/export something when augmenting a type
export {};
