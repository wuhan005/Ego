package main

import (
	"html/template"
	"net/url"
	"reflect"
)

func (r *Render) LoadFunctions() error {
	r.FunctionMaps = template.FuncMap{
		"unescaped": unescaped,
		"count":     count,
		"urlencode": urlencode,
	}
	return nil
}

func unescaped(str string) template.HTML {
	return template.HTML(str)
}

func count(data interface{}) int {
	return reflect.ValueOf(data).Len()
}

func urlencode(str string) string {
	return url.QueryEscape(str)
}
