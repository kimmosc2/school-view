package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"school-walker/view/handle"
)

var (
	ssl  bool
	port string
)

func init() {
	flag.StringVar(&port, "port", "8989", "http port,just number")
	flag.BoolVar(&ssl, "SSL", false, "use https schema")
}
func main() {
	flag.Parse()
	log.Printf("pid:%v", os.Getpid())
	log.SetFlags(log.Llongfile)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/leave", handle.ViewLeave)
	http.HandleFunc("/leave_data/save", handle.DataSave)
	if ssl {
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "server.crt", "server.key", nil))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}
}
