<template>
  <article ref="cardEl" :data-nam="card?.name" class="overflow-hidden relative rounded-lg hover:shadow-lg transition-all ease-in-out duration-300 cursor-move hover:scale-105">
    <nuxt-img :src="`wonderful/material/developments/${card?.category.toLowerCase()}/${card?.image}`" :alt="card?.name" />

    <transition
      enter-active-class="transition-all ease-in-out duration-500"
      leave-active-class="transition-all ease-in-out duration-500"
      enter-from-class="opacity-0 translate-y-0"
      enter-to-class="opacity-100 translate-y-3.5"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="isHovered" id="actions" class="absolute top-3.5 left-1/2 -translate-x-1/2">
        <nuxt-button>
          <icon name="lucide:arrow-down" /> {{ pressed }}
        </nuxt-button>
      </div>
    </transition>
  </article>
</template>

<script lang="ts" setup>
import type { BaseCard } from '~/types'

defineProps<{
  card?: BaseCard
}>()

const cardEl = useTemplateRef<HTMLDivElement>('cardEl')
const isHovered = useElementHover(cardEl)

const { pressed } = useMousePressed({ target: cardEl})
</script>
