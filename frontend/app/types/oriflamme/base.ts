export enum BaseCharacters {
  Archer = 'Archer',
  Soldier = 'Soldier',
  Spy = 'Spy',
  Heir = 'Heir',
  Shapeshifter = 'Shapeshifter',
  Lord = 'Lord',
  Assassination = 'Assassination',
  RoyalDecree = 'Royal Decree',
  Conspiracy = 'Conspiracy',
  Ambush = 'Ambush'
}

export enum AblazeCharactersEnum {
  Prince = 'Prince',
  Twin = 'Twin',
  Queen = 'Queen',
  Cutthroat = 'Cutthroat',
  Apothecary = 'Apothecary',
  Felon = 'Felon',
  Schemer = 'Schemer',
  Impersonation = 'Impersonation',
  Plot = 'Plot',
  Trap = 'Trap',
  Bribery = 'Bribery'
}

export type AblazeCharacters = `${AblazeCharactersEnum}`

export enum AllianceCharactersEnum {
  Ignore = 'Ignore'
}

export type AllianceCharacters = `${AllianceCharactersEnum}`

export type AllCharacters = BaseCharacters | AblazeCharacters | AllianceCharacters
