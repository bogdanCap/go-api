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

func FromInt(value int8) Status {
    switch Status(value) {
    case StatusNew:
        return StatusNew
    case StatusProcessing:
        return StatusProcessing
    case StatusCompleted:
        return StatusCompleted
    case StatusCancelled:
        return StatusCancelled
    default:
        return Status(0)
    }
}
