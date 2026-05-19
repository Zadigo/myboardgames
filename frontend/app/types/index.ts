export type * from './flipseven'
export type * from './oriflamme'
export type * from './splendor'

export type Nullable<T> = T | null

export type Undefineable<T> = T | undefined

export type Optional<T> = T | null | undefined

export type Defined<T> = T extends null | undefined ? never : T

export type IsDefined<T> = T extends null | undefined ? false : true

export const isEmpty = <T>(value: T): value is Defined<T> => {
  return value !== null && value !== undefined && (typeof value !== 'string' || value.trim() !== '')
}
