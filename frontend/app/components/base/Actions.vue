<template>
  <div id="actions" ref="actionEl" :class="color" class="absolute top-0 right-0 backdrop-blur-2xl shadow-sm w-auto h-auto rounded-bl-2xl py-5 px-3 z-50 flex flex-col gap-2">
    <slot />
  </div>
</template>

<script lang="ts" setup>
const { color = 'bg-primary-500/20' } = defineProps<{
  color?: string
}>()
const actionEl = useTemplateRef('actionEl')
</script>
