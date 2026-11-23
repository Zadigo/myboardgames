<template>
  <section id="flip-seven" class="h-screen px-20 space-y-2 mx-auto my-10">
    <div class="grid grid-cols-12 gap-2">
      <logic-flip-deck />
    </div>

    <!-- Tables -->
    <div v-if="isDefined(tableDetails)" class="grid grid-cols-2 grid-flow-row-dense gap-2">
      <logic-flip-player v-for="(tableClient, index) in players" :key="index" :index="index" :table-client="tableClient" />
    </div>

    <div v-if="!isConnected" class="rounded-lg bg-primary-200 w-full p-10 text-center space-y-4">
      <u-button @click="() => ws.open()">
        <icon name="lucide:refresh-cw" />
        Reconnect
      </u-button>
    </div>

    <!-- Dev -->
    <lazy-logic-flip-dev hydrate-on-idle />

    <!-- Actions -->
    <lazy-base-actions color="bg-flip-seven-200" hydrate-on-idle>
      <u-button>
        Action 1
      </u-button>

      <u-button>
        Action 2
      </u-button>
    </lazy-base-actions>
  </section>
</template>

<script lang="ts" setup>
// import { useFlipSevenLiveGameComposable } from '~/composables/_flipseven'

definePageMeta({
  layout: 'game'
})

/**
 * Websocket
 */

const { wsObject, isConnected } = useWebsiteWebsocket()
const { ws, tableDetails, players } = useFlipSevenLiveGameComposable(wsObject)

onMounted(() => {
  // joinTable()
})

/**
 * Utils
 */

useStyleTag('body { background-color: var(--color-flip-seven-200); }')
</script>
