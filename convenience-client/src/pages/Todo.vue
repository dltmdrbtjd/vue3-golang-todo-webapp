<template>
  <Greeting />
  <TodoInput
    :placeholder="'TODO를 작성해주세요'"
    v-model="text"
    @keydown.enter.prevent="save"
  />
</template>
<script setup lang="ts">
import Greeting from "@/components/todo/Greeting.vue";
import TodoInput from "@/components/todo/TodoInput.vue";
import { ref } from "vue";

const text = ref<string>("");
const save = () => {
  console.log(text.value);
};
</script>
