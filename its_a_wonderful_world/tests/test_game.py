from unittest import TestCase

from its_a_wonderful_world.logic.base import Game
from its_a_wonderful_world.logic.player import Player
from faker import Faker


fake_name = Faker()


class TestGame(TestCase):
    def _generate_players(self, game_instance: Game, number_of_players: int) -> list[Player]:
        players = []
        for _ in range(number_of_players):
            player = Player(game_instance, fake_name.name())
            game_instance.board.add_player(player)
            players.append(player)
        return players

    def test_instance(self):
        instance = Game()
        print(instance.board.all_default_cards)
        self.assertEqual(len(instance.board.all_default_cards), 150)
        self.assertTrue(instance.board.game_started)

    def test_prepare_board(self):
        instance = Game()
        instance.board.prepare()
        self.assertTrue(instance.board.game_started)
        self.assertGreater(len(instance.board.all_default_cards), 0)

    def test_add_and_remove_player(self):
        instance = Game()
        player = Player(instance, fake_name.name())
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
            total_cards == (7 * len(players)),
            "Excpected 35 total cards for 5 players (7 * 5)."
        )
