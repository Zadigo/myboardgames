export type Arrayable<T> = T[]

export type BaseCard = {
  name: string
  card_type: string
  category: string
  image: string
  isDrafted: boolean
  recycling_bonus: ResourceCubesTypes
  has_construction_bonus: boolean
  has_character_token_bonus: boolean
}

export enum _ResourceCubesTypes {
  All = 'All',
  Material = 'Material',
  Energy = 'Energy',
  Science = 'Science',
  Gold = 'Gold',
  Exploration = 'Exploration',
  Krystallium = 'Krystallium'
}

export type ResourceCubesTypes = keyof typeof _ResourceCubesTypes
