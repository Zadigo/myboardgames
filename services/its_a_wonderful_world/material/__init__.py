from its_a_wonderful_world.typings import CharacterTokenTypes


from its_a_wonderful_world.material.card.research import RESEARCH_CARDS
from its_a_wonderful_world.material.card.structure import STRUCTURE_CARDS
from its_a_wonderful_world.material.card.vehicle import VEHICLE_CARDS
from its_a_wonderful_world.material.card.discovery import DISCOVERY_CARDS
from its_a_wonderful_world.material.card.project import PROJECT_CARDS
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin


class GeneralCharacter:
    name: str = CharacterTokenTypes.GENERAL.value
    points: int = 1


class FinancierCharacter:
    name: str = CharacterTokenTypes.FINANCIER.value
    points: int = 1


__all__ = [
    'BaseDevelopmentCardMixin',
    'RESEARCH_CARDS',
    'STRUCTURE_CARDS',
    'VEHICLE_CARDS',
    'DISCOVERY_CARDS',
    'PROJECT_CARDS'
]
