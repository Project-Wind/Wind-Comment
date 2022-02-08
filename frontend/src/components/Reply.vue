<script setup>
import { articleID } from '../main'
import { inject, ref } from 'vue'
import axios from 'axios'
import md5 from 'md5'

const props = defineProps({
    comment: Object
})

const getComments = inject('getComments')
const getChildComments = inject('getChildComments', undefined)
const toggleReply = inject('toggleReply')

const name = ref('')
const email = ref('')
const content = ref('')
name.value = localStorage.getItem('name')
email.value = localStorage.getItem('email')

function postComment() {
    axios.post('/api/v0/comments', {
        articleID,
        parentID: (() => {
            if (props.comment && props.comment.parentID) {
                return props.comment.parentID
            } else if (props.comment) {
                return props.comment.id
            } else {
                return 0
            }
        })(),
        avatar: `https://www.gravatar.com/avatar/${md5(email.value.toLowerCase())}?d=retro&s=48`,
        name: name.value,
        email: email.value,
        content: props.comment ? `回复 @${props.comment.name}：${content.value}` : content.value
    })
        .then(() => {
            getComments()
            if (getChildComments) {
                getChildComments()
            }

            localStorage.setItem('name', name.value)
            localStorage.setItem('email', email.value)
            content.value = ''
            if (props.comment) {
                toggleReply(props.comment.id)
            }
        })
        .catch(err => console.error(err))
}
</script>

<template>
    <form class="w-full" @submit.prevent="postComment">
        <div class="flex justify-between flex-wrap flex-col sm:flex-row">
            <input required type="text" class="flex-1 rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mb-2 sm:(mr-2 mb-0)" placeholder="称呼" v-model="name" />
            <input required type="email" class="flex-1 rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50" placeholder="邮箱" v-model.trim="email" />
        </div>
        <textarea required class="block resize-none placeholder-gray-500 mt-3 py-2 px-3 w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50" rows="5" :placeholder="comment ? `回复 @${comment.name}：` : '请发表您的见解：'" v-model="content"></textarea>
        <button class="block mt-4 py-3 px-5 border border-gray-800 hover:(bg-gray-800 text-white) font-bold text-sm rounded-lg transition duration-150 ease-in-out">提交</button>
    </form>
</template>
