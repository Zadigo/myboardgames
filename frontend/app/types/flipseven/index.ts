export type Deck = {
  value: number
  owner: string
  category: string
  isMultiplier: boolean
  isNumber: boolean
  isBonus: boolean
  isSpecial: boolean
}

export interface TableClientDetail {
  username: string
  uuid: string
  tableUuid: string
  numberOfCards: number
  isFreezed: boolean
  hasFreezeCard: boolean
  hasSecondChanceCard: boolean
  hasSevenCards: boolean
  // hasSecondChance: boolean
  cards?: Deck[]
  score: number
  isInitiator: boolean
}

/**
 * @deprecated
 */
export interface TableClient {
  details: TableClientDetail
}

export type TableDetails = {
  uuid: string
  deck: Deck[]
  deckIndex: number
  players: Record<string, TableClientDetail>
  // clients: TableClient[]
  // currentDeck?: Deck[]
  // numberOfPlayers: number
  // gameStarted: boolean
}

export enum WsActions {
  InitialConnection = 'initial_connection',
  AcceptPlayer = 'accept_player',
  GetDeck = 'get_deck',
  DeckCreated = 'deck_created',
  UpdateWaitingLobby = 'update_waiting_lobby',
  InitiateTable = 'initiate_table',
  TableInitiated = 'table_initiated',
  Error = 'error',
  Reconnect = 'reconnect',
  Reconnected = 'reconnected',
  FlipCard = 'flip_card',
  CardFlipped = 'card_flipped'
}

export type SendMessage = { action: WsActions.AcceptPlayer, tableId: string | undefined | null, username: string }
  | { action: WsActions.GetDeck, tableId: string | undefined | null }
  | { action: WsActions.InitialConnection, username: string }
  | { action: WsActions.InitiateTable, username: string }
  | { action: WsActions.Reconnect, tableId: string | undefined | null, username: string }
  | { action: WsActions.FlipCard, tableId: string | undefined | null, playerId: string }

export type ReceiveMessage = { action: WsActions.AcceptPlayer, something: string }
  | { action: WsActions.DeckCreated, deck: Deck[] }
  | { action: WsActions.InitialConnection }
  | { action: WsActions.UpdateWaitingLobby, tableDetails: TableDetails }
  | { action: WsActions.TableInitiated, tableId: string, tableDetails: TableDetails }
  | { action: WsActions.Error, message: string }
  | { action: WsActions.Reconnected, tableDetails: TableDetails }
  | { action: WsActions.CardFlipped, tableDetails: TableDetails, cardDetails: Deck }
