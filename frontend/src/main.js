import 'virtual:windi-devtools'
import 'virtual:windi.css'
import { createApp } from 'vue'
import App from './App.vue'

export const articleID = document.getElementById('comment').dataset.articleId
export const baseURL = document.getElementById('comment').dataset.baseUrl

createApp(App).mount('#comment')
