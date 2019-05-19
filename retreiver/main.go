package main

import (
	"LearnGo/retreiver/myMock"
	"LearnGo/retreiver/real"
	"fmt"
	"time"
)

//Retriever interface with Get function
type Retriever interface {
	Get(url string) string
}

//Poster interface with Post function
type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.youtube.com")
}

func post(p Poster) {
	p.Post("https://www.google.com", map[string]string{
		"contents": "google",
	})
}

//RetrieverPoster interface which is a combination of Retriever and Poster
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.appple.com"

func session(retrieverPoster RetrieverPoster) string {
	form := make(map[string]string)
	form["contents"] = "apple official url"
	retrieverPoster.Post(url, form)
	return retrieverPoster.Get(url)
}

func main() {
	var r Retriever
	r = &myMock.Retriever{Contents: "This is my mock contents"}
	inspect(r)
	// fmt.Println(download(r))
	r = &real.Retriever{UserAgent: "my agent", TimeOut: time.Minute}
	inspect(r)
	// fmt.Println(download(r))

	//Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*myMock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	retrieverPoster := &myMock.Retriever{Contents: "original contents"}
	fmt.Println("try a session")
	fmt.Println(session(retrieverPoster))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch: ")

	switch v := r.(type) {
	case *myMock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}
