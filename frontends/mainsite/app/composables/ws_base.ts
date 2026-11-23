export type BaseResponseOptions = {
  must_identify: {
    client_id: Nullable<string>
  }
  join_website: {
    action: 'join_website',
    rpc: string
    rpc_state: string,
    client_id: string,
    message: string,
    json_message: Record<string, any>,
    router: {
      routes: {
        hash: string,
        name_fr: string,
        name_en: string,
        domain: string,
        port: number,
        active: boolean,
        environment: string,
        version: number
      }[]
    }
  }
}

export type BaseActionOptions = {
  create_game: [{ playerUuid: Nullable<string> }]
  ping_rpc: [{ client_id: string }]
}

/**
 * Creates a global long-lived websocket connection
 * to the main Golang entrypoint server
 */
export const useWebsiteWebsocket = createGlobalState(() => {
  const clientId = useSessionStorage<string>('client_id', '')
  const routes = useLocalStorage<BaseResponseOptions['join_website']['router']>('routes', null, {
    serializer: {
      read(raw) {
        return JSON.parse(raw)
      },
      write(raw) {
        return JSON.stringify(raw)
      }
    }
  })

  const { encode, decode } = useWWebsocketMessages2()

  const wsObject = useWebSocket('http://127.0.0.1:9000/v1/', {
    immediate: true,
    onConnected() {},
    onMessage(_ws, event) {
      const decoder = decode<BaseResponseOptions>(event.data)

      decoder('join_website', (data) => {
        if (data) {
          if (!isDefined(clientId.value)) {
            clientId.value = data.client_id
          }
          routes.value = data.router
        }
      })
    },
    onError() {}
  })

  const isConnected = computed(() => toValue(wsObject.status) == 'OPEN')

  return {
    clientId,
    wsObject,
    isConnected,
    encode
  }
})
