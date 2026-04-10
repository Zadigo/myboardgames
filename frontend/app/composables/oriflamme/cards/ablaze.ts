import type { AblazeCharacters } from '~/types'
import type { BaseCardInformation } from '.'

/**
 * Gives a textual description of the power of each card, as well as the rules for targeting other cards with it.
 * This is used in the ablaze card scroll, and also in the card popover when hovering over a card in play.
 */
export const ablazeCards: BaseCardInformation<AblazeCharacters>[] = [
  {
    name: 'Prince',
    power: 'When you reveal your Prince, place your family’s Twin, revealed, in the Queue. Gain 1',
    description: 'When the Prince is eliminated, discard your family’s Twin, unless it is covered by another card.',
    type: 'character',
    canTarget: 'none'
  },
  {
    name: 'Twin',
    power: 'When the Twin is eliminated, discard your family’s Prince, unless it is covered by another card.',
    description: 'During setup, only shuffle the Prince with your family’s other cards. The Twin can only be played through the Prince’s ability. The Twin can be placed as usual, at the beginning or at the end of the Queue, or on another of your Family’s cards, EXCEPT ON THE PRINCE. • Some cards may cause a player to have 2 Princes or 2 Twins at the same time. Eliminating 1 copy of one of them discards every copy of the other. Example: You have 2 Princes and 1 Twin. If 1 Prince is eliminated, the Twin is discarded and the second Prince remains. If your Twin is eliminated, both Princes are discarded.',
    type: 'character',
    canTarget: 'none'
  },
  {
    name: 'Queen',
    power: 'Gain 2',
    description: 'If an opponent eliminates the Queen, they gain 1 additional.',
    type: 'character',
    canTarget: 'adjacent'
  },
  {
    name: 'Cutthroat',
    power: 'Eliminate an adjacent card and all the opponents’ revealed cards with the same name. If only one card is eliminated, do not gain the for eliminating a card.',
    description: 'EXAMPLE: If the Cutthroat eliminates a Felon, you gain no . If the Cutthroat eliminates 2 Felons, you gain 2 . If the Cutthroat eliminates an opponent’s Queen, you gain 1 . If the Cutthroat eliminates 2 opponents’ Queens, you gain 4',
    type: 'character',
    canTarget: 'none'
  },
  {
    name: 'Apothecary',
    power: 'Eliminate a card adjacent to another card of your family.',
    description: 'The other card from your Family may be revealed or face down. • The eliminated card still counts as eliminated by the Apothecary. Example: If the eliminated card is a Trap, the Apothecary is discarded. • If you have no other card from your family in the Queue, the Apothecary’s ability has no effect.',
    type: 'character',
    canTarget: 'any'
  },
  {
    name: 'Felon',
    power: 'Each player loses 1 for each adjacent card of their family.',
    description: 'GIf a card from your Family is adjacent to your own Felon, you lose 1 . If a player has 2 cards from their Family adjacent to the Felon, they lose 2',
    type: 'character',
    canTarget: 'any'
  },
  {
    name: 'Schemer',
    power: 'Discard this card if it’s adjacent to a stack of cards. If not, gain 2',
    description: 'Discard the Schemer only if it is adjacent to a stack of cards when you resolve its ability.',
    type: 'character',
    canTarget: 'any'
  },
  {
    name: 'Impersonation',
    power: 'Eliminate an adjacent card. If it is an opponent’s character, replace it with the same character, revealed, taken from your discarded, eliminated or set aside cards. Discard Impersonation',
    description: '• If the eliminated Character was on top of a stack of cards, you can not replace it with a character from your Family. • Using Impersonation on a Prince does not allow you to play your Twin, because your Prince is already revealed when it is placed. • The Twin does not count as a card you set aside at the beginning of the game. Using Impersonation on a Twin allows you to place your own Twin only if it was previously played and discarded or eliminated.',
    type: 'intrigue',
    canTarget: 'any'
  },
  {
    name: 'Plot',
    power: 'Discard Plot. Activate the ability of any revealed character of your family. Each on the Plot can either be gained or discarded to repeat the effect above',
    description: '• You can freely choose any Character of your choice for each iteration of the effect. You can choose the same character or a different one. • Fully resolve each ability before deciding to discard or gain the next.',
    type: 'intrigue',
    canTarget: 'adjacent'
  },
  {
    name: 'Trap',
    power: 'Discard all on Trap and gain 1 . Discard Trap or If Trap is eliminated by an opponent’s card, discard the opponent’s card and steal 3 from that player.',
    description: 'When a player eliminates a Trap, they gain 1 before you steal 3 from them. ADVICE: The second part of the ability is the main appeal of the Trap. If no one falls into your Trap, the first part of its ability lets you gain 1 as a consolation. EXAMPLE 1: An opponent’s Apothecary eliminates your Trap. The owner of the Apothecary gains 1 then you steal 3 from them. The Apothecary is discarded. EXAMPLE 2: Your Apothecary eliminates your Trap. You gain 1 . Your Apothecary stays in the Queue because the Trap’s ability only triggers on opponent’s cards.',
    type: 'intrigue',
    canTarget: 'none'
  },
  {
    name: 'Bribery',
    power: 'Place your family’s Bribery token on any revealed character. This character now belongs to your family. Discard Bribery.',
    description: '• If you place your Bribery token on a character which already has a Bribery token, replace the token with yours. • You can not play your Bribery on a character on top of a stack of cards. • When a character with a Bribery token is discarded or eliminated, discard the Bribery token and place the character’s card with its original owner’s eliminated cards.',
    type: 'intrigue',
    canTarget: 'any'
  }
]
