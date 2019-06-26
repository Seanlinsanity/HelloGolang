package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.indiegogo.com/projects/cubo-ai-world-s-smartest-baby-monitor?secret_perk_token=d15cf19e#/", nil)
	// request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone)")
	// res, err := http.Get("https://www.youtube.com")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	// res, err := http.DefaultClient.Do(request)
	res, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	s, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dump respones: %s\n", s)
}
