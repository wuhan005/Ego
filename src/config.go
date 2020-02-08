package main

import "github.com/BurntSushi/toml"

func (e *ego) LoadConfig() {
	c := new(config)
	_, err := toml.DecodeFile("./data/config.toml", &c)
	if err != nil {
		panic(err)
	}
	e.Config = c
}
