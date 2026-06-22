export * from './__fixtures__'

type BaseWsMessage<T = Record<string, unknown>> = {
  action: string
} & {
  [K in keyof T]: T[K]
}

export function useWebsocketMessage() {
  function decode<T extends BaseWsMessage>(message: string): T | null {
    try {
      return JSON.parse(message) as T
    } catch (error) {
      console.error('Failed to decode WebSocket message:', error)
      return null
    }
  }

  function encode<T extends BaseWsMessage>(data: T): string {
    try {
      return JSON.stringify(data)
    } catch (error) {
      console.error('Failed to encode WebSocket message:', error)
      return ''
    }
  }

  return {
    decode,
    encode
  }
}
