<template>
  <div class="flex justify-center mt-10 w-full">
    <input
      class="input"
      type="text"
      :type="type"
      :placeholder="placeholder"
      v-model="value"
    />
  </div>
</template>
<script setup lang="ts">
import { useVModel } from "@vueuse/core";

const props = withDefaults(
  defineProps<{
    label?: string;
    type?: string;
    modelValue?: string | number;
    placeholder?: string;
  }>(),
  {
    label: "",
    type: "text",
    modelValue: "",
    placeholder: "",
  }
);

const emit = defineEmits(["update:modelValue"]);
const value = useVModel(props, "modelValue", emit);

console.log(props.modelValue);
</script>
<style scoped lang="scss">
.input {
  @apply border rounded py-2 px-2 text-xl font-bold text-gray-700 dark:text-white dark:bg-sky-800 leading-tight focus:outline-none transition-colors w-1/2;
}
</style>
