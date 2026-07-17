import enum
import uuid
from abc import ABC, abstractmethod
from typing import Iterable, Optional, Sequence

from game.utils.base import MustResolve
from typings import CardResolutionActions, TypeAbstractCard, TypeGame, TypeWebsocketClient


class CardColors(enum.Enum):
    BLUE = "blue"
    RED = "red"
    YELLOW = "yellow"
    BLACK = "black"
    GREEN = "green"
    
class Cardtypes(enum.Enum):
    CHARACTER = "character"
    INTRIGUE = "intrigue"

class AbstractCardFactory(ABC):
    """An interface for creating card instances. Subclasses should 
    implement the create_cards method to generate specific card types."""

    cards_classes: Sequence[TypeAbstractCard] = []

    def __init__(self, game: TypeGame) -> None:
        self.game: TypeGame = game
    
    @abstractmethod
    def __iter__(self) -> Iterable[TypeAbstractCard]:
        pass

    def cards_from_color(self, color: CardColors) -> Iterable[TypeAbstractCard]:
        """Generate card instances of a specific color."""
        return list(filter(lambda card: card.color == color, self.__iter__()))
    

class AbstractCard(ABC):
    """Base class for all cards.

    Args:
        game (TypeGame): The game instance the card belongs to.
        name (str): The name of the card.

    Attributes:
        game (TypeGame): The game instance the card belongs to.
        name (str): The name of the card.
        partial_resolve (bool): Indicates if the card is partially resolved.
        color (Optional[CardColors]): The color of the card.
        tokens (int): The number of tokens on the card.
        owner (Optional[TypeWebsocketClient]): The owner of the card, if any.
        requires_partial_resolution (bool): Indicates if the user has to perform actions on the card in order for it to be fully resolved.
    """

    color: Optional[CardColors] = None
    stack: list[TypeAbstractCard] = []
    owner: TypeWebsocketClient | None = None
    requires_partial_resolution: bool = False
    tokens: int = 0
    card_type: Cardtypes = Cardtypes.CHARACTER

    def __init__(self, game: TypeGame, name: str) -> None:
        self.game: TypeGame = game
        self.name: str = name
        self.card_id = str(uuid.uuid4())
        self.is_resolved = False

    def __eq__(self, other: "AbstractCard" | str) -> bool:
        if isinstance(other, str):
            return self.card_id == other
        
        if not isinstance(other, AbstractCard):
            return False
        
        return all([
            self.name == other.name,
            self.game == other.game,
            self.card_id == other.card_id,
            self.color == other.color
        ])
    
    def __repr__(self) -> str:
        return f"<{self.__class__.__name__} color={self.color}>"

    @property
    def has_owner(self) -> bool:
        return self.owner is not None

    @abstractmethod
    async def resolve(self) -> bool:
        """Fully resolve the card's effect. This method should 
        be implemented by subclasses to define the specific behavior of 
        the card when it is resolved."""

    @abstractmethod
    def partial_resolve(self) -> MustResolve:
        """A partial resolution is a resolution that requires 
        further action from the player. It returns a MustResolve 
        object that describes the action required to continue."""

    def resolve_side_effects(self) -> None:
        """Resolve any side effects that occur when the card is resolved. 
        This method should be implemented by subclasses to define any 
        additional effects that occur as a result of resolving the card."""
        if self.tokens > 0:
            self.owner.total_points += self.tokens
            self.tokens = 0

        if self.game.must_resolve is not None and self.game.must_resolve.action == CardResolutionActions.REMOVE_CARD.value:
            self.owner.total_points += 1  # Award a point for eliminating a card
            
        # Intrigue cards are removed from the queue 
        # after resolution
        if self.card_type == Cardtypes.INTRIGUE:
            self.game.card_queue.remove_card(self)

    def pre_resolve(self) -> None:
        """A pre-resolution is a resolution that occurs before 
        the card is fully resolved. It can be used to set up 
        any necessary state or conditions before the card's 
        effect is applied.
        
        Raises:
            ValueError: If there is not exactly one card in the queue or if there is a pending must-resolve action.
        """
        if self.game.card_queue.number_of_cards < 1:
            raise ValueError("Resolution is only applicable when there is one card in the queue.")
        
        # if self.game.must_resolve is not None:
        #     raise ValueError("Cannot pre-resolve when there is a pending must-resolve action.")

    def add_token(self) -> None:
        """Add a token to the card."""
        self.tokens += 1

    def set_owner(self, owner: TypeWebsocketClient) -> None:
        if self.has_owner:
            raise ValueError("Card already has an owner.")
        self.owner = owner

    def add_to_stack(self, card: TypeAbstractCard) -> None:
        """Add a card to the stack for resolution."""
        self.stack.append(card)

    async def resolve_stack(self) -> None:
        """Resolve the stack of cards one by one."""
        for card in reversed(self.stack):
            await card.resolve()
            break
