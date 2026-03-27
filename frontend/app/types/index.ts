export type Arrayable<T> = T[]


export enum _CardCategory {
  DEFAULT = 'Default',
  ASCENSION_AND_CORRUPTION = 'Ascension and Corruption',
}

export type CardCategory = keyof typeof _CardCategory

export enum _DevelopmentTypes {
  STRUCTURE = "Structure",
  VEHICLE = "Vehicle",
  RESEARCH = "Research",
  PROJECT = "Project",
  DISCOVERY = "Discovery"
}

export type DevelopmentTypes = keyof typeof _DevelopmentTypes

export type BaseCard = {
  name: string
  card_type: DevelopmentTypes
  category: CardCategory
  image: string
  isDrafted: boolean
  number_of_copies: number
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
