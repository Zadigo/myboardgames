import uuid
from abc import ABC, abstractmethod

from game.utils.base import MustResolve
from typings import TypeAbstractCard, TypeCardQueue, TypeGame, TypeWebsocketClient


class AbstractCardFactory(ABC):
    """An interface for creating card instances. Subclasses should 
    implement the create_cards method to generate specific card types."""
    
    @abstractmethod
    def create_cards(self, game: TypeGame) -> list[TypeAbstractCard]:
        """Create a list of cards for the given game instance."""


class AbstractCard(ABC):
    """Base class for all cards.

    Attributes:
        game (TypeGame): The game instance the card belongs to.
        name (str): The name of the card.
        partial_resolve (bool): Indicates if the card is partially resolved.
    """

    stack: list[TypeAbstractCard] = []
    owner: TypeWebsocketClient | None = None

    def __init__(self, game: TypeGame, name: str) -> None:
        self.game: TypeGame = game
        self.name: str = name
        self.card_id = str(uuid.uuid4())
        self._partial_resolve: bool = False

    def __eq__(self, other: "AbstractCard" | str) -> bool:
        if isinstance(other, str):
            return self.card_id == other
        
        if not isinstance(other, AbstractCard):
            return False
        
        return self.name == other.name and self.game == other.game and self.card_id == other.card_id

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

    def add_owner(self, owner: TypeWebsocketClient) -> None:
        if self.has_owner:
            raise ValueError("Card already has an owner.")
        self.owner = owner

    def resolve_stack(self) -> None:
        """Resolve the stack of cards one by one."""
        for card in reversed(self.stack):
            card.resolve()
            break
