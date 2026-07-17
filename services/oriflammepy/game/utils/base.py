import dataclasses
from typing import Optional

from typings import TypeCard


@dataclasses.dataclass
class MustResolve:
    """Class to represent an action that must be resolved in the game.
    
    Attributes:
        action (str): The action that needs to be resolved.
        from_card (TypeCard): The card that initiated the action.
        on_card (Optional[TypeCard]): The card that the action is being performed on, if applicable.
        on_previous_card (Optional[TypeCard]): The previous card on which the action is being performed in the queue, if applicable.
        on_next_card (Optional[TypeCard]): The next card in the queue on which the action is being performed, if applicable.
        user_selection (Optional[str]): The card selected by the user to resolve the action, (applicable for previous and next cards).
    """
    
    action: str
    from_card: TypeCard
    on_card: Optional[TypeCard] = None
    on_previous_card : Optional[TypeCard] = None
    on_next_card : Optional[TypeCard] = None
    user_selection: Optional[str] = None
    on_first_card : Optional[TypeCard] = None
    on_last_card : Optional[TypeCard] = None
