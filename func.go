package optional

// Map transforms the value inside a [Field] using fn.
// If the field is present, fn is applied to the value and the result is wrapped in a new present Field.
// If the field is null or missing, the corresponding state is preserved in the returned Field.
//
// Map is a standalone function because Go does not support additional type parameters on methods.
func Map[T, U any](f Field[T], fn func(T) U) Field[U] {
	switch f.state {
	case statePresent:
		return Of(fn(f.value))
	case stateNull:
		return Null[U]()
	default:
		return Missing[U]()
	}
}

// FlatMap transforms the value inside a [Field] using fn, which itself returns a Field.
// If the field is present, fn is applied and its result is returned directly.
// If the field is null or missing, the corresponding state is preserved.
//
// FlatMap is a standalone function because Go does not support additional type parameters on methods.
func FlatMap[T, U any](f Field[T], fn func(T) Field[U]) Field[U] {
	switch f.state {
	case statePresent:
		return fn(f.value)
	case stateNull:
		return Null[U]()
	default:
		return Missing[U]()
	}
}

// Equal reports whether two fields are equal.
// Two fields are equal if they have the same state and, when present, the same value.
//
// Equal requires the [comparable] constraint because Go methods cannot introduce
// additional type constraints.
func Equal[T comparable](a, b Field[T]) bool {
	if a.state != b.state {
		return false
	}
	if a.state == statePresent {
		return a.value == b.value
	}
	return true
}
