package optional

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// Bool
// ---------------------------------------------------------------------------

// Bool is a three-state optional for bool values.
type Bool struct {
	Field[bool]
}

var (
	_ sql.Scanner   = (*Bool)(nil)
	_ driver.Valuer = Bool{}
)

// OfBool creates a Bool with the given value in the present state.
func OfBool(value bool) Bool {
	return Bool{Field: Of(value)}
}

// FromBoolPtr creates a Bool from a pointer.
// If ptr is nil, the Bool is null; otherwise it contains the dereferenced value.
func FromBoolPtr(ptr *bool) Bool {
	return Bool{Field: FromPtr(ptr)}
}

// NullBool creates a Bool in the null state.
func NullBool() Bool {
	return Bool{Field: Null[bool]()}
}

// MissingBool creates a Bool in the missing state.
func MissingBool() Bool {
	return Bool{Field: Missing[bool]()}
}

// Scan implements [sql.Scanner].
func (b *Bool) Scan(src any) error {
	if src == nil {
		b.Field = Null[bool]()
		return nil
	}
	v, err := convertToBool(src)
	if err != nil {
		return err
	}
	b.Field = Of(v)
	return nil
}

// ---------------------------------------------------------------------------
// Int
// ---------------------------------------------------------------------------

// Int is a three-state optional for int values.
type Int struct {
	Field[int]
}

var (
	_ sql.Scanner   = (*Int)(nil)
	_ driver.Valuer = Int{}
)

// OfInt creates an Int with the given value in the present state.
func OfInt(value int) Int {
	return Int{Field: Of(value)}
}

// FromIntPtr creates an Int from a pointer.
func FromIntPtr(ptr *int) Int {
	return Int{Field: FromPtr(ptr)}
}

// NullInt creates an Int in the null state.
func NullInt() Int {
	return Int{Field: Null[int]()}
}

// MissingInt creates an Int in the missing state.
func MissingInt() Int {
	return Int{Field: Missing[int]()}
}

// Scan implements [sql.Scanner].
func (i *Int) Scan(src any) error {
	if src == nil {
		i.Field = Null[int]()
		return nil
	}
	n, err := convertToInt64(src, 0)
	if err != nil {
		return err
	}
	i.Field = Of(int(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (i Int) Value() (driver.Value, error) {
	if !i.Field.IsPresent() {
		return i.Field.Value()
	}
	return int64(i.Field.value), nil
}

// ---------------------------------------------------------------------------
// Int8
// ---------------------------------------------------------------------------

// Int8 is a three-state optional for int8 values.
type Int8 struct {
	Field[int8]
}

var (
	_ sql.Scanner   = (*Int8)(nil)
	_ driver.Valuer = Int8{}
)

// OfInt8 creates an Int8 with the given value in the present state.
func OfInt8(value int8) Int8 {
	return Int8{Field: Of(value)}
}

// FromInt8Ptr creates an Int8 from a pointer.
func FromInt8Ptr(ptr *int8) Int8 {
	return Int8{Field: FromPtr(ptr)}
}

// NullInt8 creates an Int8 in the null state.
func NullInt8() Int8 {
	return Int8{Field: Null[int8]()}
}

// MissingInt8 creates an Int8 in the missing state.
func MissingInt8() Int8 {
	return Int8{Field: Missing[int8]()}
}

// Scan implements [sql.Scanner].
func (i *Int8) Scan(src any) error {
	if src == nil {
		i.Field = Null[int8]()
		return nil
	}
	n, err := convertToInt64(src, 8)
	if err != nil {
		return err
	}
	i.Field = Of(int8(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (i Int8) Value() (driver.Value, error) {
	if !i.Field.IsPresent() {
		return i.Field.Value()
	}
	return int64(i.Field.value), nil
}

// ---------------------------------------------------------------------------
// Int16
// ---------------------------------------------------------------------------

// Int16 is a three-state optional for int16 values.
type Int16 struct {
	Field[int16]
}

var (
	_ sql.Scanner   = (*Int16)(nil)
	_ driver.Valuer = Int16{}
)

// OfInt16 creates an Int16 with the given value in the present state.
func OfInt16(value int16) Int16 {
	return Int16{Field: Of(value)}
}

// FromInt16Ptr creates an Int16 from a pointer.
func FromInt16Ptr(ptr *int16) Int16 {
	return Int16{Field: FromPtr(ptr)}
}

// NullInt16 creates an Int16 in the null state.
func NullInt16() Int16 {
	return Int16{Field: Null[int16]()}
}

// MissingInt16 creates an Int16 in the missing state.
func MissingInt16() Int16 {
	return Int16{Field: Missing[int16]()}
}

// Scan implements [sql.Scanner].
func (i *Int16) Scan(src any) error {
	if src == nil {
		i.Field = Null[int16]()
		return nil
	}
	n, err := convertToInt64(src, 16)
	if err != nil {
		return err
	}
	i.Field = Of(int16(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (i Int16) Value() (driver.Value, error) {
	if !i.Field.IsPresent() {
		return i.Field.Value()
	}
	return int64(i.Field.value), nil
}

// ---------------------------------------------------------------------------
// Int32
// ---------------------------------------------------------------------------

// Int32 is a three-state optional for int32 values.
type Int32 struct {
	Field[int32]
}

var (
	_ sql.Scanner   = (*Int32)(nil)
	_ driver.Valuer = Int32{}
)

// OfInt32 creates an Int32 with the given value in the present state.
func OfInt32(value int32) Int32 {
	return Int32{Field: Of(value)}
}

// FromInt32Ptr creates an Int32 from a pointer.
func FromInt32Ptr(ptr *int32) Int32 {
	return Int32{Field: FromPtr(ptr)}
}

// NullInt32 creates an Int32 in the null state.
func NullInt32() Int32 {
	return Int32{Field: Null[int32]()}
}

// MissingInt32 creates an Int32 in the missing state.
func MissingInt32() Int32 {
	return Int32{Field: Missing[int32]()}
}

// Scan implements [sql.Scanner].
func (i *Int32) Scan(src any) error {
	if src == nil {
		i.Field = Null[int32]()
		return nil
	}
	n, err := convertToInt64(src, 32)
	if err != nil {
		return err
	}
	i.Field = Of(int32(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (i Int32) Value() (driver.Value, error) {
	if !i.Field.IsPresent() {
		return i.Field.Value()
	}
	return int64(i.Field.value), nil
}

// ---------------------------------------------------------------------------
// Int64
// ---------------------------------------------------------------------------

// Int64 is a three-state optional for int64 values.
type Int64 struct {
	Field[int64]
}

var (
	_ sql.Scanner   = (*Int64)(nil)
	_ driver.Valuer = Int64{}
)

// OfInt64 creates an Int64 with the given value in the present state.
func OfInt64(value int64) Int64 {
	return Int64{Field: Of(value)}
}

// FromInt64Ptr creates an Int64 from a pointer.
func FromInt64Ptr(ptr *int64) Int64 {
	return Int64{Field: FromPtr(ptr)}
}

// NullInt64 creates an Int64 in the null state.
func NullInt64() Int64 {
	return Int64{Field: Null[int64]()}
}

// MissingInt64 creates an Int64 in the missing state.
func MissingInt64() Int64 {
	return Int64{Field: Missing[int64]()}
}

// Scan implements [sql.Scanner].
func (i *Int64) Scan(src any) error {
	if src == nil {
		i.Field = Null[int64]()
		return nil
	}
	n, err := convertToInt64(src, 64)
	if err != nil {
		return err
	}
	i.Field = Of(n)
	return nil
}

// ---------------------------------------------------------------------------
// Uint
// ---------------------------------------------------------------------------

// Uint is a three-state optional for uint values.
type Uint struct {
	Field[uint]
}

var (
	_ sql.Scanner   = (*Uint)(nil)
	_ driver.Valuer = Uint{}
)

// OfUint creates a Uint with the given value in the present state.
func OfUint(value uint) Uint {
	return Uint{Field: Of(value)}
}

// FromUintPtr creates a Uint from a pointer.
func FromUintPtr(ptr *uint) Uint {
	return Uint{Field: FromPtr(ptr)}
}

// NullUint creates a Uint in the null state.
func NullUint() Uint {
	return Uint{Field: Null[uint]()}
}

// MissingUint creates a Uint in the missing state.
func MissingUint() Uint {
	return Uint{Field: Missing[uint]()}
}

// Scan implements [sql.Scanner].
func (u *Uint) Scan(src any) error {
	if src == nil {
		u.Field = Null[uint]()
		return nil
	}
	n, err := convertToUint64(src, 0)
	if err != nil {
		return err
	}
	u.Field = Of(uint(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (u Uint) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	v := u.Field.value
	if uint64(v) > math.MaxInt64 {
		return nil, fmt.Errorf("optional: uint value %d overflows int64 for SQL driver", v)
	}
	return int64(v), nil
}

// ---------------------------------------------------------------------------
// Uint8
// ---------------------------------------------------------------------------

// Uint8 is a three-state optional for uint8 values.
type Uint8 struct {
	Field[uint8]
}

var (
	_ sql.Scanner   = (*Uint8)(nil)
	_ driver.Valuer = Uint8{}
)

// OfUint8 creates a Uint8 with the given value in the present state.
func OfUint8(value uint8) Uint8 {
	return Uint8{Field: Of(value)}
}

// FromUint8Ptr creates a Uint8 from a pointer.
func FromUint8Ptr(ptr *uint8) Uint8 {
	return Uint8{Field: FromPtr(ptr)}
}

// NullUint8 creates a Uint8 in the null state.
func NullUint8() Uint8 {
	return Uint8{Field: Null[uint8]()}
}

// MissingUint8 creates a Uint8 in the missing state.
func MissingUint8() Uint8 {
	return Uint8{Field: Missing[uint8]()}
}

// Scan implements [sql.Scanner].
func (u *Uint8) Scan(src any) error {
	if src == nil {
		u.Field = Null[uint8]()
		return nil
	}
	n, err := convertToUint64(src, 8)
	if err != nil {
		return err
	}
	u.Field = Of(uint8(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (u Uint8) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	return int64(u.Field.value), nil
}

// ---------------------------------------------------------------------------
// Uint16
// ---------------------------------------------------------------------------

// Uint16 is a three-state optional for uint16 values.
type Uint16 struct {
	Field[uint16]
}

var (
	_ sql.Scanner   = (*Uint16)(nil)
	_ driver.Valuer = Uint16{}
)

// OfUint16 creates a Uint16 with the given value in the present state.
func OfUint16(value uint16) Uint16 {
	return Uint16{Field: Of(value)}
}

// FromUint16Ptr creates a Uint16 from a pointer.
func FromUint16Ptr(ptr *uint16) Uint16 {
	return Uint16{Field: FromPtr(ptr)}
}

// NullUint16 creates a Uint16 in the null state.
func NullUint16() Uint16 {
	return Uint16{Field: Null[uint16]()}
}

// MissingUint16 creates a Uint16 in the missing state.
func MissingUint16() Uint16 {
	return Uint16{Field: Missing[uint16]()}
}

// Scan implements [sql.Scanner].
func (u *Uint16) Scan(src any) error {
	if src == nil {
		u.Field = Null[uint16]()
		return nil
	}
	n, err := convertToUint64(src, 16)
	if err != nil {
		return err
	}
	u.Field = Of(uint16(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (u Uint16) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	return int64(u.Field.value), nil
}

// ---------------------------------------------------------------------------
// Uint32
// ---------------------------------------------------------------------------

// Uint32 is a three-state optional for uint32 values.
type Uint32 struct {
	Field[uint32]
}

var (
	_ sql.Scanner   = (*Uint32)(nil)
	_ driver.Valuer = Uint32{}
)

// OfUint32 creates a Uint32 with the given value in the present state.
func OfUint32(value uint32) Uint32 {
	return Uint32{Field: Of(value)}
}

// FromUint32Ptr creates a Uint32 from a pointer.
func FromUint32Ptr(ptr *uint32) Uint32 {
	return Uint32{Field: FromPtr(ptr)}
}

// NullUint32 creates a Uint32 in the null state.
func NullUint32() Uint32 {
	return Uint32{Field: Null[uint32]()}
}

// MissingUint32 creates a Uint32 in the missing state.
func MissingUint32() Uint32 {
	return Uint32{Field: Missing[uint32]()}
}

// Scan implements [sql.Scanner].
func (u *Uint32) Scan(src any) error {
	if src == nil {
		u.Field = Null[uint32]()
		return nil
	}
	n, err := convertToUint64(src, 32)
	if err != nil {
		return err
	}
	u.Field = Of(uint32(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
func (u Uint32) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	return int64(u.Field.value), nil
}

// ---------------------------------------------------------------------------
// Uint64
// ---------------------------------------------------------------------------

// Uint64 is a three-state optional for uint64 values.
type Uint64 struct {
	Field[uint64]
}

var (
	_ sql.Scanner   = (*Uint64)(nil)
	_ driver.Valuer = Uint64{}
)

// OfUint64 creates a Uint64 with the given value in the present state.
func OfUint64(value uint64) Uint64 {
	return Uint64{Field: Of(value)}
}

// FromUint64Ptr creates a Uint64 from a pointer.
func FromUint64Ptr(ptr *uint64) Uint64 {
	return Uint64{Field: FromPtr(ptr)}
}

// NullUint64 creates a Uint64 in the null state.
func NullUint64() Uint64 {
	return Uint64{Field: Null[uint64]()}
}

// MissingUint64 creates a Uint64 in the missing state.
func MissingUint64() Uint64 {
	return Uint64{Field: Missing[uint64]()}
}

// Scan implements [sql.Scanner].
func (u *Uint64) Scan(src any) error {
	if src == nil {
		u.Field = Null[uint64]()
		return nil
	}
	n, err := convertToUint64(src, 64)
	if err != nil {
		return err
	}
	u.Field = Of(n)
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
// Returns an error if the value exceeds [math.MaxInt64].
func (u Uint64) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	v := u.Field.value
	if v > math.MaxInt64 {
		return nil, fmt.Errorf("optional: uint64 value %d overflows int64 for SQL driver", v)
	}
	return int64(v), nil
}

// ---------------------------------------------------------------------------
// Uintptr
// ---------------------------------------------------------------------------

// Uintptr is a three-state optional for uintptr values.
type Uintptr struct {
	Field[uintptr]
}

var (
	_ sql.Scanner   = (*Uintptr)(nil)
	_ driver.Valuer = Uintptr{}
)

// OfUintptr creates a Uintptr with the given value in the present state.
func OfUintptr(value uintptr) Uintptr {
	return Uintptr{Field: Of(value)}
}

// FromUintptrPtr creates a Uintptr from a pointer.
func FromUintptrPtr(ptr *uintptr) Uintptr {
	return Uintptr{Field: FromPtr(ptr)}
}

// NullUintptr creates a Uintptr in the null state.
func NullUintptr() Uintptr {
	return Uintptr{Field: Null[uintptr]()}
}

// MissingUintptr creates a Uintptr in the missing state.
func MissingUintptr() Uintptr {
	return Uintptr{Field: Missing[uintptr]()}
}

// Scan implements [sql.Scanner].
func (u *Uintptr) Scan(src any) error {
	if src == nil {
		u.Field = Null[uintptr]()
		return nil
	}
	n, err := convertToUint64(src, 64)
	if err != nil {
		return err
	}
	u.Field = Of(uintptr(n))
	return nil
}

// Value implements [driver.Valuer], returning int64 for SQL compatibility.
// Returns an error if the value exceeds [math.MaxInt64].
func (u Uintptr) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	v := uint64(u.Field.value)
	if v > math.MaxInt64 {
		return nil, fmt.Errorf("optional: uintptr value %d overflows int64 for SQL driver", v)
	}
	return int64(v), nil
}

// ---------------------------------------------------------------------------
// Float32
// ---------------------------------------------------------------------------

// Float32 is a three-state optional for float32 values.
type Float32 struct {
	Field[float32]
}

var (
	_ sql.Scanner   = (*Float32)(nil)
	_ driver.Valuer = Float32{}
)

// OfFloat32 creates a Float32 with the given value in the present state.
func OfFloat32(value float32) Float32 {
	return Float32{Field: Of(value)}
}

// FromFloat32Ptr creates a Float32 from a pointer.
func FromFloat32Ptr(ptr *float32) Float32 {
	return Float32{Field: FromPtr(ptr)}
}

// NullFloat32 creates a Float32 in the null state.
func NullFloat32() Float32 {
	return Float32{Field: Null[float32]()}
}

// MissingFloat32 creates a Float32 in the missing state.
func MissingFloat32() Float32 {
	return Float32{Field: Missing[float32]()}
}

// Scan implements [sql.Scanner].
func (f *Float32) Scan(src any) error {
	if src == nil {
		f.Field = Null[float32]()
		return nil
	}
	v, err := convertToFloat64(src)
	if err != nil {
		return err
	}
	if v > math.MaxFloat32 || v < -math.MaxFloat32 {
		return fmt.Errorf("optional: value %g overflows float32", v)
	}
	f.Field = Of(float32(v))
	return nil
}

// Value implements [driver.Valuer], returning float64 for SQL compatibility.
func (f Float32) Value() (driver.Value, error) {
	if !f.Field.IsPresent() {
		return f.Field.Value()
	}
	return float64(f.Field.value), nil
}

// ---------------------------------------------------------------------------
// Float64
// ---------------------------------------------------------------------------

// Float64 is a three-state optional for float64 values.
type Float64 struct {
	Field[float64]
}

var (
	_ sql.Scanner   = (*Float64)(nil)
	_ driver.Valuer = Float64{}
)

// OfFloat64 creates a Float64 with the given value in the present state.
func OfFloat64(value float64) Float64 {
	return Float64{Field: Of(value)}
}

// FromFloat64Ptr creates a Float64 from a pointer.
func FromFloat64Ptr(ptr *float64) Float64 {
	return Float64{Field: FromPtr(ptr)}
}

// NullFloat64 creates a Float64 in the null state.
func NullFloat64() Float64 {
	return Float64{Field: Null[float64]()}
}

// MissingFloat64 creates a Float64 in the missing state.
func MissingFloat64() Float64 {
	return Float64{Field: Missing[float64]()}
}

// Scan implements [sql.Scanner].
func (f *Float64) Scan(src any) error {
	if src == nil {
		f.Field = Null[float64]()
		return nil
	}
	v, err := convertToFloat64(src)
	if err != nil {
		return err
	}
	f.Field = Of(v)
	return nil
}

// ---------------------------------------------------------------------------
// String
// ---------------------------------------------------------------------------

// String is a three-state optional for string values.
type String struct {
	Field[string]
}

var (
	_ sql.Scanner   = (*String)(nil)
	_ driver.Valuer = String{}
)

// OfString creates a String with the given value in the present state.
func OfString(value string) String {
	return String{Field: Of(value)}
}

// FromStringPtr creates a String from a pointer.
func FromStringPtr(ptr *string) String {
	return String{Field: FromPtr(ptr)}
}

// NullString creates a String in the null state.
func NullString() String {
	return String{Field: Null[string]()}
}

// MissingString creates a String in the missing state.
func MissingString() String {
	return String{Field: Missing[string]()}
}

// Scan implements [sql.Scanner].
func (s *String) Scan(src any) error {
	if src == nil {
		s.Field = Null[string]()
		return nil
	}
	v, err := convertToString(src)
	if err != nil {
		return err
	}
	s.Field = Of(v)
	return nil
}

// ---------------------------------------------------------------------------
// Bytes
// ---------------------------------------------------------------------------

// Bytes is a three-state optional for []byte values (e.g., BLOB, BYTEA columns).
type Bytes struct {
	Field[[]byte]
}

var (
	_ sql.Scanner   = (*Bytes)(nil)
	_ driver.Valuer = Bytes{}
)

// OfBytes creates a Bytes with the given value in the present state.
func OfBytes(value []byte) Bytes {
	return Bytes{Field: Of(value)}
}

// FromBytesPtr creates a Bytes from a pointer.
func FromBytesPtr(ptr *[]byte) Bytes {
	return Bytes{Field: FromPtr(ptr)}
}

// NullBytes creates a Bytes in the null state.
func NullBytes() Bytes {
	return Bytes{Field: Null[[]byte]()}
}

// MissingBytes creates a Bytes in the missing state.
func MissingBytes() Bytes {
	return Bytes{Field: Missing[[]byte]()}
}

// Scan implements [sql.Scanner].
func (b *Bytes) Scan(src any) error {
	if src == nil {
		b.Field = Null[[]byte]()
		return nil
	}
	v, err := convertToBytes(src)
	if err != nil {
		return err
	}
	b.Field = Of(v)
	return nil
}

// ---------------------------------------------------------------------------
// Time
// ---------------------------------------------------------------------------

// Time is a three-state optional for [time.Time] values.
type Time struct {
	Field[time.Time]
}

var (
	_ sql.Scanner   = (*Time)(nil)
	_ driver.Valuer = Time{}
)

// OfTime creates a Time with the given value in the present state.
func OfTime(value time.Time) Time {
	return Time{Field: Of(value)}
}

// FromTimePtr creates a Time from a pointer.
func FromTimePtr(ptr *time.Time) Time {
	return Time{Field: FromPtr(ptr)}
}

// NullTime creates a Time in the null state.
func NullTime() Time {
	return Time{Field: Null[time.Time]()}
}

// MissingTime creates a Time in the missing state.
func MissingTime() Time {
	return Time{Field: Missing[time.Time]()}
}

// Scan implements [sql.Scanner].
func (t *Time) Scan(src any) error {
	if src == nil {
		t.Field = Null[time.Time]()
		return nil
	}
	v, err := convertToTime(src)
	if err != nil {
		return err
	}
	t.Field = Of(v)
	return nil
}

// ---------------------------------------------------------------------------
// UUID
// ---------------------------------------------------------------------------

// UUID is a three-state optional for [uuid.UUID] values.
type UUID struct {
	Field[uuid.UUID]
}

var (
	_ sql.Scanner   = (*UUID)(nil)
	_ driver.Valuer = UUID{}
)

// OfUUID creates a UUID with the given value in the present state.
func OfUUID(value uuid.UUID) UUID {
	return UUID{Field: Of(value)}
}

// FromUUIDPtr creates a UUID from a pointer.
func FromUUIDPtr(ptr *uuid.UUID) UUID {
	return UUID{Field: FromPtr(ptr)}
}

// NullUUID creates a UUID in the null state.
func NullUUID() UUID {
	return UUID{Field: Null[uuid.UUID]()}
}

// MissingUUID creates a UUID in the missing state.
func MissingUUID() UUID {
	return UUID{Field: Missing[uuid.UUID]()}
}

// Scan implements [sql.Scanner].
func (u *UUID) Scan(src any) error {
	if src == nil {
		u.Field = Null[uuid.UUID]()
		return nil
	}
	switch v := src.(type) {
	case uuid.UUID:
		u.Field = Of(v)
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("optional: cannot scan string into UUID: %w", err)
		}
		u.Field = Of(parsed)
	case []byte:
		// Try parsing as string first (common for most SQL drivers).
		if parsed, err := uuid.ParseBytes(v); err == nil {
			u.Field = Of(parsed)
			return nil
		}
		// Try raw 16-byte UUID.
		if len(v) == 16 {
			var id uuid.UUID
			copy(id[:], v)
			u.Field = Of(id)
			return nil
		}
		return fmt.Errorf("optional: cannot scan []byte (len=%d) into UUID", len(v))
	default:
		return fmt.Errorf("optional: cannot scan %T into UUID", src)
	}
	return nil
}

// Value implements [driver.Valuer], returning the UUID as a string for SQL compatibility.
func (u UUID) Value() (driver.Value, error) {
	if !u.Field.IsPresent() {
		return u.Field.Value()
	}
	return u.Field.value.String(), nil
}

// ---------------------------------------------------------------------------
// Decimal
// ---------------------------------------------------------------------------

// Decimal is a three-state optional for [decimal.Decimal] values.
type Decimal struct {
	Field[decimal.Decimal]
}

var (
	_ sql.Scanner   = (*Decimal)(nil)
	_ driver.Valuer = Decimal{}
)

// OfDecimal creates a Decimal with the given value in the present state.
func OfDecimal(value decimal.Decimal) Decimal {
	return Decimal{Field: Of(value)}
}

// FromDecimalPtr creates a Decimal from a pointer.
func FromDecimalPtr(ptr *decimal.Decimal) Decimal {
	return Decimal{Field: FromPtr(ptr)}
}

// NullDecimal creates a Decimal in the null state.
func NullDecimal() Decimal {
	return Decimal{Field: Null[decimal.Decimal]()}
}

// MissingDecimal creates a Decimal in the missing state.
func MissingDecimal() Decimal {
	return Decimal{Field: Missing[decimal.Decimal]()}
}

// Scan implements [sql.Scanner].
// Supports string, []byte, float64, and int64 source types.
func (d *Decimal) Scan(src any) error {
	if src == nil {
		d.Field = Null[decimal.Decimal]()
		return nil
	}
	switch v := src.(type) {
	case decimal.Decimal:
		d.Field = Of(v)
	case string:
		parsed, err := decimal.NewFromString(v)
		if err != nil {
			return fmt.Errorf("optional: cannot scan string %q into Decimal: %w", v, err)
		}
		d.Field = Of(parsed)
	case []byte:
		parsed, err := decimal.NewFromString(string(v))
		if err != nil {
			return fmt.Errorf("optional: cannot scan []byte into Decimal: %w", err)
		}
		d.Field = Of(parsed)
	case float64:
		d.Field = Of(decimal.NewFromFloat(v))
	case int64:
		d.Field = Of(decimal.NewFromInt(v))
	default:
		return fmt.Errorf("optional: cannot scan %T into Decimal", src)
	}
	return nil
}

// Value implements [driver.Valuer], returning the decimal as a string for SQL compatibility.
func (d Decimal) Value() (driver.Value, error) {
	if !d.Field.IsPresent() {
		return d.Field.Value()
	}
	return d.Field.value.String(), nil
}

// OfNullableBool creates a Bool from a pointer.
// It is an alias for [FromBoolPtr].
func OfNullableBool(ptr *bool) Bool { return FromBoolPtr(ptr) }

// OfNullableInt creates an Int from a pointer.
// It is an alias for [FromIntPtr].
func OfNullableInt(ptr *int) Int { return FromIntPtr(ptr) }

// OfNullableInt8 creates an Int8 from a pointer.
// It is an alias for [FromInt8Ptr].
func OfNullableInt8(ptr *int8) Int8 { return FromInt8Ptr(ptr) }

// OfNullableInt16 creates an Int16 from a pointer.
// It is an alias for [FromInt16Ptr].
func OfNullableInt16(ptr *int16) Int16 { return FromInt16Ptr(ptr) }

// OfNullableInt32 creates an Int32 from a pointer.
// It is an alias for [FromInt32Ptr].
func OfNullableInt32(ptr *int32) Int32 { return FromInt32Ptr(ptr) }

// OfNullableInt64 creates an Int64 from a pointer.
// It is an alias for [FromInt64Ptr].
func OfNullableInt64(ptr *int64) Int64 { return FromInt64Ptr(ptr) }

// OfNullableUint creates a Uint from a pointer.
// It is an alias for [FromUintPtr].
func OfNullableUint(ptr *uint) Uint { return FromUintPtr(ptr) }

// OfNullableUint8 creates a Uint8 from a pointer.
// It is an alias for [FromUint8Ptr].
func OfNullableUint8(ptr *uint8) Uint8 { return FromUint8Ptr(ptr) }

// OfNullableUint16 creates a Uint16 from a pointer.
// It is an alias for [FromUint16Ptr].
func OfNullableUint16(ptr *uint16) Uint16 { return FromUint16Ptr(ptr) }

// OfNullableUint32 creates a Uint32 from a pointer.
// It is an alias for [FromUint32Ptr].
func OfNullableUint32(ptr *uint32) Uint32 { return FromUint32Ptr(ptr) }

// OfNullableUint64 creates a Uint64 from a pointer.
// It is an alias for [FromUint64Ptr].
func OfNullableUint64(ptr *uint64) Uint64 { return FromUint64Ptr(ptr) }

// OfNullableUintptr creates a Uintptr from a pointer.
// It is an alias for [FromUintptrPtr].
func OfNullableUintptr(ptr *uintptr) Uintptr { return FromUintptrPtr(ptr) }

// OfNullableFloat32 creates a Float32 from a pointer.
// It is an alias for [FromFloat32Ptr].
func OfNullableFloat32(ptr *float32) Float32 { return FromFloat32Ptr(ptr) }

// OfNullableFloat64 creates a Float64 from a pointer.
// It is an alias for [FromFloat64Ptr].
func OfNullableFloat64(ptr *float64) Float64 { return FromFloat64Ptr(ptr) }

// OfNullableString creates a String from a pointer.
// It is an alias for [FromStringPtr].
func OfNullableString(ptr *string) String { return FromStringPtr(ptr) }

// OfNullableBytes creates a Bytes from a pointer.
// It is an alias for [FromBytesPtr].
func OfNullableBytes(ptr *[]byte) Bytes { return FromBytesPtr(ptr) }

// OfNullableTime creates a Time from a pointer.
// It is an alias for [FromTimePtr].
func OfNullableTime(ptr *time.Time) Time { return FromTimePtr(ptr) }

// OfNullableUUID creates a UUID from a pointer.
// It is an alias for [FromUUIDPtr].
func OfNullableUUID(ptr *uuid.UUID) UUID { return FromUUIDPtr(ptr) }

// OfNullableDecimal creates a Decimal from a pointer.
// It is an alias for [FromDecimalPtr].
func OfNullableDecimal(ptr *decimal.Decimal) Decimal { return FromDecimalPtr(ptr) }
