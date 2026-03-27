package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/eremin-daniil/optional"
)

func ExampleOf() {
	f := optional.Of(42)
	fmt.Println(f.IsPresent()) // true
	fmt.Println(f.MustGet())   // 42
	// Output:
	// true
	// 42
}

func ExampleNull() {
	f := optional.Null[string]()
	fmt.Println(f.IsNull())    // true
	fmt.Println(f.IsPresent()) // false
	// Output:
	// true
	// false
}

func ExampleMissing() {
	f := optional.Missing[int]()
	fmt.Println(f.IsMissing()) // true
	fmt.Println(f.IsZero())    // true (same as missing)
	// Output:
	// true
	// true
}

func ExampleFromPtr() {
	v := 42
	present := optional.FromPtr(&v)
	null := optional.FromPtr[int](nil)

	fmt.Println(present.IsPresent()) // true
	fmt.Println(null.IsNull())       // true
	// Output:
	// true
	// true
}

func ExampleField_GetOr() {
	present := optional.Of(5)
	null := optional.Null[int]()

	fmt.Println(present.GetOr(99)) // 5
	fmt.Println(null.GetOr(99))    // 99
	// Output:
	// 5
	// 99
}

func ExampleField_Or() {
	a := optional.Null[int]()
	b := optional.Of(42)

	result := a.Or(b)
	fmt.Println(result.MustGet()) // 42
	// Output:
	// 42
}

func ExampleField_MarshalJSON() {
	type Request struct {
		Name optional.Field[string] `json:"name,omitzero"`
		Age  optional.Field[int]    `json:"age,omitzero"`
	}

	// All fields present.
	r1 := Request{Name: optional.Of("Alice"), Age: optional.Of(30)}
	data1, _ := json.Marshal(r1)
	fmt.Println(string(data1))

	// Age is null.
	r2 := Request{Name: optional.Of("Bob"), Age: optional.Null[int]()}
	data2, _ := json.Marshal(r2)
	fmt.Println(string(data2))

	// Age is missing (omitted from JSON).
	r3 := Request{Name: optional.Of("Carol")}
	data3, _ := json.Marshal(r3)
	fmt.Println(string(data3))

	// Output:
	// {"name":"Alice","age":30}
	// {"name":"Bob","age":null}
	// {"name":"Carol"}
}

func ExampleField_UnmarshalJSON() {
	type Request struct {
		Name optional.Field[string] `json:"name"`
		Age  optional.Field[int]    `json:"age"`
	}

	// All fields present.
	var r1 Request
	json.Unmarshal([]byte(`{"name":"Alice","age":30}`), &r1)
	fmt.Printf("name: present=%v, age: present=%v\n", r1.Name.IsPresent(), r1.Age.IsPresent())

	// Age is null.
	var r2 Request
	json.Unmarshal([]byte(`{"name":"Bob","age":null}`), &r2)
	fmt.Printf("name: present=%v, age: null=%v\n", r2.Name.IsPresent(), r2.Age.IsNull())

	// Age is missing.
	var r3 Request
	json.Unmarshal([]byte(`{"name":"Carol"}`), &r3)
	fmt.Printf("name: present=%v, age: missing=%v\n", r3.Name.IsPresent(), r3.Age.IsMissing())

	// Output:
	// name: present=true, age: present=true
	// name: present=true, age: null=true
	// name: present=true, age: missing=true
}

func ExampleMap() {
	f := optional.Of(5)
	doubled := optional.Map(f, func(n int) int { return n * 2 })
	fmt.Println(doubled.MustGet()) // 10

	null := optional.Null[int]()
	result := optional.Map(null, func(n int) int { return n * 2 })
	fmt.Println(result.IsNull()) // true
	// Output:
	// 10
	// true
}

func ExampleFlatMap() {
	safeDivide := func(divisor int) optional.Field[float64] {
		if divisor == 0 {
			return optional.Null[float64]()
		}
		return optional.Of(100.0 / float64(divisor))
	}

	result := optional.FlatMap(optional.Of(4), safeDivide)
	fmt.Println(result.MustGet()) // 25

	zero := optional.FlatMap(optional.Of(0), safeDivide)
	fmt.Println(zero.IsNull()) // true
	// Output:
	// 25
	// true
}

func ExampleEqual() {
	fmt.Println(optional.Equal(optional.Of(42), optional.Of(42)))           // true
	fmt.Println(optional.Equal(optional.Of(1), optional.Of(2)))             // false
	fmt.Println(optional.Equal(optional.Null[int](), optional.Null[int]())) // true
	// Output:
	// true
	// false
	// true
}

func ExampleBool_Scan() {
	var b optional.Bool

	// Simulate scanning NULL from database.
	b.Scan(nil)
	fmt.Println(b.IsNull()) // true

	// Simulate scanning a value.
	b.Scan(true)
	fmt.Println(b.MustGet()) // true

	// Some databases return 0/1 for booleans.
	b.Scan(int64(1))
	fmt.Println(b.MustGet()) // true
	// Output:
	// true
	// true
	// true
}
