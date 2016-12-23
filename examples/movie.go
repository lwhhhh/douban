package main

import (
    "douban"
    "fmt"
)

func main() {
    douban.SerachMovie("硅谷")
    movie := douban.MovieInfo("1764796")
    fmt.Println(movie)
}