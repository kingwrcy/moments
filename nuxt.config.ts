// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  runtimeConfig: {
    googleRecaptchaSecretKey: "",
    notifyMail: "",
    mailHost: "",
    mailPort: 465,
    mailSecure: true,
    mailName: "",
    mailPassword: "",
    mailFrom: "",
    mailFromName: "",
    aliyunTextJudgeEnable: false,
    aliyunAccessKeyId: "",
    aliyunAccessKeySecret: "",
    public: {
      momentsCommentEnable: true,
      momentsShowComment: true,
      momentsCommentMaxLength: 120,
      momentsCommentOrderBy: "desc",
      momentsToolbarEnableDouban: true,
      momentsToolbarEnableMusic163: true,
      momentsToolbarEnableVideo: true,
      momentsMaxLine: 4,
      googleRecaptchaSiteKey: "",
      siteUrl:"",
    },
  },
  modules: ["@nuxtjs/tailwindcss", "shadcn-nuxt", "@nuxtjs/color-mode"],
  css: ["~/assets/css/index.css"],
  colorMode: {
    classSuffix: "",
  },
  shadcn: {
    /**
     * Prefix for all the imported component
     */
    prefix: "",
    /**
     * Directory that the component lives in.
     * @default "./components/ui"
     */
    componentDir: "./components/ui",
  },
  nitro: {
    esbuild: {
      options: {
        target: "esnext",
      },
    },
  },
});
