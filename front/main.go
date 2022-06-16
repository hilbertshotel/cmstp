package main

import (
	"fmt"
	"net/http"
)

const HostAddr = "127.0.0.1:8888"

func main() {
	front := http.FileServer(http.Dir("./front"))
	http.Handle("/", front)
	fmt.Println("Server listening:", HostAddr)
	http.ListenAndServe(HostAddr, nil)
}
