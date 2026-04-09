<template>
  <section id="oriflamme" class="h-screen px-20 space-y-2 mx-auto my-10">
    {{ selectedCards }}
    {{ playerCards }}
    <oriflamme-base-card-scroll title="Influence queue">
      <oriflamme-card v-for="card in influenceQueue" :key="card.uuid" :card="card" :in-queue="true" />
    </oriflamme-base-card-scroll>

    <oriflamme-base-card-scroll title="Your hand">
      <oriflamme-card v-for="card in playerCards" :key="card.uuid" :card="card" />
    </oriflamme-base-card-scroll>
  </section>
</template>

<script lang="ts" setup>
import { useOriflammeActionsComposable, useOriflammeComposable } from '~/composables/oriflamme'
import type { OriflammeCard } from '~/types'

definePageMeta({
  layout: 'game'
})

const { ws, tableUuid, playerUuid, influenceQueue, playerCards } = useOriflammeComposable()
const { selectedCards } = useOriflammeActionsComposable(ws, tableUuid, playerUuid)

const theme = ['bg-primary-50', 'dark:bg-primary-900', 'dark:text-primary-50']

onMounted(() => {
  document.body.classList.add(...theme)
})

onUnmounted(() => {
  document.body.classList.remove(...theme)
})
</script>
