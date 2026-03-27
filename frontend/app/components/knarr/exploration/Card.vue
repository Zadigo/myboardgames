<template>
  <div v-for="i in 4" :id="`exploration-${i}`" :key="i" class="h-40 bg-blue-400 rounded-lg shadow-sm col-span-3 first:col-start-1">
    Exploration {{ i }}
  </div>
</template>
