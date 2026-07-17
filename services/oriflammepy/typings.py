import enum
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from game.base import BaseGame, CardQeue
    from game.cards.base import BaseCard
    from game.ws_clients import WebsocketClient

type TypeCard = "BaseCard"


type TypeWebsocketClient = "WebsocketClient"

type TypeCardQueue = "CardQeue[TypeCard]"

type TypeGame = "BaseGame"


class WebsocketActions(enum.Enum):
    MUST_IDENTIFY = "must_identify"


class CardResolutionActions(enum.Enum):
    REMOVE_CARD = "remove_card"
