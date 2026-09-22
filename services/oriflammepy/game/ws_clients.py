import uuid
from abc import ABC, abstractmethod
import pydantic
from fastapi import WebSocket

from game import models


class BaseWebsocketClient(ABC):
    total_points: int = 0
    history: list[models.HistoryEntry] = []

    def __init__(self, websocket: WebSocket) -> None:
        self.websocket: WebSocket = websocket
        self.client_id: str = str(uuid.uuid4())

    def __str__(self) -> str:
        return self.client_id

    def __hash__(self) -> int:
        return hash(self.client_id)

    def __repr__(self) -> str:
        return f"Player({self.client_id})"

    @abstractmethod
    async def send(self, message: dict[str, str]) -> None:
        return NotImplemented

    @abstractmethod
    async def receive(self) -> models.WebsocketReceive:
        return NotImplemented

    @abstractmethod
    async def close(self) -> None:
        return NotImplemented

    def destructure(self) -> dict[str, str]:
        return {
            "client_id": self.client_id,
            "total_points": self.total_points,
        }


class WebsocketClient(BaseWebsocketClient):
    """An advanced WebSocket client that manages the connection 
    and communication with the server by wrapping the FastAPI WebSocket."""

    async def send(self, message: pydantic.BaseModel | dict[str, str]) -> None:
        if isinstance(message, pydantic.BaseModel):
            message = message.model_dump()
        await self.websocket.send_json(message)

    async def receive(self) -> models.WebsocketReceive:
        data = await self.websocket.receive_json()
        return models.WebsocketReceive(**data)

    async def close(self) -> None:
        await self.websocket.close()
