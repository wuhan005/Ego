package main

import (
	"html/template"
	"reflect"
)

func unescaped(str string) template.HTML {
	return template.HTML(str)
}

func count(data interface{}) int {
	return reflect.ValueOf(data).Len()
}
