<template>
  <section id="some-game">
    <nuxt-button @click="createTable">
      Open Flip 7 Game
    </nuxt-button>

    <!-- <nuxt-button @click="create">
      Create Flip 7 Game
    </nuxt-button> -->

    <nuxt-button :to="`/games/flip-seven/live?table=${tableId}`" :disabled="!tableId">
      Join Flip 7 Game
    </nuxt-button>
  </section>
</template>

<script lang="ts" setup>
import { useFlipSevenGameComposable, useFlipSevenLiveGameComposable } from '~/composables/flipseven'

const { create, tableId } = useFlipSevenGameComposable()
const { createTable } = useFlipSevenLiveGameComposable()
</script>
