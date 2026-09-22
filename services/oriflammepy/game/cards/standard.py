from typing import Iterable, Sequence
from game.cards.generics import CardColors
from game.cards.generics import AbstractCard, AbstractCardFactory
from game.utils.base import MustResolve
from typings import TypeAbstractCard


class Soldier(AbstractCard):
    """The Soldier card is a character card that allows the player to eliminate
    the previous or next card in the queue, based on user selection."""

    requires_partial_resolution: bool = True

    def partial_resolve(self) -> MustResolve:
        instance: MustResolve = super().partial_resolve()

        instance.on_previous_card = self.game.card_queue.get_previous_card(self)
        instance.on_next_card = self.game.card_queue.get_next_card(self)   

        return instance

    async def resolve(self) -> bool:
        result: bool = await super().resolve()

        if self.requires_partial_resolution and self.game.must_resolve is None:
            must_resolve: MustResolve = self.partial_resolve()
            await self.game.set_must_resolve(must_resolve)
        else:
            if self.game.must_resolve.user_selection is None:
                raise ValueError("User selection is required for resolution.")
            
            if self.game.must_resolve.on_previous_card is None and self.game.must_resolve.on_next_card is None:
                raise ValueError("No valid cards to eliminate for resolution.")

            if self.game.must_resolve.user_selection == self.game.must_resolve.on_previous_card:
                # Eliminate the previous card from the queue
                self.game.card_queue.remove_card(self.game.must_resolve.on_previous_card)
            elif self.game.must_resolve.user_selection == self.game.must_resolve.on_next_card:
                # Eliminate the next card from the queue
                self.game.card_queue.remove_card(self.game.must_resolve.on_next_card)
            
        return result


class Archer(AbstractCard):
    """The Archer card is a character card that allows the player to eliminate
    either the first or last card in the queue, based on user selection."""

    requires_partial_resolution: bool = True

    def partial_resolve(self) -> MustResolve:
        instance = super().partial_resolve()

        instance.on_first_card = self.game.card_queue.get_first_card()
        instance.on_last_card = self.game.card_queue.get_last_card()

        return instance

    async def resolve(self) -> bool:
        """
        Raises:
            ValueError: If user selection is required but not provided, or if there are no valid cards to eliminate.
        """
        result: bool = await super().resolve()

        if self.requires_partial_resolution and self.game.must_resolve is None:
            must_resolve: MustResolve = self.partial_resolve()
            await self.game.set_must_resolve(must_resolve)
        else:
            if self.game.must_resolve.user_selection is None:
                raise ValueError("User selection is required for resolution.")
            
            if self.game.must_resolve.on_first_card is None and self.game.must_resolve.on_last_card is None:
                raise ValueError("No valid cards to eliminate for resolution.")
            
            if self.game.must_resolve.user_selection == self.game.must_resolve.on_first_card:
                # Eliminate the first card from the queue
                self.game.card_queue.remove_card(self.game.must_resolve.on_first_card)
            elif self.game.must_resolve.user_selection == self.game.must_resolve.on_last_card:
                # Eliminate the last card from the queue
                self.game.card_queue.remove_card(self.game.must_resolve.on_last_card)

        self.resolve_side_effects()
        self.game.card_queue.post_resolve()
        return result


class ChangeForm(AbstractCard):
    """The ChangeForm card is a character card that copies the effects of 
    the next or previous card in the queue based on user selection."""

    requires_partial_resolution: bool = True

    def partial_resolve(self) -> MustResolve:
        instance = super().partial_resolve()

        instance.on_first_card = self.game.card_queue.get_first_card()
        instance.on_last_card = self.game.card_queue.get_last_card()

        return instance
    
    async def resolve(self) -> bool:
        """        
        Raises:
            ValueError: If user selection is required but not provided, or if there are no valid cards to eliminate.
        """
        result: bool = await super().resolve()

        if self.requires_partial_resolution and self.game.must_resolve is None:
            must_resolve: MustResolve = self.partial_resolve()
            await self.game.set_must_resolve(must_resolve)
        else:
            if self.game.must_resolve.user_selection is None:
                raise ValueError("User selection is required for resolution.")
            
            if self.game.must_resolve.on_first_card is None and self.game.must_resolve.on_last_card is None:
                raise ValueError("No valid cards to eliminate for resolution.")
            
            card_to_copy: TypeAbstractCard = None
            if self.game.must_resolve.user_selection == self.game.must_resolve.on_first_card:
                card_to_copy = self.game.card_queue.get_next_card(self.game.must_resolve.on_first_card)
            elif self.game.must_resolve.user_selection == self.game.must_resolve.on_last_card:
                card_to_copy = self.game.card_queue.get_previous_card(self.game.must_resolve.on_last_card)

            if card_to_copy is None:
                raise ValueError("No valid card to copy from the selected card.")
        
        await card_to_copy.resolve()
        # Resolve the side effects of the copied card
        card_to_copy.resolve_side_effects()            
        # Resolve the side effects of the ChangeForm card itself
        self.resolve_side_effects()

        self.game.card_queue.post_resolve()
        return result


class Heir(AbstractCard):
    """The Heir card is a character card that attributes 2 tokens if there are no other
    cards of type Heir in the qeue."""

    requires_partial_resolution: bool = False

    def partial_resolve(self) -> MustResolve:
        pass

    async def resolve(self) -> bool:
        result: bool = await super().resolve()

        heirs: int = 0
        for card in self.game.card_queue._queue:
            if card.name == "Heir":
                heirs += 1

        if heirs == 1 and self.owner is not None:
            self.owner.total_points += 2

        return result


class StandardCardFactory(AbstractCardFactory):
    """A factory class to create standard cards for the game."""

    cards_classes: Sequence[TypeAbstractCard] = [Soldier, Archer, ChangeForm, Heir]

    def __iter__(self) -> Iterable[TypeAbstractCard]:
        """Create a list of standard cards for the given game instance."""
        colors: list[CardColors] = [color for color in CardColors.__members__.values()]
        for color in colors:
            for item in self.cards_classes:
                card_instance = item(self.game)
                card_instance.color = color
                yield card_instance
