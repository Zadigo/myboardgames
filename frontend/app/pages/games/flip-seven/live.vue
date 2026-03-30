<template>
  <section id="flip-seven" class="h-screen px-20 space-y-2 mx-auto" as="section">
    <div class="grid grid-cols-12 gap-2">
      <div class="bg-flip-seven-300/70 backdrop-blur-xl col-span-3 flex justify-center p-5 rounded-lg gap-2">
        <div id="hidden-card" class="h-50 w-30 bg-flip-seven-900 rounded-lg">
          Another 1
        </div>

        <div id="current-card" class="h-50 w-30 bg-flip-seven-900 rounded-lg">
          Another 1
        </div>
      </div>
    </div>

    <nuxt-button @click="wsObject.send(encode({ action: 'get_deck' }))">
      Get deck
    </nuxt-button>

    <!-- Tables -->
    <div v-if="isDefined(tableDetails)" class="grid grid-cols-2 grid-flow-row-dense gap-2">
      <flip-player v-for="(tableClient, index) in tableDetails.clients" :key="index" :index="index" :table-client="tableClient" />
    </div>

    <!-- Actions -->
    <lazy-base-actions color="bg-flip-seven-200" hydrate-on-idle>
      <nuxt-button>
        Action 1
      </nuxt-button>

      <nuxt-button>
        Action 2
      </nuxt-button>
    </lazy-base-actions>
  </section>
</template>

<script lang="ts" setup>
import { useFlipSevenGameComposable, useFlipSevenLiveGameComposable } from '~/composables/flipseven'

definePageMeta({
  layout: 'game'
})

const { encode } = useWebsocketMessage()

/**
 * Websocket
 */

const { tableId } = useFlipSevenGameComposable()
const { wsObject, tableDetails } = useFlipSevenLiveGameComposable(tableId)
wsObject.open()

/**
 * Utils
 */

useStyleTag('body { background-color: var(--color-flip-seven-100); }')
</script>
