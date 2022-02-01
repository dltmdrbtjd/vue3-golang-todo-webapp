<template>
  <button
    class="bg-sky-600 text-stone-50 px-2 py-1 rounded-md"
    @click="store.Increse(1)"
  >
    숫자 더하기
  </button>
  <p class="text-white">state: {{ store.testNumber }}</p>
  <h2 class="text-white">getters: {{ store.getCurrentNumber }}</h2>
</template>
<script setup lang="ts">
import { useTestStore } from "@/store/test";

const store = useTestStore();
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
