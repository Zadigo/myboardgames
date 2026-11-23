<template>
  <section id="some-game">
    <u-button @click="createGame">
      Open Game
    </u-button>

    <u-button @click="joinGame">
      Join Game
    </u-button>
  </section>
</template>

<script lang="ts" setup>
const routeParams = useRoute()
const router = useRouter()

const { clientId, encode, wsObject } = useWebsiteWebsocket()

function pingAction() {
  const pingMessage = encode<BaseActionOptions>('ping_rpc', { client_id: clientId.value })
  return wsObject.send(pingMessage)

}

function getRoute() {
  return `/games/${routeParams.params.id}/12345-table-id?r=235dc06ff2619e0500f6992bf49d5e3460565257f57d2bda03bfffeb7c572ae1`
}

function createGame() {
  const result = pingAction()
  if (result) {
    console.log(wsObject.data.value)
    router.push(getRoute())
  }
}

function joinGame() {
  const result = pingAction()
  if (result) {
    console.log(wsObject.data.value)
    router.push(getRoute())
  }
}
</script>
