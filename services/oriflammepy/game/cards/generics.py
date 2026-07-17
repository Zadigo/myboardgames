import enum
import uuid
from abc import ABC, abstractmethod
from typing import Iterable, Optional, Sequence

from game.utils.base import MustResolve
from typings import TypeAbstractCard, TypeCardQueue, TypeGame, TypeWebsocketClient


class CardColors(enum.Enum):
    BLUE = "blue"
    RED = "red"
    YELLOW = "yellow"
    BLACK = "black"
    


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

    Attributes:
        game (TypeGame): The game instance the card belongs to.
        name (str): The name of the card.
        partial_resolve (bool): Indicates if the card is partially resolved.
        color (Optional[CardColors]): The color of the card.
        owner (Optional[TypeWebsocketClient]): The owner of the card, if any.
    """

    color: Optional[CardColors] = None
    stack: list[TypeAbstractCard] = []
    owner: TypeWebsocketClient | None = None

    def __init__(self, game: TypeGame, name: str) -> None:
        self.game: TypeGame = game
        self.name: str = name
        self.card_id = str(uuid.uuid4())
        self.is_resolved = False
        self._partial_resolve: bool = False

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
    def resolve(self) -> bool:
        """Fully resolve the card's effect. This method should 
        be implemented by subclasses to define the specific behavior of 
        the card when it is resolved."""

    @abstractmethod
    def partial_resolve(self) -> MustResolve:
        """A partial resolution is a resolution that requires 
        further action from the player. It return a MustResolve 
        object that describes the action required to continue."""

    def can_resolve(self, card_queue: TypeCardQueue) -> bool:
        pass

    def set_owner(self, owner: TypeWebsocketClient) -> None:
        if self.has_owner:
            raise ValueError("Card already has an owner.")
        self.owner = owner

    def add_to_stack(self, card: TypeAbstractCard) -> None:
        """Add a card to the stack for resolution."""
        self.stack.append(card)

    def resolve_stack(self) -> None:
        """Resolve the stack of cards one by one."""
        for card in reversed(self.stack):
            card.resolve()
            break
