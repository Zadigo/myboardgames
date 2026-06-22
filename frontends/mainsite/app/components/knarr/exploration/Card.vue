<template>
  <div :id="`exploration-${id}`" ref="cardEl" :class="theme" class="h-40 bg-primary-700 rounded-lg shadow-sm col-span-3 first:col-start-1 cursor-pointer relative" @click="() => { setIsClicked() }">
    <nuxt-popover :open-delay="300" mode="hover" class="w-70" hover>
      <div class="h-30 w-30 mx-auto" />

      <template #content>
        <div class="p-10">
          Exploration {{ id }}
        </div>
      </template>
    </nuxt-popover>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string
}>()

const [isClicked, setIsClicked] = useToggle()

const theme = computed(() => {
  return [
    {
      'shadow-xl border-4 border-primary-50': isClicked.value
    }
  ]
})

const cardEl = useTemplateRef('cardEl')

onClickOutside(cardEl, () => {
  setIsClicked(false)
})
</script>
