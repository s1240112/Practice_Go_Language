package main

import (
	"fmt"
	"net/http"
)

type Server struct{}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h := `
	<html><head><title>Hello</title></head><body>
	Hello
	</body></html>
	`
	fmt.Fprint(w, h)
}

func main(){
	http.ListenAndServe(":4000", Server{})
}