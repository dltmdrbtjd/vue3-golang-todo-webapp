<template>
  <div class="text-center mt-10">
    <p>안녕하세요, {{ userInfo.name }}님</p>
  </div>
</template>
<script setup lang="ts">
import { useUserStore } from '@/store/user';
import { computed } from 'vue';
import { User } from '@/models/user.model';


const userStore = useUserStore();
const userInfo = computed<User>(() => userStore.getUserDetail);
</script>
<style scoped lang="scss">
p {
  @apply text-white font-semibold text-2xl mt-4;
}
</style>
