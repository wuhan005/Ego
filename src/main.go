package main

import "github.com/spf13/afero"

var fs afero.Fs
func main() {
	fs = afero.NewOsFs()
	stw := new(stw)

	stw.check()
	stw.loadConfig()

	stw.Render = new(Render)
	stw.Render.GlobalConfig = stw.Config
	stw.Render.Init()

	//tmpl, err := template.ParseFiles( "./templates/index.html")
	//tmpl, err = tmpl.ParseFiles("./templates/layouts/header.html")
	//tmpl, err = tmpl.ParseFiles("./templates/layouts/include.html")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//var wr bytes.Buffer
	//err = tmpl.Execute(&wr, map[string]interface{}{
	//	"Title": "Hello World!!",
	//})
	//fmt.Println(err)
	//fmt.Println(wr.String())
}
