import asyncio
from typing import NoReturn


async def l1() -> NoReturn:
    while True:
        print("Hello from l1")
        await asyncio.sleep(1)


async def l2() -> NoReturn:
    while True:
        print("Hello from l2")
        await asyncio.sleep(1)


async def main() -> None:
    task1: asyncio.Task[NoReturn] = asyncio.create_task(l1())
    task2: asyncio.Task[NoReturn] = asyncio.create_task(l2())

    await asyncio.gather(task1, task2)


if __name__ == "__main__":
    asyncio.run(main())
