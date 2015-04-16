package akobi

import (
	"github.com/akobi/akobi-server/log"
)

// Application is the interface to be implemented by all applications
type Application interface {
	log.Emitter

	// GetShortName returns a short name for the application
	GetShortName() string

	// GetDescription describes the application's function
	GetDescription() string

	// GetMessageType returns the value for the 'type' field in
	// messages exchanged between the server and client for this
	// application. The value is expected to be unique across all
	// applications for an interview
	GetMessageType() string
}
