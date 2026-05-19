<template>
  <div class="rounded-full bg-primary-500 dark:bg-primary-600/20 backdrop-blur-xl shadow-sm w-10 h-10 p-1 flex items-center justify-center">
    <nuxt-icon :name="icon" />
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  icon: string
}>()
</script>
