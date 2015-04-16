package akobi

// ClientRole is the type for client roles
type ClientRole int

// Available client roles
const (
	Interviewer ClientRole = iota
	Interviewee
)
