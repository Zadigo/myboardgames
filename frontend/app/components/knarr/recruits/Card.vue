<template>
  <div ref="cardEl" :class="theme" class="relative overflow-hidden hover:scale-101 transition-all ease-in-out duration-300 h-80 bg-primary-700 rounded-lg shadow-sm col-span-2 first:col-start-1 cursor-pointer" @click="() => { setIsClicked() }">
    <nuxt-popover :open-delay="300" mode="hover" hover>
      <div class="h-full w-full mx-auto" />

      <template #content>
        <div class="p-10">
          Recruit Card {{ isClicked }}
        </div>
      </template>
    </nuxt-popover>
  </div>
</template>

<script lang="ts" setup>
const cardEl = useTemplateRef('cardEl')

const [isClicked, setIsClicked] = useToggle()

const theme = computed(() => {
  return [
    {
      '-translate-y-2 border-5 border-dashed border-primary-500': isClicked.value
    }
  ]
})

onClickOutside(cardEl, () => {
  setIsClicked(false)
})
</script>
