<template>
  <div class="flex justify-center mt-10 w-full">
    <input
      class="input"
      :type="props.type"
      :placeholder="props.placeholder"
      :value="props.modelValue"
      @input="handleInputChange($event)"
    />
  </div>
</template>
<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    type?: string;
    modelValue?: string | number;
    placeholder?: string;
  }>(),
  {
    type: "text",
    modelValue: "",
    placeholder: "",
  }
);

const emit = defineEmits(["update:modelValue"]);

const handleInputChange = (event: Event) => {
  emit('update:modelValue', (event.target as HTMLInputElement).value)
}
</script>
<style scoped lang="scss">
.input {
  @apply border rounded py-2 px-2 text-xl font-bold text-gray-700 dark:text-white dark:bg-sky-800 leading-tight focus:outline-none transition-colors w-1/2;
}
</style>
