from collections import deque
from dataclasses import dataclass
from typing import Optional
from uuid import uuid4

from its_a_wonderful_world.typings import TypeGame, TypePlayer, TypePlayerCard


class PlayerMixin:
    def __init__(self, player: TypePlayer):
        self.player: TypePlayer = player

    def __hash__(self):
        return hash(self.player.player_id)


class PlayerHand(PlayerMixin):
    """The cards currently in a player's hand on which 
    they can do different actions during their turn. This includes:

    * Dealing development cards
    * Choosing development cards
    * Revealing chosen cards
    * Passing cards
    * Discarding leftover cards
    """

    def __init__(self, player: TypePlayer):
        super().__init__(player=player)

        self.hand_index: int = 0

    def __repr__(self):
        return f"<PlayerHand player={self.player.username} hand_index={self.hand_index}>"

    @property
    def current_hand_cards(self) -> list[TypePlayerCard]:
        if self.player.game is None:
            return []
        return self.player.game.board.hands[self.hand_index]

    def set_hand(self, hand_index: str | int):
        """Set the current hand index for the player."""
        self.hand_index = int(hand_index)

    def pass_cards(self):
        """Pass the current hand to the next player based on 
        the rotation type."""
        self.hand_index = (self.hand_index + 1) % self.player.game.board.number_of_players  # noqa: E501

    def discard_leftover_cards(self):
        """Discard all leftover cards in the player's hand."""
        pass


class PlayerSelectedCards(PlayerMixin):
    """The cards a player has selected to either build or discard on
    which they can take actions during their turn."""

    cards: deque[TypePlayerCard] = deque()

    def __repr__(self):
        return f"<PlayerSelectedCards player={self.player.username} cards={len(self.cards)}>"

    def reveal_chosen_cards(self):
        """Every player puts there selected card face down. Just before
        passing their hand to the next player, all players reveal their
        chosen cards."""
        pass

    def deal_development_cards(self, number_of_cards: int) -> list:
        """Once the first phase is completed, each player can choose to deal
        a certain number of development cards for a resource."""
        pass

    def choose_development_cards(self):
        """A player chooses development cards from their hand to be placed
        face down in front of them."""
        pass


class PlayerBuiltCards(PlayerMixin):
    """Global information about the cards a player has 
    built during the game."""

    cards: deque[TypePlayerCard] = deque()

    @property
    def total_production(self) -> int:
        return 0

    @property
    def energy_production(self) -> int:
        return 0

    @property
    def victory_points(self) -> int:
        return 0

    @property
    def character_points(self) -> int:
        return 0

    @property
    def character_general_points(self) -> int:
        return 0

    @property
    def character_financier_points(self) -> int:
        return 0


@dataclass
class Player:
    """A player in the game."""
    username: str

    firstname: str = ''
    lastname: str = ''
    player_id: str = ''

    game: Optional[TypeGame] = None
    current_hand: Optional[PlayerHand] = None
    current_card_selection: Optional[PlayerSelectedCards] = None
    current_built_cards: Optional[PlayerBuiltCards] = None

    # selected_empire: str = ''

    def __post_init__(self):
        self.player_id: str = str(uuid4())
        self.current_hand = PlayerHand(player=self)
        self.current_card_selection = PlayerSelectedCards(self)
        self.current_built_cards = PlayerBuiltCards(self)

    def __hash__(self):
        return hash(self.player_id)
