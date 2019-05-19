package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {UserAgent=%s}", r.UserAgent)
}

func (r *Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, error := httputil.DumpResponse(response, true)
	response.Body.Close()

	if error != nil {
		panic(error)
	}
	return string(result)
}
