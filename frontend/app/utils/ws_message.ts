type DefaultResponseOptions = {
  action: string
  error: string
  message: string
}

// type DecodeResponse<T> = DefaultResponseOptions & {
//   [ K in keyof T ]: T[K]
// }

export function useWWebsocketMessages2() {
  function encode<T extends Record<string, Record<string, unknown>[]>>(action: keyof T, ...args: T[keyof T]) {
    return JSON.stringify({ action, ...args[0] })
  }

  /**
   * Return a new decoder function that can be used to decode messages from the websocket.
   */
  function decode<R>(message: string) {
    /**
     * A decoder function that listens to messages from the websocket that matches the specified
     * action. The callback function wil be called with the decoded data if the action matches, or undefined if it doesn't match.
     * @param action - The action we want to decode from the message
     * @param callback - A callback function that will be called with the decoded data if the action matches, or undefined if it doesn't match
     */
    return function <K extends keyof R>(action: K, callback: (data: R & DefaultResponseOptions | undefined) => void) {
      try {
        const wsData = JSON.parse(message) as R & DefaultResponseOptions
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

  return {
    encode,
    decode
  }
}

const { decode } = useWWebsocketMessages2()
const encoder = decode<{ firstname: string }>('')
encoder('firstname', (d) => {
  return d ? d.action : ''
})
