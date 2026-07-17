from typing import Optional

import pydantic


class WebsocketReceive(pydantic.BaseModel):
    action: str
    message: Optional[str] = None


class HistoryEntry(pydantic.BaseModel):
    action: str
