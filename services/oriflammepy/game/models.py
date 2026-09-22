from typing import Optional

import pydantic


class WebsocketSend(pydantic.BaseModel):
    action: str
    message: Optional[str] = None


class WebsocketReceive(pydantic.BaseModel):
    action: str
    message: Optional[str] = None


class HistoryEntry(pydantic.BaseModel):
    action: str


class PubSubMessage(pydantic.BaseModel):
    channel: str
    data: dict[str, str]
