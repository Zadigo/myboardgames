<template>
  <article ref="cardEl" :data-uuid="card.uuid" :class="{ ' bg-primary-200 dark:bg-primary-700': !toggleSelection, 'bg-blue-500 dark:bg-blue-700': toggleSelection }" class="relative h-70 min-w-50 transition-all ease-in-out duration-500 rounded-lg p-2 hover:shadow-xl cursor-pointer overflow-hidden">
    <img :src="card.image" :alt="card.name" class="w-full h-full aspect-square object-cover rounded-lg">

    <nuxt-popover mode="hover">
      <nuxt-button class="absolute top-0 right-0 z-10">
        Help
      </nuxt-button>

      <template #content>
        <div class="w-70 has-autofill: p-4">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae ab 
          qui at maiores aperiam aut dolores voluptatum suscipit inventore fugit nesciunt corporis, 
          ut molestias sed iste et perspiciatis quos iusto.
        </div>
      </template>
    </nuxt-popover>

    <!-- Actions -->
    <transition
      mode="in-out"
      enter-active-class="transition-all ease-in-out duration-300"
      enter-from-class="scale-95"
      enter-to-class="scale-100"
      leave-from-class="scale-100"
      leave-to-class="scale-95"
    >
      <!-- v-if="isHovered" -->
      <oriflamme-actions-base :in-queue="inQueue" :card="card" />
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
  delayEnter: 200
})

const { isSelected } = useOriflameeActionsStore()
const toggleSelection = isSelected(props.card)
</script>
