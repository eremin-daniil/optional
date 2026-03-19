package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type typedWrapper[T comparable] interface {
	IsPresent() bool
	IsNull() bool
	IsMissing() bool
	MustGet() T
}

func assertTypedWrapperConstructors[T comparable, W typedWrapper[T]](
	t *testing.T,
	present W,
	presentFromPointer W,
	null W,
	nullFromPointer W,
	missing W,
	want T,
) {
	t.Helper()

	assert.True(t, present.IsPresent())
	assert.False(t, present.IsNull())
	assert.False(t, present.IsMissing())
	assert.Equal(t, want, present.MustGet())

	assert.True(t, presentFromPointer.IsPresent())
	assert.False(t, presentFromPointer.IsNull())
	assert.False(t, presentFromPointer.IsMissing())
	assert.Equal(t, want, presentFromPointer.MustGet())

	assert.True(t, null.IsNull())
	assert.False(t, null.IsPresent())
	assert.False(t, null.IsMissing())

	assert.True(t, nullFromPointer.IsNull())
	assert.False(t, nullFromPointer.IsPresent())
	assert.False(t, nullFromPointer.IsMissing())

	assert.True(t, missing.IsMissing())
	assert.False(t, missing.IsPresent())
	assert.False(t, missing.IsNull())
}

func assertTypedWrapperJSONRoundTrip[T comparable, W typedWrapper[T]](t *testing.T, original W, wantJSON string, want T) {
	t.Helper()

	data, err := json.Marshal(original)
	require.NoError(t, err)
	assert.JSONEq(t, wantJSON, string(data))

	var restored W
	require.NoError(t, json.Unmarshal(data, &restored))
	assert.True(t, restored.IsPresent())
	assert.Equal(t, want, restored.MustGet())
}

func assertTypedWrapperJSONNull[T comparable, W typedWrapper[T]](t *testing.T, original W) {
	t.Helper()

	data, err := json.Marshal(original)
	require.NoError(t, err)
	assert.Equal(t, "null", string(data))

	var restored W
	require.NoError(t, json.Unmarshal(data, &restored))
	assert.True(t, restored.IsNull())
}

func TestTypedWrappers_Constructors(t *testing.T) {
	t.Parallel()

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		value := true
		assertTypedWrapperConstructors(t, OfBool(value), OfNullableBool(&value), NullBool(), OfNullableBool(nil), MissingBool(), value)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		value := 42
		assertTypedWrapperConstructors(t, OfInt(value), OfNullableInt(&value), NullInt(), OfNullableInt(nil), MissingInt(), value)
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		value := int8(8)
		assertTypedWrapperConstructors(t, OfInt8(value), OfNullableInt8(&value), NullInt8(), OfNullableInt8(nil), MissingInt8(), value)
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		value := int16(16)
		assertTypedWrapperConstructors(t, OfInt16(value), OfNullableInt16(&value), NullInt16(), OfNullableInt16(nil), MissingInt16(), value)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		value := int32(32)
		assertTypedWrapperConstructors(t, OfInt32(value), OfNullableInt32(&value), NullInt32(), OfNullableInt32(nil), MissingInt32(), value)
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		value := int64(64)
		assertTypedWrapperConstructors(t, OfInt64(value), OfNullableInt64(&value), NullInt64(), OfNullableInt64(nil), MissingInt64(), value)
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()
		value := uint(42)
		assertTypedWrapperConstructors(t, OfUint(value), OfNullableUint(&value), NullUint(), OfNullableUint(nil), MissingUint(), value)
	})

	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		value := uint8(8)
		assertTypedWrapperConstructors(t, OfUint8(value), OfNullableUint8(&value), NullUint8(), OfNullableUint8(nil), MissingUint8(), value)
	})

	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		value := uint16(16)
		assertTypedWrapperConstructors(t, OfUint16(value), OfNullableUint16(&value), NullUint16(), OfNullableUint16(nil), MissingUint16(), value)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		value := uint32(32)
		assertTypedWrapperConstructors(t, OfUint32(value), OfNullableUint32(&value), NullUint32(), OfNullableUint32(nil), MissingUint32(), value)
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		value := uint64(64)
		assertTypedWrapperConstructors(t, OfUint64(value), OfNullableUint64(&value), NullUint64(), OfNullableUint64(nil), MissingUint64(), value)
	})

	t.Run("uintptr", func(t *testing.T) {
		t.Parallel()
		value := uintptr(123)
		assertTypedWrapperConstructors(t, OfUintptr(value), OfNullableUintptr(&value), NullUintptr(), OfNullableUintptr(nil), MissingUintptr(), value)
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		value := float32(3.25)
		assertTypedWrapperConstructors(t, OfFloat32(value), OfNullableFloat32(&value), NullFloat32(), OfNullableFloat32(nil), MissingFloat32(), value)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		value := 6.5
		assertTypedWrapperConstructors(t, OfFloat64(value), OfNullableFloat64(&value), NullFloat64(), OfNullableFloat64(nil), MissingFloat64(), value)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		value := "hello"
		assertTypedWrapperConstructors(t, OfString(value), OfNullableString(&value), NullString(), OfNullableString(nil), MissingString(), value)
	})
}

func TestTypedWrappers_JSON(t *testing.T) {
	t.Parallel()

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfBool(true), "true", true)
		assertTypedWrapperJSONNull(t, NullBool())
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfInt(42), "42", 42)
		assertTypedWrapperJSONNull(t, NullInt())
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfInt8(8), "8", int8(8))
		assertTypedWrapperJSONNull(t, NullInt8())
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfInt16(16), "16", int16(16))
		assertTypedWrapperJSONNull(t, NullInt16())
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfInt32(32), "32", int32(32))
		assertTypedWrapperJSONNull(t, NullInt32())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfInt64(64), "64", int64(64))
		assertTypedWrapperJSONNull(t, NullInt64())
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUint(42), "42", uint(42))
		assertTypedWrapperJSONNull(t, NullUint())
	})

	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUint8(8), "8", uint8(8))
		assertTypedWrapperJSONNull(t, NullUint8())
	})

	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUint16(16), "16", uint16(16))
		assertTypedWrapperJSONNull(t, NullUint16())
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUint32(32), "32", uint32(32))
		assertTypedWrapperJSONNull(t, NullUint32())
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUint64(64), "64", uint64(64))
		assertTypedWrapperJSONNull(t, NullUint64())
	})

	t.Run("uintptr", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfUintptr(123), "123", uintptr(123))
		assertTypedWrapperJSONNull(t, NullUintptr())
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfFloat32(3.25), "3.25", float32(3.25))
		assertTypedWrapperJSONNull(t, NullFloat32())
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfFloat64(6.5), "6.5", 6.5)
		assertTypedWrapperJSONNull(t, NullFloat64())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		assertTypedWrapperJSONRoundTrip(t, OfString("hello"), `"hello"`, "hello")
		assertTypedWrapperJSONNull(t, NullString())
	})
}
