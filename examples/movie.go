package main

import (
    "fmt"
    "log"
)

func main() {
    movieList, err := SerachMovie("硅谷", true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(movieList)
}