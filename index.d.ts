declare module "nuxt/schema" {
  interface RuntimeConfig {
    googleRecaptchaSecretKey: string;
    notifyMail: string;
    mailHost: string;
    mailPort: number;
    mailSecure: boolean;
    mailName: string;
    mailPassword: string;
    mailFrom: string;
    mailFromName: string;
    aliyunTextJudgeEnable: boolean;
    aliyunAccessKeyId:string;
    aliyunAccessKeySecret: string;
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
    siteUrl:string;
  }
}
// It is always important to ensure you import/export something when augmenting a type
export {};
