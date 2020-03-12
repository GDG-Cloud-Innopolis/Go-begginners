package main

import (
	"fmt"
	"io"
	"net/http"
)

const html = `<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		html,
	)
}
func main() {
	fmt.Println("Listening on :9000")
	http.HandleFunc("/hello", hello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":9000", nil)
}
