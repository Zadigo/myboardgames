<template>
  <section id="oriflamme" class="h-screen px-20 space-y-2 mx-auto my-10">
    <div class="rounded-xl bg-primary-100 w-full h-auto p-5">
      <h1 class="font-bold bg-secondary-800/20 backdrop-blur-3xl shadow-sm p-3 rounded-lg mb-3 w-auto max-w-100">
        Influence queue
      </h1>

      <div class="overflow-x-scroll w-full h-auto flex gap-2">
        <oriflamme-card v-for="card in cards" :key="card.Uuid" :card="card" />
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import type { OriflammeCard } from '~/types'

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
</script>
