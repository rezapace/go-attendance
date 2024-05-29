package controller

import (
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func UploadFile(c echo.Context) error {
	var err error

	ctx := appengine.NewContext(c.Request())
	opt := option.WithCredentialsFile(("/app/serviceAccountKey.json"))

	storageClient, err = storage.NewClient(ctx, opt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"error":   true,
		})
	}
	defer src.Close()

	bucket := "platinum-project-backend.appspot.com"

	object := storageClient.Bucket(bucket).Object(file.Filename)
	wc := object.NewWriter(ctx)

	if _, err := io.Copy(wc, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	if err := wc.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	url := "https://storage.googleapis.com/" + bucket + "/" + wc.Attrs().Name

	return c.JSON(http.StatusOK, echo.Map{
		"message": "File uploaded successfully",
		"url":     url,
		"error":   false,
	})
}
