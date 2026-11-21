package models

// Common table events that occur in most board games
const (
	TableEventIdle    string = "table_event_idle"    // The table is waiting for players to join
	TableEventWaiting string = "table_event_waiting" // The table is waiting for player(s) to take a decision
	TableEventCreated string = "table_event_created"
	TableEventUpdated string = "table_event_updated"

	TableEventAddPlayer    string = "table_event_add_player"
	TableEventRemovePlayer string = "table_event_remove_player"
	TableEventBanPlayer    string = "table_event_ban_player"

	TableEventCreateDeck  string = "table_event_create_deck"
	TableEventShuffleDeck string = "table_event_shuffle_cards"

	TableEventDistributeCards         string = "table_event_distribute_cards"
	TableEventExchangeCardsWithPlayer string = "table_event_exchange_cards_with_player"
	TableEventExchangeCardsWithDeck   string = "table_event_exchange_cards_with_deck"
	TableEventStealCard               string = "table_event_steal_card"
	TableEventDrawCard                string = "table_event_draw_card"

	TableEventNextTurn    string = "table_event_next_turn"
	TableEventReverseTurn string = "table_event_reverse_turn"
)
