<template>
  <div class="relative w-full h-full">
    <AppBar />
    <Login />
    <main class="mx-auto max-w-3xl w-full h-full">
      <router-view />
    </main>
  </div>
</template>
<script setup lang="ts">
import AppBar from "@/components/common/AppBar.vue";
import Login from "./components/login/Login.vue";
</script>
<style></style>
