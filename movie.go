package douban

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
)

const (
    RemoteAddr = "https://api.douban.com"
    SerachMovieAPI = "/v2/movie/search?q="
    MovieInfoAPI = "/v2/movie/subject/"
)

type Movie struct {
    Rating map[string]interface{} `json:"rating"` 
    Genres []string   `json:"genres"`
    Title string    `json:"title"`
    Year string        `json:"year"`
    Id string       `json:"id"`
    Summary string  `json:"summary"`
}

type MovieList struct {
    Subjects []Movie    `json:"subjects"`
}

func SerachMovie(name string) (*MovieList, error) {
    url := RemoteAddr + SerachMovieAPI + name
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    movieList := MovieList{}
    body, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &movieList)
    defer resp.Body.Close()
    return &movieList, nil
}

func MovieInfo(objId string) (*Movie, error) {
    url := RemoteAddr + MovieInfoAPI + objId
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    movie := Movie{}
    body, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &movie)
    defer resp.Body.Close()
    return &movie, nil
}

