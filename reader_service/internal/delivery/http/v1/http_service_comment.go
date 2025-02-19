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

var _ delivery.HttpDeliveryCommentReader = (*HttpService)(nil)

func (h *HttpService) GetCommentByID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.GetCommentByIDHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.GetCommentByID")
		defer span.Finish()

		commentId := ctx.UserValue(constants.ID).(string)
		commentIdNum, err := strconv.Atoi(commentId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		commentDB, err := h.commentUsecase.GetCommentById(tracingCtx, uint(commentIdNum))
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
			h.log.WarnMsg("GetCommentByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    commentDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) GetListCommentByVideoID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchCommentHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "commentsHandlers.SearchComment")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		commentDBs, err := h.commentUsecase.GetListCommentByVideoID(tracingCtx, uint(videoIDNum), fmt.Sprintf("GetListCommentByVideoID-%v", videoIDNum))
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
			h.log.WarnMsg("SearchComment", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    commentDBs,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
