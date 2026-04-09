import type { OriflammeCard } from '~/types'

const [useOriflammeActionsComposable, _useOriflameeActionsStore] = createInjectionState((ws: Ref<WebSocket | undefined>) => {
  const selectedCards = ref<string[]>([])

  function reveal() {
    ws.value?.send('a')
  }

  function applyEffect() {

  }

  function placeToken() {

  }

  function placeCard() {

  }

  function selectCards(card: OriflammeCard) {
    if (selectedCards.value.includes(card.Uuid)) {
      selectedCards.value = selectedCards.value.filter(uuid => uuid !== card.Uuid)
    } else {
      selectedCards.value.push(card.Uuid)
    }
  }

  function _isSelected(card: OriflammeCard) {
    return selectedCards.value.includes(card.Uuid)
  }

  const isSelected = reactify(_isSelected)

  return {
    selectedCards,
    reveal,
    applyEffect,
    placeToken,
    placeCard,
    selectCards,
    isSelected
  }
})

export { useOriflammeActionsComposable }

export function useOriflameeActionsStore() {
  const store = _useOriflameeActionsStore()
  if (!store) {
    throw new Error('useOriflameeActionsStore must be used within a useOriflammeActionsComposable')
  }
  return store
}
