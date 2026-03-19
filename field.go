package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Field[T any] struct {
	value T
	state state
}

func Of[T any](value T) Field[T] {
	return Field[T]{value: value, state: statePresent}
}

func OfNullable[T any](ptr *T) Field[T] {
	if ptr == nil {
		return Field[T]{state: stateNull}
	}
	return Field[T]{value: *ptr, state: statePresent}
}

func Null[T any]() Field[T] {
	return Field[T]{state: stateNull}
}

func Missing[T any]() Field[T] {
	return Field[T]{state: stateMissing}
}

func (f Field[T]) Get() (T, bool) {
	return f.value, f.IsPresent()
}

func (f Field[T]) GetOr(defaultValue T) T {
	if f.IsPresent() {
		return f.value
	}
	return defaultValue
}

func (f Field[T]) GetOrNil() *T {
	if f.IsPresent() {
		return &f.value
	}
	return nil
}

func (f Field[T]) MustGet() T {
	if !f.IsPresent() {
		panic("optional: value is not set")
	}
	return f.value
}

func (f Field[T]) IsMissing() bool {
	return f.state == stateMissing
}

func (f Field[T]) IsNull() bool {
	return f.state == stateNull
}

func (f Field[T]) IsPresent() bool {
	return f.state == statePresent
}

var _ json.Marshaler = Field[int]{}

func (f Field[T]) MarshalJSON() ([]byte, error) {
	switch f.state {
	case stateMissing:
		return nil, fmt.Errorf("optional: missing value (did you forget omitzero?)")
	case stateNull:
		return []byte("null"), nil
	case statePresent:
		return json.Marshal(f.value)
	default:
		return nil, fmt.Errorf("optional: unknown state %v", f.state)
	}
}

var _ json.Unmarshaler = &Field[int]{}

func (f *Field[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
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
