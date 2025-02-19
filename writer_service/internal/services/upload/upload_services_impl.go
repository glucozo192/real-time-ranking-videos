package upload

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/services"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type uploadServices struct {
	log logger.Logger
	cfg *config.Config
}

func NewUploadService(
	log logger.Logger,
	cfg *config.Config,
) services.IUploadServices {
	return &uploadServices{
		log: log,
		cfg: cfg,
	}
}

func (u *uploadServices) UploadFileMedia(ctx context.Context, pathFile string, fileName string, folderUpload string, description string, overwrite bool) (string, error) {
	url := fmt.Sprintf("%sapi/upload", u.cfg.ExternalService.Media)
	data := map[string]string{
		"folder_path": folderUpload,
	}
	if len(description) > 0 {
		data["description"] = description
	}
	if overwrite {
		data["overwrite"] = "2"
		data["file_overwrite_path"] = fmt.Sprintf("%s/%s", folderUpload, fileName)
	} else {
		data["file_name"] = fileName
	}

	type ResponseSuccess struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	client := resty.New()
	resp, err := client.R().SetFile("file", pathFile).SetFormData(data).
		SetHeader("Token", u.cfg.App.TokenToServer).
		SetError(models.HttpError{}).
		SetResult(&ResponseSuccess{}).
		Post(url)
	if err != nil {
		return "", err
	}
	response := string(resp.Body())
	if val, ok := resp.Result().(*ResponseSuccess); ok {
		if val.Status == "success" {
			return response, nil
		} else {
			return response, errors.New(val.Message)
		}
	}
	return response, errors.New("request not valid")
}

func (u *uploadServices) DownloadAndUploadFile(ctx context.Context, url string, fileNameNew string, pathDownload string, folderUpload string) (string, error) {
	// downloadUrl, err := utils.GetGGDriveDownloadUrl(url)
	// if err != nil {
	// 	return "", err
	// }
	err := utils.DownloadFileFromUrlV2(ctx, url, pathDownload, fileNameNew)
	if err != nil {
		return "", err
	}

	fileOpen, err := os.Open(pathDownload + "/" + fileNameNew)
	if err != nil {
		return "", err
	}
	if fileNameNew == "" {
		fileNameNew = uuid.NewString() + path.Ext(url)
	}
	pathFile, err := u.UploadFileAndGetFileNameV2(ctx, fileNameNew, fileOpen, folderUpload)
	if err != nil {
		return "", err
	}
	return pathFile, nil
}

type responseUploadResource struct {
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
}

var ErrUploadError = errors.New("upload error")

func (u *uploadServices) UploadFileAndGetFileNameV2(ctx context.Context, fileName string, r io.Reader, path string) (string, error) {
	client := resty.New()
	var errResponse interface{}
	var response responseUploadResource

	_, err := client.R().
		SetFileReader("file", fileName, r).
		SetFormData(map[string]string{"folder_path": path}).
		SetHeader("Token", u.cfg.App.TokenToServer).
		SetError(&errResponse).
		SetResult(&response).
		Post(fmt.Sprintf("%s/api/upload", u.cfg.ExternalService.Media))
	if err != nil {
		return "", err
	}
	return response.Data.Name, nil
}
