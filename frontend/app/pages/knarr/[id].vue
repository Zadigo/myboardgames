<template>
  <nuxt-container id="knarr" class="my-10 space-y-2" as="section">
    <!-- Artifacts Section -->
    <section id="artifacts">
      <div class="grid grid-cols-12 auto-cols-max gap-1">
        <lazy-knarr-artifacts-card id="artifact-1" hydrate-on-idle />
        <lazy-knarr-artifacts-card id="artifact-2" hydrate-on-idle />
      </div>
    </section>

    <!-- Exploration Section -->
    <knarr-exploration-zones v-if="showExploration" />

    <section id="map">
      <div class="grid grid-cols-12">
        <div id="artifact-1" class="h-90 bg-blue-500 rounded-lg shadow-sm col-span-8 col-start-5">
          Map
        </div>
      </div>
    </section>

    <section id="recruits">
      <div class="grid grid-cols-12 gap-2">
        <div v-for="i in 6" :id="`recruit-${i}`" :key="i" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2 first:col-start-1">
          Recruit {{ i }}
        </div>
      </div>
    </section>

    <!-- Player Hands -->
    <lazy-knarr-hands-modal hydrate-on-idle />

    <!-- Action Bar -->
    <knarr-actions />
  </nuxt-container>
</template>

<script lang="ts" setup>
import { usePlayerHandsComposable, useExplorationComposable } from '~/composables/knarr'

const tokens = ['bg-blue-200']

onMounted(() => {
  document.body.classList.add(...tokens)
})

onUnmounted(() => {
  document.body.classList.remove(...tokens)
})

/**
 * Hands
 */

const { togglePlayerHands } = usePlayerHandsComposable()

/**
 * Exploration
 */

const { showExploration, toggleExploration } = useExplorationComposable()

/**
 * Modifier Keys
 */

onKeyStroke('h', (e) => {
  e.preventDefault()
  togglePlayerHands()
})

onKeyStroke('e', (e) => {
  e.preventDefault()
  toggleExploration()
})
</script>
