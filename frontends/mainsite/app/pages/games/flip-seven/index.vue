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
import { useFlipSevenLiveGameComposable } from '~/composables/flipseven'

const { createTable, tableId } = useFlipSevenLiveGameComposable()
</script>
