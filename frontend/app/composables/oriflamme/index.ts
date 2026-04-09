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
  start_game: [{ game_uuid: string }]
  select_cards: [{ player_uuid: string, cardIds: string[] }]
  identify: [{ player_uuid: string, username: string }]
}

type DefaultResponseOptions = {
  action: string
  error: string
  message: string
}

type ResponseOptions = {
  must_identify: { player_uuid: string }
}

export function encode<T extends keyof ActionOptions>(action: T, ...args: ActionOptions[T]) {
  return JSON.stringify({ action, args })
}

type DecodeResponse<T> = DefaultResponseOptions & {
  [K in keyof T]: T[K]
}

/**
 * Return a new decoder function that can be used to decode messages from the websocket.
 */
export function newDecoder(message: string) {
  /**
   * A decoder function that listens to messages from the websocket that matches the specified
   * action. The callback function wil be called with the decoded data if the action matches, or undefined if it doesn't match.
   * @param action - The action we want to decode from the message
   * @param callback - A callback function that will be called with the decoded data if the action matches, or undefined if it doesn't match
   */
  return function<T>(action: keyof T, callback: (data: DecodeResponse<T> | undefined) => void) {
    try {
      const wsData = JSON.parse(message) as DecodeResponse<T>
      if (action === wsData.action) {
        callback(wsData)
      } else {
        callback(undefined)
      }
    } catch (error) {
      console.error('Failed to decode message:', error)
      throw new Error('Invalid message format')
    }
  }
}

export const useOriflammeComposable = createGlobalState(() => {
  const tableId = ref<string>('')
  const playerUuid = ref<string>('')

  const { ws, open, close } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(ws) {
      ws.send(encode('idle_connection'))
    },
    onMessage(_ws, event) {
      const decoder = newDecoder(event.data)

      decoder<ResponseOptions>('must_identify', (data) => {
        if (data) {
          console.log('Received must_identify:', data)
        }
      })
    }
  })

  const isOpen = computed(() => ws.value?.readyState === WebSocket.OPEN)

  function createTable() {
    open()
    ws.value?.send(encode('start_game', { game_uuid: playerUuid.value }))
  }

  function identify(username: string) {
    ws.value?.send(encode('identify', { player_uuid: playerUuid.value, username }))
  }

  function quitTable() {
    close(1000, 'User left the game')
  }

  return {
    ws,
    tableId,
    isOpen,
    identify,
    createTable,
    quitTable
  }
})
