package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"sort"
	"strings"
)

type Project struct {
	PathName   string
	Content    string              // 正文
	HistoryKey []string            // 更新历史 Key，用于解决 map 无序
	History    map[string][]string // 更新历史

	Meta
}

type Meta struct {
	Name        string   `yaml:"name"`
	Link        string   `yaml:"link"`
	Slogan      string   `yaml:"slogan"`
	Status      int      `yaml:"status"`
	Try         string   `yaml:"try"`
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

	// 初始化项目编程语言信息
	for _, project := range e.Projects {
		e.Language.AddProject(project)
	}

	return nil
}

func NewProject(pathName string) Project {
	p := new(Project)
	p.PathName = pathName
	p.ParseMeta()    // 解析元数据
	p.ParseHistory() // 解析更新历史
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

func (p *Project) ParseHistory() {
	content, err := readFile(path.Join("./data/project/", p.PathName, "history.yml"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &p.History)
	if err != nil {
		panic(err)
	}
	p.HistoryKey = make([]string, 0)
	for k := range p.History {
		p.HistoryKey = append(p.HistoryKey, k)
	}
	sort.Strings(p.HistoryKey)
	// 反转
	for i, j := 0, len(p.HistoryKey)-1; i < j; i, j = i+1, j-1 {
		p.HistoryKey[i], p.HistoryKey[j] = p.HistoryKey[j], p.HistoryKey[i]
	}
}
