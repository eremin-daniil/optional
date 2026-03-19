package optional

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestState_Constants(t *testing.T) {
	t.Parallel()

	t.Run("iota values", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, state(0), stateMissing)
		assert.Equal(t, state(1), stateNull)
		assert.Equal(t, state(2), statePresent)
	})

	t.Run("distinct values", func(t *testing.T) {
		t.Parallel()

		states := []state{stateMissing, stateNull, statePresent}
		seen := make(map[state]bool, len(states))
		for _, s := range states {
			require.False(t, seen[s], "duplicate state value: %d", s)
			seen[s] = true
		}
	})
}

func TestState_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		state    state
		expected string
	}{
		{name: "missing", state: stateMissing, expected: "missing"},
		{name: "null", state: stateNull, expected: "null"},
		{name: "present", state: statePresent, expected: "present"},
		{name: "unknown (3)", state: state(3), expected: "unknown"},
		{name: "unknown (255)", state: state(255), expected: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, tt.state.String())
		})
	}
}

func TestState_ImplementsStringer(t *testing.T) {
	t.Parallel()

	var _ fmt.Stringer = stateMissing
	var _ fmt.Stringer = stateNull
	var _ fmt.Stringer = statePresent

	assert.Equal(t, "missing", fmt.Sprintf("%s", stateMissing))
	assert.Equal(t, "null", fmt.Sprintf("%s", stateNull))
	assert.Equal(t, "present", fmt.Sprintf("%s", statePresent))
}
