package optional

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// convertToBool
// ---------------------------------------------------------------------------

func TestConvertToBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		src     any
		want    bool
		wantErr bool
	}{
		{"bool true", true, true, false},
		{"bool false", false, false, false},
		{"int64 1", int64(1), true, false},
		{"int64 0", int64(0), false, false},
		{"int64 -1", int64(-1), true, false},
		{"float64 1.0", float64(1.0), true, false},
		{"float64 0.0", float64(0.0), false, false},
		{"string true", "true", true, false},
		{"string false", "false", false, false},
		{"string 1", "1", true, false},
		{"string 0", "0", false, false},
		{"string empty", "", false, false},
		{"string invalid", "maybe", false, true},
		{"[]byte true", []byte("true"), true, false},
		{"[]byte false", []byte("false"), false, false},
		{"[]byte empty", []byte{}, false, false},
		{"[]byte invalid", []byte("xyz"), false, true},
		{"unsupported type", int32(1), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := convertToBool(tt.src)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// convertToInt64
// ---------------------------------------------------------------------------

func TestConvertToInt64(t *testing.T) {
	t.Parallel()

	t.Run("int64 direct", func(t *testing.T) {
		t.Parallel()
		n, err := convertToInt64(int64(42), 64)
		require.NoError(t, err)
		assert.Equal(t, int64(42), n)
	})

	t.Run("float64 whole number", func(t *testing.T) {
		t.Parallel()
		n, err := convertToInt64(float64(100.0), 64)
		require.NoError(t, err)
		assert.Equal(t, int64(100), n)
	})

	t.Run("float64 non-whole number", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(float64(1.5), 64)
		assert.Error(t, err)
	})

	t.Run("float64 NaN", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(math.NaN(), 64)
		assert.Error(t, err)
	})

	t.Run("float64 +Inf", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(math.Inf(1), 64)
		assert.Error(t, err)
	})

	t.Run("float64 overflows int64", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(float64(math.MaxInt64)*2, 64)
		assert.Error(t, err)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		n, err := convertToInt64("123", 64)
		require.NoError(t, err)
		assert.Equal(t, int64(123), n)
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		n, err := convertToInt64([]byte("-456"), 64)
		require.NoError(t, err)
		assert.Equal(t, int64(-456), n)
	})

	t.Run("int8 overflow", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(int64(200), 8)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflows")
	})

	t.Run("int8 underflow", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(int64(-200), 8)
		assert.Error(t, err)
	})

	t.Run("int16 max", func(t *testing.T) {
		t.Parallel()
		n, err := convertToInt64(int64(math.MaxInt16), 16)
		require.NoError(t, err)
		assert.Equal(t, int64(math.MaxInt16), n)
	})

	t.Run("int16 overflow", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(int64(math.MaxInt16+1), 16)
		assert.Error(t, err)
	})

	t.Run("string overflow int8", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64("300", 8)
		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToInt64(true, 64)
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// convertToUint64
// ---------------------------------------------------------------------------

func TestConvertToUint64(t *testing.T) {
	t.Parallel()

	t.Run("int64 positive", func(t *testing.T) {
		t.Parallel()
		n, err := convertToUint64(int64(42), 64)
		require.NoError(t, err)
		assert.Equal(t, uint64(42), n)
	})

	t.Run("int64 negative", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(int64(-1), 64)
		assert.Error(t, err)
	})

	t.Run("float64 positive whole", func(t *testing.T) {
		t.Parallel()
		n, err := convertToUint64(float64(100.0), 64)
		require.NoError(t, err)
		assert.Equal(t, uint64(100), n)
	})

	t.Run("float64 negative", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(float64(-1.0), 64)
		assert.Error(t, err)
	})

	t.Run("float64 non-whole", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(float64(1.5), 64)
		assert.Error(t, err)
	})

	t.Run("float64 NaN", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(math.NaN(), 64)
		assert.Error(t, err)
	})

	t.Run("float64 +Inf", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(math.Inf(1), 64)
		assert.Error(t, err)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		n, err := convertToUint64("255", 8)
		require.NoError(t, err)
		assert.Equal(t, uint64(255), n)
	})

	t.Run("string overflow uint8", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64("256", 8)
		assert.Error(t, err)
	})

	t.Run("uint16 overflow from int64", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(int64(70000), 16)
		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToUint64(true, 64)
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// convertToFloat64
// ---------------------------------------------------------------------------

func TestConvertToFloat64(t *testing.T) {
	t.Parallel()

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		n, err := convertToFloat64(float64(3.14))
		require.NoError(t, err)
		assert.InDelta(t, 3.14, n, 0.001)
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		n, err := convertToFloat64(int64(42))
		require.NoError(t, err)
		assert.Equal(t, float64(42), n)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		n, err := convertToFloat64("2.718")
		require.NoError(t, err)
		assert.InDelta(t, 2.718, n, 0.001)
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		n, err := convertToFloat64([]byte("1.5"))
		require.NoError(t, err)
		assert.Equal(t, 1.5, n)
	})

	t.Run("invalid string", func(t *testing.T) {
		t.Parallel()
		_, err := convertToFloat64("abc")
		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToFloat64(true)
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// convertToString
// ---------------------------------------------------------------------------

func TestConvertToString(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		s, err := convertToString("hello")
		require.NoError(t, err)
		assert.Equal(t, "hello", s)
	})

	t.Run("[]byte", func(t *testing.T) {
		t.Parallel()
		s, err := convertToString([]byte("world"))
		require.NoError(t, err)
		assert.Equal(t, "world", s)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToString(42)
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// convertToBytes
// ---------------------------------------------------------------------------

func TestConvertToBytes(t *testing.T) {
	t.Parallel()

	t.Run("[]byte copies data", func(t *testing.T) {
		t.Parallel()
		src := []byte("hello")
		got, err := convertToBytes(src)
		require.NoError(t, err)
		assert.Equal(t, []byte("hello"), got)
		// Ensure it's a copy.
		src[0] = 'X'
		assert.Equal(t, byte('h'), got[0], "should be a copy, not aliased")
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		got, err := convertToBytes("world")
		require.NoError(t, err)
		assert.Equal(t, []byte("world"), got)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToBytes(42)
		assert.Error(t, err)
	})
}

// ---------------------------------------------------------------------------
// convertToTime
// ---------------------------------------------------------------------------

func TestConvertToTime(t *testing.T) {
	t.Parallel()

	want := time.Date(2024, 6, 15, 12, 30, 0, 0, time.UTC)

	t.Run("time.Time", func(t *testing.T) {
		t.Parallel()
		got, err := convertToTime(want)
		require.NoError(t, err)
		assert.True(t, got.Equal(want))
	})

	t.Run("RFC3339 string", func(t *testing.T) {
		t.Parallel()
		got, err := convertToTime("2024-06-15T12:30:00Z")
		require.NoError(t, err)
		assert.True(t, got.Equal(want))
	})

	t.Run("datetime string", func(t *testing.T) {
		t.Parallel()
		got, err := convertToTime("2024-06-15 12:30:00")
		require.NoError(t, err)
		assert.Equal(t, 2024, got.Year())
		assert.Equal(t, time.June, got.Month())
	})

	t.Run("date only string", func(t *testing.T) {
		t.Parallel()
		got, err := convertToTime("2024-06-15")
		require.NoError(t, err)
		assert.Equal(t, 2024, got.Year())
	})

	t.Run("[]byte RFC3339", func(t *testing.T) {
		t.Parallel()
		got, err := convertToTime([]byte("2024-06-15T12:30:00Z"))
		require.NoError(t, err)
		assert.True(t, got.Equal(want))
	})

	t.Run("invalid string", func(t *testing.T) {
		t.Parallel()
		_, err := convertToTime("not-a-time")
		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		_, err := convertToTime(42)
		assert.Error(t, err)
	})
}
