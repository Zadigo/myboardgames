<template>
  <section id="exploration" class="space-y-2">
    <div id="zone-a" class="grid grid-cols-12 auto-cols-max gap-2">
      <knarr-exploration-card v-for="i in 4" :id="`exploration-${i}`" :key="i" />
    </div>

    <div id="zone-b" class="grid grid-cols-12 auto-cols-max gap-2">
      <knarr-exploration-card v-for="i in 4" :id="`exploration-${i}`" :key="i" />
    </div>
  </section>
</template>
