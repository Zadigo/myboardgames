import { baseCards } from './base'

export * from './base'
export * from './ablaze'
export * from './embrasement'

export type BaseCardInformation<T> = {
  name: T
  power: string
  description: string
  type: 'character' | 'intrigue'
  canTarget: 'adjacent' | 'first_or_last' | 'any' | 'none'
}

/**
 * A helper composable to get information about the base cards, such as their power and description.
 * This is used in the base card scroll, and also in the card popover when hovering over a card in play.
 */
export function useOriflammeCardsComposable<T>(cards: BaseCardInformation<T>[]) {
  function _getCardByName(name: BaseCardInformation<T>['name']) {
    return cards.find(card => card.name === name)
  }

  const getCardByName = reactify(_getCardByName)

  function _cardHasAbility(cardName: BaseCardInformation<T>['name'], items: BaseCardInformation<T>['name'][]) {
    return !items.includes(cardName)
  }

  const cardHasAbility = reactify(_cardHasAbility)

  return {
    cards,
    getCardByName,
    /**
     * Helper function used to toggle the visibility of certain
     * buttons in the card popover
     */
    cardHasAbility
  }
}

/**
 * Helper composable to get all the cards of the game, including base cards and expansion cards.
 * This is useful for the card details page, where we want to display all the cards of the game.
 * It is also useful for the game logic, where we need to check if a card is in the game or not.
 */
export function useOriflammeAllCards() {
  const { cards: baseGameCards } = useOriflammeCardsComposable(baseCards)

  const allCards = computed(() => {
    return [
      ...baseGameCards
    ]
  })

  const search = ref<string>('')

  const searched = computed(() => {
    return useArrayFilter(allCards.value, card => card.name.toLowerCase().includes(search.value.toLowerCase()))
  })

  return {
    allCards,
    search,
    searched
  }
}
