package main

import "github.com/BurntSushi/toml"

func (s *stw) LoadConfig() {
	c := new(config)
	_, err := toml.DecodeFile("./data/config.toml", &c)
	if err != nil {
		panic(err)
	}
	s.Config = c
}
