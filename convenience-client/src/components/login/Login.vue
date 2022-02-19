<template>
  <a href="https://accounts.google.com/o/oauth2/auth?client_id=261134792769-p58c82c8ab3oqa33nofjoglurnudlamm.apps.googleusercontent.com\u0026redirect_uri=http%3A%2F%2Flocalhost%3A8081%2Fgoogle-login%2Fcallback\u0026response_type=code\u0026scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email\u0026state=random">google login</a>
  <div>{{ user }}</div>
</template>
<script setup lang="ts">
import {ref, onMounted} from 'vue';
import { installGoogleAuth } from '@/configs/googleOAuth';
import getEnv from "@/scripts/env";
import $http from 'axios';

let gAuth: any;
const googleClientId = getEnv("VITE_GOOGLE_OAUTH2_CLIENT_ID")
const user = ref({});
const options = {
    clientId: googleClientId,
    scope: 'profile email',
    prompt: 'select_account'
};

async function login() {
   await $http.get("/google-login")
}

let url = "";
onMounted(async () => {
    const resp = await $http.get("/google-login")
    url = resp.data
})
</script>