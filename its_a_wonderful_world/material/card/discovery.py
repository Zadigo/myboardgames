from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin


@dataclass
class AlexandersTomb(BaseDevelopmentCardMixin):
    name = "Alexander's Tomb"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class AncientAstronauts(BaseDevelopmentCardMixin):
    name = "Ancient Astronauts"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class ArchOfTheCovenant(BaseDevelopmentCardMixin):
    name = "Arch Of The Covenant"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Atlantis(BaseDevelopmentCardMixin):
    name = "Atlantis"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class BermudaTriangle(BaseDevelopmentCardMixin):
    name = "Bermuda Triangle"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class BlackbeardsTreasure(BaseDevelopmentCardMixin):
    name = "Blackbeard's Treasure"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class CenterOfTheEarth(BaseDevelopmentCardMixin):
    name = "Center Of The Earth"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class CitiesOfGold(BaseDevelopmentCardMixin):
    name = "Cities Of Gold"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class CityOfAgartha(BaseDevelopmentCardMixin):
    name = "City Of Agartha"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class FountainOfYouth(BaseDevelopmentCardMixin):
    name = "Fountain Of Youth"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class GardensOfTheHesperides(BaseDevelopmentCardMixin):
    name = "Gardens Of The Hesperides"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class IslandOfAvalon(BaseDevelopmentCardMixin):
    name = "Island Of Avalon"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class KingSalomonsMines(BaseDevelopmentCardMixin):
    name = "King Salomon's Mines"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class LostContinentOfMu(BaseDevelopmentCardMixin):
    name = "Lost Continent Of Mu"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class ParallelDimension(BaseDevelopmentCardMixin):
    name = "Parallel Dimension"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Roswell(BaseDevelopmentCardMixin):
    name = "Roswell"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class TreasureOfTheTemplars(BaseDevelopmentCardMixin):
    name = "Treasure Of The Templars"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


DISCOVERY_CARDS = (
    AlexandersTomb,
    AncientAstronauts,
    ArchOfTheCovenant,
    Atlantis,
    BermudaTriangle,
    BlackbeardsTreasure,
    CenterOfTheEarth,
    CitiesOfGold,
    CityOfAgartha,
    FountainOfYouth,
    GardensOfTheHesperides,
    IslandOfAvalon,
    KingSalomonsMines,
    LostContinentOfMu,
    ParallelDimension,
    Roswell,
    TreasureOfTheTemplars,
)
