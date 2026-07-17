from unittest.mock import MagicMock, PropertyMock

import pytest

from game.base import Game
from game.cards.standard import Archer


@pytest.fixture
def game_fixture() -> MagicMock:
    mgame = MagicMock().return_value

    card_queue = MagicMock().return_value
    type(card_queue).number_of_cards = PropertyMock(return_value=2)
    
    mgame.card_queue = card_queue
    mgame.players = []

    return mgame

def test_instance(game_fixture: MagicMock) -> None:
    instance = Archer(game=game_fixture)
    assert instance.owner is None
    
    assert instance == instance
    assert instance != "some_string"


class TestArcher:
    @pytest.fixture(autouse=True)
    def setup(self) -> None:
        game = Game()
        self.instance = Archer(game)

    def test_partial_resolution(self) -> None:
        self.instance._partial_resolve = True
        result: bool = self.instance.resolve()

        assert result is True
        assert self.instance.game.must_resolve is not None
