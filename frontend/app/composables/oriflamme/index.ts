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

type DefaultResponseOptions = {
  error: string
  message: string
}

type ResponseOptions = {
  must_identify: { playerId: string }
  another_response: [{ someData: string }]
}

export function encode<T extends keyof ActionOptions>(action: T, ...args: ActionOptions[T]) {
  return JSON.stringify({ action, args })
}

type DecodeResponse<T> = DefaultResponseOptions & {
  [K in keyof T]: T[K]
}

export function decode(message: string) {
  return function<T extends keyof ResponseOptions>(action: T) {
    try {
      return JSON.parse(message) as DecodeResponse<T>
    } catch (error) {
      console.error('Failed to decode message:', error)
      throw new Error('Invalid message format')
    }
  }
}

export const useOriflammeComposable = createGlobalState(() => {
  const tableId = ref<string>('')
  const playerId = ref<string>('')

  const { ws, open, close } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(ws) {
      ws.send(encode('idle_connection'))
    },
    onMessage(_ws, event) {
      const message = decode(event.data)
      const data = message('another_response')
      // data.
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
