interface WebsocketClient {
  username: string
  uuid: string
}

export interface Detail {
  currentPlayer?: any
  currentRound: number
  isStarted: boolean
  startedAt: string
  uuid: string
}

export interface MarvelToken {
  mind: number
  power: number
  reality: number
  shield: number
  soul: number
  space: number
  time: number
}

export interface NormalToken {
  diamond: number
  emerald: number
  onyx: number
  ruby: number
  sapphire: number
}

export interface Tokens {
  marvelTokens: MarvelToken
  normalTokens: NormalToken
}

export interface PlayingTableDetails {
  cardslevelone: any[]
  cardslevelthree: any[]
  cardsleveltwo: any[]
  clients: WebsocketClient[]
  decklevelone: any[]
  decklevelthree: any[]
  deckleveltwo: any[]
  details: Detail
  isNormalGame: boolean
  tokens: Tokens
}

export interface SplendorPlayingTable {
  playerUuid: string
  playingTable: PlayingTableDetails
}
