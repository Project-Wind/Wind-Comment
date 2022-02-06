import { createApp } from 'vue'
import App from './App.vue'
import 'virtual:windi.css'
import 'virtual:windi-devtools'

export const articleID = document.getElementById('comment').dataset.articleId

createApp(App).mount('#comment')
