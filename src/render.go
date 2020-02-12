package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

type Render struct {
	Ego          *ego
	Tmpl         *template.Template
	LayoutsFile  []string
	FunctionMaps template.FuncMap
}

func (e *ego) NewRender() *Render {
	r := new(Render)
	r.Ego = e
	return r
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
	log.Println("Render main page")
	indexPage := r.NewPage("index.html", "", nil)

	indexPage.Title = r.Ego.Config.Site.Title
	indexPage.Params["Projects"] = r.Ego.Projects

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

func (r *Render) MoveAssets() error {
	return CopyDir(fs, "./templates/assets", "./public/assets", func(s string) bool { return true })
}

func (r *Render) RenderProjects() error {
	log.Println("Render project(s) page")
	err := os.MkdirAll("./public/project", os.ModePerm)
	if err != nil {
		return err
	}

	for _, project := range r.Ego.Projects {
		err = r.renderProject(project)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Render) renderProject(project Project) error {
	mainPage := r.NewPage(project.Name, "project.html", []byte(project.Content))

	mainPage.Title = fmt.Sprintf("%s - %s", project.Meta.Name, r.Ego.Config.Site.Title)

	metaType := reflect.TypeOf(project.Meta)
	metaValue := reflect.ValueOf(project.Meta)
	metaNum := metaType.NumField()
	for i := 0; i < metaNum; i++ {
		key := metaType.Field(i).Name
		value := metaValue.Field(i).Interface()
		mainPage.Params[key] = value
	}

	mainPage.Params["History"] = project.History
	mainPage.Params["HistoryKey"] = project.HistoryKey
	mainPage.Params["IntroHTML"] = ParseMarkdown(project.Content)

	_, err := mainPage.Render()
	if err != nil {
		return err
	}
	// 项目首页 /project/Cube/index.html
	mainPage.URL = path.Join("/project/", project.Name, "/index.html")
	return mainPage.Save()
}
