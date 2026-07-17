from abc import ABC, abstractmethod
from collections import deque
from typing import Iterator, Optional

from typings import TypeAbstractCard, TypeGame


class CardQueue[T = TypeAbstractCard](ABC):
    def __init__(self, game: TypeGame) -> None:
        self._queue: deque[T] = deque()
        self.game: TypeGame = game

    def __iter__(self) -> Iterator[T]:
        return iter(self._queue)

    def __getitem__(self, index: int) -> T:
        return self._queue[index]

    @property
    def number_of_cards(self) -> int:
        return len(self._queue)

    @abstractmethod
    def prepend_card(self, card: T) -> None:
        pass

    @abstractmethod
    def append_card(self, card: T) -> None:
        pass

    @abstractmethod
    def stack_card(self) -> list[T]:
        pass

    def post_resolve(self) -> None:
        """Perform any post-resolution actions."""
        self.game.must_resolve = None

    def remove_card(self, card: Optional[T]) -> None:
        if card is None:
            return
        
        card_index: int = self._queue.index(card)
        self._queue.remove(card)
        self.post_resolve()

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

    def stack_card(self) -> list[TypeAbstractCard]:
        return list(self._queue)
