import { useOriflammeBaseCardsComposable } from './base'

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
 * Helper composable to get all the cards of the game, including base cards and expansion cards.
 * This is useful for the card details page, where we want to display all the cards of the game.
 * It is also useful for the game logic, where we need to check if a card is in the game or not.
 */
export function useOriflammeAllCards() {
  const { baseCards } = useOriflammeBaseCardsComposable()

  const cards = computed(() => {
    return [
      ...baseCards
    ]
  })

  return {
    cards
  }
}
