<template>
  <section id="recruits" class="mt-5">
    <div class="grid grid-cols-12 gap-2">
      <knarr-recruits-card v-for="i in 6" :id="`recruit-${i}`" :key="i" />
    </div>
  </section>
</template>

<script lang="ts" setup>

</script>
