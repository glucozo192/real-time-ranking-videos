package v1

import (
	"encoding/json"
	"fmt"
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

var _ delivery.HttpDeliveryComment = (*HttpService)(nil)

func (h *HttpService) ImportComment() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.ImportCommentHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.ImportComment")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		rows, err := GetRowsFromImportFile(ctx, "file", "Template File Import Comment")
		if err != nil {
			h.log.WarnMsg("GetRowsFromImportFile", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		headers := rows[0]

		// Validate headers
		if !ValidateHeadersImportAPI(models.RequiredCommentHeaders, headers) {
			h.log.WarnMsg("CreateComment", nil)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, nil, h.cfg.Http.DebugErrorsResponse)
			return
		}

		// Parse rows
		var commentImports []*ent.Comments
		for idx, row := range rows {
			if row != nil {
				if idx == 0 {
					continue
				}

				if strings.TrimSpace(strings.Join(row, "")) == "" {
					continue
				}

				// Create video
				v := &ent.Comments{}

				v.VideoID = uint(videoIDNum)

				// Map row values to struct fields using tags
				for index, header := range headers {
					switch header {
					case models.CommentImportHeaderName:
						v.UserName = row[index]
					case models.CommentImportHeaderAvatar:
						url := row[index]

						if strings.Contains(url, "avatar/avatar") {
							v.Avatar = url
						} else {
							localPath := config.GetInstance().GetPathUpload()
							k5Path := h.cfg.ExternalService.K5Path

							fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
							if err != nil {
								return
							}

							v.Avatar = fileName
						}
					case models.CommentImportHeaderComment:
						v.Comment = row[index]
					case models.CommentImportHeaderTime:
						timeInt, err := strconv.Atoi(row[index])
						if err != nil {
							h.log.WarnMsg(fmt.Sprintf("Error when import row %v at field %v", index, models.CommentImportHeaderTime), err)
							h.traceErr(span, err)
							httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
							return
						}
						v.TimePoint = float64(timeInt)
					default:
						continue
					}
				}

				// Append comment
				commentImports = append(commentImports, v)
			}
		}

		if err = h.commentUsecase.CreateInBulk(tracingCtx, commentImports); err != nil {
			h.log.WarnMsg("commentUsecase.CreateInBulk", err)
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

// CreateComment
// @Tags Objects
// @Summary Update Object
// @Description Update Object item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Objects
// @Router /objects [post]
func (h *HttpService) CreateComment() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateCommentHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.CreateComment")
		defer span.Finish()

		var createDto models.CreateCommentRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var avatarValue string
		url := createDto.Avatar

		if strings.Contains(url, "avatar/avatar") {
			avatarValue = url
		} else {
			localPath := config.GetInstance().GetPathUpload()
			k5Path := h.cfg.ExternalService.K5Path

			fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
			if err != nil {
				return
			}

			avatarValue = fileName
		}

		commentDB, err := h.commentUsecase.Create(tracingCtx, &ent.Comments{
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			Comment:     createDto.Comment,
			UserName:    createDto.UserName,
			Avatar:      avatarValue,
			TimePoint:   createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("CreateComment", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    commentDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// UpdateComment
// @Tags Objects
// @Summary Update Object
// @Description Update Object item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Objects
// @Router /objects [post]
func (h *HttpService) UpdateComment() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateCommentHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.UpdateComment")
		defer span.Finish()

		commentId := ctx.UserValue(constants.ID).(string)
		commentIdNum, err := strconv.Atoi(commentId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var createDto models.UpdateCommentRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var avatarValue string
		url := createDto.Avatar

		if strings.Contains(url, "avatar/avatar") {
			avatarValue = url
		} else {
			localPath := config.GetInstance().GetPathUpload()
			k5Path := h.cfg.ExternalService.K5Path

			fileName, err := h.CloneAndPushResourceByURL(ctx, url, localPath, k5Path)
			if err != nil {
				return
			}

			avatarValue = fileName
		}

		commentDB, err := h.commentUsecase.Update(tracingCtx, &ent.Comments{
			ID:          uint(commentIdNum),
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			Comment:     createDto.Comment,
			UserName:    createDto.UserName,
			Avatar:      avatarValue,
			TimePoint:   createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("UpdateComment", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    commentDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// DeleteComment
// @Tags Objects
// @Summary Delete Object
// @Description Delete Object item
// @Accept json
// @Produce json
// @Success 201 {object} success
// @Router /objects [post]
func (h *HttpService) DeleteComment() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteCommentHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.DeleteComment")
		defer span.Finish()

		commentId := ctx.UserValue(constants.ID).(string)
		commentIdNum, err := strconv.Atoi(commentId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.commentUsecase.Delete(tracingCtx, uint(commentIdNum))
		if err != nil {
			h.log.WarnMsg("DeleteComment", err)
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
