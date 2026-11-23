<template>
  <div ref="devEl" :style="style" class="fixed top-0 left-0 rounded-lg max-w-100 h-auto p-5 bg-primary-100/10 backdrop-blur-3xl shadow-2xl z-50">
    <u-badge :label="isOpen ? 'Connected' : 'Disconnected'" :color="isOpen ? 'success' : 'error'" />

    <div class="w-full space-x-2 space-y-2 mt-5">
      <u-button block @click="openConnection">
        Open
      </u-button>

      <u-button block @click="createGame">
        Create game
      </u-button>

      <u-button block :to="`/games/oriflamme/live?table=${tableUuid}`" :disabled="!tableUuid" @click="startGame">
        Join
      </u-button>

      <u-button block color="error" @click="quitTable">
        Close
      </u-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useOriflammeComposable } from '~/composables/_oriflamme'

const { isOpen, createGame, openConnection, quitTable, startGame, tableUuid } = useOriflammeComposable()

const windowLocation = useLocalStorage<{ x: number, y: number }>('dev-location', { x: 20, y: 20 }, {
  serializer: {
    write: value => JSON.stringify(value),
    read: raw => JSON.parse(raw)
  }
})

const devEl = useTemplateRef('devEl')
const { style, x, y } = useDraggable(devEl, {
  initialValue: windowLocation
})

watch([x, y], ([newX, newY]) => {
  windowLocation.value = { x: newX, y: newY }
})
</script>
