import { WsActions, type Deck, type ReceiveMessage, type SendMessage, type TableDetails } from '~/types'

export const useFlipSevenGameComposable = createGlobalState(() => {
  const _tableId = useSessionStorage<string | null>('tableId', null)
  const params = useUrlSearchParams('history') as { table: string }

  async function create() {
    const data = await $fetch<{ tableId: string }>('/v1/flip-seven/create', {
      method: 'POST',
      baseURL: 'http://127.0.0.1:9000',
      headers: { 'Content-Type': 'application/json' },
      body: {
        username: 'Player1'
      }
    })

    _tableId.value = data.tableId
    params.table = data.tableId
  }

  const tableId = computed(() => _tableId.value || params.table || null)

  return {
    tableId,
    create
  }
})

export const useFlipSevenLiveGameComposable = createGlobalState((tableId: MaybeRef<string | null>) => {
  const _tableId = toValue(tableId)

  const initialDeck = ref<Deck[]>([])
  const tableDetails = ref<TableDetails>()

  const { decode, encode } = useWebsocketMessage()
  const wsObject = useWebSocket(`ws://127.0.0.1:9000/ws/flip-seven?table=${_tableId}`, {
    immediate: false,
    onConnected(ws) {
      ws.send(encode<SendMessage>({
        action: WsActions.WaitingLobby,
        tableId: _tableId,
        username: 'Player1'
      }))
    },
    onDisconnected(ws, event) {
      console.log('WebSocket disconnected:', event)
    },
    onError(ws, event) {
      console.error('WebSocket error:', event)
    },
    onMessage(ws, event) {
      const message = decode<ReceiveMessage>(event.data)
      console.log('Received WebSocket message:', message)

      if (message) {
        switch (message.action) {
          case WsActions.InitialConnection:
            console.log('Received initial connection response:', message)
            break

          case WsActions.WaitingLobby:
            console.log('Received initial connection response:', message)
            break

          case WsActions.DeckCreated:
            console.log('Received deck created message:', message)
            initialDeck.value = message.deck
            break

          case WsActions.UpdateWaitingLobby:
            tableDetails.value = message.tableDetails
            break

          default:
            console.warn('Unknown WebSocket message action:', message)
        }
      } else {
        console.warn('Received invalid WebSocket message:', event.data)
      }
    }
  })

  return {
    wsObject,
    initialDeck,
    tableDetails
  }
})
