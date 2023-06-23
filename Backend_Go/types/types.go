package types

import "time"

const EVENT_COLLECTION = "events"

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Venue       string    `json:"venue"`
	Link        string    `json:"link"`
}
