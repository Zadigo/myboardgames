

export const useGameWebsocket = createGlobalState(() => {
  const { ws } = useWebSocket('ws://0.0.0.0:8000/ws/game/1234',{
    immediate: true,
    // autoReconnect: true,
    onMessage: (event) => {
      console.log('WebSocket message received:', event.data)
    }
  })

  const isConnected = computed(() => ws.value?.readyState === WebSocket.OPEN)

  return {
    wsObject: ws,
    isConnected
  }
})
