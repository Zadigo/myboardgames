import unittest
from unittest.mock import MagicMock, PropertyMock
from its_a_wonderful_world.logic.base import Board, Game
from its_a_wonderful_world.logic.player import Player, PlayerHand


class TestPlayer(unittest.TestCase):
    def test_instance(self):
        items = {
            0: ['Card A', 'Card B'],
            1: ['Card C', 'Card D'],
            2: ['Card E', 'Card F'],
            3: ['Card G', 'Card H'],
        }

        player = Player('testuser')

        player.current_hand.set_hand('2')
        self.assertEqual(player.current_hand.hand_index, 2)

        mocked_board = MagicMock(spec=Board)
        type(mocked_board).number_of_players = PropertyMock(return_value=4)
        type(mocked_board).hands = PropertyMock(return_value=items)

        mocked_game = MagicMock(spec=Game)
        type(mocked_game).board = PropertyMock(return_value=mocked_board)

        player.game = mocked_game

        self.assertEqual(player.username, 'testuser')
        self.assertIsInstance(player.current_hand, PlayerHand)

        player.current_hand.pass_cards()

        self.assertListEqual(
            player.current_hand.current_hand_cards,
            items[player.current_hand.hand_index]
        )
