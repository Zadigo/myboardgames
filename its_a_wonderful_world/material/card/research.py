from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin


@dataclass
class Aquaculture(BaseDevelopmentCardMixin):
    # name: str = "Aquaculture"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.GOLD,
    #         ResourceCubesTypes.GOLD,
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.SCIENCE

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    # victory_points: int = 1
    # is_combo_victory_points: bool = True
    # combo_victory_points_per_token_type: list[CharacterTokenTypes] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class ArtificialIntelligence(BaseDevelopmentCardMixin):
    # name: str = "Artificial Intelligence"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE

    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.SCIENCE

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    # multiplies_construction_bonus: bool = True
    # multiply_construction_bonus_constraint: Optional[DevelopmentTypes] = DevelopmentTypes.VEHICLE

    # victory_points: int = 1

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class ArtificialSun(BaseDevelopmentCardMixin):
    # name: str = "Artificial Sun"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.GOLD,
    #         ResourceCubesTypes.GOLD,
    #         ResourceCubesTypes.KRYSTALLIUM,
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.SCIENCE

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    # multiplies_construction_bonus: bool = True
    # multiply_construction_bonus_constraint: Optional[DevelopmentTypes] = DevelopmentTypes.VEHICLE

    # victory_points: int = 1

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class BionicCrafts(BaseDevelopmentCardMixin):
    # name: str = "Bionic Crafts"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.GENERAL
    #     ]
    # )

    # victory_points: int = 4

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class ClimateControl(BaseDevelopmentCardMixin):
    # name: str = "Bionic Crafts"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.ENERGY

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.ENERGY,
    #         ResourceCubesTypes.GOLD
    #     ]
    # )

    # victory_points: int = 2

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class CryoPreservation(BaseDevelopmentCardMixin):
    # name: str = "Cryo Preservation"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.GOLD

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    # victory_points: int = 1
    # combo_victory_points_per_token_type: list[CharacterTokenTypes] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER
    #     ]
    # )

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class GeneticUpgrades(BaseDevelopmentCardMixin):
    # name: str = "Genetic Upgrade"
    # card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    # construction_cost: list[ResourceCubesTypes] = field(
    #     default_factory=lambda: [
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE,
    #         ResourceCubesTypes.SCIENCE
    #     ]
    # )

    # recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.GOLD

    # construction_bonus: list[TypeConstructionBonuses] = field(
    #     default_factory=lambda: [
    #         CharacterTokenTypes.FINANCIER,
    #         CharacterTokenTypes.FINANCIER,
    #     ]
    # )

    # victory_points: int = 3

    @classmethod
    def number_of_copies(cls):
        return 1


@dataclass
class GravityInverter(BaseDevelopmentCardMixin):
    name = "Gravity Inverter"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class HumanCloning(BaseDevelopmentCardMixin):
    name = "Human Cloning"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class MegaBomb(BaseDevelopmentCardMixin):
    name = "Mega Bomb"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Neuroscience(BaseDevelopmentCardMixin):
    name = "Neuroscience"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class PlanetaryArchives(BaseDevelopmentCardMixin):
    name = "Planetary Archives"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class QuantumGenerator(BaseDevelopmentCardMixin):
    name = "Quantum Generator"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class RobotAssistants(BaseDevelopmentCardMixin):
    name = "Robot Assistants"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class RoboticAnimals(BaseDevelopmentCardMixin):
    name = "Robotic Animals"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Satellites(BaseDevelopmentCardMixin):
    name = "Satellites"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class SecurityAutomatons(BaseDevelopmentCardMixin):
    name = "Security Automatons"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class SuperSoldiers(BaseDevelopmentCardMixin):
    name = "Super Soldiers"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class SuperSonar(BaseDevelopmentCardMixin):
    name = "Super Sonar"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class SuperComputer(BaseDevelopmentCardMixin):
    name = "Super Computer"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Teleportation(BaseDevelopmentCardMixin):
    name = "Teleportation"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class TimeTravel(BaseDevelopmentCardMixin):
    name = "Time Travel"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class Transmutation(BaseDevelopmentCardMixin):
    name = "Transmutation"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class UnknownTechnology(BaseDevelopmentCardMixin):
    name = "Unknown Technology"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


@dataclass
class VirtualSecurity(BaseDevelopmentCardMixin):
    name = "Virtual Security"

    @classmethod
    def number_of_copies(cls) -> int:
        return 1


RESEARCH_CARDS = (
    Aquaculture,
    ArtificialIntelligence,
    ArtificialSun,
    BionicCrafts,
    ClimateControl,
    CryoPreservation,
    GeneticUpgrades,
    GravityInverter,
    HumanCloning,
    MegaBomb,
    Neuroscience,
    PlanetaryArchives,
    QuantumGenerator,
    RobotAssistants,
    RoboticAnimals,
    Satellites,
    SecurityAutomatons,
    SuperSoldiers,
    SuperSonar,
    SuperComputer,
    Teleportation,
    TimeTravel,
    Transmutation,
    UnknownTechnology,
    VirtualSecurity,
)
