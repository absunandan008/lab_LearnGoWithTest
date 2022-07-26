package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	//fn("I still can't believe that this is a thing")
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
