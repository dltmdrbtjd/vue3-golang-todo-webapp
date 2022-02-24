<template>
  <header>
    <div class="mx-auto max-w-3xl w-full flex justify-between">
      <ul class="flex items-center">
        <li>BLOG</li>
        <li v-if="authentication">TODO</li>
        <li v-if="authentication">ACCOUNT-BOOK</li>
      </ul>
      <div class="flex items-center">
        <p class="mr-4 text-white">{{ userInfo.name }}</p>
        <button type="button" class="round-btn bg-gray-700" 
            v-if="userInfo.picture">
          <img
            :src="userInfo.picture"
            alt="user profile image"
            class="rounded-full shadow-lg bg-slate-500"
          />
          <div v-if="!userInfo" class="rounded-full shadow-lg bg-slate-500"></div>
        </button>
        <button  v-else type="button" class="text-white bg-[#4285F4] hover:bg-[#4285F4]/90 focus:ring-4 focus:ring-[#4285F4]/50 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-[#4285F4]/55 mr-2 mb-2">
          <svg class="mr-2 -ml-1 w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fab" data-icon="google" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 488 512"><path fill="currentColor" d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"></path></svg>
          Sign in with Google
        </button>
      </div>
    </div>
  </header>
</template>
<script setup lang="ts">
import { User } from '@/models/user.model';
import { useUserStore } from '@/store/user';
import { computed } from 'vue';

const userStore = useUserStore();
const authentication = computed<boolean>(() => userStore.getTokenVerification)
const userInfo = computed<User>(() => userStore.getUserDetail)
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
