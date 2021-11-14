package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(req *http.Request, paramName string) (os.File, error) {
	req.ParseMultipartForm(10 << 20)
	file, multipartFileHeader, err := req.FormFile("image")
	if err != nil {
		fmt.Println("error while getting the File")
		fmt.Println(err)
	}
	defer file.Close()

	extension := filepath.Ext(multipartFileHeader.Filename)
	tempFile, err := ioutil.TempFile("static", "upload-*"+extension)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	return *tempFile, err
}
