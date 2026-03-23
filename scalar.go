package optional

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Bool struct {
	Field[bool]
}

func OfBool(value bool) Bool {
	return Bool{Field: Of[bool](value)}
}

func OfNullableBool(ptr *bool) Bool {
	return Bool{Field: OfNullable[bool](ptr)}
}

func NullBool() Bool {
	return Bool{Field: Null[bool]()}
}

func MissingBool() Bool {
	return Bool{Field: Missing[bool]()}
}

type Float32 struct {
	Field[float32]
}

func OfFloat32(value float32) Float32 {
	return Float32{Field: Of[float32](value)}
}

func OfNullableFloat32(ptr *float32) Float32 {
	return Float32{Field: OfNullable[float32](ptr)}
}

func NullFloat32() Float32 {
	return Float32{Field: Null[float32]()}
}

func MissingFloat32() Float32 {
	return Float32{Field: Missing[float32]()}
}

type Float64 struct {
	Field[float64]
}

func OfFloat64(value float64) Float64 {
	return Float64{Field: Of[float64](value)}
}

func OfNullableFloat64(ptr *float64) Float64 {
	return Float64{Field: OfNullable[float64](ptr)}
}

func NullFloat64() Float64 {
	return Float64{Field: Null[float64]()}
}

func MissingFloat64() Float64 {
	return Float64{Field: Missing[float64]()}
}

type Int struct {
	Field[int]
}

func OfInt(value int) Int {
	return Int{Field: Of[int](value)}
}

func OfNullableInt(ptr *int) Int {
	return Int{Field: OfNullable[int](ptr)}
}

func NullInt() Int {
	return Int{Field: Null[int]()}
}

func MissingInt() Int {
	return Int{Field: Missing[int]()}
}

type Int8 struct {
	Field[int8]
}

func OfInt8(value int8) Int8 {
	return Int8{Field: Of[int8](value)}
}

func OfNullableInt8(ptr *int8) Int8 {
	return Int8{Field: OfNullable[int8](ptr)}
}

func NullInt8() Int8 {
	return Int8{Field: Null[int8]()}
}

func MissingInt8() Int8 {
	return Int8{Field: Missing[int8]()}
}

type Int16 struct {
	Field[int16]
}

func OfInt16(value int16) Int16 {
	return Int16{Field: Of[int16](value)}
}

func OfNullableInt16(ptr *int16) Int16 {
	return Int16{Field: OfNullable[int16](ptr)}
}

func NullInt16() Int16 {
	return Int16{Field: Null[int16]()}
}

func MissingInt16() Int16 {
	return Int16{Field: Missing[int16]()}
}

type Int32 struct {
	Field[int32]
}

func OfInt32(value int32) Int32 {
	return Int32{Field: Of[int32](value)}
}

func OfNullableInt32(ptr *int32) Int32 {
	return Int32{Field: OfNullable[int32](ptr)}
}

func NullInt32() Int32 {
	return Int32{Field: Null[int32]()}
}

func MissingInt32() Int32 {
	return Int32{Field: Missing[int32]()}
}

type Int64 struct {
	Field[int64]
}

func OfInt64(value int64) Int64 {
	return Int64{Field: Of[int64](value)}
}

func OfNullableInt64(ptr *int64) Int64 {
	return Int64{Field: OfNullable[int64](ptr)}
}

func NullInt64() Int64 {
	return Int64{Field: Null[int64]()}
}

func MissingInt64() Int64 {
	return Int64{Field: Missing[int64]()}
}

type Uint struct {
	Field[uint]
}

func OfUint(value uint) Uint {
	return Uint{Field: Of[uint](value)}
}

func OfNullableUint(ptr *uint) Uint {
	return Uint{Field: OfNullable[uint](ptr)}
}

func NullUint() Uint {
	return Uint{Field: Null[uint]()}
}

func MissingUint() Uint {
	return Uint{Field: Missing[uint]()}
}

type Uint8 struct {
	Field[uint8]
}

func OfUint8(value uint8) Uint8 {
	return Uint8{Field: Of[uint8](value)}
}

func OfNullableUint8(ptr *uint8) Uint8 {
	return Uint8{Field: OfNullable[uint8](ptr)}
}

func NullUint8() Uint8 {
	return Uint8{Field: Null[uint8]()}
}

func MissingUint8() Uint8 {
	return Uint8{Field: Missing[uint8]()}
}

type Uint16 struct {
	Field[uint16]
}

func OfUint16(value uint16) Uint16 {
	return Uint16{Field: Of[uint16](value)}
}

func OfNullableUint16(ptr *uint16) Uint16 {
	return Uint16{Field: OfNullable[uint16](ptr)}
}

func NullUint16() Uint16 {
	return Uint16{Field: Null[uint16]()}
}

func MissingUint16() Uint16 {
	return Uint16{Field: Missing[uint16]()}
}

type Uint32 struct {
	Field[uint32]
}

func OfUint32(value uint32) Uint32 {
	return Uint32{Field: Of[uint32](value)}
}

func OfNullableUint32(ptr *uint32) Uint32 {
	return Uint32{Field: OfNullable[uint32](ptr)}
}

func NullUint32() Uint32 {
	return Uint32{Field: Null[uint32]()}
}

func MissingUint32() Uint32 {
	return Uint32{Field: Missing[uint32]()}
}

type Uint64 struct {
	Field[uint64]
}

func OfUint64(value uint64) Uint64 {
	return Uint64{Field: Of[uint64](value)}
}

func OfNullableUint64(ptr *uint64) Uint64 {
	return Uint64{Field: OfNullable[uint64](ptr)}
}

func NullUint64() Uint64 {
	return Uint64{Field: Null[uint64]()}
}

func MissingUint64() Uint64 {
	return Uint64{Field: Missing[uint64]()}
}

type Uintptr struct {
	Field[uintptr]
}

func OfUintptr(value uintptr) Uintptr {
	return Uintptr{Field: Of[uintptr](value)}
}

func OfNullableUintptr(ptr *uintptr) Uintptr {
	return Uintptr{Field: OfNullable[uintptr](ptr)}
}

func NullUintptr() Uintptr {
	return Uintptr{Field: Null[uintptr]()}
}

func MissingUintptr() Uintptr {
	return Uintptr{Field: Missing[uintptr]()}
}

type String struct {
	Field[string]
}

func OfString(value string) String {
	return String{Field: Of[string](value)}
}

func OfNullableString(ptr *string) String {
	return String{Field: OfNullable[string](ptr)}
}

func NullString() String {
	return String{Field: Null[string]()}
}

func MissingString() String {
	return String{Field: Missing[string]()}
}

type UUID struct {
	Field[uuid.UUID]
}

func OfUUID(value uuid.UUID) UUID {
	return UUID{Field: Of[uuid.UUID](value)}
}

func OfNullableUUID(ptr *uuid.UUID) UUID {
	return UUID{Field: OfNullable[uuid.UUID](ptr)}
}

func NullUUID() UUID {
	return UUID{Field: Null[uuid.UUID]()}
}

func MissingUUID() UUID {
	return UUID{Field: Missing[uuid.UUID]()}
}

type Decimal struct {
	Field[decimal.Decimal]
}

func OfDecimal(value decimal.Decimal) Decimal {
	return Decimal{Field: Of[decimal.Decimal](value)}
}

func OfNullableDecimal(ptr *decimal.Decimal) Decimal {
	return Decimal{Field: OfNullable[decimal.Decimal](ptr)}
}

func NullDecimal() Decimal {
	return Decimal{Field: Null[decimal.Decimal]()}
}

func MissingDecimal() Decimal {
	return Decimal{Field: Missing[decimal.Decimal]()}
}

type Time struct {
	Field[time.Time]
}

func OfTime(value time.Time) Time {
	return Time{Field: Of[time.Time](value)}
}

func OfNullableTime(ptr *time.Time) Time {
	return Time{Field: OfNullable[time.Time](ptr)}
}

func NullTime() Time {
	return Time{Field: Null[time.Time]()}
}

func MissingTime() Time {
	return Time{Field: Missing[time.Time]()}
}
