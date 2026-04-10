<template>
  <div ref="devEl" :style="style" class="fixed top-0 left-0 rounded-lg max-w-100 h-auto p-5 bg-primary-100/10 backdrop-blur-3xl shadow-2xl z-50">
    <nuxt-badge :label="isOpen ? 'Connected' : 'Disconnected'" :color="isOpen ? 'success' : 'error'" />

    <div class="w-full space-x-2 space-y-2 mt-5">
      <nuxt-button block @click="openConnection">
        Open
      </nuxt-button>

      <nuxt-button block @click="createGame">
        Create game
      </nuxt-button>

      <nuxt-button block :to="`/games/oriflamme/live?table=${tableUuid}`" :disabled="!tableUuid" @click="startGame">
        Join
      </nuxt-button>

      <nuxt-button block color="error" @click="quitTable">
        Close
      </nuxt-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useOriflammeComposable } from '~/composables/oriflamme'

const { isOpen, createGame, openConnection, quitTable, startGame, tableUuid } = useOriflammeComposable()

const devEl = useTemplateRef('devEl')
const { style } = useDraggable(devEl, {
  initialValue: { x: 20, y: 20 }
})
</script>
