<template>
  <article ref="cardEl" :data-uuid="card.Uuid" class="relative h-70 min-w-50 bg-primary-200 transition-all ease-in-out duration-500 rounded-lg p-2 hover:shadow-xl cursor-pointer">
    <!-- Actions -->
    <transition>
      <oriflamme-actions-base v-if="isHovered" :in-queue="inQueue" :card="card" />
    </transition>

    <div v-if="inQueue" class="rounded-full w-10 h-10 bg-secondary-200/50 absolute bottom-2 left-2 animate-bounce transition-all duration-500" />
  </article>
</template>

<script lang="ts" setup>
import type { OriflammeCard } from '~/types'

defineProps<{
  card: OriflammeCard
  inQueue?: boolean
}>()

const cardEl = useTemplateRef('cardEl')

const isHovered = useElementHover(cardEl, {
  delayEnter: 200
})
</script>
