package copy_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/felix-kaestner/copy"
)

type String string

func NewString(s string) fmt.Stringer { return String(s) }

// String implements fmt.Stringer
func (s String) String() string { return string(s) }

func TestInterfaceNil(t *testing.T) {
	var want fmt.Stringer
	got := copy.Deep(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestInterfacePointer(t *testing.T) {
	want := NewString("test")
	got := copy.Deep(&want) // type: *fmt.Stringer

	if !reflect.DeepEqual(got, &want) || got == &want {
		t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), &want, reflect.TypeOf(&want))
	}
}

func TestStructInterfaceNil(t *testing.T) {
	want := struct{ Value fmt.Stringer }{}
	var zero fmt.Stringer
	if got := copy.Deep(want); got.Value != zero {
		t.Errorf("Test %s: got `%#v` (type %v), want zero value `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), want, reflect.TypeOf(want))
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		want map[string]int
	}{
		{
			name: "nil map",
			want: nil,
		},
		{
			name: "empty map",
			want: map[string]int{},
		},
		{
			name: "non-empty map",
			want: map[string]int{"foo": 1, "bar": 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestMapPointer(t *testing.T) {
	tests := []struct {
		name string
		want *map[string]int
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "empty map",
			want: &map[string]int{},
		},
		{
			name: "non-empty map",
			want: &map[string]int{"foo": 1, "bar": 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "nil slice",
			want: nil,
		},
		{
			name: "empty slice",
			want: []int{},
		},
		{
			name: "non-empty slice",
			want: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestSlicePointer(t *testing.T) {
	tests := []struct {
		name string
		want *[]int
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "empty slice",
			want: &[]int{},
		},
		{
			name: "non-empty slice",
			want: &[]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestArray(t *testing.T) {
	tests := []struct {
		name string
		want [3]int
	}{
		{
			name: "empty array",
			want: [3]int{},
		},
		{
			name: "non-empty array",
			want: [3]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestArrayPointer(t *testing.T) {
	tests := []struct {
		name string
		want *[3]int
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "empty array",
			want: &[3]int{},
		},
		{
			name: "non-empty array",
			want: &[3]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestStruct(t *testing.T) {
	type Foo struct {
		Value int
	}

	tests := []struct {
		name string
		want Foo
	}{
		{
			name: "empty struct",
			want: Foo{},
		},
		{
			name: "non-empty struct",
			want: Foo{Value: 42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestStructPointer(t *testing.T) {
	type Foo struct {
		Value int
	}

	tests := []struct {
		name string
		want *Foo
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "empty struct",
			want: &Foo{},
		},
		{
			name: "non-empty struct",
			want: &Foo{Value: 42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := copy.Deep(test.want); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestStructUnexportedFields(t *testing.T) {
	type Foo struct {
		value int
	}

	tests := []struct {
		name string
		want Foo
	}{
		{
			name: "empty struct",
			want: Foo{},
		},
		{
			name: "non-empty struct",
			want: Foo{value: 42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var zero int
			if got := copy.Deep(test.want); got.value != zero {
				t.Errorf("Test %s: got `%#v` (type %v), want zero value `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), test.want, reflect.TypeOf(test.want))
			}
		})
	}
}

func TestChan(t *testing.T) {
	tests := []struct {
		name string
		want chan int
	}{
		{
			name: "nil channel",
			want: nil,
		},
		{
			name: "non-empty channel",
			want: func() chan int {
				c := make(chan int, 3)
				c <- 1
				c <- 2
				c <- 3
				close(c)
				return c
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := copy.Deep(test.want)

			if cap(got) != cap(test.want) || len(got)+cap(got) != len(test.want) {
				t.Errorf("Test %s: got `%#v` (len: %d, cap: %d), want `%#v` (len: %d, cap: %d)", t.Name(), got, len(got), cap(got), test.want, len(test.want), cap(test.want))
			}

			if len(test.want) > 0 {
				i := 0
				for v := range got {
					if i++; v != i {
						t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), v, reflect.TypeOf(v), i, reflect.TypeOf(i))
					}
				}
			}
		})
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		want func(int) string
	}{
		{
			name: "nil func",
			want: nil,
		},
		{
			name: "func",
			want: strconv.Itoa,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := copy.Deep(test.want)
			if reflect.ValueOf(got).IsNil() != reflect.ValueOf(test.want).IsNil() {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), reflect.ValueOf(got), reflect.TypeOf(got), reflect.ValueOf(test.want), reflect.TypeOf(test.want))
			}
			if test.want != nil {
				for _, arg := range []int{1, 42, 100} {
					if !reflect.DeepEqual(got(arg), test.want(arg)) {
						t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), reflect.ValueOf(got), reflect.TypeOf(got), reflect.ValueOf(test.want), reflect.TypeOf(test.want))
					}
				}
			}
		})
	}
}

func TestFuncVariadic(t *testing.T) {
	want := func(x int, y ...int) int {
		sum := x
		for _, v := range y {
			sum += v
		}
		return sum
	}
	got := copy.Deep(want)
	tests := [][]int{{1, 2, 3}, {1, 10, 100}, {42, 69, 1337}}
	for _, test := range tests {
		if !reflect.DeepEqual(got(test[0], test[1:]...), want(test[0], test[1:]...)) {
			t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), reflect.ValueOf(got), reflect.TypeOf(got), reflect.ValueOf(want), reflect.TypeOf(want))
		}
	}
}

func TestPrimitive(t *testing.T) {
	tests := []any{true, 1, int8(1<<7 - 1), int16(1<<15 - 1), int32(1<<31 - 1), int64(1<<63 - 1), uint(1), uint8(1<<8 - 1), uint16(1<<16 - 1), uint32(1<<32 - 1), uint64(1<<64 - 1), float32(1.0), float64(1.0), complex64(1.0), complex128(1.0), "foo"}

	for _, want := range tests {
		t.Run(reflect.TypeOf(want).String(), func(t *testing.T) {
			if got := copy.Deep(want); got != want {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), want, reflect.TypeOf(want))
			}
		})
	}
}

func TestPrimitivePointer(t *testing.T) {
	var (
		b    bool
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		f32  float32
		f64  float64
		c64  complex64
		c128 complex128
		s    string
	)

	tests := []any{&b, &i, &i8, &i16, &i32, &i64, &ui, &ui8, &ui16, &ui32, &ui64, &f32, &f64, &c64, &c128, &s}

	for _, want := range tests {
		t.Run(reflect.TypeOf(want).String(), func(t *testing.T) {
			if got := copy.Deep(want); !reflect.DeepEqual(got, want) || got == want {
				t.Errorf("Test %s: got `%#v` (type %v), want `%#v` (type %v)", t.Name(), got, reflect.TypeOf(got), want, reflect.TypeOf(want))
			}
		})
	}
}
