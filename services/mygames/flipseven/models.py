from django.db import models


class GamePlay(models.Model):
    board_game = models.ForeignKey(
        'games.BoardGame',
        on_delete=models.CASCADE
    )
    card_moves = models.JSONField(
        default=list
    )
    created_on = models.DateTimeField(
        auto_now_add=True
    )

    def __str__(self):
        return str(self.board_game)
