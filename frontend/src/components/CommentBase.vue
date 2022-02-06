<script setup>
import { inject } from 'vue'
import Reply from './Reply.vue'

defineProps({
    comment: Object
})

const visibleReply = inject('visibleReply')
const toggleReply = inject('toggleReply')
</script>

<template>
    <div class="flex items-start">
        <img class="rounded-full" :src="comment.avatar" />
        <div class="ml-4">
            <div class="font-bold text-sm">{{ comment.name }}</div>
            <p class="my-1 text-lg">{{ comment.content }}</p>
            <div class="text-sm text-gray-500">
                <span class="mr-4">{{ comment.date }}</span>
                <button type="button" @click="toggleReply(comment.id)">回复</button>
            </div>
        </div>
    </div>
    <Reply class="mt-4" v-if="visibleReply === comment.id" :comment="comment" />
</template>
