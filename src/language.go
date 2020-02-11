package main

import (
	"fmt"
	"path"
)

type Language struct {
	List map[string][]*Project
}

func (l *Language) AddProject(project Project) {
	if l.List == nil {
		l.List = make(map[string][]*Project)
	}
	for _, lang := range project.Languages {
		if l.List[lang] == nil {
			l.List[lang] = make([]*Project, 0)
		}
		l.List[lang] = append(l.List[lang], &project)
	}
}

func (r *Render) RenderLanguage() error {
	for lang, projects := range r.Ego.Language.List{
		if err := r.renderLanguage(lang, projects); err != nil{
			return err
		}
	}
	return nil
}

func (r *Render) renderLanguage(lang string, projects []*Project) error {
	langPage := r.NewPage(lang, "language.html", nil)
	langPage.Title = fmt.Sprintf("%s 语言项目 - %s", lang, r.Ego.Config.Site.Title)

	langPage.Params["Language"] = lang
	langPage.Params["Projects"] = projects

	_, err := langPage.Render()
	if err != nil{
		return err
	}

	// 项目首页 /project/Cube/index.html
	langPage.URL = path.Join("/language/", lang, "/index.html")
	return langPage.Save()
}
