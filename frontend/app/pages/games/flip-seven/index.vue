<template>
  <section id="some-game">
    <nuxt-button @click="create">
      Create Flip 7 Game
    </nuxt-button>

    <nuxt-button :to="`/games/flip-seven/${tableId}`" :disabled="!tableId">
      Join Flip 7 Game
    </nuxt-button>
  </section>
</template>

<script lang="ts" setup>
import { useFlipSevenGameComposable } from '~/composables/flipseven'

const { create, tableId } = useFlipSevenGameComposable()
</script>
