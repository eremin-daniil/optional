package optional

import (
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOf(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		f := Of(42)
		assert.True(t, f.IsPresent())
		assert.False(t, f.IsNull())
		assert.False(t, f.IsMissing())
		v, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, 42, v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		f := Of("hello")
		v, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, "hello", v)
	})

	t.Run("zero value is present", func(t *testing.T) {
		t.Parallel()
		f := Of(0)
		assert.True(t, f.IsPresent())
		v, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, 0, v)
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type user struct {
			Name string
			Age  int
		}
		u := user{Name: "Alice", Age: 30}
		f := Of(u)
		assert.True(t, f.IsPresent())
		v, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, u, v)
	})
}

func TestFromPtr(t *testing.T) {
	t.Parallel()

	t.Run("non-nil pointer", func(t *testing.T) {
		t.Parallel()
		v := 42
		f := FromPtr(&v)
		assert.True(t, f.IsPresent())
		assert.False(t, f.IsNull())
		got, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, 42, got)
	})

	t.Run("nil pointer", func(t *testing.T) {
		t.Parallel()
		f := FromPtr[int](nil)
		assert.True(t, f.IsNull())
		assert.False(t, f.IsPresent())
		assert.False(t, f.IsMissing())
	})

	t.Run("does not alias source pointer", func(t *testing.T) {
		t.Parallel()
		v := 1
		f := FromPtr(&v)
		v = 2
		got, _ := f.Get()
		assert.Equal(t, 1, got, "value should be copied, not aliased")
	})
}

func TestOfNullable(t *testing.T) {
	t.Parallel()

	t.Run("non-nil pointer", func(t *testing.T) {
		t.Parallel()
		v := 42
		f := OfNullable(&v)
		assert.True(t, f.IsPresent())
		assert.Equal(t, 42, f.MustGet())
	})

	t.Run("nil pointer", func(t *testing.T) {
		t.Parallel()
		f := OfNullable[int](nil)
		assert.True(t, f.IsNull())
	})
}

func TestNull(t *testing.T) {
	t.Parallel()

	f := Null[string]()
	assert.True(t, f.IsNull())
	assert.False(t, f.IsPresent())
	assert.False(t, f.IsMissing())
}

func TestMissing(t *testing.T) {
	t.Parallel()

	f := Missing[string]()
	assert.True(t, f.IsMissing())
	assert.False(t, f.IsPresent())
	assert.False(t, f.IsNull())
}

func TestField_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		field  Field[int]
		wantV  int
		wantOK bool
	}{
		{"present", Of(5), 5, true},
		{"null", Null[int](), 0, false},
		{"missing", Missing[int](), 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v, ok := tt.field.Get()
			assert.Equal(t, tt.wantOK, ok)
			assert.Equal(t, tt.wantV, v)
		})
	}
}

func TestField_GetOr(t *testing.T) {
	t.Parallel()

	t.Run("present returns value", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 5, Of(5).GetOr(99))
	})

	t.Run("null returns default", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 99, Null[int]().GetOr(99))
	})

	t.Run("missing returns default", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 99, Missing[int]().GetOr(99))
	})
}

func TestField_Ptr(t *testing.T) {
	t.Parallel()

	t.Run("present returns pointer to value", func(t *testing.T) {
		t.Parallel()
		ptr := Of(7).Ptr()
		require.NotNil(t, ptr)
		assert.Equal(t, 7, *ptr)
	})

	t.Run("null returns nil", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, Null[int]().Ptr())
	})

	t.Run("missing returns nil", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, Missing[int]().Ptr())
	})
}

func TestField_MustGet(t *testing.T) {
	t.Parallel()

	t.Run("present returns value", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 42, Of(42).MustGet())
	})

	t.Run("null panics", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(t, "optional: value is not set", func() {
			Null[int]().MustGet()
		})
	})

	t.Run("missing panics", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(t, "optional: value is not set", func() {
			Missing[int]().MustGet()
		})
	})
}

func TestField_StateChecks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		field   Field[int]
		present bool
		null    bool
		missing bool
	}{
		{"present", Of(1), true, false, false},
		{"null", Null[int](), false, true, false},
		{"missing", Missing[int](), false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.present, tt.field.IsPresent())
			assert.Equal(t, tt.null, tt.field.IsNull())
			assert.Equal(t, tt.missing, tt.field.IsMissing())
		})
	}
}

func TestField_MarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("present int", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Of(42))
		require.NoError(t, err)
		assert.JSONEq(t, "42", string(data))
	})

	t.Run("present string", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Of("hello"))
		require.NoError(t, err)
		assert.JSONEq(t, `"hello"`, string(data))
	})

	t.Run("present struct", func(t *testing.T) {
		t.Parallel()
		type payload struct {
			X int `json:"x"`
		}
		data, err := json.Marshal(Of(payload{X: 1}))
		require.NoError(t, err)
		assert.JSONEq(t, `{"x":1}`, string(data))
	})

	t.Run("present uuid", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Of(uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")))
		require.NoError(t, err)
		assert.JSONEq(t, `"550e8400-e29b-41d4-a716-446655440000"`, string(data))
	})

	t.Run("present decimal", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Of(decimal.RequireFromString("123.45")))
		require.NoError(t, err)
		assert.JSONEq(t, `"123.45"`, string(data))
	})

	t.Run("present time", func(t *testing.T) {
		t.Parallel()
		value := time.Date(2024, time.January, 2, 3, 4, 5, 123456789, time.UTC)
		data, err := json.Marshal(Of(value))
		require.NoError(t, err)
		assert.JSONEq(t, `"2024-01-02T03:04:05.123456789Z"`, string(data))
	})

	t.Run("null produces null", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Null[int]())
		require.NoError(t, err)
		assert.Equal(t, "null", string(data))
	})

	t.Run("missing produces null", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(Missing[int]())
		require.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("unknown state returns error", func(t *testing.T) {
		t.Parallel()
		f := Field[int]{state: state(255)}
		_, err := f.MarshalJSON()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown state")
	})
}

func TestField_MarshalJSON_InStruct(t *testing.T) {
	t.Parallel()

	type request struct {
		Name Field[string] `json:"name"`
		Age  Field[int]    `json:"age"`
	}

	t.Run("all present", func(t *testing.T) {
		t.Parallel()
		r := request{Name: Of("Alice"), Age: Of(30)}
		data, err := json.Marshal(r)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name":"Alice","age":30}`, string(data))
	})

	t.Run("mixed states", func(t *testing.T) {
		t.Parallel()
		r := request{Name: Of("Bob"), Age: Null[int]()}
		data, err := json.Marshal(r)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name":"Bob","age":null}`, string(data))
	})
}

func TestField_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("present int", func(t *testing.T) {
		t.Parallel()
		var f Field[int]
		require.NoError(t, json.Unmarshal([]byte("42"), &f))
		assert.True(t, f.IsPresent())
		v, ok := f.Get()
		assert.True(t, ok)
		assert.Equal(t, 42, v)
	})

	t.Run("present string", func(t *testing.T) {
		t.Parallel()
		var f Field[string]
		require.NoError(t, json.Unmarshal([]byte(`"world"`), &f))
		assert.True(t, f.IsPresent())
		assert.Equal(t, "world", f.MustGet())
	})

	t.Run("present uuid", func(t *testing.T) {
		t.Parallel()
		var f Field[uuid.UUID]
		require.NoError(t, json.Unmarshal([]byte(`"550e8400-e29b-41d4-a716-446655440000"`), &f))
		assert.True(t, f.IsPresent())
		assert.Equal(t, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), f.MustGet())
	})

	t.Run("present decimal", func(t *testing.T) {
		t.Parallel()
		var f Field[decimal.Decimal]
		require.NoError(t, json.Unmarshal([]byte(`"123.45"`), &f))
		assert.True(t, f.IsPresent())
		assert.Equal(t, decimal.RequireFromString("123.45"), f.MustGet())
	})

	t.Run("present time", func(t *testing.T) {
		t.Parallel()
		var f Field[time.Time]
		want := time.Date(2024, time.January, 2, 3, 4, 5, 123456789, time.UTC)
		require.NoError(t, json.Unmarshal([]byte(`"2024-01-02T03:04:05.123456789Z"`), &f))
		assert.True(t, f.IsPresent())
		assert.True(t, f.MustGet().Equal(want))
	})

	t.Run("null sets stateNull", func(t *testing.T) {
		t.Parallel()
		var f Field[int]
		require.NoError(t, json.Unmarshal([]byte("null"), &f))
		assert.True(t, f.IsNull())
		assert.False(t, f.IsPresent())
	})

	t.Run("null with surrounding whitespace sets stateNull", func(t *testing.T) {
		t.Parallel()
		var f Field[int]
		require.NoError(t, json.Unmarshal([]byte(" \n\t null \r\n "), &f))
		assert.True(t, f.IsNull())
		assert.False(t, f.IsPresent())
	})

	t.Run("invalid json returns error", func(t *testing.T) {
		t.Parallel()
		var f Field[int]
		err := json.Unmarshal([]byte(`"not_a_number"`), &f)
		assert.Error(t, err)
	})

	t.Run("invalid uuid returns error", func(t *testing.T) {
		t.Parallel()
		var f Field[uuid.UUID]
		err := json.Unmarshal([]byte(`"not-a-uuid"`), &f)
		assert.Error(t, err)
	})

	t.Run("invalid decimal returns error", func(t *testing.T) {
		t.Parallel()
		var f Field[decimal.Decimal]
		err := json.Unmarshal([]byte(`"not-a-decimal"`), &f)
		assert.Error(t, err)
	})

	t.Run("invalid time returns error", func(t *testing.T) {
		t.Parallel()
		var f Field[time.Time]
		err := json.Unmarshal([]byte(`"not-a-time"`), &f)
		assert.Error(t, err)
	})
}

func TestField_UnmarshalJSON_InStruct(t *testing.T) {
	t.Parallel()

	type request struct {
		Name Field[string] `json:"name"`
		Age  Field[int]    `json:"age"`
	}

	t.Run("all fields present", func(t *testing.T) {
		t.Parallel()
		var r request
		require.NoError(t, json.Unmarshal([]byte(`{"name":"Alice","age":30}`), &r))
		assert.True(t, r.Name.IsPresent())
		assert.Equal(t, "Alice", r.Name.MustGet())
		assert.True(t, r.Age.IsPresent())
		assert.Equal(t, 30, r.Age.MustGet())
	})

	t.Run("explicit null", func(t *testing.T) {
		t.Parallel()
		var r request
		require.NoError(t, json.Unmarshal([]byte(`{"name":null,"age":30}`), &r))
		assert.True(t, r.Name.IsNull())
		assert.True(t, r.Age.IsPresent())
	})

	t.Run("missing field stays missing (zero value)", func(t *testing.T) {
		t.Parallel()
		var r request
		require.NoError(t, json.Unmarshal([]byte(`{"age":25}`), &r))
		assert.True(t, r.Name.IsMissing(), "omitted field should be missing (zero value)")
		assert.True(t, r.Age.IsPresent())
		assert.Equal(t, 25, r.Age.MustGet())
	})

	t.Run("all fields missing", func(t *testing.T) {
		t.Parallel()
		var r request
		require.NoError(t, json.Unmarshal([]byte(`{}`), &r))
		assert.True(t, r.Name.IsMissing())
		assert.True(t, r.Age.IsMissing())
	})
}

func TestField_JSON_RoundTrip(t *testing.T) {
	t.Parallel()

	t.Run("present value survives round-trip", func(t *testing.T) {
		t.Parallel()
		original := Of(99)
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var restored Field[int]
		require.NoError(t, json.Unmarshal(data, &restored))
		assert.True(t, restored.IsPresent())
		assert.Equal(t, 99, restored.MustGet())
	})

	t.Run("null round-trips to null", func(t *testing.T) {
		t.Parallel()
		original := Null[int]()
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var restored Field[int]
		require.NoError(t, json.Unmarshal(data, &restored))
		assert.True(t, restored.IsNull())
	})

	t.Run("struct round-trip", func(t *testing.T) {
		t.Parallel()
		type dto struct {
			Value Field[string] `json:"value"`
		}
		original := dto{Value: Of("test")}
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var restored dto
		require.NoError(t, json.Unmarshal(data, &restored))
		assert.Equal(t, "test", restored.Value.MustGet())
	})
}

func TestField_Value(t *testing.T) {
	t.Parallel()

	t.Run("present int", func(t *testing.T) {
		t.Parallel()
		got, err := Of(42).Value()
		require.NoError(t, err)
		assert.Equal(t, 42, got)
	})

	t.Run("present string", func(t *testing.T) {
		t.Parallel()
		got, err := Of("hello").Value()
		require.NoError(t, err)
		assert.Equal(t, "hello", got)
	})

	t.Run("present bool", func(t *testing.T) {
		t.Parallel()
		got, err := Of(true).Value()
		require.NoError(t, err)
		assert.Equal(t, true, got)
	})

	t.Run("present zero value", func(t *testing.T) {
		t.Parallel()
		got, err := Of(0).Value()
		require.NoError(t, err)
		assert.Equal(t, 0, got)
	})

	t.Run("present uuid", func(t *testing.T) {
		t.Parallel()
		want := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		got, err := Of(want).Value()
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("present decimal", func(t *testing.T) {
		t.Parallel()
		want := decimal.RequireFromString("123.45")
		got, err := Of(want).Value()
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("present time", func(t *testing.T) {
		t.Parallel()
		want := time.Date(2024, time.January, 2, 3, 4, 5, 123456789, time.UTC)
		got, err := Of(want).Value()
		require.NoError(t, err)

		gotTime, ok := got.(time.Time)
		require.True(t, ok)
		assert.True(t, gotTime.Equal(want))
	})

	t.Run("null returns nil", func(t *testing.T) {
		t.Parallel()
		got, err := Null[int]().Value()
		require.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("missing returns error", func(t *testing.T) {
		t.Parallel()
		got, err := Missing[int]().Value()
		require.Error(t, err)
		assert.Nil(t, got)
		assert.Contains(t, err.Error(), "missing value cannot be used in SQL")
	})

	t.Run("unknown state returns error", func(t *testing.T) {
		t.Parallel()
		got, err := Field[int]{state: state(255)}.Value()
		require.Error(t, err)
		assert.Nil(t, got)
		assert.Contains(t, err.Error(), "unknown state")
	})
}

func TestField_ImplementsJSONInterfaces(t *testing.T) {
	t.Parallel()

	var _ json.Marshaler = Field[int]{}
	var _ json.Unmarshaler = &Field[int]{}
	var _ driver.Valuer = Field[int]{}
}
