package douban

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	RemoteAddr     = "https://api.douban.com"
	SerachMovieAPI = "/v2/movie/search?q="
	MovieInfoAPI   = "/v2/movie/subject/"
)

type Movie struct {
	Rating        map[string]interface{}   `json:"rating"`
	Genres        []string                 `json:"genres"`
	Title         string                   `json:"title"`
	Casts         []map[string]interface{} `json:"casts"`
	CollectCount  int                      `json:"collect_count"`
	OriginalTitle string                   `json:"original_title"`
	Subtype       string                   `json:"subtype"`
	Directors     []map[string]interface{} `json:"directors"`
	Year          string                   `json:"year"`
	Images        map[string]string        `json:"images"`
	Alt           string                   `json:"alt"`
	Id            string                   `json:"id"`
	Summary       string                   `json:"summary"`
	ReviewsCount  int                      `json:"reviews_count"`
	WishCount     int                      `json:"wish_count"`
	DoubanSize    string                   `json"douban_size"`
	MobileUrl     string                   `json:"mobile_url"`
	DoCount       int                      `json:"do_count"`
	ShareUrl      string                   `json:"share_url"`
	ScheduleUrl   string                   `json:"schedule_url"`
}

type MovieList struct {
	Subjects []Movie `json:"subjects"`
}

func SerachMovie(name string, strict bool) (*MovieList, error) {
	url := RemoteAddr + SerachMovieAPI + name
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	movieList := MovieList{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &movieList)
	if !strict {
		return &movieList, nil
	}
	newMovieList := MovieList{[]Movie{}}
	for _, mv := range movieList.Subjects {
		if mv.Title == name {
			newMovieList.Subjects = append(newMovieList.Subjects, mv)
		}
	}
	return &newMovieList, nil
}

func MovieInfo(objId string) (*Movie, error) {
	url := RemoteAddr + MovieInfoAPI + objId
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	movie := Movie{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &movie)
	return &movie, nil
}
