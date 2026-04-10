import type { AllCharacters } from '~/types'
import { ablazeCards } from './ablaze'
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
export function useOriflammeAllCardsComposable(searchType: 'all' | 'base' | 'ablaze' = 'all') {
  const { cards: baseGameCards } = useOriflammeCardsComposable(baseCards)
  const { cards: ablazeGameCards } = useOriflammeCardsComposable(ablazeCards)

  const allCards = computed(() => {
    return [
      ...baseGameCards,
      ...ablazeGameCards
    ]
  })

  const search = ref<AllCharacters | undefined>()

  const searched = computed(() => {
    if (searchType === 'base') {
      return useArrayFilter(baseGameCards, card => card.name.toLowerCase().includes(search.value?.toLowerCase() ?? ''))
    }

    if (searchType === 'ablaze') {
      return useArrayFilter(ablazeGameCards, card => card.name.toLowerCase().includes(search.value?.toLowerCase() ?? ''))
    }

    return useArrayFilter(allCards.value, card => card.name.toLowerCase().includes(search.value?.toLowerCase() ?? ''))
  })

  function isBaseCard(cardName: AllCharacters): boolean {
    return baseGameCards.some(card => card.name === cardName)
  }

  function _getCardByName(name: MaybeRefOrGetter<AllCharacters>) {
    return allCards.value.find(card => card.name === toValue(name))
  }

  const getCardByName = reactify(_getCardByName)

  return {
    allCards,
    search,
    searched,
    /**
     * Helper function to check if a card is a base card or an
     * expansion card, used in the card details page to display
     * the correct image and description.
     * @param cardName The name of the card to check.
     */
    isBaseCard,
    /**
     * Helper function to get a card by its name, used in the card details page to display
     * the correct image and description, and also in the game logic to check if a card is in the game or not.
     * @param name The name of the card to get information about.
     */
    getCardByName
  }
}
