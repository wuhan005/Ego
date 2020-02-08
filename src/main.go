package main

import "github.com/spf13/afero"

var fs afero.Fs
func main() {
	fs = afero.NewOsFs()
	ego := new(ego)

	ego.check()
	ego.LoadConfig()

	ego.LoadProject()
	ego.Render = new(Render)
	ego.Render.GlobalConfig = ego.Config
	ego.Render.Init()
}
