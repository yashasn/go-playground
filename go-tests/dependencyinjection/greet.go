package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
Dependency Injection (DI) in Go is a design pattern where you provide a component's dependencies from the outside,
rather than the component creating them itself
*/

func Greet(name string) {
	//here, printf write to stdout. How will you test?
	//Printf internally calls fPrintf which is expects on io.writer interface. Printf passes os.Stdout which implements io.writer interface
	fmt.Printf("Hi, %s", name)
}

func GreetDI(writer io.Writer, name string) {
	//both os.stdout and Buffer implement Write()
	fmt.Fprintf(writer, "Hi, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	GreetDI(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
