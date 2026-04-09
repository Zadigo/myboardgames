<template>
  <div class="absolute top-0 left-0 w-full h-full space-y-2">
    <div class="flex flex-col content-center gap-2 space-x-2">
      <template v-if="inQueue">
        <oriflamme-actions-button icon="lucide:badge-dollar-sign" @click="placeToken" />
        <oriflamme-actions-button icon="lucide:flip-horizontal-2" @click="reveal" />
        <oriflamme-actions-button icon="lucide:cloud-lightning" @click="applyEffect" />
      </template>
      <template v-else>
        <oriflamme-actions-button icon="lucide:move-left" @click="() => placeCard(card, 'left')" />
        <oriflamme-actions-button icon="lucide:move-right" @click="() => placeCard(card, 'right')" />
        <oriflamme-actions-button icon="lucide:hand" @click="() => selectCards(card)" />
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useOriflameeActionsStore } from '~/composables/oriflamme'
import type { OriflammeCard } from '~/types'

const { inQueue = false } = defineProps<{
  card: OriflammeCard
  inQueue?: boolean
}>()

const { reveal, applyEffect, placeToken, placeCard, selectCards } = useOriflameeActionsStore()
</script>
