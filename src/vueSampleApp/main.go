package main

import (
	"fmt"
	"net/http"

	"vueSampleApp/handler"
)

var addr =  ":3000"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.Index)
	fmt.Printf("[START] server. port: %s\n", addr)
	if err := http.ListenAndServe(addr, handler.Log(router)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}