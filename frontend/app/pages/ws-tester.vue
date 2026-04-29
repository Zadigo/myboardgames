<template>
  <section id="test" class="relative">
    <div class="grid grid-cols-2">
      <div class="p-5 space-x-2">
        <nuxt-button @click="ws.open()">
          1. Open
        </nuxt-button>
        
        <nuxt-button @click="identify">
          2. Identify
        </nuxt-button>
        
        <nuxt-button @click="createTable">
          3. Create table
        </nuxt-button>

        <nuxt-button @click="startGame">
          3. Start game
        </nuxt-button>

        <nuxt-button @click="ws.close()">
          Simulate close
        </nuxt-button>
      </div>
  
      <div class="p-5 bg-primary-100/10 shadow-2xl backdrop-blur-3xl w-full h-screen overflow-y-scroll space-y-3">
        <div class="p-2">
          <nuxt-button @click="() => messages = []">
            <icon name="lucide:trash" />
          </nuxt-button>
        </div>
        
        <pre v-for="(message, idx) in messages" :key="idx" class="overflow-x-scroll p-3 bg-primary-50 dark:bg-primary-950 rounded-lg">
          {{ message }}
        </pre>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import type { SplendorPlayingTable } from '~/types/splendor'

const { decode, encode } = useWWebsocketMessages2()

type RM = {
  must_identify: {
    playerUuid: string
  }
  table_created: SplendorPlayingTable
  identified: {}
  game_started: {
    playingTable: SplendorPlayingTable
  }
}

type SM = {
  identify: [
    {
      action: string
      playerUuid: string
      username: string
    }
  ],
  create_table: [
    {
      action: string
      playerUuid: string
    }
  ],
  start_game: [
    {
      action: string
      tableUuid: string
    }
  ]
}

const playerUuid = ref<string>('')
const tableUuid = ref<string>('')
const messages = ref([])

const ws = useWebSocket('ws://127.0.0.1:9000/ws/splendor/live', {
  immediate: false,
  onMessage(_ws, event) {
    const decoder = decode<RM>(event.data)
    
    decoder('must_identify', (data) => {
      if (isDefined(data)) {
        messages.value.push(data)
        if (isDefined(data.playerUuid)) {
          playerUuid.value = data.playerUuid
        }
      }
    })

    decoder('table_created', (data) => {
      if (isDefined(data)) {
        messages.value.push(data)
        if (isDefined(data.playingTable)) {
          tableUuid.value = data.playingTable?.details?.uuid
        }
      }
    })
  }
})

function identify() {
  ws.send(
    encode<SM>('identify', {
      action: 'identify',
      playerUuid: playerUuid.value,
      username: 'alice'
    })
  )
}

function createTable() {
  ws.send(
    encode<SM>('create_table', {
      action: 'create_table',
      playerUuid: playerUuid.value
    })
  )
}

function startGame() {
  ws.send(
    encode<SM>('start_game', {
      action: 'start_game',
      tableUuid: tableUuid.value
    })
  )
}
</script>
