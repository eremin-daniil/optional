// Package optional provides a three-state generic container for Go that
// distinguishes between a present value, an explicit null, and a missing
// (omitted) field.
//
// This is especially useful for JSON APIs with PATCH semantics and for working
// with nullable SQL columns, where you need to know whether a field was
// explicitly set to null, set to a value, or not provided at all.
//
// # Three States
//
// Every [Field] is in exactly one of these states:
//
//   - Present — the field contains a value.
//   - Null — the field was explicitly set to null.
//   - Missing — the field was not provided (the zero value of [Field]).
//
// # JSON Integration
//
// Field implements [encoding/json.Marshaler] and [encoding/json.Unmarshaler].
// When unmarshalling JSON:
//   - A JSON value → present state.
//   - A JSON null → null state.
//   - A missing JSON key → missing state (zero value, never touched by unmarshaler).
//
// Use the "omitzero" tag (Go 1.24+) to omit missing fields during marshalling:
//
//	type UpdateRequest struct {
//	    Name optional.String `json:"name,omitzero"`
//	    Age  optional.Int    `json:"age,omitzero"`
//	}
//
// # SQL Integration
//
// The concrete scalar types ([Bool], [Int], [String], [UUID], etc.) implement
// [database/sql.Scanner] and [database/sql/driver.Valuer] for seamless database
// integration with any SQL driver. Scan handles the standard driver value types
// (int64, float64, bool, []byte, string, time.Time) with automatic type
// conversion and range checking.
//
// # Scalar Types
//
// The following concrete types are provided for direct use with SQL databases:
//
//   - [Bool]
//   - [Int], [Int8], [Int16], [Int32], [Int64]
//   - [Uint], [Uint8], [Uint16], [Uint32], [Uint64], [Uintptr]
//   - [Float32], [Float64]
//   - [String], [Bytes]
//   - [Time]
//   - [UUID] (requires [github.com/google/uuid])
//   - [Decimal] (requires [github.com/shopspring/decimal])
//
// # Functional Helpers
//
// Standalone generic functions are provided for common transformations:
//
//   - [Map] — transforms the value inside a Field.
//   - [FlatMap] — monadic bind for chaining optional operations.
//   - [Equal] — compares two fields for equality.
package optional
