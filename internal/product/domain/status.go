package domain

type Status int8

const (
	StatusNew Status = iota + 1
	StatusProcessing
	StatusCompleted
	StatusCancelled
)

func (s Status) String() string {
	switch s {
	case StatusNew:
		return "new"

	case StatusProcessing:
		return "processing"

	case StatusCompleted:
		return "completed"

	case StatusCancelled:
		return "cancelled"

	default:
		return "unknown"
	}
}
