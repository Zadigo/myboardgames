<template>
  <section id="board">
    <div class="grid grid-rows-2 grid-cols-12 grid-flow-row-dense gap-2">
      <div class="col-span-12 bg-wonderful-blue-500 rounded-lg">
        Google
      </div>
      
      <div class="col-span-10 col-start-3 bg-wonderful-blue-500 rounded-lg">
        <lazy-wonderful-base-hand :cards="cardFixtures" />
      </div>
      
      <div class="col-span-2 row-span-2 bg-wonderful-violet-500 rounded-lg overflow-hidden">
        <div class="flex-row p-5 items-center">
          <div>
            <nuxt-img src="/wonderful/material/empires/aztec-empire-D.jpg" />
          </div>
        </div>
      </div>
      
      <div class="col-span-10 col-start-3 bg-wonderful-red-500 rounded-lg">
        <lazy-wonderful-base-drafted :cards="[]" />
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { cardFixtures } from '~/data/fixtures'

definePageMeta({
  layout: 'wonderful',
})

onMounted(() => {
  document.body.classList.add('bg-wonderful-blue-500/20')
})

const { wsObject } = useGameWebsocket()
</script>
