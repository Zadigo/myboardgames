from dataclasses import dataclass
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import DevelopmentTypes


@dataclass
class Aquaculture(BaseDevelopmentCardMixin):
    name: str = "Aquaculture"
    image = "aquaculture.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class ArtificialIntelligence(BaseDevelopmentCardMixin):
    name: str = "Artificial Intelligence"
    image = "artificial-intelligence.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class ArtificialSun(BaseDevelopmentCardMixin):
    name: str = "Artificial Sun"
    image = "artificial-sun.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class BionicCrafts(BaseDevelopmentCardMixin):
    name: str = "Bionic Crafts"
    image = "bionic-crafts.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class ClimateControl(BaseDevelopmentCardMixin):
    name: str = "Bionic Crafts"
    image = "climate-control.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class CryoPreservation(BaseDevelopmentCardMixin):
    name: str = "Cryo Preservation"
    image = "cryopreservation.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class GeneticUpgrades(BaseDevelopmentCardMixin):
    name: str = "Genetic Upgrade"
    image = "genetic-upgrades.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

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

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class GravityInverter(BaseDevelopmentCardMixin):
    name = "Gravity Inverter"
    image = "gravity-inverter.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class HumanCloning(BaseDevelopmentCardMixin):
    name = "Human Cloning"
    image = "human-cloning.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class MegaBomb(BaseDevelopmentCardMixin):
    name = "Mega Bomb"
    image = "mega-bomb.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Neuroscience(BaseDevelopmentCardMixin):
    name = "Neuroscience"
    image = "neuroscience.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class PlanetaryArchives(BaseDevelopmentCardMixin):
    name = "Planetary Archives"
    image = "planetary-archives.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class QuantumGenerator(BaseDevelopmentCardMixin):
    name = "Quantum Generator"
    image = "quantum-generator.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class RobotAssistants(BaseDevelopmentCardMixin):
    name = "Robot Assistants"
    image = "robot-assistants.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class RoboticAnimals(BaseDevelopmentCardMixin):
    name = "Robotic Animals"
    image = "robotic-animals.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Satellites(BaseDevelopmentCardMixin):
    name = "Satellites"
    image = "satellites.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class SecurityAutomatons(BaseDevelopmentCardMixin):
    name = "Security Automatons"
    image = "security-automatons.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class SuperSoldiers(BaseDevelopmentCardMixin):
    name = "Super Soldiers"
    image = "super-soldiers.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class SuperSonar(BaseDevelopmentCardMixin):
    name = "Super Sonar"
    image = "super-sonar.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class SuperComputer(BaseDevelopmentCardMixin):
    name = "Super Computer"
    image = "supercomputer.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Teleportation(BaseDevelopmentCardMixin):
    name = "Teleportation"
    image = "teleportation.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class TimeTravel(BaseDevelopmentCardMixin):
    name = "Time Travel"
    image = "time-travel.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class Transmutation(BaseDevelopmentCardMixin):
    name = "Transmutation"
    image = "transmutation.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class UnknownTechnology(BaseDevelopmentCardMixin):
    name = "Unknown Technology"
    image = "unknown-technology.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
        return 1


@dataclass
class VirtualReality(BaseDevelopmentCardMixin):
    name = "Virtual Reality"
    image = "virtual-reality.jpg"
    card_type = DevelopmentTypes.RESEARCH

    @staticmethod
    def number_of_copies() -> int:
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
    VirtualReality,
)
