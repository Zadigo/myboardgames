<template>
  <section id="exploration" class="space-y-2">
    <div id="zone-a" class="grid grid-cols-12 auto-cols-max gap-2">
      <div v-for="i in 4" :id="`exploration-${i}`" :key="i" class="h-40 bg-blue-400 rounded-lg shadow-sm col-span-3 first:col-start-1">
        Exploration {{ i }}
      </div>
    </div>

    <div id="zone-b" class="grid grid-cols-12 auto-cols-max gap-2">
      <div v-for="i in 4" :id="`exploration-${i}`" :key="i" class="h-40 bg-blue-400 rounded-lg shadow-sm col-span-3 first:col-start-1">
        Exploration {{ i }}
      </div>
    </div>
  </section>
</template>
