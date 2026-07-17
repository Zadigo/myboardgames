from abc import ABC, abstractmethod
from collections import deque
from typing import Iterator, Optional

from typings import TypeAbstractCard, TypeGame


class CardQueue[T = TypeAbstractCard](ABC):
    """Abstract base class for a queue of cards. In oriflamme this represents
    the order in which the cards that were played are resolved.
    
    Attributes:
        current_index (int): The index of the current card being resolved.
        _queue (deque): A deque to hold the cards in the queue.
        game (TypeGame): The game instance associated with the queue.
    """

    def __init__(self, game: TypeGame) -> None:
        self.current_index : int = 0
        self._queue: deque[T] = deque()
        self.game: TypeGame = game

    def __iter__(self) -> Iterator[T]:
        return iter(self._queue)

    def __getitem__(self, index: int) -> T:
        return self._queue[index]
    
    def __len__(self) -> int:
        return len(self._queue)

    def __repr__(self) -> str:
        items: list[T] = list(self._queue)
        list_to_display: list[T] | list[str | T] = items if len(items) <= 10 else items[:10] + ['...']
        return f'<{self.__class__.__name__}({list_to_display})>'

    @property
    def number_of_cards(self) -> int:
        return len(self._queue)
    
    @property
    def current_card(self) -> Optional[T]:
        """Return the current card being resolved, or None 
        if there are no cards left."""
        if self.current_index < len(self._queue):
            return self._queue[self.current_index]
        return None

    @abstractmethod
    def prepend_card(self, card: T) -> None:
        """Adds a card to the front of the queue."""
        return NotImplemented

    @abstractmethod
    def append_card(self, card: T) -> None:
        """Adds a card to the end of the queue."""
        return NotImplemented

    @abstractmethod
    def stack_card(self, card_to_use: T, card_to_stack: T) -> None:
        """Stacks a card on top of another card in the queue."""
        return NotImplemented
    
    @abstractmethod
    def resolve(self) -> bool:
        """Resolves each card one by one in the queue."""
        return NotImplemented
    
    def increase_current_index(self) -> None:
        """Increase the current index to point to the next card in the queue."""
        self.current_index = (self.current_index + 1) % len(self._queue)

    def post_resolve(self) -> None:
        """Perform any post-resolution actions."""
        self.game.must_resolve = None

    def remove_card(self, card: Optional[T]) -> None:
        if card is None:
            return
        
        card_index: int = self._queue.index(card)
        self._queue.remove(card)
        self.post_resolve()

    def get_first_card(self) -> T:
        """Return the first card in the queue.
        
        Raises:
            IndexError: If the queue is empty.
        """
        if self.number_of_cards == 0:
            raise IndexError("The queue is empty.")

        return self._queue[0]
    
    def get_last_card(self) -> T:
        """Return the last card in the queue.

        Raises:
            IndexError: If the queue is empty.
        """
        if self.number_of_cards == 0:
            raise IndexError("The queue is empty.")
        
        return self._queue[-1]

    def get_previous_card(self, card: Optional[T]) -> Optional[T]:
        if self.number_of_cards < 2:
            return None
        
        card_index: int = self._queue.index(card)
        if card_index == 0:
            return None
        
        return self._queue[card_index - 1]
    
    def get_next_card(self, card: Optional[T]) -> Optional[T]:
        if self.number_of_cards < 2:
            return None
        
        card_index: int = self._queue.index(card)
        if card_index == self.number_of_cards - 1:
            return None
        
        return self._queue[card_index + 1]


class Queue(CardQueue[TypeAbstractCard]):
    def append_card(self, card: TypeAbstractCard) -> None:
        self._queue.append(card)

    def prepend_card(self, card: TypeAbstractCard) -> None:
        self._queue.appendleft(card)

    def stack_card(self, card_to_use: TypeAbstractCard, card_to_stack: TypeAbstractCard) -> None:
        card_index: int = self._queue.index(card_to_use)
        card = self._queue[card_index]

        # The user can stack only his cards. He cannot stack
        # his card upon the color of another player
        if card.color != card_to_stack.color:
            raise ValueError("Cannot stack cards of different colors.")
        
        card.stack.append(card_to_stack)

    def resolve(self) -> bool:
        state: bool = self.current_card.resolve()
        self.increase_current_index()
        return state
