package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not nginx!"))
}

func main() {
	str := "web serv"
	num := 8081
	res := fmt.Sprintf("Starting %s on port %d", str, num)
	port := "localhost:8081"
	fmt.Println(res)
	http.HandleFunc("/", Handler)
	log.Printf("Start websrv %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
