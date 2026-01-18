from enum import Enum
from typing import TYPE_CHECKING, Type

if TYPE_CHECKING:
    from its_a_wonderful_world.logic.base import Game
    from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
    from its_a_wonderful_world.logic.player import Player


type TypeGame = 'Game'

type TypePlayerCard = 'BaseDevelopmentCardMixin'

type TypePlayer = 'Player'

type TypeConstructionBonuses = ResourceCubesTypes | CharacterTokenTypes


class PhaseTypes(Enum):
    """The different phases in a round 
    of It's a Wonderful World."""

    DRAFT = "Draft"
    PLANNING = "Planning"
    PRODUCTION = "Production"


class CardCategory(Enum):
    DEFAULT = "Default"
    ASCENSION_AND_CORRUPTION = "Ascension and Corruption"


class DevelopmentTypes(Enum):
    """It's a Wonderful World has different types of cards that
    represent different types of development."""

    STRUCTURE = "Structure"
    VEHICLE = "Vehicle"
    RESEARCH = "Research"
    PROJECT = "Project"
    DISCOVERY = "Discovery"


class MoveTypes(Enum):
    """Different types of moves a player can make during their turn."""

    DEAL_DEVELOPMENT_CARDS = 'Deal Development Cards'
    CHOOSE_DEVELOPMENT_CARD = 'Choose Development Card'
    REVEAL_CHOSEN_CARDS = 'Reveal Chosen Cards'
    PASS_CARDS = 'Pass Cards'
    DISCARD_LEFTOVER_CARDS = 'Discard Leftover Cards'
    START_PHASE = 'Start Phase'
    RECYCLE = 'Recycle'
    SLATE_FOR_CONSTRUCTION = 'Slate For Construction'
    PLACE_RESOURCE = 'Place Resource'
    COMPLETE_CONSTRUCTION = 'Complete Construction'
    TRANSFORM_INTO_KRYSTALLIUM = 'Transform Into Krystallium'
    TELL_YOU_ARE_READY = 'Tell You Are Ready'
    PRODUCE = 'Produce'
    RECEIVE_CHARACTER = 'Receive Character'
    PLACE_CHARACTER = 'Place Character'
    CONCEDE = 'Concede'


class ResourceCubesTypes(Enum):
    """The six types of resource cubes in the game."""

    MATERIAL = "Material"
    ENERGY = "Energy"
    SCIENCE = "Science"
    GOLD = "Gold"
    EXPLORATION = "Exploration"
    KRYSTALLIUM = "Krystallium"


class CharacterTokenTypes(Enum):
    GENERAL = "General"
    FINANCIER = "Financier"


class RotationTypes(Enum):
    """The direction the in which the cards should be distributed
    since it changes for each round."""

    CLOCKWISE = "Clockwise"
    COUNTER_CLOCKWISE = "Counter Clockwise"
