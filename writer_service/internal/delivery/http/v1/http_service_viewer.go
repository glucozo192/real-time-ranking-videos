package v1

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/delivery"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/writer_service/internal/dto"
	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryViewer = (*HttpService)(nil)

func (h *HttpService) ImportViewer() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.ImportViewerHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.ImportViewer")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		rows, err := GetRowsFromImportFile(ctx, "file", "Template File Import View")
		if err != nil {
			h.log.WarnMsg("GetRowsFromImportFile", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		headers := rows[0]

		// Validate headers
		if !ValidateHeadersImportAPI(models.RequiredViewerHeaders, headers) {
			h.log.WarnMsg("Invalid headers", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		// Parse rows
		var viewerImports []*ent.Viewers
		for idx, row := range rows {
			if row != nil {
				if idx == 0 {
					continue
				}

				if strings.TrimSpace(strings.Join(row, "")) == "" {
					continue
				}

				// Create video
				v := &ent.Viewers{}
				v.VideoID = uint(videoIDNum)

				// Map row values to struct fields using tags
				for index, header := range headers {
					switch header {
					case models.ViewerImportHeaderNumberView:
						numInt, err := strconv.Atoi(row[index])
						if err != nil {
							h.log.WarnMsg(fmt.Sprintf("Error when import row %v at field %v", index, models.ViewerImportHeaderNumberView), err)
							h.traceErr(span, err)
							httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
							return
						}
						v.Number = numInt
					case models.ViewerImportHeaderTime:
						timeInt, err := strconv.Atoi(row[index])
						if err != nil {
							h.log.WarnMsg(fmt.Sprintf("Error when import row %v at field %v", index, models.ViewerImportHeaderTime), err)
							h.traceErr(span, err)
							httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
							return
						}
						v.TimePoint = float64(timeInt)
					}
				}

				// Append viewer
				viewerImports = append(viewerImports, v)
			}
		}

		if err = h.viewerUsecase.CreateInBulk(tracingCtx, viewerImports); err != nil {
			h.log.WarnMsg("viewerUsecase.CreateInBulk", err)
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

// UpdateViewer
// @Tags Viewers
// @Summary Update Viewer
// @Description Update Viewer item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Viewers
// @Router /viewers [post]
func (h *HttpService) UpdateViewer() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateViewerHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.UpdateViewer")
		defer span.Finish()

		viewerId := ctx.UserValue(constants.ID).(string)
		viewerIdNum, err := strconv.Atoi(viewerId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var createDto models.UpdateViewerRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		viewerDB, err := h.viewerUsecase.Update(tracingCtx, &ent.Viewers{
			ID:        uint(viewerIdNum),
			VideoID:   createDto.VideoID,
			Number:    createDto.Number,
			TimePoint: createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("UpdateViewer", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    viewerDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// DeleteViewer
// @Tags Viewers
// @Summary Delete Viewer
// @Description Delete Viewer item
// @Accept json
// @Produce json
// @Success 201 {object} success
// @Router /viewers [post]
func (h *HttpService) DeleteViewer() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteViewerHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.DeleteViewer")
		defer span.Finish()

		viewerId := ctx.UserValue(constants.ID).(string)
		viewerIdNum, err := strconv.Atoi(viewerId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.viewerUsecase.Delete(tracingCtx, uint(viewerIdNum))
		if err != nil {
			h.log.WarnMsg("DeleteViewer", err)
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

// CreateViewer
// @Tags Viewers
// @Summary Create Viewer
// @Description Create new Viewer item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Viewers
// @Router /viewers [post]
func (h *HttpService) CreateViewer() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateViewerHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.CreateViewer")
		defer span.Finish()

		var createDto models.CreateViewerRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		viewerDB, err := h.viewerUsecase.Create(tracingCtx, &ent.Viewers{
			VideoID:   createDto.VideoID,
			Number:    createDto.Number,
			TimePoint: createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("CreateViewer", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    viewerDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
