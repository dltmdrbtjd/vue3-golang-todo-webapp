<template>
  <div class="mx-auto max-w-xl w-full flex justify-between">
    <button
      class="bg-sky-600 text-stone-50 px-2 py-1 rounded-md"
      @click="store.Increse(1)"
    >
      숫자 더하기
    </button>
    <p class="text-white">state: {{ store.testNumber }}</p>
    <h2 class="text-white">getters: {{ store.getCurrentNumber }}</h2>
  </div>
</template>
<script setup lang="ts">
import { useTestStore } from "@/store/test";

const store = useTestStore();
</script>

<style scoped></style>
