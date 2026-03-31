<template>
  <dev-only>
    <div ref="devEl" class="fixed top-0 left-0 rounded-lg bg-primary-200/20 backdrop-blur-2xl shadow-md p-10 cursor-grab" :style="style">
      <nuxt-button @click="getDeck">
        Get deck
      </nuxt-button>

      <div class="my-2">
        <p class="font-bol">
          Table Details
        </p>

        <pre class="min-w-50 max-w-80 max-h-60 overflow-scroll rounded-2xl inset-shadow-2xs bg-slate-100/30 p-4">
          {{ tableDetails }}
        </pre>
      </div>
    </div>
  </dev-only>
</template>

<script lang="ts" setup>
import { useFlipSevenLiveGameComposable } from '~/composables/flipseven'

const { tableDetails, getDeck } = useFlipSevenLiveGameComposable()

const devEl = useTemplateRef('devEl')
const { style } = useDraggable(devEl, {
  initialValue: { x: 0, y: 0 }
})
</script>
