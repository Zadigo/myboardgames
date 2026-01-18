export type Refeable<T> = Ref<T>

export type Games = "It's a Wonderful World" | "Dune Imperium"

export interface Player {
  id: number
  gamertag: string;
  firstname: string
  lastname: string
  gender: 'W' | 'M'
  points: number
}

export type NewPlayer = Omit<Player, 'id'> & {
  game: Games
  position?: number
}

export interface NewGame {
  game: Games
  date: string
}
