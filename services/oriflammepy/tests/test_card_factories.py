from typing import Iterable

import pytest

from game.base import Game, GameTypes
from game.cards.generics import AbstractCard, CardColors
from game.cards.standard import StandardCardFactory


@pytest.fixture
def game_fixture() -> Game:
    return Game(GameTypes.STANDARD)


def test_standard_card_factory(game_fixture: Game) -> None:
    factory = StandardCardFactory(game_fixture)
    cards: Iterable[AbstractCard] = iter(factory)

    cards_list: list[AbstractCard] = list(cards)
    # assert len(list(cards)) == 52
    assert len(cards_list) == 8

    result: Iterable[AbstractCard] = factory.cards_from_color(CardColors.RED)
    assert len(list(result)) == 2
