<template>
  <section v-if="showPlayerHands" id="players-hands" class="fixed h-200 bg-blue-500/20 w-full bottom-0 left-0 shadow-lg backdrop-blur-2xl z-50">
    <div class="relative">
      <div class="py-5 px-20 bg-blue-100/50 backdrop-blur-2xl" @click="() => { togglePlayerHands() }">
        <div class="p-5 rounded-lg hover:bg-primary-100/20 transition-colors ease-in-out duration-300 cursor-pointer">
          <h2 class="font-bold text-2xl">
            Player Name
          </h2>
        </div>
      </div>

      <div class="p-20">
        <div class="absolute left-5 top-1/2 -translate-y-1/2">
          <nuxt-button id="previous-player">
            <icon name="lucide:arrow-left" />
          </nuxt-button>
        </div>

        <div class="grid grid-cols-12 gap-2">
          <div id="player-1-hand" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2 col-start-2">
            Player 1 Hand
          </div>

          <div id="player-2-hand" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2">
            Player 2 Hand
          </div>
        </div>

        <div class="absolute right-5 top-1/2 -translate-y-1/2">
          <nuxt-button id="next-player">
            <icon name="lucide:arrow-right" />
          </nuxt-button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { usePlayerHandsStore } from '~/composables/knarr'

const { togglePlayerHands, showPlayerHands } = usePlayerHandsStore()
</script>
