<template>
  <div class="h-40 w-30 bg-flip-seven-900 rounded-lg" />
</template>

<script lang="ts" setup>
import type { Deck } from '~/types'

defineProps<{
  card: Deck
}>()
</script>
