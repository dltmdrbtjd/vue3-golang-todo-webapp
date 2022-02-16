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
import { ref, onMounted, computed } from "vue";
import { Todo } from "@/store/todo/type";

const todoStore = useTodoStore();
const todoList = computed<Todo[]>(() => todoStore.getTodoList);

let text = ref<string>("");
async function CreateTodo() {
  const newTodo = {
    username: todoList.value[0].username,
    content: text?.value,
    title: "",
  };
  await todoStore.addTodo(newTodo);
  text.value = "";
}

onMounted(() => {
  todoStore.getTodo("dltmdrbtjd1");
});
</script>
