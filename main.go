package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type remoteInfo struct {
	RemoteHostName string
	RemoteIp       string
	Time           string
	RequestUrl     string
}

func index(w http.ResponseWriter, r *http.Request) {
	info := "收到客户端请求, 请求路径为 " + r.RequestURI
	log.Print(info)
	if r.Method != "GET" {
		return
	}
	var hostName string
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "无法获取hostName"
	}

	podIp := os.Getenv("POD_IP")

	remoteInfo := remoteInfo{
		RemoteHostName: hostName,
		RemoteIp:       podIp,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		RequestUrl:     r.RequestURI,
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print(err.Error())
	}
	_ = t.Execute(w, remoteInfo)
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
