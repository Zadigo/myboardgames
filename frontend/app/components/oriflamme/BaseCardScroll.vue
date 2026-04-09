<template>
  <div class="rounded-xl bg-primary-100 w-full h-auto p-5">
    <h1 class="font-bold bg-secondary-800/20 backdrop-blur-3xl shadow-sm p-3 rounded-lg mb-3 w-auto max-w-100">
      {{ title }}
    </h1>

    <div class="overflow-x-scroll w-full h-auto flex gap-2">
      <slot />
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  title: string
}>()
</script>
