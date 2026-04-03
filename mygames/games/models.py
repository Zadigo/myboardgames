from django.db import models


class BoardGame(models.Model):
    name = models.CharField(
        max_length=100
    )
    description = models.TextField(
        max_length=1000
    )
    min_players = models.PositiveIntegerField(
        default=1
    )
    max_players = models.PositiveIntegerField(
        default=1
    )
    modified_on = models.DateTimeField(
        auto_now=True
    )
    create_on = models.DateTimeField(
        auto_now_add=True
    )

    def __str__(self):
        return self.name
