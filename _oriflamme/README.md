## Architecture

The current architecture for the server is the following:

```

                ┌────────────────────┐
                │   ServerRegistry   │
                └─────────┬──────────┘
                          │
                ┌─────────▼──────────┐
                │    GameRegistry    │  (1 per game)
                │  - clients         │
                │  - broadcast chan  │
                └─────────┬──────────┘
                          │
         ┌────────────────┼────────────────┐
         │                │                │
     Client A         Client B         Client C
```



We will be moving to a pub/sub architecture with where each **GameRoom becomes a Redis channel:**

```

          ┌──────────────┐
          │   Client A   │
          └──────┬───────┘
                 │
          ┌──────▼───────┐
          │  WS Server 1 │
          └──────┬───────┘
                 │ publish
          ┌──────▼────────────┐
          │     Redis         │
          │   Pub/Sub Bus     │
          └──────┬────────────┘
                 │ subscribe
          ┌──────▼───────┐
          │  WS Server 2 │
          └──────┬───────┘
                 │
          ┌──────▼───────┐
          │   Client B   │
          └──────────────┘
```
