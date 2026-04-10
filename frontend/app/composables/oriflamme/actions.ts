import type { Nullable, OriflammeCard, Undefineable } from '~/types'
import { useOriflammeAllCardsComposable, type ActionOptions } from '.'

const [useOriflammeActionsComposable, _useOriflameeActionsStore] = createInjectionState((ws: Ref<WebSocket | undefined>, tableUuid: MaybeRefOrGetter<Nullable<string>>, playerUuid: MaybeRefOrGetter<Nullable<string>>) => {
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

export {
  /**
   * A composable that returns actions that can be performed on the cards, such as revealing a card, applying an effect, placing a token, placing a card, and selecting cards.
   * It also returns a function to check if a card is selected or not, which is used to highlight the selected cards in the UI.
   * @param ws The WebSocket connection to send the actions to the server.
   * @param tableUuid The UUID of the table, used to identify which game the action is for.
   * @param playerUuid The UUID of the player, used to identify which player is performing the action.
   */
  useOriflammeActionsComposable
}

export function useOriflammeActionsStore() {
  const store = _useOriflameeActionsStore()
  if (!store) {
    throw new Error('useOriflameeActionsStore must be used within a useOriflammeActionsComposable')
  }
  return store
}

const [useQueuePossibleActions, _useQueuePossibleActionsStore] = createInjectionState((queue: MaybeRefOrGetter<OriflammeCard[]>) => {
  const selectedIndex = ref<Undefineable<number>>()

  function displayOptions(index: number) {
    selectedIndex.value = index
  }

  const selectedCard = computed(() => {
    if (isDefined(selectedIndex)) {
      return toValue(queue)[selectedIndex.value]
    } else {
      return undefined
    }
  })

  const cardInfo = useOriflammeAllCardsComposable('all')
  const selectedCarInfo = computed(() => {
    if (isDefined(selectedCard)) {
      return toValue(cardInfo.getCardByName(selectedCard.value.name))
    } else {
      return undefined
    }
  })

  const playOptions = computed(() => {
    if (isDefined(selectedCard) && isDefined(selectedCarInfo) && isDefined(selectedIndex)) {
      if (selectedCarInfo.value.canTarget === 'adjacent') {
        return [
          toValue(queue)[selectedIndex.value - 1],
          toValue(queue)[selectedIndex.value + 1]
        ]
      }

      if (selectedCarInfo.value.canTarget === 'first_or_last') {
        return [
          toValue(queue)[0],
          toValue(queue)[toValue(queue).length - 1]
        ]
      }

      if (selectedCarInfo.value.canTarget === 'any') {
        return toValue(queue).filter((_, index) => index !== selectedIndex.value)
      }
    }
    return []
  })

  function isHighlighted(card: OriflammeCard) {
    return playOptions.value.some(option => isDefined(option) ? option.uuid === card.uuid : false)
  }

  return {
    /**
     * The index of the selected card in the queue, used to
     * determine which card's options to display when hovering
     * over a card in the queue.
     */
    selectedIndex,
    /**
     * The selected card in the queue, used to determine which
     * card's options to display when hovering over a card in the queue.
     */
    selectedCard,
    /**
     * The information of the selected card, used to determine which options to
     * display when hovering over a card in the queue, and also to display the card's
     * description in the popover.
     */
    playOptions,
    /**
     * Helper function to set the selected index, which is used to
     * display the possible actions for a card in the queue when hovering
     * over it. It is also used to determine which card is selected and
     * which options to display for that card.
     */
    displayOptions,
    /**
     * Helps indicate whether the card should be highlighted in the queue
     */
    isHighlighted
  }
})

export {
  /**
   * A composable that helps the user know the possible actions that they can do with a card in the queue by
   * highlighting the possible targets of the card
   * @param queue The queue of cards to check the possible actions for. E.g. it can be from the player's hand, but it can be any queue of cards.
   */
  useQueuePossibleActions
}

export function useQueuePossibleActionsStore() {
  const store = _useQueuePossibleActionsStore()
  if (!store) {
    throw new Error('useQueuePossibleActionsStore must be used within a useQueuePossibleActionsComposable')
  }
  return store
}
