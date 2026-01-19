from dataclasses import dataclass, field
from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import CardCategory, CharacterTokenTypes, DevelopmentTypes, ResourceCubesTypes, TypeConstructionBonuses


@dataclass
class AirborneLaboratory(BaseDevelopmentCardMixin):
    name: str = "Airborne Laboratory"
    image: str = "airborne-laboratory.jpg"
    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.SCIENCE
    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.SCIENCE,
            ResourceCubesTypes.ENERGY
        ]
    )

    @staticmethod
    def number_of_copies():
        return 3


@dataclass
class AircraftCarrier(BaseDevelopmentCardMixin):
    name: str = "Aircraft Carrier"
    image: str = "aircraft-carrier.jpg"

    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    multiplies_construction_bonus: bool = True
    multiply_construction_bonus_constraint = DevelopmentTypes.VEHICLE

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION
        ]
    )

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class IceBreaker(BaseDevelopmentCardMixin):
    name: str = "Ice Breaker"
    image: str = "icebreaker.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    card_category: CardCategory = CardCategory.DEFAULT

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.SCIENCE
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.EXPLORATION

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION,
            ResourceCubesTypes.EXPLORATION
        ]
    )

    @staticmethod
    def number_of_copies():
        return 4


@dataclass
class Juggernaut(BaseDevelopmentCardMixin):
    name: str = "Juggernaut"
    image: str = "juggernaut.jpg"
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    card_category: CardCategory = CardCategory.DEFAULT

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.KRYSTALLIUM,
        ]
    )

    construction_bonus: list[TypeConstructionBonuses] = field(
        default_factory=lambda: [
            CharacterTokenTypes.GENERAL,
            CharacterTokenTypes.GENERAL
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION,
            ResourceCubesTypes.EXPLORATION
        ]
    )

    is_combo_victory_points: bool = True
    victory_points: int = 1
    combo_victory_points_per_card_type: list[DevelopmentTypes] = field(
        default_factory=lambda: [
            DevelopmentTypes.VEHICLE
        ]
    )

    @staticmethod
    def number_of_copies():
        return 1


@dataclass
class MegaDrill(BaseDevelopmentCardMixin):
    name: str = "Mega Drill"
    image = "mega-drill.jpg"

    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION
        ]
    )

    @staticmethod
    def number_of_copies():
        return 4


@dataclass
class SaucerSquadron(BaseDevelopmentCardMixin):
    name: str = "Saucer Squadron"
    image: str = "saucer-squadron.jpg"
    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.SCIENCE,
            ResourceCubesTypes.SCIENCE
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.SCIENCE

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION,
            ResourceCubesTypes.EXPLORATION,
            ResourceCubesTypes.EXPLORATION,
        ]
    )

    @staticmethod
    def number_of_copies():
        return 2


@dataclass
class Submarine(BaseDevelopmentCardMixin):
    name: str = "Submarine"
    image: str = "submarine.jpg"
    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION,
            ResourceCubesTypes.EXPLORATION
        ]
    )

    construction_bonus: list[TypeConstructionBonuses] = field(
        default_factory=lambda: [
            CharacterTokenTypes.GENERAL
        ]
    )

    @staticmethod
    def number_of_copies():
        return 3


@dataclass
class TankDivision(BaseDevelopmentCardMixin):
    name: str = "Tank Division"
    image: str = "tank-division.jpg"
    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.MATERIAL,
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION
        ]
    )

    construction_bonus: list[TypeConstructionBonuses] = field(
        default_factory=lambda: [
            CharacterTokenTypes.GENERAL
        ]
    )

    @staticmethod
    def number_of_copies():
        return 7


@dataclass
class Zeppelin(BaseDevelopmentCardMixin):
    name: str = "Zeppelin"
    image: str = "zeppelin.jpg"

    card_category: CardCategory = CardCategory.DEFAULT
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.ENERGY,
            ResourceCubesTypes.ENERGY,
        ]
    )

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.EXPLORATION

    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=lambda: [
            ResourceCubesTypes.EXPLORATION
        ]
    )

    @staticmethod
    def number_of_copies():
        return 6


VEHICLE_CARDS = (
    AirborneLaboratory,
    AircraftCarrier,
    IceBreaker,
    Juggernaut,
    MegaDrill,
    SaucerSquadron,
    Submarine,
    TankDivision,
    Zeppelin,
)
