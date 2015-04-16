package akobi

import (
	"errors"
	"time"

	"github.com/akobi/akobi-server/log"
	"github.com/akobi/akobi-server/message"
)

// Interview is the main interview interface
type Interview interface {
	log.Emitter
	// message.Emitter
	// message.Receiver

	// Start the interview. Calling Start() implies that a client
	// has successfully authenticated and entered the interview,
	// and that all applications can be initialized and appropriate
	// websockets opened
	Start() error

	// End the interview. End() should be called if all clients
	// have successfully left the interview and the interview's
	// scheduled time has completed.
	End() error

	// ExpectClient adds a client's details to the database for
	// retrieval when the client attempts to join the interview
	ExpectClient(role ClientRole, name, email string) error

	// ExpectApplication adds an application to the list of
	// applications for this interview. The list is maintained
	// in the database
	ExpectApplication(app Application) error

	// SetStartTime adds the interview start time to the database
	SetStartTime(t *time.Time) error

	// SetDuration adds the expected duration to the database
	SetDuration(d time.Duration) error
}

type akobiInterview struct {
	id string

	expectedClients      map[ClientRole]map[string]string
	expectedApplications map[string]Application
	startTime            *time.Time
	duration             time.Duration
}

var activeInterviews map[string]Interview

const (
	interviewPrefix   string = "interview:"
	interviewIDLength int    = 12
)

var (
	errInvalidClientRole  = errors.New("Invalid client role")
	errAlreadyExpectedApp = errors.New("Already expected application")
)

// NewInterview returns a new instance of an interview
func NewInterview() (Interview, error) {
	interview := &akobiInterview{
		id:                   "blah",
		expectedClients:      make(map[ClientRole]map[string]string),
		expectedApplications: make(map[string]Application),
		startTime:            nil,
		duration:             nil,
	}

	interview.expectedClients[Interviewer] = make(map[string]string)
	interview.expectedClients[Interviewee] = make(map[string]string)

	return interview, nil
}

func (i *akobiInterview) getPrefixString() string {
	return interviewPrefix + i.id
}

// GetLogLevel returns the minimum level for log entries
// by an interview
func (i *akobiInterview) GetLogLevel() log.Level {
	return log.INFO
}

// GetLogPrefix returns the string which will prefix every
// log entry emitted by an interview
func (i *akobiInterview) GetLogPrefix() string {
	return i.getPrefixString()
}

func (i *akobiInterview) Start() error {
	return nil
}

func (i *akobiInterview) End() error {
	return nil
}

func (i *akobiInterview) ExpectClient(role ClientRole, name, email string) error {
	if role != Interviewer || role != Interviewee {
		return errInvalidClientRole
	}

	i.expectedClients[role]["name"] = name
	i.expectedClients[role]["email"] = email
	return nil
}

func (i *akobiInterview) ExpectApplication(app Application) error {
	appKey := i.GetMessageType()

	if _, ok := i.expectedApplications[appKey]; ok {
		return errAlreadyExpectedApp
	}

	i.expectedApplications[appKey] = app
	return nil
}

func (i *akobiInterview) SetStartTime(t *time.Time) error {
    i.startTime = t
    return nil
}

func (i *akobiInterview) SetDuration(d time.Duration) error {
    i.duration = d
    return nil
}
