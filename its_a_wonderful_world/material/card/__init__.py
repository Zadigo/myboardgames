from dataclasses import dataclass, field
from typing import Optional
from uuid import uuid4

from its_a_wonderful_world.typings import (CardCategory, CharacterTokenTypes,
                                           DevelopmentTypes,
                                           ResourceCubesTypes,
                                           TypeConstructionBonuses,
                                           TypePlayerCard)


@dataclass
class BaseDevelopmentCardMixin:
    """
    Base class for all development cards.

    Args:
        name (str): The name of the card.
        image (str): The image filename for the card.
        card_id (str): The unique identifier for the card.
        card_type (DevelopmentTypes): The type of development card.
        card_category (CardCategory): The category of the card.
        construction_cost (list[ResourceCubesTypes]): The resources required to construct the card.
        construction_cost_resource_types (list[ResourceCubesTypes]): The types of resources provided by the card.
        construction_bonus (list[TypeConstructionBonuses]): Bonuses provided upon construction.
        multiplies_construction_bonus (bool): Whether this card multiplies construction bonuses.
        multiply_construction_bonus_constraint (Optional[DevelopmentTypes]): Constraint for multiplying bonuses.
        is_duo_scoring_card (bool): Whether this card scores for duo combinations.
        is_combo_victory_points (bool): Whether this card provides combo victory points.
        combo_victory_points_per_card_type (list[DevelopmentTypes]): Card types for combo victory points.
        combo_victory_points_per_token_type (list[CharacterTokenTypes]): Token types for combo victory points.
        victory_points (int): Victory points provided by the card.
        recycling_bonus (ResourceCubesTypes): Resource gained when recycling the card.
        corrupted_material_resource (list[ResourceCubesTypes]): Resources lost due to corruption.
        is_face_down (bool): Whether the card is face down.
        construction_progress (list[ResourceCubesTypes]): Progress made towards constructing the card.
    """
    name: str = ""
    image: str = ""
    card_id: str = field(default_factory=lambda: str(uuid4()))
    card_type: DevelopmentTypes = DevelopmentTypes.VEHICLE
    card_category: CardCategory = CardCategory.DEFAULT

    construction_cost: list[ResourceCubesTypes] = field(
        default_factory=list
    )
    construction_cost_resource_types: list[ResourceCubesTypes] = field(
        default_factory=list
    )
    construction_bonus: list[TypeConstructionBonuses] = field(
        default_factory=list
    )

    multiplies_construction_bonus: bool = False
    multiply_construction_bonus_constraint: Optional[DevelopmentTypes] = None

    is_duo_scoring_card: bool = False
    is_combo_victory_points: bool = False
    combo_victory_points_per_card_type: list[DevelopmentTypes] = field(
        default_factory=list
    )
    combo_victory_points_per_token_type: list[CharacterTokenTypes] = field(
        default_factory=list
    )
    victory_points: int = 0

    recycling_bonus: ResourceCubesTypes = ResourceCubesTypes.MATERIAL

    corrupted_material_resource: list[ResourceCubesTypes] = field(
        default_factory=list
    )

    is_face_down: bool = False

    construction_progress: list[ResourceCubesTypes] = field(
        default_factory=list
    )

    def __hash__(self):
        return hash(self.card_id)

    def __repr__(self):
        return f"<{self.__class__.__name__} name={self.name}>"

    @classmethod
    def number_of_copies(cls) -> int:
        """The total number of cards of this 
        type in the game."""
        return 0

    @property
    def construction_cost_count(self) -> int:
        """The total resource cost to build this card."""
        return len(self.construction_cost)

    @property
    def construction_cost_resource_types_count(self) -> int:
        """The total resource types this card provides."""
        return len(self.construction_cost_resource_types)

    @property
    def construction_progress_percentage(self) -> float:
        """The percentage increase in construction bonuses this card provides."""
        try:
            return len(self.construction_progress) / len(self.construction_cost)
        except ZeroDivisionError:
            return 0.0

    @property
    def card_difficulty(self):
        """The difficulty of constructing this card based on its construction cost."""
        return 0

    def construction_bonus_count(self, cards: list[TypePlayerCard]) -> int:
        """The total construction bonuses this card provides by multiplying
        the amount of cards the player has that meet the constraint."""
        count: int = 1  # Start with base 1 to represent the card itself

        for card in cards:
            if self.multiply_construction_bonus_constraint is None:
                continue

            if card.card_type == self.multiply_construction_bonus_constraint:
                count += 1
        return count

    def add_resource(self, resource: ResourceCubesTypes):
        """Add a resource to the construction progress of this card.

        Args:
            resource (ResourceCubesTypes): The resource to add.
        """
        if not self.construction_cost:
            raise ValueError("This card has no construction cost.")

        if self.is_face_down:
            raise ValueError("Cannot add resources to a face-down card.")

        if self.construction_progress_percentage >= 1.0:
            raise ValueError("Construction already complete.")

        if resource not in self.construction_cost:
            raise ValueError("Resource not in construction cost.")

        self.construction_progress.append(resource)
