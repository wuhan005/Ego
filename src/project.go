package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"strings"
)

type Project struct {
	PathName    string
	Content string       // 正文

	Meta
}

type Meta struct {
	Name        string   `yaml:"name"`
	Link        string   `yaml:"link"`
	Slogan      string   `yaml:"slogan"`
	Status      string   `yaml:"status"`
	Logo        string   `yaml:"logo"`
	Description string   `yaml:"description"`
	Languages   []string `yaml:"language"`
	Tags        []string `yaml:"tags"`
	Progress    struct {
		Version string `yaml:"version"`
		Percent uint   `yaml:"percent"`
	} `yaml:"progress"`
}

func (e *ego) LoadProject() error {
	projectList, err := ioutil.ReadDir("./data/project")
	if err != nil {
		return err
	}
	e.Projects = make([]Project, 0)
	for _, project := range projectList {
		e.Projects = append(e.Projects, NewProject(project.Name()))
	}

	return nil
}

func NewProject(pathName string) Project {
	p := new(Project)
	p.PathName = pathName
	p.ParseMeta() // 解析元数据
	return *p
}

func (p *Project) ParseMeta() {
	content, err := readFile(path.Join("./data/project/", p.PathName, p.PathName+".md"))
	if err != nil {
		panic(err)
	}
	metaSlice := strings.Split(string(content), "---\n")
	p.Content = metaSlice[2]

	meta := metaSlice[1]
	err = yaml.Unmarshal([]byte(meta), &p.Meta)
	if err != nil {
		panic(err)
	}
}
