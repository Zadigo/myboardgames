<template>
  <nuxt-container id="knarr" class="my-10 space-y-2" as="section">
    <!-- Artifacts Section -->
    <section id="artifacts">
      <div class="grid grid-cols-12 auto-cols-max gap-1">
        <knarr-artifacts-card id="artifact-1" />
        <knarr-artifacts-card id="artifact-2" />
      </div>
    </section>

    <!-- Exploration Section -->
    <knarr-exploration-zones v-if="showExploration" />

    <section id="map">
      <div class="grid grid-cols-12">
        <div id="artifact-1" class="h-90 bg-blue-500 rounded-lg shadow-sm col-span-8 col-start-5">
          Map
        </div>
      </div>
    </section>

    <section id="recruits">
      <div class="grid grid-cols-12 gap-2">
        <div v-for="i in 6" :id="`recruit-${i}`" :key="i" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2 first:col-start-1">
          Recruit {{ i }}
        </div>
      </div>
    </section>

    <!-- Player Hands -->
    <section v-if="showPlayerHands" id="players-hands" class="fixed h-200 bg-blue-500/20 w-full bottom-0 left-0 shadow-lg backdrop-blur-2xl z-50">
      <div class="relative">
        <div class="py-5 px-20 bg-blue-100/50 backdrop-blur-2xl" @click="() => { togglePlayerHands() }">
          <div class="p-5 rounded-lg hover:bg-primary-100/20 transition-colors ease-in-out duration-300 cursor-pointer">
            <h2 class="font-bold text-2xl">
              Player Name
            </h2>
          </div>
        </div>

        <div class="p-20">
          <div class="absolute left-5 top-1/2 -translate-y-1/2">
            <nuxt-button id="previous-player">
              <icon name="lucide:arrow-left" />
            </nuxt-button>
          </div>

          <div class="grid grid-cols-12 gap-2">
            <div id="player-1-hand" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2 col-start-2">
              Player 1 Hand
            </div>
  
            <div id="player-2-hand" class="h-80 bg-blue-400 rounded-lg shadow-sm col-span-2">
              Player 2 Hand
            </div>
          </div>

          <div class="absolute right-5 top-1/2 -translate-y-1/2">
            <nuxt-button id="next-player">
              <icon name="lucide:arrow-right" />
            </nuxt-button>
          </div>
        </div>
      </div>
    </section>

    <!-- Action Bar -->
    <div id="actions" ref="actionEl" class="fixed bg-primary-500/20 backdrop-blur-2xl shadow-lg w-auto h-auto rounded-xl py-5 px-3 z-50 flex flex-col gap-2" :style="style">
      <nuxt-button variant="subtle" @click="() => { togglePlayerHands() }">
        <icon name="lucide:hand" />
        Hands
      </nuxt-button>

      <nuxt-button variant="subtle" @click="() => { toggleExploration() }">
        <icon name="lucide:map" />
        Exploration
      </nuxt-button>
    </div>
  </nuxt-container>
</template>

<script lang="ts" setup>
const tokens = ['bg-blue-200']

onMounted(() => {
  document.body.classList.add(...tokens)
})

onUnmounted(() => {
  document.body.classList.remove(...tokens)
})

const [showPlayerHands, togglePlayerHands] = useToggle(false)

const actionEl = useTemplateRef('actionEl')

const { style } = useDraggable(actionEl, {
  initialValue: { x: 1400, y: 20 }
})

const [showExploration, toggleExploration] = useToggle(false)
</script>
