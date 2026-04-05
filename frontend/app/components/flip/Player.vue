<template>
  <div :class="{ 'bg-flip-seven-300/70': index === 2, 'bg-flip-seven-200/70': index !== 2 }" class="w-auto gap-2 shadow-sm backdrop-blur-xl rounded-lg p-5 relative flex flex-col">
    <div id="infos" class="flex gap-1 place-self-end">
      <div id="score" class="p-2 rounded-lg bg-[#fffceb] font-semibold text-center flex items-center">
        {{ liveScore }}
      </div>

      <div id="score" class="p-2 rounded-lg bg-[#fffceb] font-bold text-center flex items-center gap-2">
        <nuxt-avatar />
        <span>{{ tableClient.username }}</span>
      </div>
    </div>

    <div class="flex gap-2">
      <flip-card v-for="card in tableClient.cards" :key="card.value" :card="card" class="h-40 w-30 bg-flip-seven-900 rounded-lg" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { TransitionPresets } from '@vueuse/core'
import { TEST_USERNAME, useFlipSevenLiveGameComposable } from '~/composables/flipseven'
import type { TableClientDetail } from '~/types'

defineProps<{
  index: number
  tableClient: TableClientDetail
}>()

/**
 * Live Score
 */

const { tableDetails } = useFlipSevenLiveGameComposable()

const _liveScore = computed(() => tableDetails.value?.deck?.filter((card) => {
  return card.owner === TEST_USERNAME
}).reduce((acc, card) => {
  return acc + card.value
}, 0) || 0)

const liveScore = useTransition(_liveScore, {
  duration: 500,
  easing: TransitionPresets.easeInOutBack
})
</script>
