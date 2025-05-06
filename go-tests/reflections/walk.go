package reflections

import "reflect"

/*
Type safety in Go i.e that Go has offered us in terms of functions that work with known types, such as string, int and our own types like BankAccount.
If we violate, compiler will throw an error.

What if don't know the type during compile time? --- Use interface{}, don't get confused with interface.
*/
func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	//TODO- Handle pointers, slices, arrays, maps
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}
