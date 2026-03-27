<template>
  <lazy-nuxt-modal v-model:open="showHelpModal" hydrate-on-visible>
    <template #body>
      <h2 class="text-2xl font-bold mb-4">Help</h2>
      <p class="mb-2">This is the help modal content. You can add instructions, tips, or any information you want to assist users.</p>
    </template>
  </lazy-nuxt-modal>
</template>

<script lang="ts" setup>
const { showHelpModal } = useHelpComposable()
</script>
