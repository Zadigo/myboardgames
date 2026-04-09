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

type DefaultResponseOptions = {
  action: string
  error: string
  message: string
}

type ResponseOptions = {
  must_identify: { player_uuid: string }
  create_game: { table_uuid: string }
}

export function encode<T extends keyof ActionOptions>(action: T, ...args: ActionOptions[T]) {
  return JSON.stringify({ action, ...args[0] })
}

type DecodeResponse<T> = DefaultResponseOptions & {
  [K in keyof T]: T[K]
}

/**
 * Return a new decoder function that can be used to decode messages from the websocket.
 */
export function useWsMessageDecoder<R>(message: string) {
  /**
   * A decoder function that listens to messages from the websocket that matches the specified
   * action. The callback function wil be called with the decoded data if the action matches, or undefined if it doesn't match.
   * @param action - The action we want to decode from the message
   * @param callback - A callback function that will be called with the decoded data if the action matches, or undefined if it doesn't match
   */
  return function <T extends keyof R>(action: T, callback: (data: DecodeResponse<R[T]> | undefined) => void) {
    try {
      const wsData = JSON.parse(message) as DecodeResponse<R[T]>
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
  const tableUuid = ref<string>('')
  const playerUuid = ref<string>('')

  const { ws, open, close, send } = useWebSocket('ws://127.0.0.1:9000/oriflamme/live', {
    immediate: false,
    onConnected(_ws) {
    },
    onMessage(_ws, event) {
      const decoder = useWsMessageDecoder<ResponseOptions>(event.data)

      decoder('must_identify', (data) => {
        if (data) {
          playerUuid.value = data.player_uuid
          send(encode('identify', { player_uuid: playerUuid.value, username: 'pauline' }))
        }
      })

      decoder('create_game', (data) => {
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
    send(encode('create_game', { player_uuid: playerUuid.value }))
  }

  function startGame() {
    send(encode('start_game', { game_uuid: playerUuid.value }))
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
