package main

import (
	"log"
)

func (e *ego) check() {
	err := checkTemplates()
	if err != nil {
		panic(err)
	}
}

func checkTemplates() error {
	log.Println("Check templates")
	// 检查 templates
	err := isExist("./templates", true)
	if err != nil {
		return err
	}

	// 检查 layouts
	err = isExist("./templates/layouts", true)
	if err != nil {
		return err
	}
	return nil
}
