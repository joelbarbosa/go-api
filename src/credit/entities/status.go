package entities

type Status int64

const (
	Accepted Status = iota
	Rejected
)

func (s Status) String() string {
	switch s {
	case Accepted:
		return "credit line authorized"
	case Rejected:
		return "credit line rejected"
	}
	return "credit line under analysis"
}
