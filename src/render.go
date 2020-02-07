package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Render struct {
	GlobalConfig *config
	Tmpl         *template.Template
	LayoutsFile  []string
	FunctionMaps template.FuncMap
}

func (r *Render) Init() {
	r.Tmpl = template.New("index.html")

	// 载入 Layouts 模板
	r.LoadLayouts()
	// 载入模板函数
	r.LoadFunctions()
	// 删除 public 文件夹的内容
	r.CleanPublic()
	// 复制 Assets 静态文件
	r.MoveAssets()
	// 渲染主页
	fmt.Println(r.RenderIndex())
}

func (r *Render) LoadLayouts() error {
	path := "./templates/layouts"
	r.LayoutsFile = []string{}
	layouts, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range layouts {
		if !file.IsDir() {
			r.LayoutsFile = append(r.LayoutsFile, filepath.Join(path, file.Name()))
		}
	}
	return nil
}

func (r *Render) LoadFunctions() error{
	r.FunctionMaps = template.FuncMap{
		"unescaped": unescaped,
	}
	return nil
}

func (r *Render) RenderIndex() error {
	indexPage := r.NewPage("index.html", "")

	indexPage.Title = "John's Projects"
	indexPage.Params["Avatar"] = r.GlobalConfig.Profile.Avatar
	indexPage.Params["NickName"] = r.GlobalConfig.Profile.NickName
	indexPage.Params["Site"] = r.GlobalConfig.Profile.Site
	indexPage.Params["Intro"] = r.GlobalConfig.Profile.Intro

	_, err := indexPage.Render()
	if err != nil {
		return err
	}
	indexPage.URL = "/index.html"
	return indexPage.Save()
}

func (r *Render) CleanPublic() error {
	err := os.RemoveAll("./public/")
	if err != nil {
		return err
	}
	return os.Mkdir("./public/", os.ModePerm)
}

func (r *Render) MoveAssets() error{
	return CopyDir(fs, "./templates/assets", "./public/assets", func(s string) bool { return true })
}