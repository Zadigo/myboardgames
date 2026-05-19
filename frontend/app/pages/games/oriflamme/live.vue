<template>
  <section id="oriflamme" class="h-screen px-20 space-y-2 mx-auto my-10">
    <lazy-oriflamme-base-card-scroll title="Influence queue" hydrate-on-idle>
      <oriflamme-card v-for="(card, idx) in influenceQueue" :key="card.uuid" :card="card" :in-queue="true" :index-in-queue="idx" />
    </lazy-oriflamme-base-card-scroll>

    <oriflamme-base-card-scroll title="Your hand">
      <oriflamme-card v-for="(card, idx) in playerCards" :key="card.uuid" :card="card" :index-in-queue="idx" />
    </oriflamme-base-card-scroll>

    <dev-only>
      <lazy-oriflamme-dev-tool hydrate-on-idle />
    </dev-only>
  </section>
</template>

<script lang="ts" setup>
import { useOriflammeActionsComposable, useOriflammeComposable, useQueuePossibleActions } from '~/composables/oriflamme'

definePageMeta({
  layout: 'game'
})

/**
 * Websocket
 */

const { ws, tableUuid, playerUuid, influenceQueue, playerCards } = useOriflammeComposable()
useOriflammeActionsComposable(ws, tableUuid, playerUuid)

/**
 * Play Options
 */

useQueuePossibleActions(influenceQueue)

/**
 * Page Theme
 */

const theme = ['bg-primary-50', 'dark:bg-primary-900', 'dark:text-primary-50']

onMounted(() => {
  document.body.classList.add(...theme)
})

onUnmounted(() => {
  document.body.classList.remove(...theme)
})
</script>
