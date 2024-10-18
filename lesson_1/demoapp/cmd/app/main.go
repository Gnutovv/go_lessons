package main

import (
	"github.com/Gnutovv/go_lessons/lesson_1/demoapp/pkg/string_utils"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			s := req.URL.Query()["s"]
			rev := string_utils.RevertStr(s[0])
			w.Write([]byte(rev))
		}))
}
