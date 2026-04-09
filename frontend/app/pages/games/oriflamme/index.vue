<template>
  <section id="some-game">
    Is open {{ isOpen }} Player: {{ playerUuid }} Table: {{ tableUuid }}

    <nuxt-button @click="openConnection">
      Open connection
    </nuxt-button>

    <nuxt-button @click="createGame">
      Create game
    </nuxt-button>

    <nuxt-button color="error" @click="quitTable">
      Close Oriflamme Game
    </nuxt-button>

    <nuxt-button :to="`/games/oriflamme/live?table=${tableUuid}`" :disabled="!tableUuid">
      Join Oriflamme Game
    </nuxt-button>
  </section>
</template>

<script lang="ts" setup>
import { useOriflammeComposable } from '~/composables/oriflamme'

const { createGame, openConnection, quitTable, tableUuid, isOpen, playerUuid } = useOriflammeComposable()
</script>
