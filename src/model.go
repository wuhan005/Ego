package main

type ego struct {
	Config   *config
	Render   *Render
	Projects []Project // 获取到的所有项目
	Language *Language
}

type config struct {
	Site    site    `toml:"site"`
	Profile profile `toml:"profile"`
}

type site struct {
	Title string `toml:"title"`
}

type profile struct {
	NickName string `toml:"nickname"`
	Site     string `toml:"site"`
	Avatar   string `toml:"avatar"`
	Intro    string `toml:"intro"`
}
