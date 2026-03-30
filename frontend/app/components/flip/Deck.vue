<template>
  <div class="bg-flip-seven-300/70 backdrop-blur-xl col-span-3 flex justify-center p-5 rounded-lg gap-2">
    <div id="hidden-card" class="h-50 w-30 bg-flip-seven-900 rounded-lg cursor-pointer hover:scale-105 transition-all ease-in-out duration-200" @click="newCard">
      Another 1
    </div>

    <div id="card-reveal" class="h-50 w-30 bg-flip-seven-100 rounded-lg">
      Another 1
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useFlipSevenGameComposable, useFlipSevenLiveGameComposable } from '~/composables/flipseven'

const { tableId } = useFlipSevenGameComposable()
const { wsObject } = useFlipSevenLiveGameComposable(tableId)

const { encode } = useWebsocketMessage()

/**
 * Flip New Card
 */

function newCard() {
  wsObject.send(encode({ action: 'flip_card', playerId: "Player 1" }))
}
</script>
