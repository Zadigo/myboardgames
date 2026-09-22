import dataclasses
from typing import Optional

import redis

from typings import TypeAbstractCard


@dataclasses.dataclass
class MustResolve:
    """Class to represent an action that must be resolved in the game.
    
    Attributes:
        action (str): The action that needs to be resolved.
        from_card (TypeAbstractCard): The card that initiated the action.
        on_previous_card (Optional[TypeAbstractCard]): The previous card on which the action is being performed in the queue, if applicable.
        on_next_card (Optional[TypeAbstractCard]): The next card in the queue on which the action is being performed, if applicable.
        user_selection (Optional[str]): The card selected by the user to resolve the action.
    """
    
    action: str
    from_card: TypeAbstractCard
    on_previous_card : Optional[TypeAbstractCard] = None
    on_next_card : Optional[TypeAbstractCard] = None
    user_selection: Optional[str] = None
    on_first_card : Optional[TypeAbstractCard] = None
    on_last_card : Optional[TypeAbstractCard] = None


def create_redis_connection() -> redis.Redis:
    instance = redis.Redis(host='localhost', port=6379, db=0)
    try:
        instance.ping()
        print("Connected to Redis")
    except redis.ConnectionError:
        print("Failed to connect to Redis")
    return instance
