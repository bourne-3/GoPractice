package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string

}

func movie() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		{Title: "Game of the Thorn", Year: 1996, Color: false,
			Actors: []string{"tyrion", "snow"}},
	}
	fmt.Println(movies)

	// Marshal

	// 将struct转换为json
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshaling failed : %s", data)
	}
	fmt.Printf("%s \n", data)

	// 为了更加好看
	data2, err := json.MarshalIndent(movies, "", "     ")
	if err != nil {
		log.Fatalf("json marshaling failed : %s", data)
	}
	fmt.Printf("%s \n", data2)

	// 反向
	var titles []struct{ Title string}
	if err := json.Unmarshal(data, &titles);err!=nil{
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}