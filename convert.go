package optional

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// convertToBool converts a SQL driver value to bool.
// Supported source types: bool, int64, float64, string, []byte.
func convertToBool(src any) (bool, error) {
	switch v := src.(type) {
	case bool:
		return v, nil
	case int64:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case string:
		if v == "" {
			return false, nil
		}
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf("optional: cannot convert string %q to bool: %w", v, err)
		}
		return b, nil
	case []byte:
		if len(v) == 0 {
			return false, nil
		}
		b, err := strconv.ParseBool(string(v))
		if err != nil {
			return false, fmt.Errorf("optional: cannot convert []byte to bool: %w", err)
		}
		return b, nil
	default:
		return false, fmt.Errorf("optional: cannot convert %T to bool", src)
	}
}

// convertToInt64 converts a SQL driver value to int64 with range checking.
// bitSize specifies the target integer size: 0 (platform int), 8, 16, 32, or 64.
// Supported source types: int64, float64, string, []byte.
func convertToInt64(src any, bitSize int) (int64, error) {
	effectiveBits := bitSize
	if effectiveBits == 0 {
		effectiveBits = strconv.IntSize
	}

	var n int64
	needsRangeCheck := true

	switch v := src.(type) {
	case int64:
		n = v
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, fmt.Errorf("optional: cannot convert non-finite float64 %v to int", v)
		}
		if v != math.Trunc(v) {
			return 0, fmt.Errorf("optional: cannot convert float64 %g to int: not a whole number", v)
		}
		if v < math.MinInt64 || v > math.MaxInt64 {
			return 0, fmt.Errorf("optional: value %g overflows int64", v)
		}
		n = int64(v)
	case string:
		var err error
		n, err = strconv.ParseInt(v, 10, effectiveBits)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert string %q to int%d: %w", v, effectiveBits, err)
		}
		needsRangeCheck = false // ParseInt already checks range
	case []byte:
		var err error
		n, err = strconv.ParseInt(string(v), 10, effectiveBits)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert []byte to int%d: %w", effectiveBits, err)
		}
		needsRangeCheck = false
	default:
		return 0, fmt.Errorf("optional: cannot convert %T to int%d", src, effectiveBits)
	}

	if needsRangeCheck && effectiveBits < 64 {
		minVal := int64(-1) << (effectiveBits - 1)
		maxVal := int64(1)<<(effectiveBits-1) - 1
		if n < minVal || n > maxVal {
			return 0, fmt.Errorf("optional: value %d overflows int%d", n, effectiveBits)
		}
	}

	return n, nil
}

// convertToUint64 converts a SQL driver value to uint64 with range checking.
// bitSize specifies the target unsigned integer size: 0 (platform uint), 8, 16, 32, or 64.
// Supported source types: int64, float64, string, []byte.
func convertToUint64(src any, bitSize int) (uint64, error) {
	effectiveBits := bitSize
	if effectiveBits == 0 {
		effectiveBits = strconv.IntSize
	}

	var n uint64
	needsRangeCheck := true

	switch v := src.(type) {
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("optional: negative value %d cannot be converted to uint%d", v, effectiveBits)
		}
		n = uint64(v)
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, fmt.Errorf("optional: cannot convert non-finite float64 %v to uint", v)
		}
		if v < 0 {
			return 0, fmt.Errorf("optional: negative value %g cannot be converted to uint%d", v, effectiveBits)
		}
		if v != math.Trunc(v) {
			return 0, fmt.Errorf("optional: cannot convert float64 %g to uint: not a whole number", v)
		}
		if v > math.MaxUint64 {
			return 0, fmt.Errorf("optional: value %g overflows uint64", v)
		}
		n = uint64(v)
	case string:
		var err error
		n, err = strconv.ParseUint(v, 10, effectiveBits)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert string %q to uint%d: %w", v, effectiveBits, err)
		}
		needsRangeCheck = false
	case []byte:
		var err error
		n, err = strconv.ParseUint(string(v), 10, effectiveBits)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert []byte to uint%d: %w", effectiveBits, err)
		}
		needsRangeCheck = false
	default:
		return 0, fmt.Errorf("optional: cannot convert %T to uint%d", src, effectiveBits)
	}

	if needsRangeCheck && effectiveBits < 64 {
		maxVal := uint64(1)<<effectiveBits - 1
		if n > maxVal {
			return 0, fmt.Errorf("optional: value %d overflows uint%d", n, effectiveBits)
		}
	}

	return n, nil
}

// convertToFloat64 converts a SQL driver value to float64.
// Supported source types: float64, int64, string, []byte.
func convertToFloat64(src any) (float64, error) {
	switch v := src.(type) {
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert string %q to float64: %w", v, err)
		}
		return f, nil
	case []byte:
		f, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			return 0, fmt.Errorf("optional: cannot convert []byte to float64: %w", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("optional: cannot convert %T to float64", src)
	}
}

// convertToString converts a SQL driver value to string.
// Supported source types: string, []byte.
func convertToString(src any) (string, error) {
	switch v := src.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		return "", fmt.Errorf("optional: cannot convert %T to string", src)
	}
}

// convertToBytes converts a SQL driver value to []byte.
// The returned slice is always a copy to avoid aliasing the driver's buffer.
// Supported source types: []byte, string.
func convertToBytes(src any) ([]byte, error) {
	switch v := src.(type) {
	case []byte:
		cp := make([]byte, len(v))
		copy(cp, v)
		return cp, nil
	case string:
		return []byte(v), nil
	default:
		return nil, fmt.Errorf("optional: cannot convert %T to []byte", src)
	}
}

// convertToTime converts a SQL driver value to time.Time.
// Supported source types: time.Time, string, []byte.
func convertToTime(src any) (time.Time, error) {
	switch v := src.(type) {
	case time.Time:
		return v, nil
	case string:
		return parseTime(v)
	case []byte:
		return parseTime(string(v))
	default:
		return time.Time{}, fmt.Errorf("optional: cannot convert %T to time.Time", src)
	}
}

// timeLayouts lists the time formats attempted by parseTime, in order of priority.
var timeLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02 15:04:05.999999999Z07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

// parseTime tries multiple common time formats.
func parseTime(s string) (time.Time, error) {
	for _, layout := range timeLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("optional: cannot parse %q as time.Time", s)
}
