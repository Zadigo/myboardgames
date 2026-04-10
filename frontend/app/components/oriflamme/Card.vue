<template>
  <div class="wrapper">
    <article ref="cardEl" :data-uuid="card.uuid" :class="cardTheme" class="relative h-70 min-w-50 transition-all bg-center bg-no-repeat bg-cover ease-in-out duration-500 rounded-lg p-2 hover:shadow-xl cursor-pointer overflow-hidden">
      <nuxt-img :src="cardImage" :alt="card.name" class="w-full h-full absolute top-0 left-0 z-9 rounded-lg" @click="displayOptions(indexInQueue)" />

      <lazy-nuxt-popover mode="hover" hydrate-on-idle>
        <nuxt-button class="absolute top-0 right-0 z-10">
          Help
        </nuxt-button>

        <template #content>
          <div class="w-70 h-auto p-4">
            <nuxt-badge label="Power" color="info" />
            <h3 class="font-bold italic mb-3">
              {{ cardInfo?.power }}
            </h3>

            <nuxt-badge label="Description" color="info" />
            <p class="font-light">
              {{ cardInfo?.description }}
            </p>
          </div>
        </template>
      </lazy-nuxt-popover>

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

    <div class="w-full flex justify-center">
      <nuxt-badge :label="card.name" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useOriflammeActionsStore, useOriflammeAllCardsComposable, useQueuePossibleActionsStore } from '~/composables/oriflamme'
import type { AllCharacters, OriflammeCard } from '~/types'

const props = defineProps<{
  indexInQueue: number
  card: OriflammeCard<AllCharacters>
  inQueue?: boolean
}>()

/**
 * Actions Hover
 */

const cardEl = useTemplateRef('cardEl')

const isHovered = useElementHover(cardEl, {
  delayEnter: 200
})

/**
 * Card Actions
 */

const { isSelected } = useOriflammeActionsStore()
const toggleSelection = isSelected(props.card)

/**
 * Card Theme
 */

const cardTheme = computed(() => {
  return [
    {
      'bg-primary-200 dark:bg-primary-700': !toggleSelection.value,
      'bg-blue-500 dark:bg-blue-700': toggleSelection.value
    }
  ]
})

/**
 * Card Information
 */

const { getCardByName } = useOriflammeAllCardsComposable('all')
const cardInfo = getCardByName(props.card.name)

/**
 * Card Highlight
 */

const { isHighlighted, displayOptions, selectedIndex, playOptions } = useQueuePossibleActionsStore()
const highlight = isHighlighted(props.card)

/**
 * Card Image
 */

const cardImage = computed(() => {
  if (!props.card.inQueue) return cardInfo.value?.image

  if (props.card.isRevealed) {
    return cardInfo.value?.image
  }

  return '/images/oriflamme/back.jpg'
})
</script>

<style lang="css" scoped>
/* article::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(315deg, var(--color-green-50), var(--color-green-100));
}

article::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(315deg, var(--color-green-50), var(--color-green-300));
  filter: blur(30px);
} */
</style>
