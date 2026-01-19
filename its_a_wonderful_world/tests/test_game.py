from unittest import TestCase

from its_a_wonderful_world.logic.base import Game
from its_a_wonderful_world.logic.player import Player
from faker import Faker


fake_name = Faker()


class TestGame(TestCase):
    def _generate_players(self, game_instance: Game, number_of_players: int) -> list[Player]:
        players = []
        for _ in range(number_of_players):
            player = Player(game=game_instance, username=fake_name.name())
            game_instance.board.add_player(player)
            players.append(player)
        return players

    def test_instance(self):
        instance = Game()

        for item in instance.board.all_default_cards:
            print(item.name, item.number_of_copies())

        self.assertEqual(len(instance.board.all_default_cards), 150)
        self.assertTrue(instance.board.game_started)

    def test_prepare_board(self):
        instance = Game()
        instance.board.prepare()
        self.assertTrue(instance.board.game_started)
        self.assertGreater(len(instance.board.all_default_cards), 0)

    def test_add_and_remove_player(self):
        instance = Game()
        player = Player(game=instance, username=fake_name.name())
        instance.board.add_player(player)
        self.assertIn(player.player_id, instance.board.players)

        instance.board.remove_player(player)
        self.assertNotIn(player.player_id, instance.board.players)

    def test_total_cards_to_select(self):
        number_of_players = 5

        instance = Game()
        players = self._generate_players(instance, number_of_players)

        for player in players:
            instance.board.add_player(player)

        self.assertEqual(instance.board.number_of_players, number_of_players)

        total_cards = instance.board.total_cards_to_select
        self.assertTrue(
            total_cards == (6 * len(players)),
            "Excpected 30 total cards for 5 players (6 * 5)."
        )

    def test_pick_cards_with_pop(self):
        number_of_players = 4

        instance = Game()
        self._generate_players(instance, number_of_players)

        total_cards_before = len(instance.board.all_default_cards)
        instance.board.pick_cards_with_pop()
        total_cards_after = len(instance.board.all_default_cards)

        expected_cards_picked = instance.board.total_cards_to_select
        actual_cards_picked = total_cards_before - total_cards_after

        self.assertEqual(
            expected_cards_picked,
            actual_cards_picked,
            f"Expected {expected_cards_picked} cards to be picked, "
            f"but got {actual_cards_picked}."
        )

        self.assertTrue(len(instance.board.hands.keys()) > 0)

        for key, value in instance.board.hands.items():
            with self.subTest(player_id=key):
                self.assertEqual(
                    len(value),
                    sum(instance.board.number_of_cards_to_select),
                    f"Player {key} should have "
                    f"{sum(instance.board.number_of_cards_to_select)} "
                    f"cards in hand."
                )

        player = instance.board.players[list(instance.board.players.keys())[0]]
        self.assertTrue(player.current_hand is not None)
        self.assertIsInstance(player.current_hand.hand_index, int)  # noqa: E501

    def test_number_of_cards_to_select(self):
        items = [
            (2, 8),  # Number of players, expected cards
            (5, 6)
        ]

        for number_of_players, expected in items:
            instance = Game()
            self._generate_players(instance, number_of_players)

            result, ac_cards = instance.board.number_of_cards_to_select
            self.assertEqual(
                result,
                expected,
                f"For {number_of_players} players, expected {expected} cards to select, but got {result}."
            )
