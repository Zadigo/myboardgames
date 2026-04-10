<template>
  <div class="rounded-xl bg-primary-100 dark:bg-primary-800 w-full h-auto px-5 py-10">
    <h1 class="font-bold bg-secondary-800/20 dark:bg-secondary-900/20 backdrop-blur-3xl shadow-sm p-3 rounded-lg mb-3 w-auto max-w-100">
      {{ title }}
    </h1>

    <div id="scrollable-area" class="overflow-x-scroll w-full h-auto flex gap-2">
      <slot />
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  title: string
}>()
</script>

<style scoped>
#scrollable-area::-webkit-scrollbar {
  display: none;
}
</style>
