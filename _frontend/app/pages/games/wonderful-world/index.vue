<template>
  <div class="ranking">
    <header-block image="/wonderful.jpeg" />
    <nuxt-link to="/games/wonderful-world/ranking">Go to Ranking</nuxt-link>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  title: 'Index - Wonderful World',
  layout: 'game'
})
</script>
