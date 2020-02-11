package main

import (
	"log"
)

func (e *ego) Check() {
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
	log.Println("Check layouts")
	err = isExist("./templates/layouts", true)
	if err != nil {
		return err
	}

	// 检查 data
	log.Println("Check data")
	err = isExist("./data/project", true)
	if err != nil {
		return err
	}
	return nil
}
