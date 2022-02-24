<template>
  <Spinner />
</template>
<script setup lang="ts">
import { onMounted } from 'vue';
import Spinner from '@/components/common/Spinner.vue';
import { useRoute } from 'vue-router';
import $http from 'axios';
import router from '@/routes';
import { setLocalStorage } from '@/scripts/localStorage';
import { useUserStore } from '@/store/user';

const route = useRoute();
let authQuery = route.fullPath.split('?')[1];

const userStore = useUserStore();

onMounted(async () => {
  const resp = await $http.get(`/ google-login/callback?${authQuery}`);
  setLocalStorage('google-access-token', resp.data.data);
  userStore.getUserInfo();
  userStore.googleTokenVerification(true);
  router.push('/');
});
</script>
