package optional

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Map
// ---------------------------------------------------------------------------

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("present transforms value", func(t *testing.T) {
		t.Parallel()
		result := Map(Of(5), func(n int) string {
			return strconv.Itoa(n * 2)
		})
		assert.True(t, result.IsPresent())
		assert.Equal(t, "10", result.MustGet())
	})

	t.Run("null preserves null", func(t *testing.T) {
		t.Parallel()
		result := Map(Null[int](), func(n int) string {
			return strconv.Itoa(n)
		})
		assert.True(t, result.IsNull())
	})

	t.Run("missing preserves missing", func(t *testing.T) {
		t.Parallel()
		result := Map(Missing[int](), func(n int) string {
			return strconv.Itoa(n)
		})
		assert.True(t, result.IsMissing())
	})
}

// ---------------------------------------------------------------------------
// FlatMap
// ---------------------------------------------------------------------------

func TestFlatMap(t *testing.T) {
	t.Parallel()

	safeDiv := func(a int) Field[float64] {
		if a == 0 {
			return Null[float64]()
		}
		return Of(100.0 / float64(a))
	}

	t.Run("present with successful fn", func(t *testing.T) {
		t.Parallel()
		result := FlatMap(Of(4), safeDiv)
		assert.True(t, result.IsPresent())
		assert.Equal(t, 25.0, result.MustGet())
	})

	t.Run("present with fn returning null", func(t *testing.T) {
		t.Parallel()
		result := FlatMap(Of(0), safeDiv)
		assert.True(t, result.IsNull())
	})

	t.Run("null preserves null", func(t *testing.T) {
		t.Parallel()
		result := FlatMap(Null[int](), safeDiv)
		assert.True(t, result.IsNull())
	})

	t.Run("missing preserves missing", func(t *testing.T) {
		t.Parallel()
		result := FlatMap(Missing[int](), safeDiv)
		assert.True(t, result.IsMissing())
	})
}

// ---------------------------------------------------------------------------
// Equal
// ---------------------------------------------------------------------------

func TestEqual(t *testing.T) {
	t.Parallel()

	t.Run("both present same value", func(t *testing.T) {
		t.Parallel()
		assert.True(t, Equal(Of(42), Of(42)))
	})

	t.Run("both present different values", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Equal(Of(1), Of(2)))
	})

	t.Run("both null", func(t *testing.T) {
		t.Parallel()
		assert.True(t, Equal(Null[int](), Null[int]()))
	})

	t.Run("both missing", func(t *testing.T) {
		t.Parallel()
		assert.True(t, Equal(Missing[int](), Missing[int]()))
	})

	t.Run("present vs null", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Equal(Of(1), Null[int]()))
	})

	t.Run("present vs missing", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Equal(Of(1), Missing[int]()))
	})

	t.Run("null vs missing", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Equal(Null[int](), Missing[int]()))
	})

	t.Run("string equal", func(t *testing.T) {
		t.Parallel()
		assert.True(t, Equal(Of("hello"), Of("hello")))
	})

	t.Run("string not equal", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Equal(Of("hello"), Of("world")))
	})
}

// ---------------------------------------------------------------------------
// Or
// ---------------------------------------------------------------------------

func TestField_Or(t *testing.T) {
	t.Parallel()

	t.Run("present returns self", func(t *testing.T) {
		t.Parallel()
		result := Of(1).Or(Of(2))
		assert.Equal(t, 1, result.MustGet())
	})

	t.Run("null returns other", func(t *testing.T) {
		t.Parallel()
		result := Null[int]().Or(Of(2))
		assert.Equal(t, 2, result.MustGet())
	})

	t.Run("missing returns other", func(t *testing.T) {
		t.Parallel()
		result := Missing[int]().Or(Of(3))
		assert.Equal(t, 3, result.MustGet())
	})

	t.Run("null or null", func(t *testing.T) {
		t.Parallel()
		result := Null[int]().Or(Null[int]())
		assert.True(t, result.IsNull())
	})
}

// ---------------------------------------------------------------------------
// OrElse
// ---------------------------------------------------------------------------

func TestField_OrElse(t *testing.T) {
	t.Parallel()

	t.Run("present returns value", func(t *testing.T) {
		t.Parallel()
		result := Of(42).OrElse(func() int { return 99 })
		assert.Equal(t, 42, result)
	})

	t.Run("null calls fn", func(t *testing.T) {
		t.Parallel()
		called := false
		result := Null[int]().OrElse(func() int {
			called = true
			return 99
		})
		assert.True(t, called)
		assert.Equal(t, 99, result)
	})

	t.Run("missing calls fn", func(t *testing.T) {
		t.Parallel()
		result := Missing[int]().OrElse(func() int { return 77 })
		assert.Equal(t, 77, result)
	})
}

// ---------------------------------------------------------------------------
// IsZero
// ---------------------------------------------------------------------------

func TestField_IsZero(t *testing.T) {
	t.Parallel()

	t.Run("missing is zero", func(t *testing.T) {
		t.Parallel()
		assert.True(t, Missing[int]().IsZero())
	})

	t.Run("null is not zero", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Null[int]().IsZero())
	})

	t.Run("present is not zero", func(t *testing.T) {
		t.Parallel()
		assert.False(t, Of(0).IsZero())
	})

	t.Run("zero value of Field is zero", func(t *testing.T) {
		t.Parallel()
		var f Field[int]
		assert.True(t, f.IsZero())
	})
}

// ---------------------------------------------------------------------------
// String / GoString
// ---------------------------------------------------------------------------

func TestField_String(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "42", Of(42).String())
	assert.Equal(t, "hello", Of("hello").String())
	assert.Equal(t, "null", Null[int]().String())
	assert.Equal(t, "missing", Missing[int]().String())
}

func TestField_GoString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "optional.Of(42)", Of(42).GoString())
	assert.Equal(t, `optional.Of("hello")`, Of("hello").GoString())
	assert.Equal(t, "optional.Null()", Null[int]().GoString())
	assert.Equal(t, "optional.Missing()", Missing[int]().GoString())
}

// ---------------------------------------------------------------------------
// Ensure new methods satisfy interfaces
// ---------------------------------------------------------------------------

func TestField_ImplementsStringer(t *testing.T) {
	t.Parallel()

	var _ interface{ String() string } = Of(1)
	var _ interface{ GoString() string } = Of(1)

	// Verify via require to actually call the methods
	require.NotEmpty(t, Of(1).String())
	require.NotEmpty(t, Of(1).GoString())
}
