<template>
  <nuxt-container id="knarr" class="h-screen space-y-2" as="section">
    <!-- Artifacts Section -->
    <section id="artifacts">
      <div class="grid grid-cols-12 auto-cols-max gap-1">
        <lazy-knarr-artifacts-card id="artifact-1" hydrate-on-idle />
        <lazy-knarr-artifacts-card id="artifact-2" hydrate-on-idle />
      </div>
    </section>

    <!-- Exploration Section -->
    <transition
      mode="out-in"
      enter-active-class="transition-all duration-800"
      leave-active-class="transition-all duration-800"
      enter-from-class="-translate-x-5 opacity-0"
      enter-to-class="translate-x-0 opacity-100"
      leave-from-class="translate-x-0 opacity-100"
      leave-to-class="-translate-x-10 opacity-0"
    >
      <knarr-exploration-zones v-if="showExploration" />
    </transition>

    <!-- Map -->
    <knarr-map />

    <!-- Recruits Section -->
    <knarr-recruits-grid />

    <!-- Player Hands -->
    <lazy-knarr-hands-modal hydrate-on-idle />

    <!-- Action Bar -->
    <knarr-actions />
  </nuxt-container>
</template>

<script lang="ts" setup>
import { usePlayerHandsComposable, useExplorationComposable } from '~/composables/knarr'

definePageMeta({
  layout: 'game'
})

/**
 * Background
 */

// useStyleTag('body { background-color: var(--color-primary-100); }')

// useStyleTag(`
//   html {
//     background-image: url('/images/back.jpg');
//     background-size: cover;
//     background-position: center;
//     background-repeat: no-repeat;
//     backdrop-filter: blur(20px);
//   }
// `)

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
