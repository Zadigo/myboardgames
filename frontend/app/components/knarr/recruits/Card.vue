<template>
  <div ref="cardEl" :class="theme" class="hover:scale-101 transition-all ease-in-out duration-300 h-80 bg-blue-400 rounded-lg shadow-sm col-span-2 first:col-start-1 cursor-pointer" @click="() => { setIsClicked() }">
    Recruit Card {{ isClicked }}
  </div>
</template>

<script lang="ts" setup>
const cardEl = useTemplateRef('cardEl')

const [isClicked, setIsClicked] = useToggle()

const theme = computed(() => {
  return [
    {
      '-translate-y-2 border-5 border-dashed border-blue-200': isClicked.value
    }
  ]
})

onClickOutside(cardEl, () => {
  setIsClicked(false)
})
</script>
