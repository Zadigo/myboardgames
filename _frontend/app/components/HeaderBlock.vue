<template>
  <section id="header" ref="headerEl" class="bg-primary p-20 h-[80vh] text-white relative flex items-center">
    <div class="absolute inset-0 opacity-80 bg-gradient-to-t from-black to-transparent" />
    <div class="z-40">
      <h1 class="text-8xl font-bold">Dune imperium</h1>
      <p class="text-2xl">A game of strategy and resource management</p>
      <nuxt-button size="xl" color="info" variant="solid" to="/games/dune-imperium/ranking" class="mt-10">
        Rankings
      </nuxt-button>
    </div>
  </section>
</template>

<script setup lang="ts">
const { image } = defineProps<{ image: string }>()
const headerEl = useTemplateRef('headerEl')

onMounted(() => {
  if (headerEl.value) {
    headerEl.value.style.backgroundImage = `url('${image}')`
  }
})
</script>
