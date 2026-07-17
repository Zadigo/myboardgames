from unittest.mock import MagicMock, PropertyMock

import pytest


@pytest.fixture
def game_fixture() -> MagicMock:
    mgame = MagicMock().return_value

    card_queue = MagicMock().return_value
    type(card_queue).number_of_cards = PropertyMock(return_value=2)
    
    mgame.card_queue = card_queue
    mgame.players = []

    return mgame

