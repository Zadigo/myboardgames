export type Arrayable<T> = T[]

export type BaseCard = {
  name: string
  type: string
  category: string
  image: string
  isDrafted: boolean
  has_construction_bonus: boolean
  has_character_token_bonus: boolean
}
