<template>
  <header>
    <div class="mx-auto max-w-3xl w-full flex justify-between">
      <ul class="flex items-center">
        <li v-if="authentication">TODO</li>
      </ul>
      <div class="flex items-center">
        <p class="mr-4 text-white">{{ userInfo.name }}</p>
        <button
          type="button"
          class="round-btn bg-gray-700"
          v-if="authentication"
        >
          <img
            :src="userInfo.picture"
            alt="user profile image"
            class="rounded-full shadow-lg bg-slate-500"
          />
          <div
            v-if="!userInfo"
            class="rounded-full shadow-lg bg-slate-500"
          ></div>
        </button>
        <GoogleLoginButton v-else />
      </div>
    </div>
  </header>
</template>
<script setup lang="ts">
import { User } from '@/models/user.model';
import { useUserStore } from '@/store/user';
import { computed } from 'vue';
import GoogleLoginButton from './LoginButton.vue';

const userStore = useUserStore();
const authentication = computed<boolean>(() => userStore.getTokenVerification);
const userInfo = computed<User>(() => userStore.getUserDetail);
</script>
<style scoped lang="scss">
header {
  @apply px-4 py-4 flex sticky top-0 left-0 right-0 z-10 bg-white dark:bg-gray-800 transition-all outline-none border-none;
}
li {
  @apply pr-4 text-white font-bold cursor-pointer hover:text-lime-300;
}
.round-btn {
  @apply w-10 h-10 rounded-full;
}
</style>
