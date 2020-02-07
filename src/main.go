package main

import "github.com/spf13/afero"

var fs afero.Fs
func main() {
	fs = afero.NewOsFs()
	stw := new(stw)

	stw.check()
	stw.LoadConfig()

	stw.LoadProject()
	stw.Render = new(Render)
	stw.Render.GlobalConfig = stw.Config
	stw.Render.Init()
}
