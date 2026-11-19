export type Nullable<T> = T | null

export type Undefineable<T> = T | undefined

export type Optional<T> = T | null | undefined

export type Defined<T> = T extends null | undefined ? never : T

export type IsDefined<T> = T extends null | undefined ? false : true
