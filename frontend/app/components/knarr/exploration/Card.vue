<template>
  <div :id="`exploration-${id}`" class="h-40 bg-primary-400 rounded-lg shadow-sm col-span-3 first:col-start-1">
    Exploration {{ id }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string
}>()
</script>
