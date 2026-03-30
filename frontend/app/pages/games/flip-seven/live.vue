<template>
  <section id="flip-seven" class="h-screen px-20 space-y-2 mx-auto my-10">
    <div class="grid grid-cols-12 gap-2">
      <flip-deck />
    </div>

    <!-- Tables -->
    <div v-if="isDefined(tableDetails)" class="grid grid-cols-2 grid-flow-row-dense gap-2">
      <flip-player v-for="(tableClient, index) in tableDetails.clients" :key="index" :index="index" :table-client="tableClient" />
    </div>

    <div v-if="!isConnected" class="rounded-lg bg-primary-200 w-full p-10 text-center space-y-4">
      <nuxt-button @click="wsObject.open()">
        <icon name="lucide:refresh-cw" />
        Reconnect
      </nuxt-button>
    </div>

    <!-- Dev -->
    <lazy-flip-dev hydrate-on-idle />

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

/**
 * Websocket
 */

const { tableId } = useFlipSevenGameComposable()
const { wsObject, tableDetails, isConnected } = useFlipSevenLiveGameComposable(tableId)
wsObject.open()

/**
 * Utils
 */

useStyleTag('body { background-color: var(--color-flip-seven-200); }')
</script>
