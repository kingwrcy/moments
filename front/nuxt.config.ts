// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ["@nuxt/ui",'@vueuse/nuxt','dayjs-nuxt'],
  ssr:false,
  dayjs:{
    locales:['zh'],
    defaultLocale:'zh'
  },
  ui: {
    icons: ['carbon']
  },
  tailwindcss:{

  },
  vite:{
    server: {
      proxy: {
        "/api": {
          target: "http://localhost:37892",
          changeOrigin: true,
          // rewrite: (path) => path.replace(/^\/api/, ""),
        },
      },
    },
  }
})