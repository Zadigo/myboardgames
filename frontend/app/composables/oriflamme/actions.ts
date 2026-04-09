import type { OriflammeCard } from '~/types'

export function useOriflammeActionsComposable(ws: Ref<WebSocket | undefined>) {
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

  function isSelected(card: OriflammeCard) {
    return selectedCards.value.includes(card.Uuid)
  }

  return {
    reveal,
    applyEffect,
    placeToken,
    placeCard,
    selectCards,
    isSelected
  }
}
