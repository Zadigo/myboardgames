import uvicorn
from fastapi import FastAPI, websockets

from its_a_wonderful_world.logic.base import Game

app = FastAPI()


@app.websocket('/ws')
async def game(websocket: websockets.WebSocket):
    await websocket.accept()
    instance = Game()
    await websocket.send_json({'message': 'Game instance created', 'game_id': instance.board.board_id})

    while True:
        await websocket.receive_json()


if __name__ == '__main__':
    uvicorn.run(app, host='127.0.0.1', port=8000)
