import { WsActions, type Deck, type ReceiveMessage, type SendMessage, type TableDetails } from '~/types'

export const TEST_USERNAME = 'Player 1'

export const useFlipSevenLiveGameComposable = createGlobalState(() => {
  const params = useUrlSearchParams('history') as { table: string }
  const tableId = useSessionStorage<string | null>('tableId', null)

  watchEffect(() => {
    if (params.table) {
      tableId.value = params.table
    }
  })

  watch(tableId, (newTableId) => {
    if (newTableId) {
      params.table = newTableId
    }
  })

  const initialDeck = ref<Deck[]>([])
  const tableDetails = ref<TableDetails>()

  const { decode, encode } = useWebsocketMessage()
  const wsObject = useWebSocket(`ws://127.0.0.1:9000/ws/flip-seven`, {
    immediate: false,
    onConnected(_ws) {},
    onDisconnected(_ws, _event) {},
    onError(_ws, _event) {},
    onMessage(ws, event) {
      const message = decode<ReceiveMessage>(event.data)

      console.log('Received WebSocket message:', message)

      if (message) {
        switch (message.action) {
          case WsActions.InitialConnection:
            console.log('Received initial connection response:', message)
            break

          case WsActions.TableInitiated:
            params.table = message.tableId
            break

          case WsActions.WaitingLobby:
            console.log('Received waiting lobby message:', message)
            break

          case WsActions.DeckCreated:
            console.log('Received deck created message:', message)

            initialDeck.value = message.deck
            if (isDefined(tableDetails)) {
              tableDetails.value.currentDeck = initialDeck.value
            }
            break

          case WsActions.CardFlipped:
            console.log('Received card flipped message:', message)

            tableDetails.value = message.tableDetails
            break

          case WsActions.UpdateWaitingLobby:
            tableDetails.value = message.tableDetails
            break

          case WsActions.Reconnected:
            tableDetails.value = message.tableDetails
            break

          case WsActions.Error:
            console.error('Received error message from WebSocket:', message.message)
            break

          default:
            console.warn('Unknown WebSocket message action:', message)
        }
      } else {
        console.warn('Received invalid WebSocket message:', event.data)
      }
    }
  })

  const isConnected = computed(() => wsObject.status.value === 'OPEN')

  function createTable() {
    if (!isConnected.value) {
      wsObject.open()
      wsObject.send(encode<SendMessage>({
        action: WsActions.InitiateTable,
        username: TEST_USERNAME
      }))
    }
  }

  function joinTable() {
    if (isConnected.value) {
      wsObject.send(encode<SendMessage>({
        action: WsActions.WaitingLobby,
        tableId: toValue(tableId),
        username: TEST_USERNAME
      }))
    }
  }

  function getDeck() {
    wsObject.send(encode<SendMessage>({
      action: WsActions.GetDeck,
      tableId: toValue(tableId)
    }))
  }

  function reconnect() {
    if (wsObject.status.value === 'CLOSED') {
      wsObject.open()
      wsObject.send(encode<SendMessage>({
        action: WsActions.Reconnect,
        tableId: toValue(tableId),
        username: TEST_USERNAME
      }))
    }
  }

  function flipCard() {
    wsObject.send(encode<SendMessage>({
      action: WsActions.FlipCard,
      tableId: tableId.value,
      playerId: TEST_USERNAME
    }))
  }

  return {
    tableId,
    wsObject,
    initialDeck,
    tableDetails,
    isConnected,
    /**
     * Create a new table by sending a message to the server.
     * If the WebSocket is not connected, it will first establish the connection and then
     * send the message. This function can be called when the user clicks a button
     * to create a new game table.
     */
    createTable,
    getDeck,
    joinTable,
    reconnect,
    flipCard
  }
})
