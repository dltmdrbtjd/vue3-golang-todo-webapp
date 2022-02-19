<template>
    <Spinner />
</template>
<script setup lang="ts">
import { onMounted } from "vue";
import Spinner from '@/components/common/Spinner.vue';
import { useRoute } from "vue-router";
import $http from "axios";

const route = useRoute();
let authQuery = route.fullPath.split("?")[1];
onMounted(async () => {
    const resp = await $http.get(`/google-login/callback?${authQuery}`)
    console.log(resp.data)
})
</script>