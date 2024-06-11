// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  modules: ["@nuxtjs/tailwindcss", "shadcn-nuxt", "@nuxtjs/color-mode"],
  colorMode: {
    classSuffix: "",
    storageKey:'vueuse-color-scheme'
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
  app: {
    head: {
      link: [
        { href: `/css/APlayer.min.css`, rel: 'stylesheet' },
      ],
      script: [
        { src: `/js/APlayer.min.js`, type: 'text/javascript', async: true, defer: true },
        { src: `/js/Meting.min.js`, type: 'text/javascript', async: true, defer: true },
      ]
    }
  },
  plugins: [
    '~/plugins/meting.ts'
  ],
  vue: {  
    compilerOptions: {
      isCustomElement: (tag:string) => ['meting-js'].includes(tag),
    },
  }
});
