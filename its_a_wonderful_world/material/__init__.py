from its_a_wonderful_world.typings import CharacterTokenTypes


class GeneralCharacter:
    name: str = CharacterTokenTypes.GENERAL.value
    points: int = 1


class FinancierCharacter:
    name: str = CharacterTokenTypes.FINANCIER.value
    points: int = 1
