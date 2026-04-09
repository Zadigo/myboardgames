<template>
  <section id="some-game">
    Is open {{ isOpen }}

    <nuxt-button @click="createTable">
      Open Oriflamme Game
    </nuxt-button>

    <nuxt-button color="error" @click="quitTable">
      Close Oriflamme Game
    </nuxt-button>

    <nuxt-button :to="`/games/oriflamme/live?table=${tableId}`" :disabled="!tableId">
      Join Oriflamme Game
    </nuxt-button>
  </section>
</template>

<script lang="ts" setup>
import { useOriflammeComposable } from '~/composables/oriflamme'

const { createTable, quitTable, tableId, isOpen } = useOriflammeComposable()
</script>
