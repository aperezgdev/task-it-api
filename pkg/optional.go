package pkg

type Optional[T any] struct {
	Value T
	IsPresent bool
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{value, true}
}

func EmptyOptional[T any]() Optional[T]{
	return Optional[T]{}
}
