import type { Nullable, OriflammeCard } from '~/types'
import type { ActionOptions } from '.'

const [ useOriflammeActionsComposable, _useOriflameeActionsStore ] = createInjectionState((ws: Ref<WebSocket | undefined>, tableUuid: MaybeRefOrGetter<Nullable<string>>, playerUuid: MaybeRefOrGetter<Nullable<string>>) => {
  const selectedCards = ref<string[]>([])

  const { encode } = useWWebsocketMessages2()

  function reveal() {
    ws.value?.send('a')
  }

  function applyEffect() {

  }

  function placeToken() {

  }

  function placeCard(card: OriflammeCard, direction: 'left' | 'right') {
    if (direction === 'left') {
      ws.value?.send(
        encode<ActionOptions>('play_card', {
          playerUuid: playerUuid.value,
          tableUuid: tableUuid.value,
          cardUuid: card.uuid,
          cardAction: 'place_card_left'
        })
      )
    }

    if (direction === 'right') {
      ws.value?.send(
        encode<ActionOptions>('play_card', {
          playerUuid: playerUuid.value,
          tableUuid: tableUuid.value,
          cardUuid: card.uuid,
          cardAction: 'place_card_right'
        })
      )
    }
  }

  /**
   * Card selection
   */

  const maxSelectionReached = computed(() => selectedCards.value.length >= 7)

  function selectCards(card: OriflammeCard) {
    if (selectedCards.value.includes(card.uuid)) {
      selectedCards.value = selectedCards.value.filter(uuid => uuid !== card.uuid)
    } else {
      if (maxSelectionReached.value) return
      selectedCards.value.push(card.uuid)
    }

    useTimeoutFn(() => {
      ws.value?.send(
        encode<ActionOptions>('select_cards', {
          tableUuid: tableUuid.value,
          playerUuid: playerUuid.value,
          selectedCards: selectedCards.value
        })
      )
    }, 2000)
  }

  function _isSelected(card: OriflammeCard) {
    return selectedCards.value.includes(card.uuid)
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
