package main

import (
	"fmt"
	"net/http"

	"ascii-art-web-stylize/functions"
)

func main() {
	http.HandleFunc("/static/", functions.StyleFunc)
	http.HandleFunc("/", functions.HandlerMainFunc)
	http.HandleFunc("/asciiart", functions.HandlerArtFunc)
	http.HandleFunc("/export", functions.HandlerExport)
	fmt.Println("runing server : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
}
