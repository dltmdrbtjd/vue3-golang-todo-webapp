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
import {computed, onMounted, reactive} from "vue";
import router from "./routes";
import { useRoute } from "vue-router";

const userStore = useUserStore();

async function userTokenVerification(useremail: string) {
  try {
      await $http.get(`/google-token-verification/${useremail}`)
      userStore.googleTokenVerification(true)
    } catch(error) {
    console.error(error)
  }
}
function getUserInfoDetail() {
  userStore.getUserInfo()
}

const path = useRoute().path
const loginPathCheck = path === "/google-login/callback" ? true : false;
const callbackPathCheck2 = path === "/login"  ? true : false;

const useremail = getLocalStorage("convenience-tools-email");
const reactiveRef = reactive({ useremail })

onMounted(() => {
  if (reactiveRef.useremail) {
    userTokenVerification(String(reactiveRef.useremail))
    getUserInfoDetail()
  } else if(loginPathCheck && callbackPathCheck2){
    router.push("/login")
    userStore.googleTokenVerification(false)
  }
})
</script>
<style></style>
