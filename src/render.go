package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

type Render struct {
	GlobalConfig *config
	Projects     []Project
	Tmpl         *template.Template
	LayoutsFile  []string
	FunctionMaps template.FuncMap
}

func NewRender(config *config, projects []Project) *Render {
	r := new(Render)
	r.GlobalConfig = config
	r.Projects = projects
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

func (r *Render) LoadFunctions() error {
	r.FunctionMaps = template.FuncMap{
		"unescaped": unescaped,
	}
	return nil
}

func (r *Render) RenderIndex() error {
	indexPage := r.NewPage("index.html", "", nil)

	indexPage.Title = "John's Projects"
	indexPage.Params["Avatar"] = r.GlobalConfig.Profile.Avatar
	indexPage.Params["NickName"] = r.GlobalConfig.Profile.NickName
	indexPage.Params["Site"] = r.GlobalConfig.Profile.Site
	indexPage.Params["Intro"] = r.GlobalConfig.Profile.Intro
	indexPage.Params["Projects"] = r.Projects

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
	err := os.MkdirAll("./public/project", os.ModePerm)
	if err != nil {
		return err
	}

	for _, project := range r.Projects {
		err = r.renderProject(project)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Render) renderProject(project Project) error {
	mainPage := r.NewPage(project.Name, "project.html", []byte(project.Content))

	mainPage.Title = fmt.Sprintf("%s - John's Project", project.Meta.Name)

	metaType := reflect.TypeOf(project.Meta)
	metaValue := reflect.ValueOf(project.Meta)
	metaNum := metaType.NumField()
	for i := 0; i < metaNum; i++ {
		key := metaType.Field(i).Name
		value := metaValue.Field(i).Interface()
		mainPage.Params[key] = value
	}

	_, err := mainPage.Render()
	if err != nil {
		return err
	}
	// 项目首页 /project/Cube/index.html
	mainPage.URL = path.Join("/project/", project.Name, "/index.html")
	fmt.Println(mainPage.URL)
	return mainPage.Save()
}
