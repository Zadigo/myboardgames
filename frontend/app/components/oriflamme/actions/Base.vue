<template>
  <div class="top-0 w-auto">
    <template v-if="inQueue">
      <oriflamme-actions-button icon="lucide:badge-dollar-sign" @click="placeToken" />
      <oriflamme-actions-button icon="lucide:flip-horizontal-2" @click="reveal" />
      <oriflamme-actions-button icon="lucide:cloud-lightning" @click="applyEffect" />
    </template>
    <template v-else>
      <oriflamme-actions-button icon="lucide:hand" @click="placeCard" />
      <oriflamme-actions-button :class="{ selected: isSelected(card) }" icon="lucide:hand" @click="selectCards(card)" />
    </template>
  </div>
</template>

<script lang="ts" setup>
import { useOriflammeComposable, useOriflammeActionsComposable } from '~/composables/oriflamme'
import type { OriflammeCard } from '~/types'

const { inQueue = false } = defineProps<{
  card: OriflammeCard
  inQueue?: boolean
}>()

const { ws } = useOriflammeComposable()
const { reveal, applyEffect, placeToken, placeCard, selectCards, isSelected } = useOriflammeActionsComposable(ws)
</script>
