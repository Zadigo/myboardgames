<template>
  <div :id="`artifact-${id}`" class="h-30 bg-primary-500 rounded-lg shadow-sm col-span-3 first:col-start-4">
    Artifact {{ id }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string
}>()
</script>
