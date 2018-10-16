// +build ignore

// Copyright 2015,2016,2017,2018 SeukWon Kang (kasworld@gmail.com)

package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", ":8080", "Serve port")
	folder := flag.String("dir", ".", "Serve Dir")
	fmt.Printf("dir2http dir=%v port=%v , http://localhost%v/ \n",
		*folder, *port, *port)

	webMux := http.NewServeMux()
	webMux.Handle("/",
		http.FileServer(http.Dir(*folder)),
	)
	if err := http.ListenAndServe(*port, webMux); err != nil {
		fmt.Println(err.Error())
	}
}
