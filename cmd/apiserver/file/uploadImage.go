package file

import (
	"io/ioutil"
	"net/http"
)

func UploadImage(response http.ResponseWriter, request *http.Request) {

	// Limit 10 MB
	_ = request.ParseMultipartForm(10 * 1024 * 1024)

	file, _, err := request.FormFile("image")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	//upload Image
	tempImage, err := ioutil.TempFile("images", "image-*.jpg")
	if err != nil {
		panic(err)
	}
	defer tempImage.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	tempImage.Write(fileBytes)
}
