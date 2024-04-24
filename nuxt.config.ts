// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  runtimeConfig: {    
    recaptchaV3secretKey: process.env.NUXT_GOOGLE_RECAPTCHA_SECRET_KEY || "",
    public: {
      commentEnable:process.env.NUXT_MOMENTS_COMMENT_ENABLE || "true",
      showComment:process.env.NUXT_MOMENTS_SHOW_COMMENT || "true",
      commentMaxLength: process.env.NUXT_MOMENTS_COMMENT_MAX_LENGTH || "120",
      commentOrderBy: process.env.NUXT_MOMENTS_COMMENT_ORDER_BY || "desc",
      toolbarEnableDouban: process.env.NUXT_MOMENTS_TOOLBAR_ENABLE_DOUBAN || "true",
      toolbarEnableMusic163: process.env.NUXT_MOMENTS_TOOLBAR_ENABLE_MUSIC163 || "true",
      toolbarEnableVideo: process.env.NUXT_MOMENTS_TOOLBAR_ENABLE_VIDEO || "true",
      maxLine: process.env.NUXT_MOMENTS_MAX_LINE || "4",
      recaptchaV3SiteKey: process.env.NUXT_GOOGLE_RECAPTCHA_SITE_KEY || "",
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
