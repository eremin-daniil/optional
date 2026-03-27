package optional

import (
	"database/sql/driver"
	"math"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Bool Scan
// ---------------------------------------------------------------------------

func TestBool_Scan(t *testing.T) {
	t.Parallel()

	t.Run("nil sets null", func(t *testing.T) {
		t.Parallel()
		var b Bool
		require.NoError(t, b.Scan(nil))
		assert.True(t, b.IsNull())
	})

	t.Run("bool true", func(t *testing.T) {
		t.Parallel()
		var b Bool
		require.NoError(t, b.Scan(true))
		assert.True(t, b.IsPresent())
		assert.True(t, b.MustGet())
	})

	t.Run("bool false", func(t *testing.T) {
		t.Parallel()
		var b Bool
		require.NoError(t, b.Scan(false))
		assert.True(t, b.IsPresent())
		assert.False(t, b.MustGet())
	})

	t.Run("int64 1", func(t *testing.T) {
		t.Parallel()
		var b Bool
		require.NoError(t, b.Scan(int64(1)))
		assert.True(t, b.MustGet())
	})

	t.Run("int64 0", func(t *testing.T) {
		t.Parallel()
		var b Bool
		require.NoError(t, b.Scan(int64(0)))
		assert.False(t, b.MustGet())
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var b Bool
		assert.Error(t, b.Scan(struct{}{}))
	})
}

// ---------------------------------------------------------------------------
// Int Scan + Value
// ---------------------------------------------------------------------------

func TestInt_Scan(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var i Int
		require.NoError(t, i.Scan(nil))
		assert.True(t, i.IsNull())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		var i Int
		require.NoError(t, i.Scan(int64(42)))
		assert.Equal(t, 42, i.MustGet())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var i Int
		require.NoError(t, i.Scan("99"))
		assert.Equal(t, 99, i.MustGet())
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var i Int
		assert.Error(t, i.Scan(true))
	})
}

func TestInt_Value(t *testing.T) {
	t.Parallel()

	t.Run("present returns int64", func(t *testing.T) {
		t.Parallel()
		v, err := OfInt(42).Value()
		require.NoError(t, err)
		assert.Equal(t, int64(42), v)
	})

	t.Run("null returns nil", func(t *testing.T) {
		t.Parallel()
		v, err := NullInt().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("missing returns error", func(t *testing.T) {
		t.Parallel()
		_, err := MissingInt().Value()
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// Int8 Scan + Value
// ---------------------------------------------------------------------------

func TestInt8_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var i Int8
		require.NoError(t, i.Scan(int64(127)))
		assert.Equal(t, int8(127), i.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var i Int8
		assert.Error(t, i.Scan(int64(200)))
	})

	t.Run("underflow", func(t *testing.T) {
		t.Parallel()
		var i Int8
		assert.Error(t, i.Scan(int64(-200)))
	})
}

func TestInt8_Value(t *testing.T) {
	t.Parallel()
	v, err := OfInt8(100).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(100), v)
}

// ---------------------------------------------------------------------------
// Int16 Scan + Value
// ---------------------------------------------------------------------------

func TestInt16_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var i Int16
		require.NoError(t, i.Scan(int64(32000)))
		assert.Equal(t, int16(32000), i.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var i Int16
		assert.Error(t, i.Scan(int64(40000)))
	})
}

func TestInt16_Value(t *testing.T) {
	t.Parallel()
	v, err := OfInt16(1000).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(1000), v)
}

// ---------------------------------------------------------------------------
// Int32 Scan + Value
// ---------------------------------------------------------------------------

func TestInt32_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var i Int32
		require.NoError(t, i.Scan(int64(2_000_000)))
		assert.Equal(t, int32(2_000_000), i.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var i Int32
		assert.Error(t, i.Scan(int64(math.MaxInt32+1)))
	})
}

func TestInt32_Value(t *testing.T) {
	t.Parallel()
	v, err := OfInt32(12345).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(12345), v)
}

// ---------------------------------------------------------------------------
// Int64 Scan
// ---------------------------------------------------------------------------

func TestInt64_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var i Int64
		require.NoError(t, i.Scan(int64(9_000_000_000)))
		assert.Equal(t, int64(9_000_000_000), i.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var i Int64
		require.NoError(t, i.Scan(nil))
		assert.True(t, i.IsNull())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var i Int64
		require.NoError(t, i.Scan("12345"))
		assert.Equal(t, int64(12345), i.MustGet())
	})
}

// ---------------------------------------------------------------------------
// Uint Scan + Value
// ---------------------------------------------------------------------------

func TestUint_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uint
		require.NoError(t, u.Scan(int64(42)))
		assert.Equal(t, uint(42), u.MustGet())
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		var u Uint
		assert.Error(t, u.Scan(int64(-1)))
	})
}

func TestUint_Value(t *testing.T) {
	t.Parallel()

	t.Run("present", func(t *testing.T) {
		t.Parallel()
		v, err := OfUint(100).Value()
		require.NoError(t, err)
		assert.Equal(t, int64(100), v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})
}

// ---------------------------------------------------------------------------
// Uint8 Scan + Value
// ---------------------------------------------------------------------------

func TestUint8_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uint8
		require.NoError(t, u.Scan(int64(255)))
		assert.Equal(t, uint8(255), u.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var u Uint8
		assert.Error(t, u.Scan(int64(300)))
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		var u Uint8
		assert.Error(t, u.Scan(int64(-1)))
	})
}

func TestUint8_Value(t *testing.T) {
	t.Parallel()
	v, err := OfUint8(200).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(200), v)
}

// ---------------------------------------------------------------------------
// Uint16 Scan + Value
// ---------------------------------------------------------------------------

func TestUint16_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uint16
		require.NoError(t, u.Scan(int64(65535)))
		assert.Equal(t, uint16(65535), u.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var u Uint16
		assert.Error(t, u.Scan(int64(70000)))
	})
}

func TestUint16_Value(t *testing.T) {
	t.Parallel()
	v, err := OfUint16(50000).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(50000), v)
}

// ---------------------------------------------------------------------------
// Uint32 Scan + Value
// ---------------------------------------------------------------------------

func TestUint32_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uint32
		require.NoError(t, u.Scan(int64(4_000_000_000)))
		assert.Equal(t, uint32(4_000_000_000), u.MustGet())
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var u Uint32
		assert.Error(t, u.Scan(int64(5_000_000_000)))
	})
}

func TestUint32_Value(t *testing.T) {
	t.Parallel()
	v, err := OfUint32(3_000_000_000).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(3_000_000_000), v)
}

// ---------------------------------------------------------------------------
// Uint64 Scan + Value
// ---------------------------------------------------------------------------

func TestUint64_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uint64
		require.NoError(t, u.Scan(int64(42)))
		assert.Equal(t, uint64(42), u.MustGet())
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		var u Uint64
		assert.Error(t, u.Scan(int64(-1)))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var u Uint64
		require.NoError(t, u.Scan("18446744073709551615"))
		assert.Equal(t, uint64(math.MaxUint64), u.MustGet())
	})
}

func TestUint64_Value(t *testing.T) {
	t.Parallel()

	t.Run("fits in int64", func(t *testing.T) {
		t.Parallel()
		v, err := OfUint64(42).Value()
		require.NoError(t, err)
		assert.Equal(t, int64(42), v)
	})

	t.Run("overflows int64", func(t *testing.T) {
		t.Parallel()
		_, err := OfUint64(math.MaxUint64).Value()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflows int64")
	})
}

// ---------------------------------------------------------------------------
// Uintptr Scan + Value
// ---------------------------------------------------------------------------

func TestUintptr_Scan(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var u Uintptr
		require.NoError(t, u.Scan(int64(123)))
		assert.Equal(t, uintptr(123), u.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var u Uintptr
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})
}

func TestUintptr_Value(t *testing.T) {
	t.Parallel()
	v, err := OfUintptr(999).Value()
	require.NoError(t, err)
	assert.Equal(t, int64(999), v)
}

// ---------------------------------------------------------------------------
// Float32 Scan + Value
// ---------------------------------------------------------------------------

func TestFloat32_Scan(t *testing.T) {
	t.Parallel()

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		var f Float32
		require.NoError(t, f.Scan(float64(3.25)))
		assert.InDelta(t, float32(3.25), f.MustGet(), 0.001)
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()
		var f Float32
		assert.Error(t, f.Scan(float64(math.MaxFloat64)))
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var f Float32
		require.NoError(t, f.Scan(nil))
		assert.True(t, f.IsNull())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var f Float32
		require.NoError(t, f.Scan("1.5"))
		assert.InDelta(t, float32(1.5), f.MustGet(), 0.001)
	})
}

func TestFloat32_Value(t *testing.T) {
	t.Parallel()
	v, err := OfFloat32(3.25).Value()
	require.NoError(t, err)
	assert.Equal(t, float64(3.25), v)
}

// ---------------------------------------------------------------------------
// Float64 Scan
// ---------------------------------------------------------------------------

func TestFloat64_Scan(t *testing.T) {
	t.Parallel()

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		var f Float64
		require.NoError(t, f.Scan(float64(6.5)))
		assert.Equal(t, 6.5, f.MustGet())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		var f Float64
		require.NoError(t, f.Scan(int64(42)))
		assert.Equal(t, float64(42), f.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var f Float64
		require.NoError(t, f.Scan(nil))
		assert.True(t, f.IsNull())
	})
}

// ---------------------------------------------------------------------------
// String Scan
// ---------------------------------------------------------------------------

func TestString_Scan(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var s String
		require.NoError(t, s.Scan("hello"))
		assert.Equal(t, "hello", s.MustGet())
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		var s String
		require.NoError(t, s.Scan([]byte("world")))
		assert.Equal(t, "world", s.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var s String
		require.NoError(t, s.Scan(nil))
		assert.True(t, s.IsNull())
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var s String
		assert.Error(t, s.Scan(42))
	})
}

// ---------------------------------------------------------------------------
// Bytes Scan
// ---------------------------------------------------------------------------

func TestBytes_Scan(t *testing.T) {
	t.Parallel()

	t.Run("[]byte copies data", func(t *testing.T) {
		t.Parallel()
		var b Bytes
		src := []byte("hello")
		require.NoError(t, b.Scan(src))
		assert.Equal(t, []byte("hello"), b.MustGet())
		// Mutate source to verify copy.
		src[0] = 'X'
		assert.Equal(t, byte('h'), b.MustGet()[0])
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var b Bytes
		require.NoError(t, b.Scan("data"))
		assert.Equal(t, []byte("data"), b.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var b Bytes
		require.NoError(t, b.Scan(nil))
		assert.True(t, b.IsNull())
	})
}

// ---------------------------------------------------------------------------
// Time Scan
// ---------------------------------------------------------------------------

func TestTime_Scan(t *testing.T) {
	t.Parallel()

	want := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)

	t.Run("time.Time", func(t *testing.T) {
		t.Parallel()
		var tm Time
		require.NoError(t, tm.Scan(want))
		assert.True(t, tm.MustGet().Equal(want))
	})

	t.Run("string RFC3339", func(t *testing.T) {
		t.Parallel()
		var tm Time
		require.NoError(t, tm.Scan("2024-06-15T12:00:00Z"))
		assert.True(t, tm.MustGet().Equal(want))
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		var tm Time
		require.NoError(t, tm.Scan([]byte("2024-06-15T12:00:00Z")))
		assert.True(t, tm.MustGet().Equal(want))
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var tm Time
		require.NoError(t, tm.Scan(nil))
		assert.True(t, tm.IsNull())
	})

	t.Run("invalid string", func(t *testing.T) {
		t.Parallel()
		var tm Time
		assert.Error(t, tm.Scan("not-a-time"))
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var tm Time
		assert.Error(t, tm.Scan(42))
	})
}

// ---------------------------------------------------------------------------
// UUID Scan + Value
// ---------------------------------------------------------------------------

func TestUUID_Scan(t *testing.T) {
	t.Parallel()

	want := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var u UUID
		require.NoError(t, u.Scan("550e8400-e29b-41d4-a716-446655440000"))
		assert.Equal(t, want, u.MustGet())
	})

	t.Run("uuid.UUID", func(t *testing.T) {
		t.Parallel()
		var u UUID
		require.NoError(t, u.Scan(want))
		assert.Equal(t, want, u.MustGet())
	})

	t.Run("[]byte string format", func(t *testing.T) {
		t.Parallel()
		var u UUID
		require.NoError(t, u.Scan([]byte("550e8400-e29b-41d4-a716-446655440000")))
		assert.Equal(t, want, u.MustGet())
	})

	t.Run("[]byte 16 raw bytes", func(t *testing.T) {
		t.Parallel()
		var u UUID
		require.NoError(t, u.Scan(want[:]))
		assert.Equal(t, want, u.MustGet())
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var u UUID
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})

	t.Run("invalid string", func(t *testing.T) {
		t.Parallel()
		var u UUID
		assert.Error(t, u.Scan("not-a-uuid"))
	})

	t.Run("invalid []byte length", func(t *testing.T) {
		t.Parallel()
		var u UUID
		assert.Error(t, u.Scan([]byte{1, 2, 3}))
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var u UUID
		assert.Error(t, u.Scan(42))
	})
}

func TestUUID_Value(t *testing.T) {
	t.Parallel()

	t.Run("present returns string", func(t *testing.T) {
		t.Parallel()
		id := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		v, err := OfUUID(id).Value()
		require.NoError(t, err)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", v)
	})

	t.Run("null returns nil", func(t *testing.T) {
		t.Parallel()
		v, err := NullUUID().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})
}

// ---------------------------------------------------------------------------
// Decimal Scan + Value
// ---------------------------------------------------------------------------

func TestDecimal_Scan(t *testing.T) {
	t.Parallel()

	want := decimal.RequireFromString("123.45")

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan("123.45"))
		assert.True(t, want.Equal(d.MustGet()))
	})

	t.Run("decimal.Decimal", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan(want))
		assert.True(t, want.Equal(d.MustGet()))
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan([]byte("123.45")))
		assert.True(t, want.Equal(d.MustGet()))
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan(float64(99.5)))
		assert.True(t, d.IsPresent())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan(int64(42)))
		assert.True(t, decimal.NewFromInt(42).Equal(d.MustGet()))
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		require.NoError(t, d.Scan(nil))
		assert.True(t, d.IsNull())
	})

	t.Run("invalid string", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		assert.Error(t, d.Scan("abc"))
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var d Decimal
		assert.Error(t, d.Scan(true))
	})
}

func TestDecimal_Value(t *testing.T) {
	t.Parallel()

	t.Run("present returns string", func(t *testing.T) {
		t.Parallel()
		v, err := OfDecimal(decimal.RequireFromString("42.5")).Value()
		require.NoError(t, err)
		assert.Equal(t, "42.5", v)
	})

	t.Run("null returns nil", func(t *testing.T) {
		t.Parallel()
		v, err := NullDecimal().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})
}

// ---------------------------------------------------------------------------
// Verify all scalar types implement sql.Scanner and driver.Valuer
// ---------------------------------------------------------------------------

func TestScalar_ImplementsInterfaces(t *testing.T) {
	t.Parallel()

	// sql.Scanner (pointer receiver)
	var _ interface{ Scan(any) error } = (*Bool)(nil)
	var _ interface{ Scan(any) error } = (*Int)(nil)
	var _ interface{ Scan(any) error } = (*Int8)(nil)
	var _ interface{ Scan(any) error } = (*Int16)(nil)
	var _ interface{ Scan(any) error } = (*Int32)(nil)
	var _ interface{ Scan(any) error } = (*Int64)(nil)
	var _ interface{ Scan(any) error } = (*Uint)(nil)
	var _ interface{ Scan(any) error } = (*Uint8)(nil)
	var _ interface{ Scan(any) error } = (*Uint16)(nil)
	var _ interface{ Scan(any) error } = (*Uint32)(nil)
	var _ interface{ Scan(any) error } = (*Uint64)(nil)
	var _ interface{ Scan(any) error } = (*Uintptr)(nil)
	var _ interface{ Scan(any) error } = (*Float32)(nil)
	var _ interface{ Scan(any) error } = (*Float64)(nil)
	var _ interface{ Scan(any) error } = (*String)(nil)
	var _ interface{ Scan(any) error } = (*Bytes)(nil)
	var _ interface{ Scan(any) error } = (*Time)(nil)
	var _ interface{ Scan(any) error } = (*UUID)(nil)
	var _ interface{ Scan(any) error } = (*Decimal)(nil)

	// driver.Valuer (value receiver)
	var _ driver.Valuer = Bool{}
	var _ driver.Valuer = Int{}
	var _ driver.Valuer = Int8{}
	var _ driver.Valuer = Int16{}
	var _ driver.Valuer = Int32{}
	var _ driver.Valuer = Int64{}
	var _ driver.Valuer = Uint{}
	var _ driver.Valuer = Uint8{}
	var _ driver.Valuer = Uint16{}
	var _ driver.Valuer = Uint32{}
	var _ driver.Valuer = Uint64{}
	var _ driver.Valuer = Uintptr{}
	var _ driver.Valuer = Float32{}
	var _ driver.Valuer = Float64{}
	var _ driver.Valuer = String{}
	var _ driver.Valuer = Bytes{}
	var _ driver.Valuer = Time{}
	var _ driver.Valuer = UUID{}
	var _ driver.Valuer = Decimal{}
}

// ---------------------------------------------------------------------------
// Bytes constructors
// ---------------------------------------------------------------------------

func TestBytes_Constructors(t *testing.T) {
	t.Parallel()

	t.Run("OfBytes", func(t *testing.T) {
		t.Parallel()
		b := OfBytes([]byte("hello"))
		assert.True(t, b.IsPresent())
		assert.Equal(t, []byte("hello"), b.MustGet())
	})

	t.Run("FromBytesPtr non-nil", func(t *testing.T) {
		t.Parallel()
		data := []byte("world")
		b := FromBytesPtr(&data)
		assert.True(t, b.IsPresent())
		assert.Equal(t, []byte("world"), b.MustGet())
	})

	t.Run("FromBytesPtr nil", func(t *testing.T) {
		t.Parallel()
		b := FromBytesPtr(nil)
		assert.True(t, b.IsNull())
	})

	t.Run("NullBytes", func(t *testing.T) {
		t.Parallel()
		b := NullBytes()
		assert.True(t, b.IsNull())
	})

	t.Run("MissingBytes", func(t *testing.T) {
		t.Parallel()
		b := MissingBytes()
		assert.True(t, b.IsMissing())
	})
}

func TestScalar_OfNullableAliases(t *testing.T) {
	t.Parallel()

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v := true
		assert.Equal(t, FromBoolPtr(&v), OfNullableBool(&v))
		assert.Equal(t, FromBoolPtr(nil), OfNullableBool(nil))
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v := 42
		assert.Equal(t, FromIntPtr(&v), OfNullableInt(&v))
		assert.Equal(t, FromIntPtr(nil), OfNullableInt(nil))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v := "hello"
		assert.Equal(t, FromStringPtr(&v), OfNullableString(&v))
		assert.Equal(t, FromStringPtr(nil), OfNullableString(nil))
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		v := []byte("hello")
		assert.Equal(t, FromBytesPtr(&v), OfNullableBytes(&v))
		assert.Equal(t, FromBytesPtr(nil), OfNullableBytes(nil))
	})

	t.Run("uuid", func(t *testing.T) {
		t.Parallel()
		v := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		assert.Equal(t, FromUUIDPtr(&v), OfNullableUUID(&v))
		assert.Equal(t, FromUUIDPtr(nil), OfNullableUUID(nil))
	})

	t.Run("decimal", func(t *testing.T) {
		t.Parallel()
		v := decimal.RequireFromString("123.45")
		assert.Equal(t, FromDecimalPtr(&v), OfNullableDecimal(&v))
		assert.Equal(t, FromDecimalPtr(nil), OfNullableDecimal(nil))
	})
}

// ---------------------------------------------------------------------------
// Value() null/missing paths for all overridden types
// ---------------------------------------------------------------------------

func TestScalar_Value_NullAndMissing(t *testing.T) {
	t.Parallel()

	t.Run("Int8 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullInt8().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Int8 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingInt8().Value()
		assert.Error(t, err)
	})

	t.Run("Int16 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullInt16().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Int16 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingInt16().Value()
		assert.Error(t, err)
	})

	t.Run("Int32 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullInt32().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Int32 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingInt32().Value()
		assert.Error(t, err)
	})

	t.Run("Uint null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uint missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUint().Value()
		assert.Error(t, err)
	})

	t.Run("Uint8 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint8().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uint8 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUint8().Value()
		assert.Error(t, err)
	})

	t.Run("Uint16 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint16().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uint16 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUint16().Value()
		assert.Error(t, err)
	})

	t.Run("Uint32 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint32().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uint32 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUint32().Value()
		assert.Error(t, err)
	})

	t.Run("Uint64 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUint64().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uint64 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUint64().Value()
		assert.Error(t, err)
	})

	t.Run("Uintptr null", func(t *testing.T) {
		t.Parallel()
		v, err := NullUintptr().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Uintptr missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUintptr().Value()
		assert.Error(t, err)
	})

	t.Run("Float32 null", func(t *testing.T) {
		t.Parallel()
		v, err := NullFloat32().Value()
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("Float32 missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingFloat32().Value()
		assert.Error(t, err)
	})

	t.Run("UUID missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingUUID().Value()
		assert.Error(t, err)
	})

	t.Run("Decimal missing", func(t *testing.T) {
		t.Parallel()
		_, err := MissingDecimal().Value()
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// Additional Scan nil paths for types not covered above
// ---------------------------------------------------------------------------

func TestScalar_Scan_Nil(t *testing.T) {
	t.Parallel()

	t.Run("Int8 nil", func(t *testing.T) {
		t.Parallel()
		var i Int8
		require.NoError(t, i.Scan(nil))
		assert.True(t, i.IsNull())
	})

	t.Run("Int16 nil", func(t *testing.T) {
		t.Parallel()
		var i Int16
		require.NoError(t, i.Scan(nil))
		assert.True(t, i.IsNull())
	})

	t.Run("Int32 nil", func(t *testing.T) {
		t.Parallel()
		var i Int32
		require.NoError(t, i.Scan(nil))
		assert.True(t, i.IsNull())
	})

	t.Run("Uint nil", func(t *testing.T) {
		t.Parallel()
		var u Uint
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})

	t.Run("Uint8 nil", func(t *testing.T) {
		t.Parallel()
		var u Uint8
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})

	t.Run("Uint16 nil", func(t *testing.T) {
		t.Parallel()
		var u Uint16
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})

	t.Run("Uint32 nil", func(t *testing.T) {
		t.Parallel()
		var u Uint32
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})

	t.Run("Uint64 nil", func(t *testing.T) {
		t.Parallel()
		var u Uint64
		require.NoError(t, u.Scan(nil))
		assert.True(t, u.IsNull())
	})
}

// ---------------------------------------------------------------------------
// Scan error paths for types not covered above
// ---------------------------------------------------------------------------

func TestScalar_Scan_Errors(t *testing.T) {
	t.Parallel()

	t.Run("Int8 unsupported type", func(t *testing.T) {
		t.Parallel()
		var i Int8
		assert.Error(t, i.Scan(true))
	})

	t.Run("Int16 unsupported type", func(t *testing.T) {
		t.Parallel()
		var i Int16
		assert.Error(t, i.Scan(true))
	})

	t.Run("Int32 unsupported type", func(t *testing.T) {
		t.Parallel()
		var i Int32
		assert.Error(t, i.Scan(true))
	})

	t.Run("Int64 unsupported type", func(t *testing.T) {
		t.Parallel()
		var i Int64
		assert.Error(t, i.Scan(true))
	})

	t.Run("Uint unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uint
		assert.Error(t, u.Scan(true))
	})

	t.Run("Uint8 unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uint8
		assert.Error(t, u.Scan(true))
	})

	t.Run("Uint16 unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uint16
		assert.Error(t, u.Scan(true))
	})

	t.Run("Uint32 unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uint32
		assert.Error(t, u.Scan(true))
	})

	t.Run("Uint64 unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uint64
		assert.Error(t, u.Scan(true))
	})

	t.Run("Uintptr unsupported type", func(t *testing.T) {
		t.Parallel()
		var u Uintptr
		assert.Error(t, u.Scan(true))
	})

	t.Run("Float32 unsupported type", func(t *testing.T) {
		t.Parallel()
		var f Float32
		assert.Error(t, f.Scan(true))
	})

	t.Run("Float64 unsupported type", func(t *testing.T) {
		t.Parallel()
		var f Float64
		assert.Error(t, f.Scan(true))
	})

	t.Run("Bytes unsupported type", func(t *testing.T) {
		t.Parallel()
		var b Bytes
		assert.Error(t, b.Scan(42))
	})
}

// ---------------------------------------------------------------------------
// Edge cases for Uint/Uintptr Value overflow and converter edge cases
// ---------------------------------------------------------------------------

func TestUint_Value_Overflow(t *testing.T) {
	t.Parallel()

	// On 64-bit platforms, this test verifies the overflow check.
	// On 32-bit platforms, uint can't exceed MaxInt64 so this is a no-op.
	if ^uint(0) > uint(math.MaxInt64) {
		_, err := OfUint(uint(math.MaxInt64) + 1).Value()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflows int64")
	}
}

func TestUintptr_Value_Overflow(t *testing.T) {
	t.Parallel()

	if ^uintptr(0) > uintptr(math.MaxInt64) {
		_, err := OfUintptr(uintptr(math.MaxInt64) + 1).Value()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflows int64")
	}
}

func TestConvertToUint64_EdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("[]byte valid", func(t *testing.T) {
		t.Parallel()
		n, err := convertToUint64([]byte("42"), 64)
		require.NoError(t, err)
		assert.Equal(t, uint64(42), n)
	})

	t.Run("[]byte overflow", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64([]byte("300"), 8)
		assert.Error(t, err)
	})

	t.Run("int64 overflow for uint8", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(int64(300), 8)
		assert.Error(t, err)
	})
}

func TestConvertToFloat64_ByteSlice_Invalid(t *testing.T) {
	t.Parallel()
	_, err := convertToFloat64([]byte("abc"))
	assert.Error(t, err)
}

func TestDecimal_Scan_InvalidBytes(t *testing.T) {
	t.Parallel()
	var d Decimal
	assert.Error(t, d.Scan([]byte("not-a-decimal")))
}
