package message

import (
	"time"
)

// DateTimeFormat is the format for the DateTime field
const DateTimeFormat = time.RFC3339

// Message represents the format for server-client
// communication over a WebSocket
type Message struct {
	DateTime    *time.Time `json:"datetime"`
	Type        string     `json:"type"`
	ClientID    string     `json:"clientID"`
	InterviewID string     `json:"interviewID"`
	Data        []byte     `json:"data"`
}

// Emitter is the interface anything sending a
// message directly over a websocket must implement
type Emitter interface {
	EmitMessage() (*Message, error)
}

// Listener is the interface anything listening
// directly over a websocket must implement
type Listener interface {
	HandleMessage() (*Message, error)
}
