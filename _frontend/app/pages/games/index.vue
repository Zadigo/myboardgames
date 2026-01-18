<template>
  <section id="all-games">
    <nuxt-container>
      <div class="games">
        <h1>All Games</h1>
        <p>This is the all games page.</p>

        <div class="grid grid-cols-3 gap-2">
          <nuxt-card v-for="game in games" :key="game.name">
            <nuxt-link :to="game.to">
              {{ game.name }}
            </nuxt-link>
          </nuxt-card>
        </div>
      </div>
    </nuxt-container>
  </section>
</template>

<script setup lang="ts">
const games = [
  { name: 'Dune Imperium', to: '/games/dune-imperium' },
  { name: 'Wonderful World', to: '/games/wonderful-world' }
  
]
</script>
