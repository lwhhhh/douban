package main

import (
    "douban"
    "fmt"
    "log"
)

func main() {
    movieList, err := douban.SerachMovie("硅谷", true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(movieList)
    fmt.Println()
    movieInfo, err := douban.MovieInfo(movieList.Subjects[0].Id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(movieInfo)
}