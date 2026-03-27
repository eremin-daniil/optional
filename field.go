package optional

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var nullBytes = []byte("null")

// Field is a generic three-state container that distinguishes between
// a present value, an explicit null, and a missing (omitted) field.
// The zero value of Field is the missing state.
type Field[T any] struct {
	value T
	state state
}

// Of creates a Field containing the given value in the present state.
func Of[T any](value T) Field[T] {
	return Field[T]{value: value, state: statePresent}
}

// FromPtr creates a Field from a pointer. If ptr is nil, the Field is null.
// If ptr is non-nil, the Field contains the dereferenced value.
func FromPtr[T any](ptr *T) Field[T] {
	if ptr == nil {
		return Field[T]{state: stateNull}
	}
	return Field[T]{value: *ptr, state: statePresent}
}

// OfNullable creates a Field from a pointer.
// It is an alias for [FromPtr].
func OfNullable[T any](ptr *T) Field[T] {
	return FromPtr(ptr)
}

// Null creates a Field in the null state.
func Null[T any]() Field[T] {
	return Field[T]{state: stateNull}
}

// Missing creates a Field in the missing state.
// This is equivalent to the zero value of Field[T].
func Missing[T any]() Field[T] {
	return Field[T]{state: stateMissing}
}

// Get returns the value and a boolean indicating whether the value is present.
func (f Field[T]) Get() (T, bool) {
	return f.value, f.IsPresent()
}

// GetOr returns the value if present, otherwise returns defaultValue.
func (f Field[T]) GetOr(defaultValue T) T {
	if f.IsPresent() {
		return f.value
	}
	return defaultValue
}

// Ptr returns a pointer to the value if present, otherwise nil.
func (f Field[T]) Ptr() *T {
	if f.IsPresent() {
		return &f.value
	}
	return nil
}

// MustGet returns the value if present, otherwise panics.
func (f Field[T]) MustGet() T {
	if !f.IsPresent() {
		panic("optional: value is not set")
	}
	return f.value
}

// Or returns f if it contains a present value, otherwise returns other.
func (f Field[T]) Or(other Field[T]) Field[T] {
	if f.IsPresent() {
		return f
	}
	return other
}

// OrElse returns the value if present, otherwise calls fn and returns its result.
func (f Field[T]) OrElse(fn func() T) T {
	if f.IsPresent() {
		return f.value
	}
	return fn()
}

// IsMissing reports whether the field is in the missing state.
func (f Field[T]) IsMissing() bool {
	return f.state == stateMissing
}

// IsNull reports whether the field is in the null state.
func (f Field[T]) IsNull() bool {
	return f.state == stateNull
}

// IsPresent reports whether the field contains a value.
func (f Field[T]) IsPresent() bool {
	return f.state == statePresent
}

// IsZero reports whether the field is in the missing state (zero value of Field).
// This method supports the omitzero JSON tag introduced in Go 1.24+.
func (f Field[T]) IsZero() bool {
	return f.state == stateMissing
}

var _ fmt.Stringer = Field[int]{}

// String returns a human-readable string representation of the field.
func (f Field[T]) String() string {
	switch f.state {
	case stateMissing:
		return "missing"
	case stateNull:
		return "null"
	case statePresent:
		return fmt.Sprintf("%v", f.value)
	default:
		return "unknown"
	}
}

var _ fmt.GoStringer = Field[int]{}

// GoString returns a Go-syntax representation of the field for use with %#v.
func (f Field[T]) GoString() string {
	switch f.state {
	case stateMissing:
		return "optional.Missing()"
	case stateNull:
		return "optional.Null()"
	case statePresent:
		return fmt.Sprintf("optional.Of(%#v)", f.value)
	default:
		return fmt.Sprintf("optional.Field{state: %d}", f.state)
	}
}

var _ json.Marshaler = Field[int]{}

// MarshalJSON implements [json.Marshaler].
// Returns an error for missing values — use the omitzero tag to omit them.
func (f Field[T]) MarshalJSON() ([]byte, error) {
	switch f.state {
	case stateMissing:
		return nil, fmt.Errorf("optional: missing value (did you forget omitzero?)")
	case stateNull:
		return nullBytes, nil
	case statePresent:
		return json.Marshal(f.value)
	default:
		return nil, fmt.Errorf("optional: unknown state %v", f.state)
	}
}

var _ json.Unmarshaler = &Field[int]{}

// UnmarshalJSON implements [json.Unmarshaler].
// JSON null sets the field to the null state. Any other valid JSON sets it to present.
// Fields not present in JSON remain in the missing state (zero value).
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)

	if bytes.Equal(data, nullBytes) {
		*f = Field[T]{state: stateNull}
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = Field[T]{value: v, state: statePresent}
	return nil
}

var _ driver.Valuer = Field[int]{}

// Value implements [driver.Valuer].
// Returns nil for null, an error for missing, and the raw value for present.
//
// Note: for scalar wrapper types (Int, UUID, Decimal, etc.),
// the Value method is overridden to return a valid [driver.Value] type.
func (f Field[T]) Value() (driver.Value, error) {
	switch f.state {
	case stateMissing:
		return nil, fmt.Errorf("optional: missing value cannot be used in SQL")
	case stateNull:
		return nil, nil
	case statePresent:
		return f.value, nil
	default:
		return nil, fmt.Errorf("optional: unknown state %v", f.state)
	}
}
