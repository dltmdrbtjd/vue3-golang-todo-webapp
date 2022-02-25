<template>
  <div v-for="(todo, idx) in computedRef.todoList" :key="todo._id">
    <div
      class="mt-10 card border w-2/3 mx-auto bg-gray-800 text-white flex items-center justify-between"
      v-if="!todo.isEdit"
    >
      <div class="w-3/4">
        <strong class="font-bold text-2xl block">{{ todo.content }}</strong>
      </div>
      <div class="flex justify-end">
        <div class="flex">
          <svg
            @click="deleteTodo(String(todo._id))"
            xmlns="http://www.w3.org/2000/svg"
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            stroke="red"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            />
          </svg>
          <svg
            @click="onTodoEdit(String(todo._id), todo.content)"
            xmlns="http://www.w3.org/2000/svg"
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            stroke="#0ea5e9"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
            />
          </svg>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            @click="todoStatusChange(String(todo._id), idx)"
            :stroke="todo.status === 'success' ? 'green' : 'gray'"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
        </div>
      </div>
    </div>
    <div v-else>
      <TodoInput v-model="text" @keydown.enter.prevent="editTodo(todo, idx)" />
    </div>
  </div>
</template>
<script setup lang="ts">
import { useTodoStore } from "@/store/todo";
import TodoInput from "./TodoInput.vue";
import { Todo } from "@/models/todo.model";
import { computed, reactive, ref } from "vue";

const todoStore = useTodoStore();
const todoList = computed<Todo[]>(() => todoStore.getTodoList);

const computedRef = reactive({ todoList });

function deleteTodo(todoId: string) {
  todoStore.deleteTodo(todoId);
}

let text = ref<string>("");

function onTodoEdit(todoId: string, content: string) {
  todoStore.isEditTodo(todoId, true);
  text.value = content;
}

function editTodo(todo: Todo, index: number) {
  todoStore.editTodo(String(todo._id), text.value, index);
}

function todoStatusChange(todoId: string, index: number) {
  todoStore.checkTodoStatus(todoId, index);
}

</script>
<style scoped lang="scss">
svg {
  @apply ml-2 cursor-pointer;
}
</style>
