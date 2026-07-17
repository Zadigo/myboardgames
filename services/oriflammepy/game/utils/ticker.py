import asyncio
from typing import NoReturn

from typings import TypeWebsocketClient


async def ticker(clients: set[TypeWebsocketClient]) -> NoReturn:
    """
    A simple ticker that sends a "tick" message to all connected clients every second.
    This function runs indefinitely and should be run in an asyncio event loop.
    """
    while True:
        for client in clients:
            try:
                await client.send("tick")
            except Exception:
                clients.remove(client)
        await asyncio.sleep(40)  # Adjust the sleep time as needed
