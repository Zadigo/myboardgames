package logic

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Action   string `json:"action"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
	// The unique identifier for the player/client sending
	// the message. This can be used to associate messages with
	// specific players and manage game state accordingly.
	PlayerUuid  string  `json:"playerUuid,omitempty"`
	CardAction  string  `json:"cardAction,omitempty"`
	CardsInPlay []*Card `json:"cardsInPlay,omitempty"`
	// The unique identifier for the game table that the player is trying to join or create.
	// This can be used to associate players with specific games and manage game state accordingly.
	TableUuid string `json:"tableUuid,omitempty"`
	// Indicates the player/client is the one that initiated the new game
	Initiator        bool          `json:"initiator,omitempty"`
	SelectedCards    []string      `json:"selectedCards,omitempty"`
	CardUuid         string        `json:"cardUuid,omitempty"`
	GameRegistry     *GameRegistry `json:"gameRegistry,omitempty"`
	CardQueue        *CardQueue    `json:"queue,omitempty"`
	GiveCardToPlayer string        `json:"giveToPlayer,omitempty"`
	GiveCardUuid     string        `json:"giveCardUuid,omitempty"`
}

// The WebsocketClient struct represents a client that is connected to
// the server via a websocket connection.
type WebsocketClient struct {
	// Unique identifier for the player
	Uuid string `json:"uuid"`
	// Player's username
	Username string `json:"username"`
	// Indicates that the player is the initiator of the game
	Initiator bool `json:"initiator"`
	// Whether the player was stopped by the freeze card
	// or cannot continue flipping cards because he has
	// two similar card numbers
	IsFreezed bool `json:"isFreezed"`
	// Cumulative score across all rounds
	// played by the player
	TotalScore int `json:"totalScore"`
	// Indicates that the player is the initiator of the game
	IsInitiator bool `json:"isInitiator"`
	// The player's websocket connection, not serialized to JSON
	conn *websocket.Conn `json:"-"`
	// The send channel is used to queue messages
	// that need to be sent to the client.
	send chan WebsocketMessage `json:"-"`
	// Mutex to protect access to the player's data, not serialized to JSON
	mu sync.Mutex `json:"-"`
}

func (client *WebsocketClient) SendJsonMessage(message WebsocketMessage) error {
	return client.conn.WriteJSON(message)
}

func (client *WebsocketClient) ReceiveJsonMessage() (WebsocketMessage, error) {
	var message WebsocketMessage
	err := client.conn.ReadJSON(&message)
	return message, err
}

func (c *WebsocketClient) GetCards(queue *CardQueue) []*Card {
	cards := []*Card{}
	for _, card := range queue.Queue {
		if card.Owner != nil && card.Owner.Uuid == c.Uuid {
			cards = append(cards, card)
		}
	}
	return cards
}

func (c *WebsocketClient) GetCurrentScore(queue *CardQueue) int {
	currentHand := c.GetCards(queue)

	score := 0
	for _, card := range currentHand {
		score += card.Value
	}
	return score
}

func (c *WebsocketClient) NumberOfCards(queue *CardQueue) int {
	return len(c.GetCards(queue))
}

func (c *WebsocketClient) HasSevenCards(queue *CardQueue) bool {
	return c.NumberOfCards(queue) == 7
}

func (c *WebsocketClient) SetFlipCard(queue *CardQueue, selectedCard *Card) bool {
	cards := c.GetCards(queue)
	selectedCard.Owner = c

	var secondChanceCard *Card
	var hasNumber bool

	for _, card := range cards {
		if card.IsSpecial && !card.IsDiscarded {
			// Find the first occurence of a "Second Chance"
			// card in the player's hand
			if card.Category == "Second Chance" {
				secondChanceCard = card
				break
			}
		}
	}

	for _, card := range cards {
		if selectedCard.Value == card.Value {
			hasNumber = true

			if secondChanceCard != nil {
				hasNumber = false
				secondChanceCard.IsDiscarded = true
				break
			}
		}
	}

	return hasNumber
}

func NewWebsocketClient(conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		Uuid:     uuid.NewString(),
		Username: "",
		TotalScore: 0,
		IsInitiator: false,
		conn:     conn,
		send: make(chan WebsocketMessage),
	}
}
