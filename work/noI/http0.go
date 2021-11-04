package main

import (
	"net/http"
	"strings"
)

func test(w http.ResponseWriter, r *http.Request) {
	//req.Header中Header本质是:type Header map[string][]string
	header := r.Header
	for h := range header {
		// fmt.Fprintln(w, h, header[h])
		w.Header().Set(h, strings.Join(header[h], `,`))
	}
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", test)

	http.ListenAndServe(":8000", nil)
}
