package main

type ego struct {
	Config *config
	Render *Render
}

type config struct {
	Profile profile `toml:"profile"`
}

type profile struct {
	NickName string `toml:"nickname"`
	Site     string `toml:"site"`
	Avatar   string `toml:"avatar"`
	Intro    string `toml:"intro"`
}
