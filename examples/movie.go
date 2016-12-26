package main

import (
	"douban"
	"fmt"
	"log"
)

func main() {
	movieList, err := douban.SerachMovie("公司的力量", true)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", movieList)
	// fmt.Println()
	movieInfo, err := douban.MovieInfo(movieList.Subjects[0].Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movieInfo.Genres)
	fmt.Println()
	
}
