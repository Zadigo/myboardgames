from typing import Iterable, Sequence
from game.cards.generics import CardColors
from game.cards.generics import AbstractCard, AbstractCardFactory
from game.utils.base import MustResolve
from typings import CardResolutionActions, TypeAbstractCard, TypeGame


class Soldier(AbstractCard):
    def __init__(self, game: TypeGame) -> None:
        super().__init__(game=game, name="Soldier")

    def partial_resolve(self) -> MustResolve:
        pass

    def resolve(self) -> None:
        if self.partial_resolve:
            previous_card: TypeAbstractCard | None = self.game.card_queue.get_previous_card(self)
            next_card: TypeAbstractCard | None = self.game.card_queue.get_next_card(self)
            
            options = MustResolve(
                action=CardResolutionActions.REMOVE_CARD.value,
                from_card=self
            )
            
            if previous_card is not None:
                options.on_previous_card = previous_card

            if next_card is not None:
                options.on_next_card = next_card

            self.game.must_resolve_action(options)
            return


class Archer(AbstractCard):
    def __init__(self, game: TypeGame) -> None:
        super().__init__(game=game, name="Archer")

    def partial_resolve(self) -> MustResolve:
        instance = MustResolve(
            CardResolutionActions.REMOVE_CARD.value, 
            from_card=self
        )

        first_card: TypeAbstractCard | None = self.game.card_queue.get_previous_card(self)
        last_card: TypeAbstractCard | None = self.game.card_queue.get_next_card(self)

        instance.on_previous_card = first_card
        instance.on_next_card = last_card

        return instance

    def resolve(self) -> bool:
        if self.game.card_queue.number_of_cards > 1:
            if self._partial_resolve:
                must_resolve: MustResolve = self.partial_resolve()
                self.game.must_resolve_action(must_resolve)
                return True
            else:
                if self.game.must_resolve is not None:
                    # Eliminate the previous card from the queue
                    if self.game.must_resolve.on_card == self.game.must_resolve.on_first_card:
                        self.game.must_resolve_action(self.game.must_resolve.on_first_card)

                    # Eliminate the next card from the queue
                    if self.game.must_resolve.on_card == self.game.must_resolve.on_last_card:
                        self.game.must_resolve_action(self.game.must_resolve.on_last_card)

                return True
        return False
    

class StandardCardFactory(AbstractCardFactory):
    """A factory class to create standard cards for the game."""

    cards_classes: Sequence[TypeAbstractCard] = [Soldier, Archer]

    def __iter__(self) -> Iterable[TypeAbstractCard]:
        """Create a list of standard cards for the given game instance."""
        colors: list[CardColors] = [color for color in CardColors.__members__.values()]
        for color in colors:
            for item in self.cards_classes:
                card_instance = item(self.game)
                card_instance.color = color
                yield card_instance
