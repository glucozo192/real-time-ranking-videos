package services

import (
	"context"
)

type IUploadServices interface {
	UploadFileMedia(ctx context.Context, pathFile string, fileName string, folderUpload string, description string, overwrite bool) (string, error)
	DownloadAndUploadFile(ctx context.Context, url string, fileNameNew string, pathDownload string, folderUpload string) (string, error)
}
