package main

import (
    "fmt"
    "go-postgres/routers"
    "log"
    "net/http"
)

func main() {
    r := routers.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
    fmt.Println("Starting server on the port 3000...")

    log.Fatal(http.ListenAndServe(":3000", r))
}