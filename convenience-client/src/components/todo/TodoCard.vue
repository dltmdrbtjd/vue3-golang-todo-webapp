<template>
    <div class="mt-10 card border w-2/3 mx-auto bg-gray-800 text-white flex items-center justify-between">
        <div class="w-3/4">
            <strong class="font-bold text-2xl block">오늘 todo카드 까지 완료하기</strong>
        </div>
        <div class="flex justify-end">
            <div class="flex">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="red">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
                <svg @click="onClickAddTodo(todoMocks)" xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="#0ea5e9">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                </svg>
                <svg v-if="status" xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="green">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="gray">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { useTodoStore } from '@/store/todo';
import { Todo } from '@/store/todo/type';

let status = true
const todoStore = useTodoStore();
const todoMocks: Todo = {
    username: "hi",
    content: "hi",
    title: "hi"
}

function onClickAddTodo (todo: Todo) {
    todoStore.addTodo(todo)
}
</script>
<style scoped lang="scss">
svg {
    @apply ml-2 cursor-pointer
}
</style>