export const useFlipSevenGameComposable = createGlobalState(() => {
  const tableId = useSessionStorage<string | null>('tableId', null)
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

    tableId.value = data.tableId
    params.table = data.tableId
  }

  return {
    tableId,
    create
  }
})

enum WsActions {
  InitialConnection = 'initial_connection',
  InitialConnectionResponse = 'idle_connect'
}

type SendMessage = { action: WsActions.InitialConnection, tableId: string | null }

type ReceiveMessage = { action: WsActions.InitialConnectionResponse, something: string }

export const useFlipLiveGameComposable = createGlobalState((tableId: string | null) => {
  const { decode, encode } = useWebsocketMessage()

  const wsObject = useWebSocket('ws://127.0.0.1:9000/flip-sevenn', {
    immediate: false,
    onConnected(ws) {
      ws.send(encode<SendMessage>({ action: WsActions.InitialConnection, tableId }))
    },
    onDisconnected(ws, event) {
      console.log('WebSocket disconnected:', event)
    },
    onError(ws, event) {
      console.error('WebSocket error:', event)
    },
    onMessage(ws, event) {
      const message = decode<ReceiveMessage>(event.data)
      if (message) {
        switch (message.action) {
          case WsActions.InitialConnectionResponse:
            console.log('Received initial connection response:', message)
            break
          default:
            console.warn('Unknown WebSocket message action:', message.action)
        }
      } else {
        console.warn('Received invalid WebSocket message:', event.data)
      }
    }
  })

  return {
    wsObject
  }
})
