<script setup>
import { articleID } from './main'
import { provide, ref } from 'vue'
import Comment from './components/Comment.vue'
import Reply from './components/Reply.vue'
import axios from 'axios'

const comments = ref([])
const getComments = () => axios.get(`/api/v0/comments/${articleID}?parentID=0&sort=id&orderBy=DESC`)
    .then(resp => comments.value = resp.data)
    .catch(err => console.error(err))
getComments()
provide('getComments', getComments)

const visibleReply = ref(0)
const toggleReply = id => visibleReply.value = visibleReply.value === id ? 0 : id

provide('visibleReply', visibleReply)
provide('toggleReply', toggleReply)
</script>

<template>
    <div class="bg-white rounded-xl border border-solid border-[#ddd] shadow shadow-[#ddd] px-8 py-2 divide-y divide-dashed divide-gray-300">
        <Reply class="my-8" />
        <Comment v-for="comment in comments" :key="comment.id" :comment="comment" />
    </div>
</template>
