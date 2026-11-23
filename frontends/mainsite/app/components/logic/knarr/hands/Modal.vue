<template>
  <u-slideover v-model:open="showPlayerHands" side="bottom" inset title="Slideover with inset">
    <template #content>
      <div class="relative">
        <div class="py-10 px-20 bg-primary-100/50 backdrop-blur-2xl">
          <h2 class="font-bold text-xl">
            Player Name
          </h2>
        </div>

        <div class="p-20">
          <div class="absolute left-5 top-1/2 -translate-y-1/2">
            <u-button id="previous-player">
              <icon name="lucide:arrow-left" />
            </u-button>
          </div>

          <div class="grid grid-cols-12 gap-2">
            <div id="player-1-hand" class="h-80 bg-primary-400 rounded-lg shadow-sm col-span-2 col-start-2">
              Player 1 Hand
            </div>

            <div id="player-2-hand" class="h-80 bg-primary-400 rounded-lg shadow-sm col-span-2">
              Player 2 Hand
            </div>
          </div>

          <div class="absolute right-5 top-1/2 -translate-y-1/2">
            <u-button id="next-player">
              <icon name="lucide:arrow-right" />
            </u-button>
          </div>
        </div>
      </div>
    </template>
  </u-slideover>
</template>

<script setup lang="ts">
import { usePlayerHandsStore } from '~/composables/_knarr'

const { togglePlayerHands, showPlayerHands } = usePlayerHandsStore()

/**
 * Keys
 */

const stopLeft = onKeyStroke('ArrowLeft', (e) => {
  e.preventDefault()
  console.log('Previous Player')
})

const stopRight = onKeyStroke('ArrowRight', (e) => {
  e.preventDefault()
  console.log('Next Player')
})

watch(showPlayerHands, (value) => {
  if (!value) {
    stopLeft()
    stopRight()
  }
})
</script>
