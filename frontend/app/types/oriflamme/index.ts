import type { Nullable } from '..'
import type { BlueBoxCharacters } from './bluebox'

export type * from './bluebox'

export type WebsocketClient = {
  uuid: string
  username: string
  initiator: boolean
  tokens: number
  discardPile: OriflammeCard[] | null
}

export type OriflammeCard = {
  uuid: string
  positionInQueue: number
  name: `${BlueBoxCharacters}`
  type: string
  stack: OriflammeCard[] | null
  color: string
  owner: WebsocketClient
  tokens: number
  image: string
  isSelected: boolean
  isRemoved: boolean
  isDiscarded: boolean
  isRevealed: boolean
  inQueue: boolean
}

export type PlayerChoices = {
  /**
   * Whether to apply an effect on the card immediately before
   * or immediately after the card being resolved.
   */
  cardBefore: boolean
  /**
   * Whether to apply an effect on the first or last card in the queue.
   */

  firstCard: boolean
  /**
   * The index of the card in the queue on which to apply an effect.
   * This is used for the Assassination card, which can eliminate any card in the queue.
   */
  atIndex: number
  /**
   * The choice to apply to the card that the shapeshifter is copying. This is used when
   * the shapeshifter copies the effect of a soldier, which can eliminate either
   * the card immediately before or immediately after the soldier.
   */
  shapeShifterCardBefore: boolean
  /**
   * The choice to apply to the card that the shapeshifter is copying. This is used when
   * the shapeshifter copies the effect of an archer, which can eliminate either
   * the first or last card in the queue.
   */
  shapeShifterFirstCard: boolean
  /**
   * The index of the card in the queue on which to apply an effect. This is used when
   * the shapeshifter copies the effect of the assassination card, which can eliminate any card in the queue.
   */
  shapeShifterAtIndex: number
  /**
   * When the Shapesifter copies the effect of a card, we need to resolve the effect based
   * on the position of the card being copied in the queue, not the position of the shapeshifter itself.
   * This field is used to temporarily store the index of the card being copied during the resolution of
   * the shapeshifter's effect.
   */
  temporaryResolutionIndex: number
  /**
   * Indicates whether the card is being temporarily controlled by the
   * Shapeshifter's effect.
   */
  isRemote: boolean
  /**
   * The card that is remotely controlling the effect of
   * another card (e.g. the shapeshifter copying the effect of another card).
   */
  remoteCard: OriflammeCard
  /**
   * The card that the Royal Decree wants to move. This is used when the player chooses to apply the Royal Decree's effect
   * to move a card from its current position in the queue to another position.
   */
  royalDecreeTargetCard: OriflammeCard
  /**
   * The index of the card that the Royal Decree wants to move.
   */
  royalDecreeFrom: number
  /**
   * The target position in the queue to which the Royal Decree wants to move the card.
   */
  royalDecreeTo: number
}

export enum CardActionsEnum {
  PlaceCard = 'place_card',
  Reveal = 'reveal',
  PlaceToken = 'place_token',
  StackCard = 'stack_card'
}

export type CardActions = `${CardActionsEnum}`

/**
 * Holds the state of a game, including the players,
 * the cards in play, the influence queue, and other relevant information.
 */
export type GameRegistry = {
  uuid: string
  clients: Record<string, WebsocketClient>
  influenceQueueLayer: {
    queue: OriflammeCard[]
    resolutionIndex: number
  }
  isRunning: boolean
  isStarted: boolean
  startedAt: string
  cardsInPlay: Nullable<OriflammeCard[]>
}
