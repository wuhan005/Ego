package main

import (
	"fmt"
	"github.com/spf13/afero"
)

var fs afero.Fs

func main() {
	fs = afero.NewOsFs()
	ego := new(ego)

	ego.Check()      // 检查文件完整性
	ego.LoadConfig() // 加载配置

	ego.LoadProject() // 加载项目配置
	ego.Render = NewRender(ego.Config, ego.Projects)
	ego.Render.Init()
	ego.Render.RenderProjects() // 渲染项目

	fmt.Println(ego.Render.RenderIndex()) // 渲染主页

}
