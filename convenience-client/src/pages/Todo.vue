<template>
  <Greeting />
  <TodoInput
    :placeholder="'TODO를 작성해주세요'"
    v-model="text"
    @keydown.enter.prevent="CreateTodo"
  />
  <TodoCard />
</template>
<script setup lang="ts">
import Greeting from "@/components/todo/Greeting.vue";
import TodoInput from "@/components/todo/TodoInput.vue";
import TodoCard from "@/components/todo/TodoCard.vue";

import { useTodoStore } from "@/store/todo";
import { ref, onMounted } from "vue";
import { getLocalStorage } from "@/scripts/localStorage";

const todoStore = useTodoStore();
const userName = String(getLocalStorage("username"))

let text = ref<string>("");
async function CreateTodo() {
  const newTodo = {
    username: userName,
    content: text?.value,
    title: "",
  };
  await todoStore.addTodo(newTodo);
  text.value = "";
}

onMounted(() => {
  todoStore.getTodo(userName);
});
</script>