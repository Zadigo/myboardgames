# Booard Games

## Requirements & Assumptions 🟠

### Clarifying Questions

*Questions that need to be answered to better understand the requirements and constraints of the system*

- **Channels** Email ? SMS ? Push Notifications ?
- **User Authentication** Email/Password ? Social Login ?
- **Players** How many players simultanuously?

### Functional Requirements 🟢

*Describes the specific features and functionalities that the system must provide*

- **Mainsite** application for accessing all the games for the platform
- **Game service** Golang application that hosts the logic for the board game

## Capacity Planning ⏰

### Database

*Estimates the expected load on the system, such as the number of users, transactions, or requests per second. This helps in designing a system that can handle the anticipated traffic and scale as needed*

Estimation for one user campaign:

- **Players** 4 to 6 players per game

### Storage

*Estimates the storage requirements for the system, such as the amount of data that needs to be stored and the growth rate over time. This helps in designing a system that can handle the anticipated storage needs and scale as required.*

## High Level Architecture 🏗️

*Describes the overall structure of the system, including the main components and how they interact with each other. This can be illustrated using diagrams such as component diagrams or architecture diagrams.*

```mermaid
flowchart

A[Nuxt] --> B(Game A)
A --> C(Game B)
A --> D(Game C)

B --> E(Redis)
C --> E(Redis)
D --> E(Redis)

```

## System Workflow 🔄

*Explains the sequence of interactions between different components of the system, such as how a user request flows through the application, how data is processed, and how responses are generated. This can be illustrated using sequence diagrams or flowcharts.*

```mermaid
sequenceDiagram
autonumber

box Frontend
actor P as Other User
actor U as User
participant N@{type: "entity"} as Nuxt
end 

box Backend
participant G@{type: "entity"} as Golang
participant R@{type: "database"} as Redis
participant D@{type: "boundary"} as Django
participant PG@{type: "database"} as PostgreSQL
end

U->> N: Create game
N->> G: New Game

par storage
G-->>R: Save game
R-->>G: Game details

G->>()N: Create lobby
N()<<-->>()U: Wait for players
P->>N: Join lobby

alt Number of players
N->>G: Start game
else Not enough players
G->>N: Stop lobby
end

par Playing
U<<-->>G: Game active
G<<-->>R: Saving game details
G-->>D: Saving game details
D<<-->>PG: Saving game details
end
end
```

* Emailing service architecture: [services/goemailer/ARCHITECTURE.md](services/goemailer/ARCHITECTURE.md)

## Api Design 🛠️

*Describes the design of the APIs that will be used for communication between different components of the system, such as the frontend and backend. This includes the endpoints, request and response formats, authentication mechanisms, and any other relevant details about how the APIs will function.*

> Determines also whether the system will be using RESTful APIs or GraphQL, and how the frontend will interact with these APIs to fetch and manipulate data.
> If the system uses microservices architecture, the API design will also include details about how different microservices will communicate with each other, such as using RESTful APIs, gRPC, or message queues.

| Endpoint           | Method | Description               | Request Body        | Response Body                                                      |
| ------------------ | ------ | ------------------------- | ------------------- | ------------------------------------------------------------------ |
| /create            | POST   | Create a new game         | { userId: number } | { tableId: string, currentPlayers: object, gameDetails: object }  |
| /{tableId}/join    | POST   | Join a table (Websocket)  | { userId: number } | { tableId: string, currentPlayers: object, gameDetails: object } |
| /{tableId}/observe | POST   | Observe a game on a table | { userId: number }  | { tableId: string, currentPlayers: object, gameDetails: object }  |

## Data storage

*Describes how the system will store and manage data, including the choice of database (e.g., relational, NoSQL), data models, and how data will be accessed and manipulated by the application.*

### Amazon S3

*Explains the the manner in which the system will use Amazon S3 for storing and retrieving files, including the structure of the S3 buckets, access control policies, and how the application will interact with S3 for file uploads and downloads.*

### Database

*Explains the choice of database (e.g., relational, NoSQL) and how it will be used to store and manage data for the application. This includes the data models, relationships between entities, and how the application will perform CRUD (Create, Read, Update, Delete) operations on the database.*

```mermaid
erDiagram

```

## Caching

*Describes the caching strategy for the application, including what data will be cached, how it will be cached (e.g., in-memory cache, distributed cache), and how the cache will be invalidated when data changes. For example, product data that is frequently accessed but infrequently updated can be cached to improve performance and reduce load on the database.*

Caching will be almost exlusively done with Redis as an in-memory data store.

## Scalability

*Describes how the system will be designed to handle increasing loads and scale as needed. This includes strategies for horizontal scaling (adding more servers) and vertical scaling (upgrading existing servers), as well as any load balancing techniques that will be used to distribute traffic across multiple servers.*

```mermaid
flowchart LR

A[User Requests] --> B(Load Balancer)
B --> C[Server 1]
B --> D[Server 2]
B --> E[Server 3]

C --> G[Game Service]
D --> H[Game Service]
E --> I[Game Service]
```

---

## References ⏰

*List of services and components that will be part of the system, along with their respective technologies and descriptions. This can be presented in a tabular format for clarity.*

| Service   | Language/Framework | Description                            |
| --------- | ------------------ | -------------------------------------- |
| Mainsite  | Nuxt               | Allows users to find and play games    |
| Flipseven | Golang             | Adaptation for the card game Flipseven |
| Splendor  | Golang             | Adaptation for the card game Splendor  |
| Oriflamme | Golang             | Adaptation for the card game Oriflamme |
| Knarr     | Golang             | Adaptation for the card game Knarr     |

## Technologies Used 🌳

*List of the main technologies used in the system, along with their purpose and version. This can help in understanding the technical stack of the application and how different components are implemented.*

| Technology        | Purpose/Usage            | Version |
| ----------------- | ------------------------ | ------- |
| Django            | Web framework            | ✅ 6.X  |
| PostgreSQL        | Database                 | ✅ 13.X |
| Redis             | Caching, message broker  | ✅ -    |
| RabbitMQ          | Message broker           | ✅ -    |
| Docker            | Containerization         | ✅ 20.X |
| Nuxt 4            | Frontend framework       | ✅ 4.X  |
| Firebase          | Authentication, database | ✅ -    |
| AWS S3            | Static and media storage | ✅ -    |
| Cloudfront        | CDN for static files     | ✅ -    |
| Google Analytics  | Traffic analysis         | ✅ -    |
| Facebook Pixels   | Traffic analysis         | ✅ -    |
| Microsoft Clarity | Traffic analysis         | ✅ -    |
