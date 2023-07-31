package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("input-image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// open file
		src, errSrc := file.Open()
		if errSrc != nil {
			return c.JSON(http.StatusBadRequest, errSrc.Error())
		}
		
		defer src.Close()

		tmpFile, errTmp := ioutil.TempFile("uploads", "image-*.png")
		if errTmp != nil {
			return c.JSON(http.StatusBadRequest, errTmp.Error())
		}
		
		defer tmpFile.Close()

		writenCopy, errWriten := io.Copy(tmpFile, src)

		if errWriten != nil {
			c.JSON(http.StatusBadRequest, errWriten.Error())
		}

		fmt.Println("written copy :", writenCopy)

		data := tmpFile.Name()
		fileName := data[8:]

		c.Set("dataFile", fileName) // Image-12342342.png

		return next(c)
	}
}