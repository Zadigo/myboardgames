<template>
  <div>
    <component :is="activeComponent" :rpc-hash="query.r" />
  </div>
</template>


<script lang="ts" setup>
import FlipGamePage from '~/components/logic/flip/GamePage.vue'
import KnarrGamePage from '~/components/logic/knarr/GamePage.vue'
import OriflammeGamePage from '~/components/logic/oriflamme/GamePage.vue'

const route = useRoute()
const query = useUrlSearchParams('history') as { r: string }

let activeComponent: Component

switch (route.params.id) {
  case 'knarr':
    activeComponent = KnarrGamePage
    break
  case 'oriflamme':
    activeComponent = OriflammeGamePage
    break
  case 'flip-seven':
    activeComponent = FlipGamePage
    break
}
</script>
