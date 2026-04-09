import type { OriflammeCard } from '~/types'

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
  start_game: [{ playerUuid: string, tableUuid: string }]
  select_cards: [{ playerUuid: string, selectedCards: string[] }]
  identify: [{ playerUuid: string, username: string }]
}

export type ResponseOptions = {
  must_identify: { playerUuid: string }
  create_game: { tableUuid: string }
}

export const useOriflammeComposable = createGlobalState(() => {
  const tableUuid = ref<string>('')
  const playerUuid = ref<string>('')
  const influenceQueue = ref<OriflammeCard[]>([])

  const { encode, decode } = useWWebsocketMessages2()

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
          tableUuid.value = data.tableUuid
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
    openConnection,
    createGame,
    startGame,
    quitTable
  }
})
