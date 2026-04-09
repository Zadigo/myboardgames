export type WebscoketClient = {
  Uuid: string
}

export type OriflammeCard = {
  Uuid: string
  PositionInQueue: number
  Name: string
  Type: string
  Stack: OriflammeCard[]
  Color: string
  Owner: WebscoketClient
  IsSelected: boolean
  IsRemoved: boolean
  IsDiscarded: boolean
  Tokens: number
  IsRevealed: boolean
  Image: string
}

export type PlayerChoices = {
  /**
   * Whether to apply an effect on the card immediately before
   * or immediately after the card being resolved.
   */
  CardBefore: boolean
  /**
   * Whether to apply an effect on the first or last card in the queue.
   */

  FirstCard: boolean
  /**
   * The index of the card in the queue on which to apply an effect.
   * This is used for the Assassination card, which can eliminate any card in the queue.
   */
  AtIndex: number
  /**
   * The choice to apply to the card that the shapeshifter is copying. This is used when
   * the shapeshifter copies the effect of a soldier, which can eliminate either
   * the card immediately before or immediately after the soldier.
   */
  ShapeShifterCardBefore: boolean
  /**
   * The choice to apply to the card that the shapeshifter is copying. This is used when
   * the shapeshifter copies the effect of an archer, which can eliminate either
   * the first or last card in the queue.
   */
  ShapeShifterFirstCard: boolean
  /**
   * The index of the card in the queue on which to apply an effect. This is used when
   * the shapeshifter copies the effect of the assassination card, which can eliminate any card in the queue.
   */
  ShapeShifterAtIndex: number
  /**
   * When the Shapesifter copies the effect of a card, we need to resolve the effect based
   * on the position of the card being copied in the queue, not the position of the shapeshifter itself.
   * This field is used to temporarily store the index of the card being copied during the resolution of
   * the shapeshifter's effect.
   */
  TemporaryResolutionIndex: number
  /**
   * Indicates whether the card is being temporarily controlled by the
   * Shapeshifter's effect.
   */
  IsRemote: boolean
  /**
   * The card that is remotely controlling the effect of
   * another card (e.g. the shapeshifter copying the effect of another card).
   */
  RemoteCard: OriflammeCard
  /**
   * The card that the Royal Decree wants to move. This is used when the player chooses to apply the Royal Decree's effect
   * to move a card from its current position in the queue to another position.
   */
  RoyalDecreeTargetCard: OriflammeCard
  /**
   * The index of the card that the Royal Decree wants to move.
   */
  RoyalDecreeFrom: number
  /**
   * The target position in the queue to which the Royal Decree wants to move the card.
   */
  RoyalDecreeTo: number
}
