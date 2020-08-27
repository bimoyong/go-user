package auth

type Kind int8

const (
	// None kind. Not authenticated.
	None Kind = iota
	// System kind. System user.
	System
	// Registered kind. Registered user.
	Registered
)

func (k Kind) String() string {
	switch k {
	case None:
		return "none"
	case System:
		return "system"
	case Registered:
		return "registered"
	}
	return ""
}
