<template>
  <div :id="`artifact-${id}`" class="h-30 bg-primary-600 rounded-lg shadow-sm col-span-3 first:col-start-4">
    <u-popover :open-delay="300" mode="hover" class="w-70" hover>
      <div class="h-30 w-full mx-auto" />

      <template #content>
        <div class="p-10">
          Artifact {{ id }}
        </div>
      </template>
    </u-popover>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string
}>()
</script>
