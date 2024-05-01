package entities

type NonLogable[T any] struct {
	Value T
}

func (v *NonLogable[T]) String() string {
	return "[REDACTED]"
}
