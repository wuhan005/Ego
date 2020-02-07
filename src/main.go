package main

func main() {
	stw := new(stw)

	stw.check()
	stw.loadConfig()

	stw.Render = new(Render)
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
