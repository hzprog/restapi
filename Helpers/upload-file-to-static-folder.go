package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(req *http.Request, paramName string) (string, error) {
	req.ParseMultipartForm(10 << 20)

	file, multipartFileHeader, err := req.FormFile("image")
	if err != nil {
		fmt.Println("error while getting the File")
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	extension := filepath.Ext(multipartFileHeader.Filename)
	tempFile, err := ioutil.TempFile("static", "upload-*"+extension)

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}

func DeleteFile(path string) {
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("File deleted %s successfully", path)
}
