<template>
  <div class="global-ranking h-screen">
    <nuxt-container class="space-y-4">
      <div class="text-center">
        <h1 class="text-6xl font-bold mb-6 uppercase">Global Rankings</h1>
      </div>

      {{ currentTab }}

      <nuxt-button to="/rankings/add">
        Add New Ranking
      </nuxt-button>

      <nuxt-tabs v-model="currentTab" :items="items" :default-value="0">
        <template #content="{ item }">
          <ranking-block class="bg-white rounded-xl p-10" />
        </template>
      </nuxt-tabs>
    </nuxt-container>
  </div>
</template>

<script setup lang="ts">
import type { TabsItem } from '@nuxt/ui'

const items: TabsItem[] = [
  {
    label: 'Global Rankings',
  },
  {
    label: "It's a Wonderful World",
  },
  {
    label: '7 Wonders Duel',
  },
  {
    label: 'Terraforming Mars',
  }
]

const tokens = [
  'bg-gradient-to-b',
  'from-primary-200',
  'via-primary-100',
  'to-primary-50'
]

onMounted(() => {
  document.body.classList.add(...tokens)
})

onUnmounted(() => {
  document.body.classList.remove(...tokens)
})

const currentTab = ref<number>(0)

const queryParams = useUrlSearchParams('history') as { ranking: string}
if (isDefined(queryParams.ranking)) {
  currentTab.value = useToNumber(queryParams.ranking).value || 0
}
</script>
