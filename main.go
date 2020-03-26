package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	//"github.com/nareix/joy4"
)

var hl string

func main() {
	flag.StringVar(&hl,"hl","0.0.0.0:6768","http server监听地址")
	flag.Parse()



	http.HandleFunc("/screen", func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type","image/png")
		cmd:=exec.Command("adb","shell","screencap","-p")
		cmd.Stdout=writer
		cmd.Run()
	})

	http.HandleFunc("/input/key", func(writer http.ResponseWriter, request *http.Request) {
		key:=request.URL.Query().Get("key")
		cmd:=exec.Command("adb","shell","input","keyevent", key)
		cmd.Stdout=writer
		cmd.Run()
	})

	http.HandleFunc("/input/swipe", func(writer http.ResponseWriter, request *http.Request) {
		x1:=request.URL.Query().Get("x1")
		y1:=request.URL.Query().Get("y1")
		x2:=request.URL.Query().Get("x2")
		y2:=request.URL.Query().Get("y2")
		cmd:=exec.Command("adb","shell","input","swipe", x1,y1,x2,y2)
		cmd.Stdout=writer
		cmd.Run()
	})

	http.HandleFunc("/input/tap", func(writer http.ResponseWriter, request *http.Request) {
		x1:=request.URL.Query().Get("x1")
		y1:=request.URL.Query().Get("y1")
		cmd:=exec.Command("adb","shell","input","tap", x1,y1)
		cmd.Stdout=writer
		cmd.Run()
	})

	http.HandleFunc("/on-off", func(writer http.ResponseWriter, request *http.Request) {
		cmd:=exec.Command("adb","shell","input","keyevent", "26")
		cmd.Stdout=writer
		cmd.Run()
	})

	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))

	fmt.Println("服务启动，请访问：",fmt.Sprintf("http://%s",hl))
	http.ListenAndServe(hl,nil)
}
