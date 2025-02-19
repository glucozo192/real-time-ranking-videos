package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (u *videoUsecase) ZipById(ctx context.Context, videoId uint, resource *models.VideoItem, versionNew int) error {
	var version int
	var item models.VideoItem
	var videoInfo *ent.Videos
	var folderWord, fileNameZip string
	gData, gDataCtx := errgroup.WithContext(ctx)
	// get version word
	gData.Go(func() error {
		if versionNew == 0 {
			versionMax, err := u.GetMaxVersion(gDataCtx)
			if err != nil {
				return err
			}
			version = int(versionMax) + 1
		} else {
			version = versionNew
		}
		return nil
	})
	gData.Go(func() error {
		// get word
		var err error
		//wordInfo, err = u.repository.FindById(gDataCtx, wordId)
		videoInfo, err = u.videoRepository.GetVideoById(gDataCtx, videoId)
		if err != nil {
			return err
		}
		var resourceArray models.VideoItem
		if resource == nil {
			if errUnmarshal := json.Unmarshal([]byte(videoInfo.Config), &resourceArray); errUnmarshal != nil {
				return errUnmarshal
			}
			item = resourceArray
		} else {
			item = *resource
		}
		return nil
	})
	if gDataErr := gData.Wait(); gDataErr != nil {
		return gDataErr
	}

	folderWord, fileNameZip = u.getPathVideoResource(videoInfo.PathResource, version)
	item.PathResource = fileNameZip
	// zip word
	gZ, gZCtx := errgroup.WithContext(ctx)
	//gZ.Go(func() error {
	//	if err := u.processZipResource(gZCtx, videoInfo, item, folderWord, fileNameZip, models.FolderHd); err != nil {
	//		return errors.New("zip word hd failed: " + err.Error())
	//	}
	//	return nil
	//})
	gZ.Go(func() error {
		if err := u.processZipResource(gZCtx, videoInfo, item, folderWord, fileNameZip, models.FolderHdr); err != nil {
			return errors.New("zip word hdr failed: " + err.Error())
		}
		return nil
	})
	if gZErr := gZ.Wait(); gZErr != nil {
		return gZErr
	}

	// save path and version word
	itemString, err := json.Marshal(item)
	if err != nil {
		return err
	}
	videoInfo.Config = string(itemString)
	videoInfo.PathResource = fileNameZip
	videoInfo.Version = uint(version)
	if _, errSaveWord := u.videoRepository.UpdateVideo(ctx, videoInfo); errSaveWord != nil {
		return errors.New("save word failed: " + errSaveWord.Error())
	}

	return nil
}

func (u *videoUsecase) getPathVideoResource(pathVideoOld string, version int) (string, string) {
	pathWord := ""
	if pathVideoOld != "" {
		path := strings.Replace(pathVideoOld, ".zip", "", -1)
		arrPathWord := strings.Split(path, "_")
		pathWord = arrPathWord[0]
	} else {
		pathWord = uuid.NewString()
	}
	pathWord = fmt.Sprintf("%s_%d", pathWord, version)
	pathWord = strings.ToLower(pathWord)
	return pathWord, fmt.Sprintf("%s.zip", pathWord)
}

func (p *videoUsecase) getListSaveResourceVideo(videoInfo *ent.Videos, item models.VideoItem, format string) ([]SaveResourceVideo, error) {
	listResource := make([]SaveResourceVideo, 0)
	listResource = append(listResource, &SaveResourceVideoItem{Video: item})
	dataString, err := json.Marshal(item)
	if err != nil {
		return listResource, err
	}
	for _, image := range item.Comments {
		listResource = append(listResource, &SaveResourceCommentImage{Format: format, Image: image})
	}
	listResource = append(listResource, &SaveResourceVideoItemJson{
		FileName: fmt.Sprintf("%d.json", videoInfo.ID),
		DataFile: dataString,
	})
	return listResource, nil
}

func (u *videoUsecase) processZipResource(ctx context.Context, videoInfo *ent.Videos, item models.VideoItem, pathWord string, fileNameZip string, format string) error {
	//folder storage
	folderStorage := fmt.Sprintf("%s/%s/%s", config.GetInstance().GetPathUpload(), format, models.FolderResource)
	folderUpload := fmt.Sprintf("%s/%s/%s", models.FolderZip, format, models.FolderResource)
	folderWord := fmt.Sprintf("%s/%s", folderStorage, pathWord)

	// download resource and save file word
	listResource, err := u.getListSaveResourceVideo(videoInfo, item, format)
	if err != nil {
		return err
	}
	if err := u.saveResource(ctx, folderWord, listResource); err != nil {
		return err
	}
	fileZip := fmt.Sprintf("%s/%s", folderStorage, fileNameZip)
	errZip := u.zipAndUploadFileWord(ctx, folderWord, fileZip, folderUpload, fileNameZip)
	if errZip != nil {
		return errZip
	}
	//remove
	_ = os.RemoveAll(folderWord)
	if format != models.HDR {
		_ = os.RemoveAll(fileZip)
	}
	return nil
}

func (u *videoUsecase) zipAndUploadFileWord(ctx context.Context, folderSource string, fileZip string, folderUpload string, fileName string) error {
	// zip file
	if err := utils.ZipFolder2(folderSource, fileZip); err != nil {
		return err
	}
	// upload file zip to media
	if _, err := u.uploadServices.UploadFileMedia(ctx, fileZip, fileName, folderUpload, "", true); err != nil {
		return err
	}
	return nil
}

func (u *videoUsecase) saveResource(ctx context.Context, folderWord string, listResource []SaveResourceVideo) error {
	//mkdir
	if err := utils.MkdirAll(folderWord); err != nil {
		return err
	}
	// download resource and save file word
	g, _ := errgroup.WithContext(ctx)
	g.SetLimit(10)
	for _, dataItemResource := range listResource {
		itemResource := dataItemResource
		g.Go(func() error {
			if err := itemResource.SaveResource(ctx, folderWord, u.cfg.Storage.Disk); err != nil {
				return err
			}
			return nil
		})
	}
	if gDErr := g.Wait(); gDErr != nil {
		return gDErr
	}
	return nil
}
