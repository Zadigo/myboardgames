<template>
  <div class="ranking">
    <div class="flex justify-between mb-10">
      <nuxt-button variant="outline" color="info" size="xl" class="rounded-full">
        <icon name="i-lucide-filter" />
      </nuxt-button>

      <nuxt-button variant="outline" color="info" size="xl" class="rounded-full">
        <icon name="i-lucide-search" />
      </nuxt-button>
    </div>

    <div class="rounded-lg">
      <div v-for="(item, index) in rankingsFixture" :key="index" class="p-4 border-b border-gray-200 last:border-0 flex items-center justify-between gap-5">
        <div class="flex items-center justify-start gap-3">
          <h2 class="text-4xl font-bold">{{ item.rank }}.</h2>
          <nuxt-avatar :src="``" size="3xl" />
          <p class="font-semibold text-2xl">{{ item.player }}</p>
        </div>

        <p class="text-4xl font-bold">{{ item.points }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const rankingsFixture = [
  { rank: 1, player: 'Alice', gender: 'W', games_played: 10, points: 150 },
  { rank: 2, player: 'Bob', gender: 'M', games_played: 10, points: 120 },
  { rank: 3, player: 'Charlie', gender: 'M', games_played: 10, points: 100 }
]
</script>
