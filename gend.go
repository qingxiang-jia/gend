package main

import (
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir("C:/Users/ybjqx/Music/"))
	http.Handle("/", http.StripPrefix("/", handler))
	http.ListenAndServe(":8081", nil)
}
