package main

import (
	"fmt"
	"github.com/spf13/afero"
)

var fs afero.Fs

func main() {
	fs = afero.NewOsFs()
	ego := new(ego)
	ego.Language = new(Language)

	ego.Check()      // 检查文件完整性
	ego.LoadConfig() // 加载配置

	ego.LoadProject() // 加载项目配置

	ego.Render = ego.NewRender(ego.Config, ego.Projects)
	ego.Render.Init()
	fmt.Println(ego.Render.RenderLanguage())
	ego.Render.RenderProjects() // 渲染项目

	ego.Render.RenderIndex() // 渲染主页

}
