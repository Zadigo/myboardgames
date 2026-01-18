<template>
  <nuxt-card>
    <nuxt-badge v-if="index === 0" variant="solid">Winner</nuxt-badge>
    <nuxt-badge v-else-if="index === 1" variant="soft">Runner-up</nuxt-badge>
    <nuxt-badge v-else-if="index === 2" variant="soft">Second Runner-up</nuxt-badge>
    <nuxt-input v-model="player.gamertag" placeholder="Player gamertag" />
    <nuxt-input v-model="player.firstname" placeholder="Player firstname" />
    <nuxt-input v-model="player.lastname" placeholder="Player lastname" />
    <nuxt-select v-model="player.gender" :items="['W', 'M']" placeholder="Select a gender" />
    <nuxt-input-number v-model="player.points" :min="1" :max="50" placeholder="Player points" />
    <nuxt-select v-model="player.game" :items="Array.from(gameNamesFixtures)" placeholder="Select a game" />
    <nuxt-switch label="Winner" />
    <nuxt-switch label="Runner-up" />
    <nuxt-switch label="Second Runner-up" />
  </nuxt-card>
</template>

<script setup lang="ts">
import { gameNamesFixtures } from '~/data'
import type { NewPlayer } from '~/types'

const props = defineProps<{ modelValue: NewPlayer, index: number, count: number }>()
const emit = defineEmits<{ 'update:modelValue': NewPlayer }>()
const player = useVModel(props, 'modelValue', emit)
</script>
