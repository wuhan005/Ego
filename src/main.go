package main

import (
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

	ego.Render = ego.NewRender()
	ego.Render.Init()
	ego.Render.RenderLanguage()
	ego.Render.RenderProjects() // 渲染项目
	ego.Render.RenderAbout()	// 关于页面
	ego.Render.RenderIndex() // 渲染主页
}
