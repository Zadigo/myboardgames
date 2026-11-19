<template>
  <section id="some-game">
    <u-button @click="createTable">
      Open Flip 7 Game
    </u-button>

    <!-- <u-button @click="create">
      Create Flip 7 Game
    </u-button> -->

    <u-button :to="`/games/flip-seven/live?table=${tableId}`" :disabled="!tableId">
      Join Flip 7 Game
    </u-button>
  </section>
</template>

<script lang="ts" setup>
import { useFlipSevenLiveGameComposable } from '~/composables/flipseven'

const { createTable, tableId } = useFlipSevenLiveGameComposable()
</script>
