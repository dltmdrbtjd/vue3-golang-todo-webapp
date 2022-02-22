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
import { ref, onMounted, computed, reactive } from "vue";
import { useUserStore } from "@/store/user";
import { User } from "@/models/user.model";

const todoStore = useTodoStore();
const userStore = useUserStore();
const userInfo = computed<User>(() => userStore.getUserDetail)
const computedRef = reactive({ userInfo })

let text = ref<string>("");
async function CreateTodo() {
  const newTodo = {
    username: computedRef.userInfo.name,
    content: text?.value,
    title: "",
  };
  await todoStore.addTodo(newTodo);
  text.value = "";
}

onMounted(() => {
  todoStore.getTodo(computedRef.userInfo.name);
});
</script>
