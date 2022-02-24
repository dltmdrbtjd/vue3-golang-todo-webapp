<template>
  <div class="relative w-full h-full">
    <AppBar />
    <main class="mx-auto max-w-3xl w-full h-full">
      <router-view />
    </main>
  </div>
</template>
<script setup lang="ts">
import $http from "axios";
import AppBar from "@/components/common/AppBar.vue";
import { getLocalStorage } from "./scripts/localStorage";
import { useUserStore } from "./store/user";
import { onMounted, reactive } from "vue";
import router from "./routes";
import { useRoute } from "vue-router";

const userStore = useUserStore();

async function userTokenVerification(googleAccessToken: string) {
  try {
      const resp = await $http.get(`/google-token-verification/${googleAccessToken}`)
      userStore.googleTokenVerification(resp.data.verification)
    } catch(error) {
    console.error(error)
  }
}

const path = useRoute().path
const loginPathCheck = path === "/google-login/callback" ? true : false;
const callbackPathCheck = path === "/login"  ? true : false;

const token = getLocalStorage("google-access-token");
const reactiveRef = reactive({ token })

onMounted(() => {
  if (reactiveRef.token) {
    userTokenVerification(String(reactiveRef.token));
    userStore.getUserInfo();
  } else if(loginPathCheck && callbackPathCheck){
    router.push("/login")
    userStore.googleTokenVerification(false)
  }
})
</script>
