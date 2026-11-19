<template>
  <section id="some-game">
    Is open {{ isOpen }} Player: {{ playerUuid }} Table: {{ tableUuid }}

    <u-button @click="openConnection">
      Open connection
    </u-button>

    <u-button @click="createGame">
      Create game
    </u-button>

    <u-button color="error" @click="quitTable">
      Close Oriflamme Game
    </u-button>

    <u-button :to="`/games/oriflamme/live?table=${tableUuid}`" :disabled="!tableUuid" @click="startGame">
      Join Oriflamme Game
    </u-button>
  </section>
</template>

<script lang="ts" setup>
import { useOriflammeComposable } from '~/composables/oriflamme'

const { createGame, openConnection, quitTable, startGame, tableUuid, isOpen, playerUuid } = useOriflammeComposable()
</script>

