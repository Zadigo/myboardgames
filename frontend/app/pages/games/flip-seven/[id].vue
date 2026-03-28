<template>
  <section id="flip-seven" class="h-screen px-20 space-y-2 mx-auto" as="section">
    <div class="grid grid-cols-12 gap-2">
      <div class="bg-flip-seven-300/70 backdrop-blur-xl col-span-3 flex justify-center p-5 rounded-lg gap-2">
        <div id="hidden-card" class="h-50 w-30 bg-flip-seven-900 rounded-lg">
          Another 1
        </div>

        <div id="current-card" class="h-50 w-30 bg-flip-seven-900 rounded-lg">
          Another 1
        </div>
      </div>
    </div>

    <!-- Tables -->
    <div class="grid grid-cols-2 grid-flow-row-dense gap-2">
      <div v-for="i in 12" :key="i" :class="{ 'bg-flip-seven-300/70': i === 2, 'bg-flip-seven-200/70': i !== 2 }" class="w-auto gap-2 shadow-sm backdrop-blur-xl rounded-lg p-5 relative flex flex-col">
        <div id="infos" class="flex gap-1 place-self-end">
          <div id="score" class="p-2 rounded-lg bg-[#fffceb] font-semibold text-center flex items-center">
            15
          </div>

          <div id="score" class="p-2 rounded-lg bg-[#fffceb] font-bold text-center flex items-center gap-2">
            <nuxt-avatar />
            <span>My username</span>
          </div>
        </div>

        <div class="flex gap-2">
          <flip-card v-for="j in 7" :key="j" class="h-40 w-30 bg-flip-seven-900 rounded-lg" />
        </div>
      </div>
    </div>

    <!-- Actions -->
    <lazy-base-actions color="bg-flip-seven-200" hydrate-on-idle>
      <nuxt-button>
        Action 1
      </nuxt-button>

      <nuxt-button>
        Action 2
      </nuxt-button>
    </lazy-base-actions>
  </section>
</template>

<script lang="ts" setup>
definePageMeta({
  layout: 'game'
})

useStyleTag('body { background-color: var(--color-flip-seven-100); }')

// Secondary: #fade34

// Test Websocket

const ws = useWebSocket('ws://127.0.0.1:9000/ws/flip-seven')
ws.open()
</script>
