<template>
  <section id="oriflamme" class="h-screen px-20 space-y-2 mx-auto my-10">
    <oriflamme-base-card-scroll title="Influence queue">
      <oriflamme-card v-for="card in cards" :key="card.Uuid" :card="card" :in-queue="true" />
    </oriflamme-base-card-scroll>

    <oriflamme-base-card-scroll title="Your hand">
      <oriflamme-card v-for="card in cards" :key="card.Uuid" :card="card" />
    </oriflamme-base-card-scroll>
  </section>
</template>

<script lang="ts" setup>
import type { OriflammeCard } from '~/types'

definePageMeta({
  layout: 'game'
})

const cards = ref<OriflammeCard[]>([
  {
    Uuid: '1',
    Name: 'Archer',
    Color: 'red',
    Image: '/images/oriflamme/card1.png',
    IsDiscarded: false,
    IsRemoved: false,
    IsRevealed: false,
    IsSelected: false,
    Owner: {
      Uuid: 'player 1'
    },
    PositionInQueue: 1,
    Stack: [],
    Tokens: 1,
    Type: 'Character'
  }
])

const theme = ['bg-primary-50']

onMounted(() => {
  document.body.classList.add(...theme)
})

onUnmounted(() => {
  document.body.classList.remove(...theme)
})
</script>
