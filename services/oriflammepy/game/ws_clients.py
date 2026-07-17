import uuid
from abc import ABC, abstractmethod

from fastapi import WebSocket


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
        pass

    @abstractmethod
    async def receive(self) -> str:
        pass

    @abstractmethod
    async def close(self) -> None:
        pass


class WebsocketClient(BaseWebsocketClient):
    async def send(self, message: str) -> None:
        await self.websocket.send_text(message)

    async def receive(self) -> str:
        return await self.websocket.receive_text()

    async def close(self) -> None:
        await self.websocket.close()
