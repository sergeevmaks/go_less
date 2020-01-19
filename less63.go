package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	name := req.URL.Query().Get("name")
	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello World</title>
	</head>
	<body>`)
	io.WriteString(res, "param name="+string(name))
	io.WriteString(res, `
		</body>
</html>`)
}

func main() {

	http.HandleFunc("/", hello)

	http.ListenAndServe(":3333", nil)
}
