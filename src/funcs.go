package main

import "html/template"

func unescaped(str string) template.HTML {
	return template.HTML(str)
}
