from game.cards.base import BaseCard
from game.utils.base import MustResolve
from typings import CardResolutionActions, TypeGame


class Soldier(BaseCard):
    def __init__(self, game: TypeGame) -> None:
        super().__init__(game=game, name="Soldier")

    def resolve(self) -> None:
        if self.partial_resolve:
            previous_card: BaseCard | None = self.game.card_queue.get_previous_card(self)
            next_card: BaseCard | None = self.game.card_queue.get_next_card(self)
            
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


class Archer(BaseCard):
    def __init__(self, game: TypeGame) -> None:
        super().__init__(game=game, name="Archer")

    def partial_resolve(self) -> MustResolve:
        instance = MustResolve(
            CardResolutionActions.REMOVE_CARD.value, 
            from_card=self
        )

        first_card: BaseCard | None = self.game.card_queue.get_previous_card(self)
        last_card: BaseCard | None = self.game.card_queue.get_next_card(self)

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
