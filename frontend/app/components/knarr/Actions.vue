<template>
  <div id="actions" ref="actionEl" class="absolute top-0 right-0 bg-primary-500/20 backdrop-blur-2xl w-auto h-auto rounded-bl-2xl py-5 px-3 z-50 flex flex-col gap-2">
    <nuxt-button variant="subtle" @click="() => { togglePlayerHands() }">
      <icon name="lucide:hand" />
      Hands
    </nuxt-button>

    <nuxt-button variant="subtle" @click="() => { toggleExploration() }">
      <icon name="lucide:map" />
      Exploration
    </nuxt-button>

    <nuxt-button variant="subtle" @click="() => { toggleExploration() }">
      <icon name="lucide:message-circle" />
      Discussion
    </nuxt-button>
  </div>
</template>

<script lang="ts" setup>
import { useExplorationStore, usePlayerHandsStore } from '~/composables/knarr'

const { togglePlayerHands } = usePlayerHandsStore()

const actionEl = useTemplateRef('actionEl')

// const { style } = useDraggable(actionEl, {
//   initialValue: { x: 1400, y: 20 }
// })

const { toggleExploration } = useExplorationStore()
</script>
