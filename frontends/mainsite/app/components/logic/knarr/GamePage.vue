<template>
  <u-container id="knarr" class="h-screen space-y-2" as="section">
    <!-- Artifacts Section -->
    <section id="artifacts">
      <div class="grid grid-cols-12 auto-cols-max gap-1">
        <lazy-logic-knarr-artifacts-card id="artifact-1" hydrate-on-idle />
        <lazy-logic-knarr-artifacts-card id="artifact-2" hydrate-on-idle />
      </div>
    </section>

    <!-- Exploration Section -->
    <transition mode="out-in" enter-active-class="transition-all duration-800" leave-active-class="transition-all duration-800" enter-from-class="-translate-x-5 opacity-0" enter-to-class="translate-x-0 opacity-100" leave-from-class="translate-x-0 opacity-100" leave-to-class="-translate-x-10 opacity-0">
      <lazy-logic-knarr-exploration-zones v-if="showExploration" />
    </transition>

    <!-- Map -->
    <logic-knarr-map />

    <!-- Recruits Section -->
    <logic-knarr-recruits-grid />

    <!-- Player Hands -->
    <lazy-logic-knarr-hands-modal hydrate-on-idle />

    <!-- Action Bar -->
    <logic-knarr-actions />
  </u-container>
</template>

<script lang="ts" setup>
import { usePlayerHandsComposable, useExplorationComposable } from '~/composables/_knarr'

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
