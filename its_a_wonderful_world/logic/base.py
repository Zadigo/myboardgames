import random
from collections import defaultdict, deque
from typing import Optional, Type
from uuid import uuid4
from itertools import chain
from its_a_wonderful_world.actions import Actions
from its_a_wonderful_world.material.card.research import RESEARCH_CARDS
from its_a_wonderful_world.material.card.vehicle import VEHICLE_CARDS
from its_a_wonderful_world.typings import (PhaseTypes, RotationTypes,
                                           TypePlayer, TypePlayerCard)


class DiscardedCards:
    """The cards that have been discardeded throughout the game."""


class Board:
    """A game board represents the current state of the game for all players
    on that board and the different overall actions that can be taken. A game
    board is associated with a single game instance.

    The board can do the following things:

    * Keep track of the players on the board
    * Keep track of the cards on the board
    * Prepare the board for a new game
    * Add and remove players from the board
    * Keep track of the current game state (e.g. current round, current phase)
    """

    all_default_cards: deque[TypePlayerCard] = deque()
    players: dict[str, TypePlayer] = {}
    # Dictionnary of the index to the the hands that
    # were created for each player.
    hands: defaultdict[int, list[TypePlayerCard]] = defaultdict(list)

    # current_hands: PlayerHand = PlayerHand()

    def __init__(self):
        self.board_id = str(uuid4())
        self.game_started: bool = False
        self.current_game: Optional['Game'] = None

        self.current_round: int = 1
        self.current_phase: PhaseTypes = PhaseTypes.PLANNING

        self.first_player: Optional[TypePlayer] = None
        self.current_first_player: Optional[TypePlayer] = None

    def __get__(self, instance: 'Game', cls: Type['Game']):
        self.current_game = instance
        return self

    def __hash__(self):
        return hash(self.board_id)

    @property
    def number_of_players(self) -> int:
        """Number of players currently on the board."""
        return len(self.players.keys())

    @property
    def number_of_cards_to_select(self) -> tuple[int, int]:
        """Number of cards that a player can select from based on 
        player count and expansions. In the default game, players
        select 7 cards with no cards to discard. In the Ascension
        and Corruption expansion, players select 6 cards and discard
        3 cards with more than 2 players, or select 10 cards and
        discard 2 cards with 2 or fewer players."""
        if self.current_game is None:
            raise ValueError(
                "Cannot determine number of cards to select without a game instance.")

        if self.current_game.is_ascension_and_corruption_expansion:
            if self.number_of_players <= 2:
                return (10, 2)
            else:
                return (6, 3)
        else:
            if self.number_of_players <= 2:
                return (10, 0)
            else:
                return (7, 0)

    @property
    def total_cards_to_select(self) -> int:
        """Total number of cards a player selects."""
        default_cards, extension_cards = self.number_of_cards_to_select
        return sum((default_cards, extension_cards)) * self.number_of_players

    def prepare(self):
        """Prepare the board for a new game:

        * Intiialize the cards that will be used in the game (e.g. shuffle the deck)
        * Set `game_started` to `True`
        """
        if self.current_game is None:
            raise ValueError("Cannot prepare board without a game instance.")

        available_cards = chain(*self.current_game.available_cards)
        
        for item in available_cards:
            card_instance = item()

            for _ in range(card_instance.number_of_copies()):
                self.all_default_cards.append(card_instance)

        if len(self.all_default_cards) != 150:
            # raise ValueError("The total number of cards in the deck is incorrect.")
            print("The total number of cards in the deck is incorrect")

        self.game_started = True

    def add_player(self, player: TypePlayer):
        self.players[player.player_id] = player

    def remove_player(self, player: TypePlayer):
        if player.player_id in self.players:
            del self.players[player.player_id]

    def pick_cards_with_pop(self):
        """Pick a set of cards for the players to choose from by selecting
        from the top of the deck."""
        selected_cards: deque[TypePlayerCard] = deque()

        if self.total_cards_to_select > len(self.all_default_cards):
            raise ValueError(
                "Not enough cards in the deck to "
                "select the required number of cards."
            )

        for _ in range(self.total_cards_to_select):
            selected_cards.append(self.all_default_cards.popleft())

        cards_per_player = len(selected_cards) // self.number_of_players

        for index in range(self.number_of_players):
            for _ in range(cards_per_player):
                self.hands[index].append(selected_cards.popleft())

    def pick_cards_with_random(self):
        """Pick a set of cards for the players to choose from by randomly
        selecting from the deck."""
        selected_cards = random.sample(
            self.all_default_cards, self.total_cards_to_select)

        cards_per_player = len(selected_cards) // self.number_of_players

        for index in range(self.number_of_players):
            start = index * cards_per_player
            end = start + cards_per_player
            self.hands[index] = selected_cards[start:end]


class Game:
    """The overall game state."""

    board = Board()
    actions = Actions()

    # boards: dict[str, Board] = {}

    number_of_rounds: int = 4
    available_cards: list[tuple[TypePlayerCard, ...]] = [
        VEHICLE_CARDS,
        RESEARCH_CARDS
    ]

    def __init__(self):
        self.discarded_cards: DiscardedCards = DiscardedCards()
        self.is_ascension_and_corruption_expansion: bool = False
        self.current_rotation: RotationTypes = RotationTypes.CLOCKWISE

        self.board.prepare()

    def __repr__(self):
        return f"<Game id={id(self)} players={self.board.number_of_players} rounds={self.number_of_rounds}>"
