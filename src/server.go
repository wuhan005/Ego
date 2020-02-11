package main

import (
	"log"
	"net/http"
)

func (e *ego) DoServer(port string){
	if port == ""{
		port = ":8000"
	}
	if port[0] != ':'{
		port = ":" + port
	}

	http.Handle("/", http.FileServer(http.Dir("./public/")))
	log.Println("HTTP server runs on ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatalln(err)
	}
}