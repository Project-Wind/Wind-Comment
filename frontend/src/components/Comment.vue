<script setup>
import axios from 'axios'
import { provide, ref } from 'vue'
import { baseURL } from "../main"
import CommentBase from './CommentBase.vue'

const props = defineProps({
    comment: Object
})

const childComments = ref([])
const getChildComments = () => axios.get(`${baseURL}/api/v0/comments/${props.comment.articleID}?parentID=${props.comment.id}`)
    .then(resp => childComments.value = resp.data)
    .catch(err => console.error(err))
getChildComments()
provide('getChildComments', getChildComments)
</script>

<template>
    <div class="py-4">
        <div class="my-4">
            <CommentBase :comment="comment" />
        </div>
        <div class="ml-4 sm:ml-[calc(48px+1rem)] my-4" v-for="childComment in childComments" :key="childComment.id">
            <CommentBase :comment="childComment" />
        </div>
    </div>
</template>
