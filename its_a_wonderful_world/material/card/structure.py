from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin


@dataclass
class FinancialCenter(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 5


@dataclass
class GoldMine(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class IndustrialComplex(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 6


@dataclass
class MilitaryBase(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 6


@dataclass
class NuclearPlant(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 5


@dataclass
class OffshoreOilRig(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 5


@dataclass
class RecyclingPlant(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 7


@dataclass
class ResearchCenter(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 7


@dataclass
class TransportatonNetwork(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class WindTurbines(BaseDevelopmentCardMixin):
    @classmethod
    def number_of_copies(cls) -> int:
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
