from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import DevelopmentTypes


@dataclass
class FinancialCenter(BaseDevelopmentCardMixin):
    name = "Financial Center"
    image = "financial-center.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 5


@dataclass
class GoldMine(BaseDevelopmentCardMixin):
    name = "Gold Mine"
    image = "gold-mine.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class IndustrialComplex(BaseDevelopmentCardMixin):
    name = "Industrial Complex"
    image = "industrial-complex.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 6


@dataclass
class MilitaryBase(BaseDevelopmentCardMixin):
    name = "Military Base"
    image = "military-base.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 6


@dataclass
class NuclearPlant(BaseDevelopmentCardMixin):
    name = "Nuclear Plant"
    image = "nuclear-plant.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 5


@dataclass
class OffshoreOilRig(BaseDevelopmentCardMixin):
    name = "Offshore Oil Rig"
    image = "offshore-oil-rig.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 5


@dataclass
class RecyclingPlant(BaseDevelopmentCardMixin):
    name = "Recycling Plant"
    image = "recycling-plant.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 7


@dataclass
class ResearchCenter(BaseDevelopmentCardMixin):
    name = "Research Center"
    image = "research-center.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 7


@dataclass
class TransportatonNetwork(BaseDevelopmentCardMixin):
    name = "Transportation Network"
    image = "transportation-network.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class WindTurbines(BaseDevelopmentCardMixin):
    name = "Wind Turbines"
    image = "wind-turbines.jpg"
    card_type = DevelopmentTypes.STRUCTURE

    @staticmethod
    def number_of_copies() -> int:
        return 7


STRUCTURE_CARDS = (
    FinancialCenter,
    GoldMine,
    IndustrialComplex,
    MilitaryBase,
    NuclearPlant,
    OffshoreOilRig,
    RecyclingPlant,
    ResearchCenter,
    TransportatonNetwork,
    WindTurbines,
)
