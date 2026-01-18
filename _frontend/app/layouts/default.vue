<template>
  <section id="site">
    <nav>
      <ul>
        <li v-for="link in links" :key="link.url">
          <nuxt-link :to="link.url">{{ link.title }}</nuxt-link>
        </li>
      </ul>
    </nav>

    <slot />
  </section>
</template>

<script lang="ts" setup>
const links = [
  { title: 'Home', url: '/' },
  { title: 'Games', url: '/games' },
  { title: 'Rankings', url: '/rankings' }
]
</script>
