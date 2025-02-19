package usecase

import (
	"context"
	"fmt"
	"path"
	"strings"

	storage_service "github.com/glu/video-real-time-ranking/core/pkg/storage"
	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
)

type SaveResourceVideo interface {
	SaveResource(ctx context.Context, folderStorage string, diskStorage string) error
}

type DownloadVideoItemUrlDownload struct{}

func (u *DownloadVideoItemUrlDownload) UrlDownload(isCDN int) string {
	linkDownload := config.GetInstance().GetMediaS3() + "/upload/cms_platform/video_k5_test"
	if isCDN == 1 {
		linkDownload = config.GetInstance().GetMediaDisplay() + "/platform/uploads"
	}
	if isCDN == 2 {
		linkDownload = "https://monkeymediadev.s3.ap-southeast-1.amazonaws.com" + "/upload/cms_platform"
	}
	return linkDownload
}

type DownloadResourceVideo struct {
	DownloadVideoItemUrlDownload
}

func (u *DownloadResourceVideo) DownloadDefault(ctx context.Context, folderStorage string, filePath string, isCDN int) error {
	fileName := path.Base(filePath)
	linkDownload := fmt.Sprintf("%s/%s", u.UrlDownload(isCDN), filePath)
	err := utils.DownloadFileFromUrl(ctx, linkDownload, folderStorage, fileName)
	if err != nil {
		if isCDN < 2 {
			return u.DownloadDefault(ctx, folderStorage, filePath, isCDN+1)
		}
		return err
	}
	return nil
}

type SaveResourceCommentImage struct {
	DownloadResourceVideo
	Folder   string
	FilePath string
	Format   string
	Image    models.CommentItem
}

func (u *SaveResourceCommentImage) SetData() {
	u.Folder = models.FolderImage
	u.FilePath = u.Image.Avatar
}

func (u *SaveResourceCommentImage) SaveResource(ctx context.Context, folderStorage string, diskStorage string) error {
	u.SetData()
	if strings.Contains(u.FilePath, "avatar/avatar") {
		return nil
	}
	fileName := path.Base(u.FilePath)
	folderStorage = fmt.Sprintf("%s/%s", folderStorage, u.Folder)
	filePath := fmt.Sprintf("%s", fileName)
	return u.DownloadDefault(ctx, folderStorage, filePath, 0)
}

type SaveResourceVideoItem struct {
	DownloadResourceVideo
	Folder   string
	FilePath string
	Video    models.VideoItem
}

func (u *SaveResourceVideoItem) SetData() {
	u.Folder = models.FolderResource
	u.FilePath = u.Video.VideoURL
}

func (u *SaveResourceVideoItem) SaveResource(ctx context.Context, folderStorage string, diskStorage string) error {
	u.SetData()
	fileName := path.Base(u.FilePath)
	folderStorage = fmt.Sprintf("%s/%s", folderStorage, u.Folder)
	filePath := fmt.Sprintf("%s", fileName)
	return u.DownloadDefault(ctx, folderStorage, filePath, 0)
}

type SaveResourceVideoItemJson struct {
	FileName string
	DataFile []byte
}

func (u *SaveResourceVideoItemJson) SaveResource(ctx context.Context, folderStorage, diskStorage string) error {
	if errPut := storage_service.Put(ctx, diskStorage, folderStorage, u.FileName, u.DataFile); errPut != nil {
		return errPut
	}
	return nil
}
