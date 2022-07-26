package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	//fn("I still can't believe that this is a thing")
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
