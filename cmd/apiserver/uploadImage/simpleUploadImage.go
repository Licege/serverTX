package uploadImage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func SimpleUploadImage(c *gin.Context)  {
	// извлекаем id продукта и создаем директорию
	productId, _ := c.GetPostForm("product_id")
	{
		if len(productId) == 0 {
			productId = "default"
		}
		err := os.MkdirAll(fmt.Sprintf("%s/product/%s", IMAGE_DIR, productId), os.ModePerm)
		if err != nil {
			httpError(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.MkdirAll error: %s", err))
			return
		}
	}
	path := fmt.Sprintf("%s/product/%s", IMAGE_DIR, productId)

	// извлекаем файл из парамeтров post запроса
	form, _ = c.MultipartForm()
	var fileName string
	// берем первое имя файла из присланного списка
	imgExt := "jpeg"
	for key := range form.File {
		fileName = key
		// извлекаем расширение файла
		arr := strings.Split(fileName, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr) - 1]
		}
		continue
	}
	//извлекаем содержание присланного файла по названию файла
	file, _, err := c.Request.FormFile(fileName)
	if err != nil {
		httpError(c, http.StatusBadRequest, fmt.Sprintf("UploadXml c.Request.FormFile error %s", err.Error()))
		return
	}
	defer file.Close()

	//Читаем содержание присланного файла в []byte
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}

	fullFileName := fmt.Sprintf("%s/%s", randomFileName(), imgExt)
	//Открываем файл для сохранения картинки
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	if err != nil {
		httpError(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.Create err: %s", err))
		return
	}
	defer fileOnDisk.Close()

	_, err = fileOnDisk.Write(fileBytes)
	if err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}

	//возвращаем ссылку на файл
	httpSuccess(c, map[string]string{"file": fmt.Sprintf("%s/%s", strings.Replace(path, IMAGE_DIR, STAT_IMAGE_PATH, 1), fullFileName)})
}