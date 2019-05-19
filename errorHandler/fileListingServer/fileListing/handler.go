package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

//HandleFileList function
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	absPath, _ := filepath.Abs("../" + path)
	file, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
