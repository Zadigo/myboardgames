from itertools import chain
from typing import Any

import uvicorn
from fastapi import FastAPI, websockets
from fastapi.middleware.cors import CORSMiddleware

from its_a_wonderful_world.logic.base import Game
from its_a_wonderful_world.models import GameCardsResponse, WebsocketMessage
from its_a_wonderful_world.typings import WebsocketActionsTypes

app = FastAPI(
    title="It's a Wonderful World Backend",
    description="Backend for It's a Wonderful World board game implementation",
    version="0.1.0",
    debug=True,
)

origins = [
    "http://localhost:3000",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")  # Don't forget this!
async def root():
    return {"message": "Hello World"}


@app.websocket('/ws/game/{game_id}')
async def game(websocket: websockets.WebSocket, game_id: str):
    await websocket.accept()
    instance = Game()
    await websocket.send_json({'message': 'Game instance created', 'game_id': instance.board.board_id})

    while True:
        data: dict[str, Any] = await websocket.receive_json()
        validated_data = WebsocketMessage(**data)

        if validated_data.action.value == WebsocketActionsTypes.START_GAME.value:
            instance.board.prepare()


@app.get('/cards', response_model=GameCardsResponse)
async def cards() -> GameCardsResponse:
    items = []

    game = Game()
    available_cards = chain(*game.available_cards)

    for card in available_cards:
        item = {
            'name': card.name,
            'card_type': card.card_type.value,
            'category': card.card_category.value,
            'image': card.image,
            'recycling_bonus': card.recycling_bonus.value
            # 'has_construction_bonus': card.has_construction_bonus,
            # 'has_character_token_bonus': card.has_character_token_bonus,
        }
        items.append(item)

    print(items[:3])

    return GameCardsResponse(data=items)


if __name__ == '__main__':
    uvicorn.run(app, host='127.0.0.1', port=8000)
