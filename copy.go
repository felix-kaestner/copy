package copy

import "reflect"

// Deep returns a deep copy of v.
//
// Deep will not copy unexported struct fields (lowercase field names).
func Deep[T any](v T) T {
	if rv := reflect.ValueOf(v); rv.IsValid() {
		return deep(rv).Interface().(T)
	}
	return v
}

func deep(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Interface:
		// If the value is an interface, create a deep copy of the interface's value.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		return deep(v.Elem()).Convert(v.Type())
	case reflect.Pointer:
		// If the value is a pointer, create a deep copy of the value the pointer references.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		v = v.Elem()
		pv := reflect.New(v.Type())
		pv.Elem().Set(deep(v))
		return pv
	case reflect.Map:
		// If the value is a map, create a new map with the same size and deep copy all entries.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		vt := v.Type()
		mt := reflect.MapOf(vt.Key(), vt.Elem())
		mv := reflect.MakeMapWithSize(mt, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			mv.SetMapIndex(deep(iter.Key()), deep(iter.Value()))
		}
		return mv
	case reflect.Slice:
		// If the value is a slice, create a new slice with the same size and capacity and deep copy all values.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		sv := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			sv.Index(i).Set(deep(v.Index(i)))
		}
		return sv
	case reflect.Array:
		// If the value is an array, create a new array and deep copy all values.
		av := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			av.Index(i).Set(deep(v.Index(i)))
		}
		return av
	case reflect.Struct:
		// If the value is a struct, create a new struct and deep copy all fields.
		sv := reflect.New(v.Type()).Elem()
		for i := 0; i < v.NumField(); i++ {
			if sf := sv.Field(i); sf.CanSet() {
				sf.Set(deep(v.Field(i)))
			}
		}
		return sv
	case reflect.Chan:
		// If the value is a channel, create a new channel with the same capacity and send deep copies of all values until the channel is closed.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		cv := reflect.MakeChan(v.Type(), v.Cap())
		go func() {
			for {
				vv, ok := v.Recv()
				if !ok {
					break
				}
				cv.Send(deep(vv))
			}
			cv.Close()
		}()
		return cv
	case reflect.Func:
		// If the value is a function, create a new function which takes in the same arguments and delegates the call.
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		fn := v.Call
		vt := v.Type()
		if vt.IsVariadic() {
			fn = v.CallSlice
		}
		return reflect.MakeFunc(vt, func(args []reflect.Value) []reflect.Value { return fn(args) })
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String, reflect.Uintptr:
		// Check if v is not the zero value and the value can be read (exported).
		if v.CanInterface() {
			// Create a new instace of the type and set it to the value we want to copy.
			pv := reflect.New(v.Type()).Elem()
			pv.Set(v)
			return pv
		}
		// Just return the zero value for the primitive.
		return reflect.Zero(v.Type())
	case reflect.Invalid:
		// If the value is invalid, e.g. an uninitialized interface variable, just return it as is.
		// In this case v.IsValid() == false and all other methods on v except String panic.
		return v
	default:
		panic("unsupported type: " + v.String())
	}
}
