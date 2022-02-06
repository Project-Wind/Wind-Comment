import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import WindiCSS from 'vite-plugin-windicss'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        WindiCSS()
    ],
    server: {
        proxy: {
            '/api': 'http://127.0.0.1:8080'
        }
    },
    build: {
        lib: {
            entry: './src/main.js',
            name: 'WindComment'
        }
    }
})
