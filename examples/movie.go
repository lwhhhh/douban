package main

import (
    "douban"
    "fmt"
    "log"
)

func main() {
    movieList, err := douban.SerachMovie("20", true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(movieList)
}