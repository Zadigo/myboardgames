from typing import Iterable, Sequence
from game.cards.generics import CardColors
from game.cards.generics import AbstractCard, AbstractCardFactory
from game.utils.base import MustResolve
from typings import CardResolutionActions, TypeAbstractCard, TypeGame


class Soldier(AbstractCard):
    requires_partial_resolution: bool = False

    def __init__(self, game: TypeGame) -> None:
        super().__init__(game, 'Soldier')

    def partial_resolve(self) -> MustResolve:
        pass

    def resolve(self) -> None:
        if self.requires_partial_resolution:
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
    requires_partial_resolution: bool = True

    def __init__(self, game: TypeGame) -> None:
        super().__init__(game, 'Archer')

    def partial_resolve(self) -> MustResolve:
        instance = MustResolve(
            CardResolutionActions.REMOVE_CARD.value, 
            from_card=self
        )

        first_card: TypeAbstractCard | None = self.game.card_queue.get_first_card()
        last_card: TypeAbstractCard | None = self.game.card_queue.get_last_card()

        instance.on_first_card = first_card
        instance.on_last_card = last_card

        return instance

    async def resolve(self) -> bool:
        """
        Resolve the Archer card by eliminating the first or last card in the queue 
        if user selection is provided.
        
        Raises:
            ValueError: If user selection is required but not provided, or if there are no valid cards to eliminate.
        """
        super().pre_resolve()

        result = True

        if self.requires_partial_resolution and self.game.must_resolve is None:
            must_resolve: MustResolve = self.partial_resolve()
            await self.game.set_must_resolve(must_resolve)
        else:
            if self.game.must_resolve.user_selection is None:
                raise ValueError("User selection is required for resolution.")
            
            if self.game.must_resolve.on_first_card is None and self.game.must_resolve.on_last_card is None:
                raise ValueError("No valid cards to eliminate for resolution.")
            
            # Eliminate the previous card from the queue
            if self.game.must_resolve.user_selection == self.game.must_resolve.on_first_card:
                self.game.card_queue.remove_card(self.game.must_resolve.on_first_card)

            # Eliminate the next card from the queue
            if self.game.must_resolve.user_selection == self.game.must_resolve.on_last_card:
                self.game.card_queue.remove_card(self.game.must_resolve.on_last_card)

        self.resolve_side_effects()
        self.game.card_queue.post_resolve()
        return result
    

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
