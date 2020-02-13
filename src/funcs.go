package main

import (
	"html/template"
	"net/url"
	"reflect"
	"strings"
)

func (r *Render) LoadFunctions() error {
	r.FunctionMaps = template.FuncMap{
		"unescaped":   unescaped,
		"count":       count,
		"urlencode":   urlencode,
		"symbol2text": symbol2text,
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

func symbol2text(str string) string {
	dict := map[string]string{
		"#": "Sharp",
		"+": "Plus",
		"=": "Equal",
		"-": "Minus",
	}
	for k, v := range dict {
		str = strings.ReplaceAll(str, k, "_"+v)
	}
	return str
}
