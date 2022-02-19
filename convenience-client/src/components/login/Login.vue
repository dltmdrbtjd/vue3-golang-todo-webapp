<template>
  <a @click="login">google login</a>
</template>
<script setup lang="ts">
import {onMounted} from 'vue';
import $http from 'axios';

async function login() {
    window.location.replace(url);
}

let url = "";
onMounted(async () => {
    const resp = await $http.get("/google-login")
    url = resp.data.data
    console.log(url)
})
</script>