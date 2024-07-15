import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'

import UnoCSS from 'unocss/vite'
import {fileURLToPath} from "node:url";

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    server: {
        proxy: {
            '^/api': {
                target: 'http://127.0.0.1:37892/',
                changeOrigin: true,
                // rewrite: (path) => path.replace(/^\/api/, '')

            },
        },
    },
    plugins: [
        vue(),
        AutoImport({
            dts: true,
            imports: [
                'vue',

                {
                    '@vueuse/core':[
                        'useStorage'
                    ],
                    'vue-sonner':[
                        'toast',
                        'Toaster'
                    ],
                    'naive-ui': [
                        'useDialog',
                        'useMessage',
                        'useNotification',
                        'useLoadingBar',
                    ],
                },
            ],
        }),
        Components({
            resolvers: [NaiveUiResolver()],
        }),
        UnoCSS(),
    ],
})
