package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
	fmt.Fprintln(w, "hello cloudOS")
}

func main() {
	fmt.Println("server start....")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("server start err:" + err.Error())
	}
	fmt.Println("server start success!!!")
}
