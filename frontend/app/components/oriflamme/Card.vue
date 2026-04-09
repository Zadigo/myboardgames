<template>
  <article ref="cardEl" :data-uuid="card.Uuid" :class="{ ' bg-primary-200': !toggleSelection, 'bg-blue-500': toggleSelection }" class="relative h-70 min-w-50 transition-all ease-in-out duration-500 rounded-lg p-2 hover:shadow-xl cursor-pointer">
    <!-- Actions -->
    <transition
      mode="in-out"
      enter-active-class="transition-all ease-in-out duration-300"
      enter-from-class="scale-95"
      enter-to-class="scale-100"
      leave-from-class="scale-100"
      leave-to-class="scale-95"
    >
      <oriflamme-actions-base v-if="isHovered" :in-queue="inQueue" :card="card" />
    </transition>

    <div v-if="inQueue" class="rounded-full w-10 h-10 bg-secondary-200/50 absolute bottom-2 left-2 animate-bounce transition-all duration-500" />
  </article>
</template>

<script lang="ts" setup>
import { useOriflameeActionsStore } from '~/composables/oriflamme'
import type { OriflammeCard } from '~/types'

const props = defineProps<{
  card: OriflammeCard
  inQueue?: boolean
}>()

const cardEl = useTemplateRef('cardEl')

const isHovered = useElementHover(cardEl, {
  delayEnter: 200,
  delayLeave: 500
})

const { isSelected } = useOriflameeActionsStore()
const toggleSelection = isSelected(props.card)
</script>
