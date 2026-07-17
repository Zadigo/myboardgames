import uuid
from abc import ABC, abstractmethod
from collections import deque
from typing import Iterator, Optional

from game.utils.base import MustResolve
from typings import TypeCard, TypeWebsocketClient


class CardQueue[T = TypeCard](ABC):
    def __init__(self, game: "BaseGame") -> None:
        self._queue: deque[T] = deque()
        self.game: "BaseGame" = game

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

    @abstractmethod
    def build_cards(self) -> list[T]:
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


class Queue(CardQueue[TypeCard]):
    def append_card(self, card: TypeCard) -> None:
        self._queue.append(card)

    def prepend_card(self, card: TypeCard) -> None:
        self._queue.appendleft(card)

    def stack_card(self) -> list[TypeCard]:
        return list(self._queue)


class BaseGame(ABC):
    def __init__(self) -> None:
        self.card_queue: CardQueue[TypeCard] = Queue(self)
        self.game_id: str = str(uuid.uuid4())
        self.must_resolve: Optional[MustResolve] = None
        self.is_running: bool = False
        self.current_player_index: int = 0

    def __hash__(self) -> int:
        return hash(self.game_id)
    
    @property
    def current_player(self) -> Optional[TypeWebsocketClient]:
        if not self.players:
            return None
        return self.players[self.current_player_index]

    @abstractmethod
    def start(self) -> None:
        pass

    @abstractmethod
    def end(self) -> None:
        pass

    @abstractmethod
    def add_player(self, player: TypeWebsocketClient) -> None:
        pass

    @abstractmethod
    def remove_player(self, player: TypeWebsocketClient) -> None:
        pass

    @abstractmethod
    def send_message(self, message: str) -> None:
        pass

    @abstractmethod
    def next_player(self) -> None:
        pass

    def must_resolve_action(self, options: "MustResolve") -> None:
        self.must_resolve = options
        self.send_message(
            f"Action required: {options.action} for card: {options.from_card.name}")


class Game(BaseGame):
    players: list[TypeWebsocketClient] = []

    @property
    def can_start(self) -> bool:
        return len(self.players) >= 2
    
    async def end(self) -> None:
        self.is_running = False
        await self.send_message("Game has ended.")

    async def start(self) -> None:
        self.is_running = True

        while self.is_running:
            if not self.can_start:
                continue

            # Game logic here

    def add_player(self, player: TypeWebsocketClient) -> None:
        pass

    def remove_player(self, player: TypeWebsocketClient) -> None:
        pass

    def next_player(self) -> None:
        self.current_player_index = (
            self.current_player_index + 1
        ) % len(self.players)

        self.current_player.send(f"It's your turn, {self.current_player}!")

    async def send_message(self, message: str) -> None:
        for player in self.players:
            player.send(message)
