<template>
  <div class="rounded-full bg-primary-500 w-15 h-15 flex items-center justify-center">
    <icon :name="icon" />
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  icon: string
}>()
</script>
