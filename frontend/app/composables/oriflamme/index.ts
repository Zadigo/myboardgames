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

type ActionOptions = {
  create_game: [{ player_uuid: string }]
  start_game: [{ game_uuid: string }]
  select_cards: [{ player_uuid: string, cardIds: string[] }]
  identify: [{ player_uuid: string, username: string }]
}

type ResponseOptions = {
  must_identify: { player_uuid: string }
  create_game: { table_uuid: string }
}

export const useOriflammeComposable = createGlobalState(() => {
  const tableUuid = ref<string>('')
  const playerUuid = ref<string>('')

  const { encode, decode } = useWWebsocketMessages2()

  const { ws, open, close, send } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(_ws) {
    },
    onMessage(_ws, event) {
      decode<ResponseOptions>(event.data)('must_identify', (data) => {
        if (data) {
          playerUuid.value = data.player_uuid
          send(encode<ActionOptions>('identify', { player_uuid: playerUuid.value, username: 'pauline' }))
        }
      })

      decode<ResponseOptions>(event.data)('create_game', (data) => {
        if (data) {
          tableUuid.value = data.table_uuid
        }
      })
    }
  })

  const isOpen = computed(() => ws.value?.readyState === WebSocket.OPEN)

  async function openConnection() {
    open()
  }

  async function createGame() {
    send(encode<ActionOptions>('create_game', { player_uuid: playerUuid.value }))
  }

  function startGame() {
    send(encode<ActionOptions>('start_game', { game_uuid: playerUuid.value }))
  }

  function quitTable() {
    close(1000, 'User left the game')
  }

  return {
    ws,
    playerUuid,
    tableUuid,
    isOpen,
    openConnection,
    createGame,
    startGame,
    quitTable
  }
})
