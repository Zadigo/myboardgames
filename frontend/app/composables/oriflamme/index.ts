export enum OriflammeCardAction {
  IdleConnection = 'idle_connection',
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
  idle_connection: []
  start_game: [{ gameId: string }]
  select_cards: [{ playerId: string, cardIds: string[] }]
  must_identify: [{ playerId: string }]
}

export function encode<T extends keyof ActionOptions>(action: T, ...args: ActionOptions[T]) {
  return JSON.stringify({ action, args })
}

export const useOriflammeComposable = createGlobalState(() => {
  const tableId = ref<string>('')

  const { ws, open, close } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(ws) {
      ws.send(encode('idle_connection'))
    }
  })

  const isOpen = computed(() => ws.value?.readyState === WebSocket.OPEN)

  function createTable() {
    open()
    ws.value?.send(encode('start_game', { gameId: '1234' }))
  }

  function quitTable() {
    close(1000, 'User left the game')
  }

  return {
    ws,
    tableId,
    isOpen,
    createTable,
    quitTable
  }
})
