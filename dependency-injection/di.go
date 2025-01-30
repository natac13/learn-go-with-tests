package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Chris")

	if err := http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

	fmt.Println("Running on port 5001")
}
