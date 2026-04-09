export function useOriflammeActionsComposable(ws: Ref<WebSocket | undefined>) {
  function reveal() {
    ws.value?.send('a')
  }

  function applyEffect() {

  }

  function placeToken() {

  }

  return {
    reveal,
    applyEffect,
    placeToken
  }
}
