package main

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type Render struct {
	Tmpl        *template.Template
	LayoutsFile []string
}

func (r *Render) Init() {
	r.Tmpl = template.New("index.html")
	// 载入 Layouts 模板
	r.LoadLayouts()
	// 渲染主页
	r.RenderIndex()
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

func (r *Render) RenderIndex() error {
	indexPage := r.NewPage("index.html", "")
	//indexPage
	indexPage.Title = "John's Projects"
	_, err := indexPage.Render()
	if err == nil {
		indexPage.URL = "/index.html"
		indexPage.Save()
	}
	return nil
}
