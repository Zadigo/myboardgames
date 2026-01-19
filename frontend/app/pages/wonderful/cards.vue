<template>
  <section id="cards">
    <h1>Cards</h1>

    <nuxt-container>
      <nuxt-card>
        <nuxt-input v-model="search" placeholder="Search cards..." />
        <!-- <nuxt-switch v-model="showOnlyFavorites" label="Show only favorites" /> -->
        <nuxt-switch v-model="hasConstructionBonus" label="Construction bonus tokens" />
        <nuxt-switch v-model="hasCharacterBonus" label="Character bonus tokens" />
        <nuxt-switch v-model="hasMaterialBonus" label="Material bonus tokens" />
      </nuxt-card>

      <div class="grid grid-cols-4 gap-3">
        <transition-group
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
        >
          <nuxt-card v-for="(card, idx) in filteredCards" :key="idx">
           <div class="rounded-lg h-full w-full">
             <nuxt-img :src="`wonderful/material/developments/${card?.category.toLowerCase()}/${card?.image}`" :alt="card.name" :width="325" />
           </div>
 
           <template #footer>
             <div class="p-2">
               <h2 class="text-lg font-bold">{{ card.name }}</h2>
             </div>
           </template>
          </nuxt-card>
        </transition-group>
      </div>
    </nuxt-container>
  </section>
</template>

<script lang="ts" setup>
import type { BaseCard, Arrayable } from '~/types'

definePageMeta({
  layout: 'default',
})

type CardsResponseApi = {
  data: Arrayable<BaseCard>
} 

const { load } = useMemoize(async () => {
  const cards = await $fetch<CardsResponseApi>('/cards', {
    method: 'GET',
    baseURL: useRuntimeConfig().public.apiDomain
  })

  return cards
})

const cards = await load()

const search = ref('')
const hasConstructionBonus = ref(false)
const hasCharacterBonus = ref(false)
const hasMaterialBonus = ref(false)

const searchedCards = computed(() => {
  return cards.data.filter((card) => card.name.toLowerCase().includes(search.value.toLowerCase()))
})

const filteredCards = computed(() => {
  return searchedCards.value.filter((card) => {
    if (hasConstructionBonus.value && !card.has_construction_bonus) {
      return false
    }

    if (hasCharacterBonus.value && !card.has_character_token_bonus) {
      return false
    }
    
    if (hasMaterialBonus.value && !card.isDrafted) {
      return false
    }
    
    return true
  })
})
</script>
