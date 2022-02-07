<template>
  <div class="relative w-full h-full">
    <AppBar />
    <main class="mx-auto max-w-3xl w-full h-full">
      <router-view />
    </main>
  </div>
</template>
<script setup lang="ts">
import AppBar from "@/components/common/AppBar.vue";
</script>
<style></style>
