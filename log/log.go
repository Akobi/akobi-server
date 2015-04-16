package log

// Level is the type for log levels
type Level int

// All available log levels
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

// Emitter is the interface to be implemented by
// anything that wants to write out a log entry
type Emitter interface {
	GetLogLevel() Level
	GetLogPrefix() string
}
