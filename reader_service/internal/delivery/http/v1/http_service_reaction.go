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

var _ delivery.HttpDeliveryReactionReader = (*HttpService)(nil)

func (h *HttpService) GetReactionByID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.GetReactionByIDHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.GetReactionByID")
		defer span.Finish()

		reactionId := ctx.UserValue(constants.ID).(string)
		reactionIdNum, err := strconv.Atoi(reactionId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		reactionDB, err := h.reactionUsecase.GetReactionById(tracingCtx, uint(reactionIdNum))
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
			h.log.WarnMsg("GetReactionByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    reactionDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) GetListReactionByVideoID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchReactionHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.SearchReaction")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		reactionDBs, err := h.reactionUsecase.GetListReactionByVideoID(tracingCtx, uint(videoIDNum), fmt.Sprintf("GetListReactionByVideoID-%v", videoIDNum))
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
			h.log.WarnMsg("SearchReaction", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    reactionDBs,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
