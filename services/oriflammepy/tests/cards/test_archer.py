import pytest

from game.base import Game, GameTypes
from game.cards.generics import CardColors
from game.cards.standard import Archer
from game.utils.base import MustResolve


@pytest.fixture
def game_fixture() -> Game:
    instance = Game(GameTypes.STANDARD)

    # Cards for testing
    card1 = Archer(instance)
    card1.color = CardColors.RED

    card2 = Archer(instance)
    card2.color = CardColors.BLUE

    card3 = Archer(instance)
    card3.color = CardColors.GREEN

    instance.card_queue.append_card(card1)
    instance.card_queue.append_card(card2)
    instance.card_queue.append_card(card3)

    return instance


def test_partial_resolution_in_first_position(game_fixture: Game) -> None:
    card = game_fixture.card_queue._queue[0]

    assert card.requires_partial_resolution is True
    assert card is not None

    must_resolve: MustResolve = card.partial_resolve()
    assert must_resolve is not None
    assert must_resolve.action == "remove_card"
    assert must_resolve.from_card == card
    # Can eliminate self and last card
    assert must_resolve.on_first_card is not None
    assert must_resolve.on_last_card is not None


def test_partial_resolution_in_middle_position(game_fixture: Game) -> None:
    card = game_fixture.card_queue._queue[1]

    assert card.requires_partial_resolution is True
    assert card is not None

    must_resolve: MustResolve = card.partial_resolve()
    assert must_resolve is not None
    assert must_resolve.action == "remove_card"
    assert must_resolve.from_card == card
    # Can eliminate the first and last card
    assert must_resolve.on_first_card is not None
    assert must_resolve.on_last_card is not None


def test_partial_resolution_in_last_position(game_fixture: Game) -> None:
    card = game_fixture.card_queue._queue[-1]

    assert card.requires_partial_resolution is True
    assert card is not None

    must_resolve: MustResolve = card.partial_resolve()
    assert must_resolve is not None
    assert must_resolve.action == "remove_card"
    assert must_resolve.from_card == card
    # Can eliminate the first and last card
    assert must_resolve.on_first_card is not None
    assert must_resolve.on_last_card is not None


async def test_resolve(game_fixture: Game) -> None:
    card = game_fixture.card_queue._queue[1]

    assert card.requires_partial_resolution is True
    assert card is not None

    # Simulate the resolution process
    must_resolve: MustResolve = card.partial_resolve()
    assert must_resolve is not None

    # Simulate user selection of the first card to eliminate
    must_resolve.user_selection = must_resolve.on_first_card

    # Call resolve and check if it returns True (indicating further action is required)
    result = await card.resolve()
    assert result is True
