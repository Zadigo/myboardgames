<template>
  <div class="top-0 w-auto">
    <oriflamme-actions-button icon="lucide:badge-dollar-sign" @click="placeToken" />
    <oriflamme-actions-button icon="lucide:flip-horizontal-2" @click="reveal" />
    <oriflamme-actions-button icon="lucide:cloud-lightning" @click="applyEffect" />
  </div>
</template>

<script lang="ts" setup>
import { useOriflammeComposable, useOriflammeActionsComposable } from '~/composables/oriflamme'

const { ws } = useOriflammeComposable()
const { reveal, applyEffect, placeToken } = useOriflammeActionsComposable(ws)
</script>
