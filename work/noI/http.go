package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func param(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//req.Header中Header本质是:type Header map[string][]string
	header := req.Header
	fmt.Fprintln(res, "Header全部数据:", header)
	for h := range header {
		fmt.Fprintln(res, h, header[h])
		res.Header().Set(h, strings.Join(header[h], `,`))
	}
	res.Header().Set("cjx", "name")
	res.WriteHeader(200)

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/param", param)

	log.Fatal(http.ListenAndServe(":8080", router))
}
