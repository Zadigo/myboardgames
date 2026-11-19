<template>
  <section id="game">
    <main>
      <slot />
    </main>

    <!-- Modals -->
    <modals-help />

    <!-- Fabs -->
    <u-button size="xl" class="rounded-full fixed bottom-4 right-4" @click="() => { toggleShowHelpModal() }">
      <icon name="lucide:circle-question-mark" /> {{ showHelpModal }}
    </u-button>
  </section>
</template>

<script lang="ts" setup>
/**
 * Help Modal State
 */

const { toggleShowHelpModal, showHelpModal } = useHelpComposable()
</script>
