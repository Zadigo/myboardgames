<template>
  <div class="ranking">
    <nuxt-container>
      <ranking-block />
    </nuxt-container>
  </div>
</template>

<script setup lang="ts">
onMounted(() => {
  document.body.classList.add('bg-primary', 'text-white')
})

onUnmounted(() => {
  document.body.classList.remove('bg-primary', 'text-white')
})
</script>
