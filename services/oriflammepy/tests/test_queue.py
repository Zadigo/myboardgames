import random
from collections import deque
from typing import Iterable

import pytest

from game.base import Game, GameTypes
from game.card_queue import Queue
from game.cards.generics import AbstractCard
from game.cards.standard import StandardCardFactory


@pytest.fixture
def game_fixture() -> Game:
    return Game(GameTypes.STANDARD)


@pytest.fixture
def factory_fixture(game_fixture: Game) -> StandardCardFactory:
    return StandardCardFactory(game_fixture)


def test_instance(game_fixture: Game) -> None:
    queue = Queue(game_fixture)

    assert isinstance(queue._queue, deque)

    # Test __iter__
    assert len(list(queue)) == 0

    # Test __getitem__
    with pytest.raises(IndexError):
        assert queue[0]

    assert queue.number_of_cards == 0

    value: str = repr(queue)
    assert isinstance(value, str)
    assert value.startswith('<Queue(') and value.endswith(')>')


def test_adding_actions(game_fixture: Game, factory_fixture: StandardCardFactory) -> None:
    queue = Queue(game_fixture)

    # factory = StandardCardFactory(game_fixture)
    # cards: Iterable[AbstractCard] = list(iter(factory))

    cards: Iterable[AbstractCard] = list(iter(factory_fixture))
    selected_cards: list[AbstractCard] = cards[:3]

    queue.append_card(selected_cards[0])
    queue.prepend_card(selected_cards[1])
    queue.prepend_card(selected_cards[2])

    assert len(list(queue)) == 3

    previous_card: AbstractCard | None = queue.get_previous_card(selected_cards[1])
    assert previous_card == queue._queue[0]

    next_card: AbstractCard | None = queue.get_next_card(selected_cards[1])
    assert next_card == queue._queue[2]


def test_stack_actions(game_fixture: Game, factory_fixture: StandardCardFactory) -> None:
    queue = Queue(game_fixture)

    cards: Iterable[AbstractCard] = list(iter(factory_fixture))

    queue.append_card(cards[0])
    queue.stack_card(cards[0], cards[1])

    assert cards[0].stack[-1] == cards[1]


def test_remove_card(game_fixture: Game, factory_fixture: StandardCardFactory) -> None:
    queue = Queue(game_fixture)

    cards: Iterable[AbstractCard] = list(iter(factory_fixture))
    selected_cards: list[AbstractCard] = random.sample(cards, 3)

    queue.append_card(selected_cards[0])
    queue.prepend_card(selected_cards[1])
    queue.prepend_card(selected_cards[2])

    queue.remove_card(selected_cards[0])

    assert len(list(queue)) == 2


def test_increase_current_index(game_fixture: Game, factory_fixture: StandardCardFactory) -> None:
    queue = Queue(game_fixture)

    cards: Iterable[AbstractCard] = list(iter(factory_fixture))
    selected_cards: list[AbstractCard] = random.sample(cards, 3)

    queue.append_card(selected_cards[0])
    queue.append_card(selected_cards[1])
    queue.append_card(selected_cards[2])

    assert queue.current_index == 0

    queue.increase_current_index()
    assert queue.current_index == 1

    queue.increase_current_index()
    assert queue.current_index == 2

    queue.increase_current_index()
    assert queue.current_index == 0

    current_card: AbstractCard | None = queue.current_card
    assert current_card is not None
