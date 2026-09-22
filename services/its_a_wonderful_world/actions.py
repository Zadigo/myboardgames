from its_a_wonderful_world.typings import TypeGame


class Actions:
    """Defines all the actions that were made during the game."""

    actions: list[tuple[str, dict]] = []  # (action_name, action_data)

    def __get__(self, instance: TypeGame, cls: TypeGame):
        self.current_game = instance
        return self
