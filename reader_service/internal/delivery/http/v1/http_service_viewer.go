package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/glu/video-real-time-ranking/ent"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/delivery"
	"github.com/glu/video-real-time-ranking/reader_service/internal/dto"

	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryViewerReader = (*HttpService)(nil)

func (h *HttpService) GetViewerByID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.GetViewerByIDHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.GetViewerByID")
		defer span.Finish()

		viewerId := ctx.UserValue(constants.ID).(string)
		viewerIdNum, err := strconv.Atoi(viewerId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		viewerDB, err := h.viewerUsecase.GetViewerById(tracingCtx, uint(viewerIdNum))
		if err != nil {
			if ent.IsNotFound(err) {
				responseJSON, _ := json.Marshal(dto.HttpResponse{
					Status:  fasthttp.StatusOK,
					Message: "Not Found!",
					Data:    nil,
				})
				ctx.SetContentType("application/json")
				ctx.Write(responseJSON)
				return
			}
			h.log.WarnMsg("GetViewerByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    viewerDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) GetListViewerByVideoID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchViewerHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "viewersHandlers.SearchViewer")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		viewerDBs, err := h.viewerUsecase.GetListViewerByVideoID(tracingCtx, uint(videoIDNum), fmt.Sprintf("GetListViewerByVideoID-%v", videoIDNum))
		if err != nil {
			if ent.IsNotFound(err) {
				responseJSON, _ := json.Marshal(dto.HttpResponse{
					Status:  fasthttp.StatusOK,
					Message: "Not Found!",
					Data:    nil,
				})
				ctx.SetContentType("application/json")
				ctx.Write(responseJSON)
				return
			}
			h.log.WarnMsg("SearchViewer", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    viewerDBs,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
