# optional

[![Go Reference](https://pkg.go.dev/badge/github.com/eremin-daniil/optional.svg)](https://pkg.go.dev/github.com/eremin-daniil/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/eremin-daniil/optional)](https://goreportcard.com/report/github.com/eremin-daniil/optional)

A production-ready **three-state generic optional** for Go that distinguishes between:

| State       | Meaning                              | JSON behaviour            | SQL behaviour    |
|-------------|--------------------------------------|---------------------------|------------------|
| **Present** | Field contains a value               | marshals the value        | sends the value  |
| **Null**    | Field was explicitly set to `null`   | marshals as `null`        | sends `NULL`     |
| **Missing** | Field was not provided at all        | omitted with `omitzero`   | returns an error |

This is essential for **JSON PATCH** semantics, nullable SQL columns, and any API where you need to know whether a client explicitly set a field to `null`, provided a value, or omitted it entirely.

## Installation

```bash
go get github.com/eremin-daniil/optional
```

Requires **Go 1.24+** (for `omitzero` JSON tag support).

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/eremin-daniil/optional"
)

type UpdateUserRequest struct {
    Name  optional.String `json:"name,omitzero"`
    Email optional.String `json:"email,omitzero"`
    Age   optional.Int    `json:"age,omitzero"`
}

func main() {
    input := `{"name": "Alice", "email": null}`

    var req UpdateUserRequest
    if err := json.Unmarshal([]byte(input), &req); err != nil {
        panic(err)
    }

    fmt.Println(req.Name.IsPresent()) // true  → client sent "Alice"
    fmt.Println(req.Email.IsNull())   // true  → client explicitly set null
    fmt.Println(req.Age.IsMissing())  // true  → client didn't send age at all
}
```

## API Overview

### Constructors

```go
// Generic constructors (for any type T)
present := optional.Of(42)
fromPtr := optional.FromPtr(&value)
nullable := optional.OfNullable(&value) // alias for FromPtr
nullValue := optional.Null[int]()
missingValue := optional.Missing[int]()

// Typed constructors (for SQL-compatible scalar types)
b := optional.OfBool(true)
i := optional.OfInt(42)
s := optional.OfString("hello")
u := optional.OfUUID(uuid.New())
d := optional.OfDecimal(decimal.NewFromFloat(99.99))
t := optional.OfTime(time.Now())

// From pointers (nil → null)
nullableBool := optional.OfNullableBool(boolPtr)   // alias: FromBoolPtr
nullableInt := optional.OfNullableInt(intPtr)      // alias: FromIntPtr
nullableString := optional.OfNullableString(strPtr) // alias: FromStringPtr
```

### Accessors

```go
v, ok := f.Get()        // value + presence check
fallbackValue := f.GetOr(fallback)
mustValue := f.MustGet() // value or panic
ptr := f.Ptr()           // *T or nil
resolved := f.Or(other)  // self if present, otherwise other
lazyValue := f.OrElse(func() T { return compute() })
```

### State Checks

```go
f.IsPresent() // has a value
f.IsNull()    // explicitly null
f.IsMissing() // not provided
f.IsZero()    // same as IsMissing (supports omitzero)
```

### Formatting

```go
fmt.Println(f)         // "42", "null", or "missing"
fmt.Printf("%#v", f)   // "optional.Of(42)", "optional.Null()", etc.
```

## JSON Integration

`Field[T]` implements `json.Marshaler` and `json.Unmarshaler`:

```go
type Request struct {
    Name optional.Field[string] `json:"name,omitzero"`
    Age  optional.Field[int]    `json:"age,omitzero"`
}
```

| JSON input               | Name state | Age state |
|--------------------------|------------|-----------|
| `{"name":"Alice","age":30}` | present    | present   |
| `{"name":null,"age":30}`    | null       | present   |
| `{"age":30}`                | **missing**| present   |
| `{}`                        | **missing**| **missing**|

The `omitzero` tag (Go 1.24+) ensures missing fields are omitted during marshalling.

## SQL Integration

All scalar types implement `sql.Scanner` and `driver.Valuer` for seamless database usage with **any SQL driver** (PostgreSQL, MySQL, SQLite, etc.):

```go
type User struct {
    ID    int
    Name  optional.String
    Email optional.String
    Age   optional.Int
}

// Reading from database
var user User
row := db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", 1)
err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)

// Writing to database
_, err = db.Exec(
    "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
    user.Name, user.Email, user.Age,
)
```

### Scan Type Conversion

Scan methods handle automatic type conversion from standard SQL driver types:

| Target type | Accepted SQL driver types                    |
|-------------|----------------------------------------------|
| `Bool`      | `bool`, `int64`, `float64`, `string`, `[]byte` |
| `Int*`      | `int64`, `float64`, `string`, `[]byte` (with overflow checking) |
| `Uint*`     | `int64`, `float64`, `string`, `[]byte` (with overflow and sign checking) |
| `Float*`    | `float64`, `int64`, `string`, `[]byte`       |
| `String`    | `string`, `[]byte`                           |
| `Bytes`     | `[]byte` (copied), `string`                  |
| `Time`      | `time.Time`, `string`, `[]byte` (multiple formats) |
| `UUID`      | `string`, `[]byte` (string format or raw 16 bytes) |
| `Decimal`   | `string`, `[]byte`, `float64`, `int64`       |

### Value Type Mapping

Value methods return proper `driver.Value` types:

| Scalar type          | `driver.Value` type |
|----------------------|---------------------|
| `Bool`               | `bool`              |
| `Int`, `Int8`–`Int64`| `int64`            |
| `Uint*`              | `int64` (overflow error if > MaxInt64) |
| `Float32`            | `float64`           |
| `Float64`            | `float64`           |
| `String`             | `string`            |
| `Bytes`              | `[]byte`            |
| `Time`               | `time.Time`         |
| `UUID`               | `string`            |
| `Decimal`            | `string`            |

## Scalar Types

| Type       | Go type            | External dependency |
|------------|--------------------|---------------------|
| `Bool`     | `bool`             | —                   |
| `Int`      | `int`              | —                   |
| `Int8`     | `int8`             | —                   |
| `Int16`    | `int16`            | —                   |
| `Int32`    | `int32`            | —                   |
| `Int64`    | `int64`            | —                   |
| `Uint`     | `uint`             | —                   |
| `Uint8`    | `uint8`            | —                   |
| `Uint16`   | `uint16`           | —                   |
| `Uint32`   | `uint32`           | —                   |
| `Uint64`   | `uint64`           | —                   |
| `Uintptr`  | `uintptr`          | —                   |
| `Float32`  | `float32`          | —                   |
| `Float64`  | `float64`          | —                   |
| `String`   | `string`           | —                   |
| `Bytes`    | `[]byte`           | —                   |
| `Time`     | `time.Time`        | —                   |
| `UUID`     | `uuid.UUID`        | [google/uuid](https://github.com/google/uuid) |
| `Decimal`  | `decimal.Decimal`  | [shopspring/decimal](https://github.com/shopspring/decimal) |

## Functional Helpers

```go
// Map transforms the value (standalone function due to Go generics limitations).
result := optional.Map(optional.Of(5), func(n int) string {
    return fmt.Sprintf("value: %d", n)
})
// result = Field[string]{present, "value: 5"}

// FlatMap for chaining optional operations.
result := optional.FlatMap(optional.Of(10), func(n int) optional.Field[float64] {
    if n == 0 {
        return optional.Null[float64]()
    }
    return optional.Of(100.0 / float64(n))
})

// Equal compares two fields (requires comparable types).
optional.Equal(optional.Of(42), optional.Of(42)) // true
optional.Equal(optional.Null[int](), optional.Null[int]()) // true
```

## PATCH Request Example

A complete example of using optional fields for a PATCH endpoint:

```go
type UpdateUserRequest struct {
    Name  optional.String `json:"name,omitzero"`
    Email optional.String `json:"email,omitzero"`
    Age   optional.Int    `json:"age,omitzero"`
}

func handlePatchUser(req UpdateUserRequest, userID int, db *sql.DB) error {
    // Build dynamic UPDATE query based on provided fields.
    var setClauses []string
    var args []any
    argIdx := 1

    if !req.Name.IsMissing() {
        setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIdx))
        args = append(args, req.Name)
        argIdx++
    }
    if !req.Email.IsMissing() {
        setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIdx))
        args = append(args, req.Email)
        argIdx++
    }
    if !req.Age.IsMissing() {
        setClauses = append(setClauses, fmt.Sprintf("age = $%d", argIdx))
        args = append(args, req.Age)
        argIdx++
    }

    if len(setClauses) == 0 {
        return nil // nothing to update
    }

    query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d",
        strings.Join(setClauses, ", "), argIdx)
    args = append(args, userID)

    _, err := db.Exec(query, args...)
    return err
}
```

## License

MIT License — see [LICENSE](LICENSE) for details.

