package controller

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func (app *App) UploadToCloud(file io.Reader, newfileName string) (*uploader.UploadResult, error) {

	ctx := context.Background()
	resp, err := app.CloudService.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: newfileName})

	if err != nil {
		fmt.Println("Couldnt save upload to cloud")
		return nil, err
	}

	fmt.Println("Secure Url")
	fmt.Println(resp.SecureURL)
	fmt.Println("")
	fmt.Println(resp)
	return resp, nil

}
