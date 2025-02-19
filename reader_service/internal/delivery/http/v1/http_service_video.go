package v1

import (
	"encoding/json"
	"strconv"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/delivery"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/dto"
	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryVideoReader = (*HttpService)(nil)

func (h *HttpService) GetVideoByID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.GetVideoByIDHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.GetVideoByID")
		defer span.Finish()

		videoId := ctx.UserValue(constants.ID).(string)
		videoIdNum, err := strconv.Atoi(videoId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		videoDB, err := h.videoUsecase.GetVideoById(tracingCtx, uint(videoIdNum))
		if err != nil {
			if ent.IsNotFound(err) {
				responseJSON, _ := json.Marshal(dto.HttpResponse{
					Status:  fasthttp.StatusOK,
					Message: "Success",
					Data:    nil,
				})
				ctx.SetContentType("application/json")
				ctx.Write(responseJSON)
				return
			}
			h.log.WarnMsg("GetVideoByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    videoDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) SearchVideo() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchVideoHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "videosHandlers.SearchVideo")
		defer span.Finish()

		pq := utils.NewPaginationFromQueryParams(string(ctx.QueryArgs().Peek(constants.Size)), string(ctx.QueryArgs().Peek(constants.Page)))

		var createDto models.SearchVideoRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		if createDto.Pagination == nil {
			createDto.Pagination = pq
		}

		videoDBs, err := h.videoUsecase.SearchVideo(tracingCtx, createDto, string(ctx.Request.Body()))
		if err != nil {
			h.log.WarnMsg("SearchVideo", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    videoDBs,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
