package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"strings"
)

type Projects []Project

type Project struct {
	PathName    string
	Name        string   `yaml:"name"`
	Link        string   `yaml:"link"`
	Slogan      string   `yaml:"slogan"`
	Description string   `yaml:"description"`
	Languages   []string `yaml:"language"`
	Tags        []string `yaml:"tags"`
	Progress    struct {
		Version string `yaml:"version"`
		Percent uint   `yaml:"percent"`
	} `yaml:"progress"`
	Content string       // 正文
}

func (s *stw) LoadProject() error {
	projectList, err := ioutil.ReadDir("./data/project")
	if err != nil {
		return err
	}

	for _, project := range projectList {
		NewProject(project.Name())
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
	err = yaml.Unmarshal([]byte(meta), &p)
	if err != nil {
		panic(err)
	}
}
