package entities

type Ticket struct {
	ID        int  `json:"id"`
	UserID    int  `json:"user_id"`
	EventID   int  `json:"event_id"`
	Companion bool `json:"companion"`
}
