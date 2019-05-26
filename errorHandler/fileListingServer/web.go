package main

import (
	filelisting "LearnGo/errorHandler/filelistingServer/fileListing"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.New(os.Stderr, "【Recover】", log.Ldate|log.Ltime|log.Lshortfile).Printf("recover panic: %v", rec)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.New(os.Stderr, "【Warning】", log.Ldate|log.Ltime|log.Lshortfile).Printf("handle error: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error //Error() string
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":3011", nil)
	if err != nil {
		panic(err)
	}
}
