package main

type ego struct {
	Config   *config
	CLI      *CLI
	Render   *Render
	Projects Projects // 获取到的所有项目
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
	GitHub   string `toml:"github"`
	Site     string `toml:"site"`
	Avatar   string `toml:"avatar"`
	Intro    string `toml:"intro"`
}
