import type { CardActions, GameRegistry, Nullable, OriflammeCard } from '~/types'

export * from './actions'

export enum OriflammeCardAction {
  StartGame = 'start_game',
  SelectCards = 'select_cards',
  MustIdentify = 'must_identify',
  PlayCard = 'play_card',
  PlaceCard = 'place_card',
  Reveal = 'reveal',
  PlaceToken = 'place_token',
  StackCard = 'stack_card',
  ResolveQueue = 'resolve_queue',
  EndGame = 'end_game'
}

export type ActionOptions = {
  create_game: [{ playerUuid: string }]
  start_game: [{ playerUuid: string, tableUuid: Nullable<string> }]
  select_cards: [{ tableUuid: Nullable<string>, playerUuid: string, selectedCards: string[] }]
  identify: [{ playerUuid: string, username: string }]
  play_card: [{ playerUuid: string, tableUuid: Nullable<string>, cardAction: CardActions }]
}

export type ResponseOptions = {
  must_identify: { playerUuid: string }
  create_game: { gameRegistry: GameRegistry }
  place_card: { queue: null }
  start_game: { cardsInPlay: OriflammeCard[] }
}

export const useOriflammeComposable = createGlobalState(() => {
  const { encode, decode } = useWWebsocketMessages2()

  const playerUuid = useLocalStorage<string>('oriflamme-player-uuid', null)
  const gameRegistry = useLocalStorage<GameRegistry>('oriflamme-game-registry', null, {
    serializer: {
      read: raw => JSON.parse(raw),
      write: value => JSON.stringify(value)
    }
  })
  const tableUuid = computed(() => isDefined(gameRegistry) ? gameRegistry.value.uuid : null)

  const influenceQueue = computed(() => isDefined(gameRegistry) ? gameRegistry.value.influenceQueueLayer.queue : [])

  const cardsInPlay = computed(() => {
    if (!isDefined(gameRegistry)) {
      return []
    } else {
      return gameRegistry.value.cardsInPlay || []
    }
  })
  const playerCards = computed(() => cardsInPlay.value.filter(card => card.owner.uuid === playerUuid.value))

  const { ws, open, close, send } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(_ws) {
    },
    onMessage(_ws, event) {
      decode<ResponseOptions>(event.data)('must_identify', (data) => {
        if (data) {
          playerUuid.value = data.playerUuid
          send(encode<ActionOptions>('identify', { playerUuid: playerUuid.value, username: 'pauline' }))
        }
      })

      decode<ResponseOptions>(event.data)('create_game', (data) => {
        if (data) {
          gameRegistry.value = data.gameRegistry
        }
      })

      decode<ResponseOptions>(event.data)('place_card', (data) => {
        if (data) {
          // Do something
        }
      })

      decode<ResponseOptions>(event.data)('start_game', (data) => {
        if (data) {
          // Do something
        }
      })
    }
  })

  const isOpen = computed(() => ws.value?.readyState === WebSocket.OPEN)

  async function openConnection() {
    open()
  }

  async function createGame() {
    send(encode<ActionOptions>('create_game', { playerUuid: playerUuid.value }))
  }

  function startGame() {
    send(encode<ActionOptions>('start_game', { playerUuid: playerUuid.value, tableUuid: tableUuid.value }))
  }

  function quitTable() {
    close(1000, 'User left the game')
  }

  return {
    ws,
    playerUuid,
    tableUuid,
    isOpen,
    influenceQueue,
    cardsInPlay,
    playerCards,
    openConnection,
    createGame,
    startGame,
    quitTable
  }
})
