package pagination

import "strconv"

const (
    DefaultLimit  = 20
    DefaultOffset = 0
    MaxLimit      = 100
)

func Limit(value *string) (int) {
	if value == nil || *value == "" {
		return DefaultLimit
	}

	limit, err := strconv.Atoi(*value)
	
	if err != nil || limit <= 0 {
		return DefaultLimit
	}

    if limit > MaxLimit {
        return MaxLimit
    }

    return limit
}

func Offset(value *string) (int) {
	if value == nil || *value == "" {
		return DefaultOffset
	}

	offset, err := strconv.Atoi(*value)

	if err != nil || offset <= 0 {
		return DefaultOffset
	}

    return offset
}