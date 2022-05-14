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
	str := "String"
	num := 42
	res := fmt.Sprintf("return %s +  %d", str, num)
	port := ":8081"
	fmt.Print(res)
	http.HandleFunc("/", Handler)
	log.Printf("Start websrv %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
