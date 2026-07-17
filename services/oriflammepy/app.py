import asyncio

from fastapi import FastAPI
from fastapi.websockets import WebSocket

from game.base import Game
from game.ws_clients import WebsocketClient
from typings import TypeWebsocketClient, WebsocketActions

GAMES: dict[str, "Game"] = {}

CONNECTIONS: set[TypeWebsocketClient] = set()

app = FastAPI()


@app.post('/create')
async def create_game() -> dict[str, str]:
    instance = Game()
    GAMES[instance.game_id] = instance
    
    if instance.can_start and not instance.is_running:
        # await instance.start()
        async with asyncio.TaskGroup() as tg:
            tg.create_task(instance.start())

    return {"game_id": instance.game_id}


@app.websocket("/ws/{game_id}")
async def join_game(game_id: str, websocket: WebSocket) -> None:
    await websocket.accept()

    game: Game | None = GAMES.get(game_id)
    if game is None:
        await websocket.close()
        return

    client = WebsocketClient(websocket)
    CONNECTIONS.add(client)
    game.add_player(client)

    # Here you would handle the game logic with the websocket
    # For now, just keep the connection open
    while True:
        try:
            message = await client.receive()
        except Exception as e:
            print(f"Error receiving message: {e}")

            CONNECTIONS.remove(client)
            game.remove_player(client)

            await client.close()
            break
        else:
            match message.action:
                case WebsocketActions.MUST_IDENTIFY.value:
                    pass
                case _:
                    pass
