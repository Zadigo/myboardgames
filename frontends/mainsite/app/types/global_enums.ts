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
}

export enum FlipSevenWsActions {
  FlipCard = 'flip_card',
  CardFlipped = 'card_flipped'
}
