import uuid
from abc import ABC, abstractmethod

from fastapi import WebSocket

from game import models


class BaseWebsocketClient(ABC):
    def __init__(self, websocket: WebSocket) -> None:
        self.websocket: WebSocket = websocket
        self.client_id: str = str(uuid.uuid4())

    def __hash__(self) -> int:
        return hash(self.client_id)

    def __repr__(self) -> str:
        return f"Player({self.client_id})"

    @abstractmethod
    async def send(self, message: str) -> None:
        return NotImplemented

    @abstractmethod
    async def receive(self) -> models.WebsocketReceive:
        return NotImplemented

    @abstractmethod
    async def close(self) -> None:
        return NotImplemented


class WebsocketClient(BaseWebsocketClient):
    total_points: int = 0
    history: list[models.HistoryEntry] = []

    async def send(self, message: str) -> None:
        await self.websocket.send_text(message)

    async def receive(self) -> models.WebsocketReceive:
        data = await self.websocket.receive_json()
        return models.WebsocketReceive(**data)

    async def close(self) -> None:
        await self.websocket.close()
