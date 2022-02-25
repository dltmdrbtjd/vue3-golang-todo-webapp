<template>
  <div v-if="authentication">
    <Greeting />
    <TodoInput
      :placeholder="'TODO를 작성해주세요'"
      v-model="text"
      @keydown.enter.prevent="CreateTodo"
    />
    <TodoCard />
  </div>
  <div v-else class="text-center mt-10">
    <p>안녕하세요, 로그인후에 이용해주세요.</p>
  </div>
</template>
<script setup lang="ts">
import Greeting from "@/components/todo/Greeting.vue";
import TodoInput from "@/components/todo/TodoInput.vue";
import TodoCard from "@/components/todo/TodoCard.vue";

import { useTodoStore } from "@/store/todo";
import { ref, onMounted, computed } from "vue";
import { getLocalStorage } from "@/scripts/localStorage";
import { useUserStore } from "@/store/user";


const todoStore = useTodoStore();
const userStore = useUserStore();

const userName = String(getLocalStorage("username"))
const authentication = computed<boolean>(() => userStore.getTokenVerification);

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
<style scoped lang="scss">
p {
  @apply text-white font-semibold text-2xl mt-4;
}
</style>
