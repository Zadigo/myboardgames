<template>
  <section id="add-player">
    <nuxt-container>
      <!-- <nuxt-card>
        <nuxt-input v-model="newPlayer.gamertag" placeholder="Player gamertag" />
        <nuxt-input v-model="newPlayer.firstname" placeholder="Player firstname" />
        <nuxt-input v-model="newPlayer.lastname" placeholder="Player lastname" />
        <nuxt-select v-model="newPlayer.gender" :items="['M', 'F']" placeholder="Select a gender" />
        <nuxt-input-number v-model="newPlayer.points" :min="1" :max="50" placeholder="Player points" />
        <nuxt-select v-model="newPlayer.game" :items="gameNamesFixtures" placeholder="Select a game" />
      </nuxt-card> -->
      <!-- <new-participant-block v-model="newPlayer" /> -->

      <nuxt-select v-model="newGame.game" :items="gameNamesFixtures" placeholder="Game name" />
      <nuxt-input v-model="newGame.date" type="date" placeholder="Game date" />

      <nuxt-input-number v-model="participantNumber" :min="1" :max="10" />
      <nuxt-button @click="addParticipant">Add participant</nuxt-button>
      
      <template v-for="(participant, index) in orderedParticipants" :key="index">
        <template v-if="orderedParticipants[index]">
          <new-participant-block v-model="orderedParticipants[index]" :index="index" :count="orderedParticipants.length" />
        </template>
      </template>
      
    </nuxt-container>
  </section>
</template>

<script setup lang="ts">
import { gameNamesFixtures } from '~/data'
import type { NewGame, NewPlayer, Refeable } from '~/types'

function useCreateGame() {
  const newGame = ref<NewGame>({
    game: 'It\'s a Wonderful World',
    date: new Date().toISOString().substring(0, 10)
  })

  return {
    newGame
  }
}

function useCreatePlayer(game: Refeable<NewGame>) {
  const participantNumber = ref<number>(1)
  const participants = reactive<NewPlayer[]>([])

  const newPlayer = ref<NewPlayer>({
    gamertag: '',
    firstname: '',
    lastname: '',
    gender: 'M',
    points: 0,
    game: 'It\'s a Wonderful World',
    position: 0
  })

  function addParticipant() {
    Array.from({ length: participantNumber.value }).forEach(() => {
      participants.push({
        gamertag: '',
        firstname: '',
        lastname: '',
        gender: 'M',
        points: 0,
        game: game.value.game,
        position: 0
      })
    })  
  }

  const orderedParticipants = useSorted(participants, (a, b) => b.points - a.points)

  return {
    newPlayer,
    addParticipant,
    participantNumber,
    participants,
    orderedParticipants
  }
}

function useGamePoints(newGame: Refeable<NewGame>) {
  const winnerPoints = ref<number>(1)
  const runnerUpPoints = ref<number>(50)
  const runnerUp2Points = ref<number>(25)
  const maxPointWinners = ref<number>(3)

  switch (newGame.value.game) {
    case 'It\'s a Wonderful World':
      winnerPoints.value = 50
      runnerUpPoints.value = 25
      runnerUp2Points.value = 5
      maxPointWinners.value = 3
      break

    case 'Dune Imperium':
      winnerPoints.value = 50
      runnerUpPoints.value = 25
      runnerUp2Points.value = 5
      maxPointWinners.value = 3
      break

    default:
      winnerPoints.value = 0
      runnerUpPoints.value = 0
      runnerUp2Points.value = 0
      maxPointWinners.value = 3
      break
  }

  return {
    winnerPoints,
    runnerUpPoints,
    runnerUp2Points,
    maxPointWinners
  }
}

const { newGame } = useCreateGame()
const { newPlayer, participantNumber, participants, orderedParticipants, addParticipant } = useCreatePlayer(newGame)
</script>
