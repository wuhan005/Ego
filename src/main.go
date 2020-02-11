package main

import (
	"github.com/spf13/afero"
	"log"
)

var fs afero.Fs

func main() {
	fs = afero.NewOsFs()
	ego := new(ego)
	ego.Language = new(Language)

	ego.Check()      // 检查文件完整性
	ego.LoadConfig() // 加载配置
	ego.DoCLI()
}

func (e *ego) DoRender() {
	err := e.LoadProject() // 加载项目配置
	if err != nil {
		log.Fatalln(err)
	}
	e.Render = e.NewRender()
	e.Render.Init()
	err = e.Render.RenderLanguage()
	if err != nil {
		log.Fatalln(err)
	}
	err = e.Render.RenderProjects() // 渲染项目
	if err != nil {
		log.Fatalln(err)
	}
	err = e.Render.RenderAbout() // 关于页面
	if err != nil {
		log.Fatalln(err)
	}
	err = e.Render.RenderIndex() // 渲染主页
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done! Enjoy it!")
}
