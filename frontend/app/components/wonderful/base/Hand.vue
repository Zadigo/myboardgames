<template>
  <div id="hand" ref="handEl" class="min-h-58.75">
    <h2>Votre main</h2>
    <div class="flex justify-center gap-4 content-center px-10 py-3">
      <wonderful-base-card v-for="card in cards" :key="card.image" :card="card" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { BaseCard } from '~/types'

defineProps<{
  cards: BaseCard[]
}>()

const el = useTemplateRef('handEl')
useMouseInElement(el, {})
</script>
