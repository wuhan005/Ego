package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Page struct {
	Title       string
	RawMarkdown []byte
	FileName    string
	Layouts     []string               // Render 预制 Layouts
	Params      map[string]interface{} // 绑定变量
	Content     []byte
	URL         string
	OutputName  string
}

func (r *Render) NewPage(fileName string, markDownPath string) Page {
	if markDownPath != "" {
		// TODO 读取 Markdown
	}

	return Page{
		Title:       "",
		RawMarkdown: nil,
		FileName:    fileName,
		Layouts:     r.LayoutsFile,
		Params:      map[string]interface{}{},
		Content:     []byte{}, // 文件内容
		URL:         "",
		OutputName:  "",
	}
}

func (p *Page) Render() ([]byte, error) {
	tpl, err := template.ParseFiles(append([]string{"./templates/"+p.FileName}, p.Layouts...)...)
	if err != nil{
		return nil, err
	}
	var wr bytes.Buffer
	p.Params["Title"] = p.Title // 设置标题
	err = tpl.Execute(&wr, p.Params)
	if err != nil {
		return nil, err
	}
	p.Content = wr.Bytes()
	return p.Content, nil
}

func (p *Page) Save() error {
	// 解析 URL
	path, name := filepath.Split(p.URL)
	p.OutputName = name
	path, err := filepath.Abs("./public/" + path) // 转换为绝对路径
	if err != nil {
		return err
	}
	if _, err = os.Stat(path); err != nil {
		err = os.MkdirAll(path, os.ModePerm) // 创建文件夹
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filepath.Join(path, name), p.Content, 0644)
}
