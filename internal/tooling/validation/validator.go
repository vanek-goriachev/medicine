package validation

type Validator[T any] interface {
	Validate(obj T) error
}
