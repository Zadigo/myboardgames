import enum
import random
import uuid
from abc import ABC, abstractmethod
from typing import Optional
from game.card_queue import CardQueue, Queue
from game.cards.generics import AbstractCardFactory, CardColors
from game.cards.standard import StandardCardFactory
from game.utils.base import MustResolve
from typings import TypeAbstractCard, TypeWebsocketClient


class GameTypes(enum.Enum):
    STANDARD = "standard"


class BaseGame(ABC):
    """
    Attributes:
        card_queue (CardQueue[TypeAbstractCard]): The queue of cards in the game.
        game_id (str): Unique identifier for the game instance.
        must_resolve (Optional[MustResolve]): The current action that must be resolved.
        game_type (GameTypes): The type of the game.
        is_running (bool): Indicates if the game is currently running.
        available_cards (list[TypeAbstractCard]): The list of available cards in the game.
        current_player_index (int): Index of the current player in the players list.
        available_colors (list[CardColors]): The list of available colors in the game.
    """
    def __init__(self, game_type: GameTypes) -> None:
        self.available_cards: list[TypeAbstractCard] = []
        self.available_colors: list[CardColors] = [color for color in CardColors]
        self.card_queue: CardQueue[TypeAbstractCard] = Queue(self)

        self.game_id: str = str(uuid.uuid4())
        self.must_resolve: Optional[MustResolve] = None
        self.game_type: GameTypes = game_type
        self.is_running: bool = False
        self.current_player_index: int = 0

        self.factories: dict[GameTypes, AbstractCardFactory] = {
            GameTypes.STANDARD: StandardCardFactory
        }

    def __hash__(self) -> int:
        return hash(self.game_id)
    
    @property
    def current_player(self) -> Optional[TypeWebsocketClient]:
        if not self.players:
            return None
        return self.players[self.current_player_index]

    @abstractmethod
    async def start(self) -> None:
        pass

    @abstractmethod
    async def end(self) -> None:
        pass

    @abstractmethod
    def add_player(self, player: TypeWebsocketClient) -> None:
        pass

    @abstractmethod
    def remove_player(self, player: TypeWebsocketClient) -> None:
        pass

    @abstractmethod
    async def send_message(self, message: str) -> None:
        pass

    @abstractmethod
    def next_player(self) -> None:
        pass

    @abstractmethod
    async def build_cards(self) -> None:
        pass

    async def must_resolve_action(self, options: "MustResolve") -> None:
        self.must_resolve = options
        await self.send_message(
            f"Action required: {options.action} for card: {options.from_card.name}")




class Game(BaseGame):
    players: list[TypeWebsocketClient] = []

    @property
    def can_start(self) -> bool:
        return len(self.players) >= 2
    
    async def build_cards(self) -> None:
        factory_class: AbstractCardFactory | None = self.factories.get(self.game_type)
        if factory_class is None:
            raise ValueError(f"No factory found for game type: {self.game_type}")
        
        factory_instance = factory_class()
        cards = factory_instance.create_cards(self)
        for card in cards:
            self.available_cards.append(card)

        random.shuffle(self.available_colors)

        for color in self.available_colors:
            player_cards: filter[TypeAbstractCard] = filter(lambda card: card.color == color, self.available_cards)
            for player in self.players:
                player.attribute_cards(player_cards)

    
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
            await player.send(message)


async def create_game(game_type: GameTypes) -> BaseGame:
    instance: BaseGame = None
    if game_type == GameTypes.STANDARD:
        instance = Game(game_type)
    else:
        raise ValueError(f"Unsupported game type: {game_type}")
    
    await instance.build_cards()
    await instance.send_message(f"Game created with ID: {instance.game_id}")
    return instance
