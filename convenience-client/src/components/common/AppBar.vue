<template>
  <header>
    <div class="mx-auto max-w-3xl w-full flex justify-between">
      <ul class="flex items-center">
        <li>TODO</li>
        <li>BLOG</li>
        <li>ACCOUNT-BOOK</li>
      </ul>
      <div class="flex items-center">
        <p class="mr-4 text-white">dltmdrbtjd</p>
        <button type="button" class="round-btn bg-gray-700">
          <img
            src="https://avatars.githubusercontent.com/u/59644518?v=4"
            alt="profile image"
            class="rounded-full shadow-lg bg-slate-500"
          />
        </button>
      </div>
    </div>
  </header>
</template>
<script setup lang="ts"></script>
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
