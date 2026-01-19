from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import DevelopmentTypes


@dataclass
class CasinoCity(BaseDevelopmentCardMixin):
    name = "Casino City"
    image = "casino-city.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class EspionageAgency(BaseDevelopmentCardMixin):
    name = "Espionage Agency"
    image = "espionage-agency.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class GiantDam(BaseDevelopmentCardMixin):
    name = "Giant Dam"
    image = "giant-dam.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class GiantTower(BaseDevelopmentCardMixin):
    name = "Giant Tower"
    image = "giant-tower.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class HarborZone(BaseDevelopmentCardMixin):
    name = "Harbor Zone"
    image = "harbor-zone.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class LunarBase(BaseDevelopmentCardMixin):
    name = "Lunar Base"
    image = "lunar-base.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class MagneticTrain(BaseDevelopmentCardMixin):
    name = "Magnetic Train"
    image = "magnetic-train.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class Museum(BaseDevelopmentCardMixin):
    name = "Museum"
    image = "museum.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class NationalMonument(BaseDevelopmentCardMixin):
    name = "National Monument"
    image = "national-monument.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class PolarBase(BaseDevelopmentCardMixin):
    name = "Polar Base"
    image = "polar-base.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


class PropagandaCenter(BaseDevelopmentCardMixin):
    name = "Propaganda Center"
    image = "propaganda-center.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class SecretLaboratory(BaseDevelopmentCardMixin):
    name = "Secret Laboratory"
    image = "secret-laboratory.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class SecretSociety(BaseDevelopmentCardMixin):
    name = "Secret Society"
    image = "secret-society.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class SolarCannon(BaseDevelopmentCardMixin):
    name = "Solar Cannon"
    image = "solar-cannon.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class SpaceElevator(BaseDevelopmentCardMixin):
    name = "Space Elevator"
    image = "space-elevator.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class UndergroundCity(BaseDevelopmentCardMixin):
    name = "Underground City"
    image = "underground-city.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class UnderwaterCity(BaseDevelopmentCardMixin):
    name = "Underwater City"
    image = "underwater-city.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 2


@dataclass
class UniversalExposition(BaseDevelopmentCardMixin):
    name = "Universal Exposition"
    image = "universal-exposition.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Univeristy(BaseDevelopmentCardMixin):
    name = "University"
    image = "university.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class WorldCongress(BaseDevelopmentCardMixin):
    name = "World Congress"
    image = "world-congress.jpg"
    card_type = DevelopmentTypes.PROJECT

    @staticmethod
    def number_of_copies() -> int:
        return 1


PROJECT_CARDS = (
    CasinoCity,
    EspionageAgency,
    GiantDam,
    GiantTower,
    HarborZone,
    LunarBase,
    MagneticTrain,
    Museum,
    NationalMonument,
    PolarBase,
    PropagandaCenter,
    SecretLaboratory,
    SecretSociety,
    SolarCannon,
    SpaceElevator,
    UndergroundCity,
    UnderwaterCity,
    UniversalExposition,
    Univeristy,
    WorldCongress,
)
