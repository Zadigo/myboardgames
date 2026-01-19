from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import DevelopmentTypes


@dataclass
class AlexandersTomb(BaseDevelopmentCardMixin):
    name = "Alexander's Tomb"
    image = "alexanders-tomb.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class AncientAstronauts(BaseDevelopmentCardMixin):
    name = "Ancient Astronauts"
    image = "ancient-astronauts.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class ArkOfTheCovenant(BaseDevelopmentCardMixin):
    name = "Ark Of The Covenant"
    image = "ark-of-the-covenant.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Atlantis(BaseDevelopmentCardMixin):
    name = "Atlantis"
    image = "atlantis.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class BermudaTriangle(BaseDevelopmentCardMixin):
    name = "Bermuda Triangle"
    image = "bermuda-triangle.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class BlackbeardsTreasure(BaseDevelopmentCardMixin):
    name = "Blackbeard's Treasure"
    image = "blackbeards-treasure.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class CenterOfTheEarth(BaseDevelopmentCardMixin):
    name = "Center Of The Earth"
    image = "center-of-the-earth.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class CitiesOfGold(BaseDevelopmentCardMixin):
    name = "Cities Of Gold"
    image = "cities-of-gold.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class CityOfAgartha(BaseDevelopmentCardMixin):
    name = "City Of Agartha"
    image = "city-of-agartha.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class FountainOfYouth(BaseDevelopmentCardMixin):
    name = "Fountain Of Youth"
    image = "fountain-of-youth.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class GardensOfTheHesperides(BaseDevelopmentCardMixin):
    name = "Gardens Of The Hesperides"
    image = "gardens-of-the-hesperides.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class IslandOfAvalon(BaseDevelopmentCardMixin):
    name = "Island Of Avalon"
    image = "island-of-avalon.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class KingSolomonsMines(BaseDevelopmentCardMixin):
    name = "King Solomon's Mines"
    image = "king-solomons-mines.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class LostContinentOfMu(BaseDevelopmentCardMixin):
    name = "Lost Continent Of Mu"
    image = "lost-continent-of-mu.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class ParallelDimension(BaseDevelopmentCardMixin):
    name = "Parallel Dimension"
    image = "parallel-dimension.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Roswell(BaseDevelopmentCardMixin):
    name = "Roswell"
    image = "roswell.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class TreasureOfTheTemplars(BaseDevelopmentCardMixin):
    name = "Treasure Of The Templars"
    image = "treasure-of-the-templars.jpg"
    card_type = DevelopmentTypes.DISCOVERY

    @staticmethod
    def number_of_copies() -> int:
        return 1


DISCOVERY_CARDS = (
    AlexandersTomb,
    AncientAstronauts,
    ArkOfTheCovenant,
    Atlantis,
    BermudaTriangle,
    BlackbeardsTreasure,
    CenterOfTheEarth,
    CitiesOfGold,
    CityOfAgartha,
    FountainOfYouth,
    GardensOfTheHesperides,
    IslandOfAvalon,
    KingSolomonsMines,
    LostContinentOfMu,
    ParallelDimension,
    Roswell,
    TreasureOfTheTemplars,
)
