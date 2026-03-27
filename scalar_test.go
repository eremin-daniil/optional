package optional

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type typedWrapper[T any] interface {
	IsPresent() bool
	IsNull() bool
	IsMissing() bool
	MustGet() T
}

func assertTypedWrapperConstructors[T any, W typedWrapper[T]](
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

func assertTypedWrapperJSONRoundTrip[T any, W typedWrapper[T]](t *testing.T, original W, wantJSON string, want T) {
	t.Helper()

	data, err := json.Marshal(original)
	require.NoError(t, err)
	assert.JSONEq(t, wantJSON, string(data))

	var restored W
	require.NoError(t, json.Unmarshal(data, &restored))
	assert.True(t, restored.IsPresent())
	assert.Equal(t, want, restored.MustGet())
}

func assertTypedWrapperJSONNull[T any, W typedWrapper[T]](t *testing.T, original W) {
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
		assertTypedWrapperConstructors(t, OfBool(value), FromBoolPtr(&value), NullBool(), FromBoolPtr(nil), MissingBool(), value)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		value := 42
		assertTypedWrapperConstructors(t, OfInt(value), FromIntPtr(&value), NullInt(), FromIntPtr(nil), MissingInt(), value)
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		value := int8(8)
		assertTypedWrapperConstructors(t, OfInt8(value), FromInt8Ptr(&value), NullInt8(), FromInt8Ptr(nil), MissingInt8(), value)
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		value := int16(16)
		assertTypedWrapperConstructors(t, OfInt16(value), FromInt16Ptr(&value), NullInt16(), FromInt16Ptr(nil), MissingInt16(), value)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		value := int32(32)
		assertTypedWrapperConstructors(t, OfInt32(value), FromInt32Ptr(&value), NullInt32(), FromInt32Ptr(nil), MissingInt32(), value)
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		value := int64(64)
		assertTypedWrapperConstructors(t, OfInt64(value), FromInt64Ptr(&value), NullInt64(), FromInt64Ptr(nil), MissingInt64(), value)
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()
		value := uint(42)
		assertTypedWrapperConstructors(t, OfUint(value), FromUintPtr(&value), NullUint(), FromUintPtr(nil), MissingUint(), value)
	})

	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		value := uint8(8)
		assertTypedWrapperConstructors(t, OfUint8(value), FromUint8Ptr(&value), NullUint8(), FromUint8Ptr(nil), MissingUint8(), value)
	})

	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		value := uint16(16)
		assertTypedWrapperConstructors(t, OfUint16(value), FromUint16Ptr(&value), NullUint16(), FromUint16Ptr(nil), MissingUint16(), value)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		value := uint32(32)
		assertTypedWrapperConstructors(t, OfUint32(value), FromUint32Ptr(&value), NullUint32(), FromUint32Ptr(nil), MissingUint32(), value)
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		value := uint64(64)
		assertTypedWrapperConstructors(t, OfUint64(value), FromUint64Ptr(&value), NullUint64(), FromUint64Ptr(nil), MissingUint64(), value)
	})

	t.Run("uintptr", func(t *testing.T) {
		t.Parallel()
		value := uintptr(123)
		assertTypedWrapperConstructors(t, OfUintptr(value), FromUintptrPtr(&value), NullUintptr(), FromUintptrPtr(nil), MissingUintptr(), value)
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		value := float32(3.25)
		assertTypedWrapperConstructors(t, OfFloat32(value), FromFloat32Ptr(&value), NullFloat32(), FromFloat32Ptr(nil), MissingFloat32(), value)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		value := 6.5
		assertTypedWrapperConstructors(t, OfFloat64(value), FromFloat64Ptr(&value), NullFloat64(), FromFloat64Ptr(nil), MissingFloat64(), value)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		value := "hello"
		assertTypedWrapperConstructors(t, OfString(value), FromStringPtr(&value), NullString(), FromStringPtr(nil), MissingString(), value)
	})

	t.Run("uuid", func(t *testing.T) {
		t.Parallel()
		value := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		assertTypedWrapperConstructors(t, OfUUID(value), FromUUIDPtr(&value), NullUUID(), FromUUIDPtr(nil), MissingUUID(), value)
	})

	t.Run("decimal", func(t *testing.T) {
		t.Parallel()
		value := decimal.RequireFromString("123.45")
		assertTypedWrapperConstructors(t, OfDecimal(value), FromDecimalPtr(&value), NullDecimal(), FromDecimalPtr(nil), MissingDecimal(), value)
	})

	t.Run("time", func(t *testing.T) {
		t.Parallel()
		value := time.Date(2024, time.January, 2, 3, 4, 5, 123456789, time.UTC)
		assertTypedWrapperConstructors(t, OfTime(value), FromTimePtr(&value), NullTime(), FromTimePtr(nil), MissingTime(), value)
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

	t.Run("uuid", func(t *testing.T) {
		t.Parallel()
		value := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		assertTypedWrapperJSONRoundTrip(t, OfUUID(value), `"550e8400-e29b-41d4-a716-446655440000"`, value)
		assertTypedWrapperJSONNull(t, NullUUID())
	})

	t.Run("decimal", func(t *testing.T) {
		t.Parallel()
		value := decimal.RequireFromString("123.45")
		assertTypedWrapperJSONRoundTrip(t, OfDecimal(value), `"123.45"`, value)
		assertTypedWrapperJSONNull(t, NullDecimal())
	})

	t.Run("time", func(t *testing.T) {
		t.Parallel()
		value := time.Date(2024, time.January, 2, 3, 4, 5, 123456789, time.UTC)

		data, err := json.Marshal(OfTime(value))
		require.NoError(t, err)
		assert.JSONEq(t, `"2024-01-02T03:04:05.123456789Z"`, string(data))

		var restored Time
		require.NoError(t, json.Unmarshal(data, &restored))
		assert.True(t, restored.IsPresent())
		assert.True(t, restored.MustGet().Equal(value))

		assertTypedWrapperJSONNull(t, NullTime())
	})
}
