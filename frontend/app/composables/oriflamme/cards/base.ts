import type { BaseCharacters } from '~/types'
import type { BaseCardInformation } from '.'

/**
 * Gives a textual description of the power of each card, as well as the rules for targeting other cards with it.
 * This is used in the base card scroll, and also in the card popover when hovering over a card in play.
 */
export const baseCards: BaseCardInformation<`${BaseCharacters}`>[] = [
  {
    name: 'Archer',
    power: 'Eliminate the first or last card from the Queue',
    description: 'The targeted card can be revealed or not. You may choose or may be forced to eliminate a card from your family, includng the Archer itself if it’s the first or last card in the Queue. You gain 1 in any case',
    type: 'character',
    canTarget: 'first_or_last'
  },
  {
    name: 'Soldier',
    power: ' Eliminate an adjacent card',
    description: 'The targeted card can be revealed or not. You may choose or may be forced to eliminate a card from your family. You gain 1 in any case.',
    type: 'character',
    canTarget: 'adjacent'
  },
  {
    name: 'Spy',
    power: 'Steal 1 from a player with an adjacent card',
    description: 'The adjacent card can be revealed or not.Take the directly from the owner of the card, not from the adjacent card (in case it’s face down and has on it). Stealing from yourself has no effect.',
    type: 'character',
    canTarget: 'adjacent'
  },
  {
    name: 'Heir',
    power: 'If there is no other card revealed with the same name, gain 2',
    description: 'As soon as another Heir is revealed and uncovered in the Queue, no Heirs gain any',
    type: 'character',
    canTarget: 'none'
  },
  {
    name: 'Shapeshifter',
    power: 'Copy the ability of an adjacent revealed character',
    description: 'The targeted card must be revealed and uncovered. TheShapeshifter copies only the ability of a card, not its name. No matter what it copies, the Shapeshifter always retains its name, Shapeshifter. EXAMPLE: If you copy an Heir, you gain 2 unless there is another Shapeshifter in the Queue. Once resolved, the Shapeshifter immediately loses the copied ability Consequently, copying a Shapeshifter with a Shapeshifter has no effect. At each Resolution Phase, the Shapeshifter can, of course, choose to copy a different character than the previous turn.',
    type: 'character',
    canTarget: 'adjacent'
  },
  {
    name: 'Lord',
    power: 'Gain 1 plus 1 for each adjacent card that is in your family.',
    description: 'Gain 1 additionnal for each adjacent card in your family, whether it’s revealed or not. If your Lord is adjacent to a stack of cards of your family, only the top card is taken into account',
    type: 'character',
    canTarget: 'none'
  },
  {
    name: 'Assassination',
    power: 'Eliminate a card anywhere in the Queue. Discard Assassination',
    description: 'The targeted card can be revealed or not. You can eliminate a card from your family, including the Assassination itself. You gain 1 in any case.',
    type: 'intrigue',
    canTarget: 'any'
  },
  {
    name: 'Royal Decree',
    power: 'Move a card anywhere in the Queue, except on another card. Discard Royal Decree',
    description: 'It doesn’t matter whether the target card is revealed. It can be from any family. Any on the card are moved as well. If you target a stack of cards, only the top card can be moved. To move a card, leave the Royal Decree where it is, make the move, then remove the Royal Decree and continue resolution with the following card in the Queue. NOTE: By moving a card before the Royal Decree, you can prevent a card from being resolved, or instead have a card resolved twice by moving it after the Royal Decree',
    type: 'intrigue',
    canTarget: 'any'
  },
  {
    name: 'Conspiracy',
    power: 'Gain double the accumulated on Conspiracy. Discard Conspiracy.',
    description: 'If there are 3 on this card when you reveal it, you gain these and you gain 3 more, for a total of 6.',
    type: 'intrigue',
    canTarget: 'none'
  },
  {
    name: 'Ambush',
    power: 'Discard all on Ambush and gain 1 . Discard Ambush.',
    description: 'If Ambush is eliminated by an opponent’s card, discard the opponent’s card and gain 4 . NOTE: The second ability is, of course, the main appeal of the Ambush. If no one has fallen into your trap, the first ability still allows you to gain a single as a consolation. EXAMPLE 1: If an opponent’s Soldier eliminates your Ambush, the Soldier’s owner gains 1 and you gain 4 . The Soldier is discarded. EXAMPLE 2: If your own Archer eliminates your Ambush, you gain 1 . The Ambush is discarded. Your Archer remains in the Queue (as the second ability of Ambush only applies to opponent’s cards).',
    type: 'intrigue',
    canTarget: 'none'
  }
]
