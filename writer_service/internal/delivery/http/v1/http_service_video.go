package v1

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/glu/video-real-time-ranking/writer_service/config"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/delivery"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/writer_service/internal/dto"

	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryVideo = (*HttpService)(nil)

func (h *HttpService) ZipVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.ZipVideo")
		defer span.Finish()

		videoId := ctx.UserValue(constants.ID).(string)
		videoIdNum, err := strconv.Atoi(videoId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.videoUsecase.CreateVideoItemConfig(tracingCtx, uint(videoIdNum))
		if err != nil {
			h.log.WarnMsg("CreateVideoItemConfig", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.videoUsecase.ZipById(ctx, uint(videoIdNum), nil, 0)
		if err != nil {
			h.log.WarnMsg("ZipById", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    nil,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) ImportVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.ImportVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.UpdateVideo")
		defer span.Finish()

		rows, err := GetRowsFromImportFile(ctx, "file", "Template File Import Video")
		if err != nil {
			h.log.WarnMsg("GetRowsFromImportFile", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		headers := rows[0]

		// Validate headers
		if !ValidateHeadersImportAPI(models.RequiredVideoHeaders, headers) {
			h.log.WarnMsg("Invalid headers", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		// Parse rows
		var videoImports []*ent.Videos
		for idx, row := range rows {
			if row != nil {
				if idx == 0 {
					continue
				}

				if strings.TrimSpace(strings.Join(row, "")) == "" {
					continue
				}

				// Create video
				v := &ent.Videos{}

				// Map row values to struct fields using tags
				for index, header := range headers {
					switch header {
					case models.VideoImportHeaderName:
						v.Name = row[index]
					case models.VideoImportHeaderAuthor:
						v.Author = row[index]
					case models.VideoImportHeaderVideoURL:
						url := row[index]

						localPath := config.GetInstance().GetPathUpload()
						k5Path := h.cfg.ExternalService.K5Path

						fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
						if err != nil {
							return
						}

						v.VideoURL = fileName
					}
				}

				// Append video
				videoImports = append(videoImports, v)
			}
		}

		if err = h.videoUsecase.CreateInBulk(tracingCtx, videoImports); err != nil {
			h.log.WarnMsg("videoUsecase.CreateInBulk", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    nil,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// UpdateVideo
// @Tags Videos
// @Summary Update Video
// @Description Update Video item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Videos
// @Router /videos [post]
func (h *HttpService) UpdateVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.UpdateVideo")
		defer span.Finish()

		var createDto models.UpdateVideoRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		url := createDto.VideoURL

		localPath := config.GetInstance().GetPathUpload()
		k5Path := h.cfg.ExternalService.K5Path

		fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
		if err != nil {
			return
		}

		videoDB, err := h.videoUsecase.Update(tracingCtx, &ent.Videos{
			ID:           createDto.ID,
			Name:         createDto.Name,
			Description:  createDto.Description,
			VideoURL:     fileName,
			Config:       createDto.Config,
			PathResource: createDto.PathResource,
			LevelSystem:  createDto.LevelSystem,
			Status:       createDto.Status,
			Note:         createDto.Note,
			Assign:       createDto.Assign,
			Author:       createDto.Author,
		})
		if err != nil {
			h.log.WarnMsg("UpdateVideo", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    videoDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// DeleteVideo
// @Tags Videos
// @Summary Delete Video
// @Description Delete Video item
// @Accept json
// @Produce json
// @Success 201 {object} success
// @Router /videos [post]
func (h *HttpService) DeleteVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.DeleteVideo")
		defer span.Finish()

		videoId := ctx.UserValue(constants.ID).(string)
		videoIdNum, err := strconv.Atoi(string(videoId))
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.videoUsecase.Delete(tracingCtx, uint(videoIdNum))
		if err != nil {
			h.log.WarnMsg("DeleteVideo", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    nil,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// CreateVideo
// @Tags Videos
// @Summary Create Video
// @Description Create new Video item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Videos
// @Router /videos [post]
func (h *HttpService) CreateVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.CreateVideo")
		defer span.Finish()

		var createDto models.CreateVideoRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		url := createDto.VideoURL

		localPath := config.GetInstance().GetPathUpload()
		k5Path := h.cfg.ExternalService.K5Path

		fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
		if err != nil {
			return
		}

		videoDB, err := h.videoUsecase.Create(tracingCtx, &ent.Videos{
			Name:         createDto.Name,
			Description:  createDto.Description,
			VideoURL:     fileName,
			Config:       createDto.Config,
			PathResource: createDto.PathResource,
			LevelSystem:  createDto.LevelSystem,
			Status:       createDto.Status,
			Note:         createDto.Note,
			Assign:       createDto.Assign,
			Author:       createDto.Author,
		})
		if err != nil {
			h.log.WarnMsg("CreateVideo", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    videoDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// ImportVideos
// @Tags Videos
// @Import Create Video
// @Description Create new Video item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Videos
// @Router /videos/import-by-excel [post
func (h *HttpService) ImportDataVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.ImportVideos")
		defer span.Finish()

		message := "Import Videos by excel"
		var input models.UpdateVideoByExcel

		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.Error("Error parsing multipart form: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		files := form.File["file"]
		if len(files) == 0 {
			ctx.Error("No file uploaded", fasthttp.StatusBadRequest)
			return
		}

		fileHeader := files[0]
		input.File = fileHeader

		if err := h.videoUsecase.BulkImportByExcel(tracingCtx, input.File); err != nil {
			h.log.WarnMsg("CreateVideo", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: message,
			//Data:    videoDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}

}
