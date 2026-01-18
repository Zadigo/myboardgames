from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin


@dataclass
class CasinoCity(BaseDevelopmentCardMixin):
    name = "Casino City"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class EspionageAgency(BaseDevelopmentCardMixin):
    name = "Espionage Agency"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class GiantDam(BaseDevelopmentCardMixin):
    name = "Espionage Agency"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class GiantTower(BaseDevelopmentCardMixin):
    name = "Giant Tower"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class HarborZone(BaseDevelopmentCardMixin):
    name = "Harbor Zone"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class LunarBase(BaseDevelopmentCardMixin):
    name = "Lunar Base"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class MagneticTrain(BaseDevelopmentCardMixin):
    name = "Magnetic Train"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class Museum(BaseDevelopmentCardMixin):
    name = "Museum"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class NationalMonument(BaseDevelopmentCardMixin):
    name = "National Monument"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class PolarBase(BaseDevelopmentCardMixin):
    name = "Polar Base"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


class PropagandaCenter(BaseDevelopmentCardMixin):
    name = "Propaganda Center"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class SecretLaboratory(BaseDevelopmentCardMixin):
    name = "Secret Laboratory"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class SecretSociety(BaseDevelopmentCardMixin):
    name = "Secret Society"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class SolarCannon(BaseDevelopmentCardMixin):
    name = "Solar Cannon"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class SpaceElevator(BaseDevelopmentCardMixin):
    name = "Space Elevator"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class UndergroundCity(BaseDevelopmentCardMixin):
    name = "Underground City"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class UnderwaterCity(BaseDevelopmentCardMixin):
    name = "Underwater City"

    @classmethod
    def number_of_copies(cls) -> int:
        return 2


@dataclass
class UniversalExposition(BaseDevelopmentCardMixin):
    name = "Universal Exposition"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Univeristy(BaseDevelopmentCardMixin):
    name = "University"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class WorldCongress(BaseDevelopmentCardMixin):
    name = "World Congress"

    @classmethod
    def number_of_copies(cls) -> int:
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
