// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  runtimeConfig: {    
    public: {
      commentEnable:process.env.MOMENTS_COMMENT_ENABLE || "true",
      showComment:process.env.MOMENTS_SHOW_COMMENT || "true",
      commentMaxLength: process.env.MOMENTS_COMMENT_MAX_LENGTH || "120",
      commentOrderBy: process.env.MOMENTS_COMMENT_ORDER_BY || "desc",
      toolbarEnableDouban: process.env.MOMENTS_TOOLBAR_ENABLE_DOUBAN || "true",
      toolbarEnableMusic163: process.env.MOMENTS_TOOLBAR_ENABLE_MUSIC163 || "true",
      toolbarEnableVideo: process.env.MOMENTS_TOOLBAR_ENABLE_VIDEO || "true",
      maxLine: process.env.MOMENTS_MAX_LINE || "4",
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
