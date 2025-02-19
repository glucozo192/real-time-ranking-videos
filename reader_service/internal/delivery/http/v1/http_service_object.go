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
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/dto"

	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryObjectReader = (*HttpService)(nil)

func (h *HttpService) GetObjectByID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.GetObjectByIDHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.GetObjectByID")
		defer span.Finish()

		objectId := ctx.UserValue(constants.ID).(string)
		objectIdNum, err := strconv.Atoi(objectId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		objectDB, err := h.objectUsecase.GetObjectById(tracingCtx, uint(objectIdNum))
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
			h.log.WarnMsg("GetObjectByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusOK)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    objectDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) GetListObjectByVideoID() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchObjectHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.SearchObject")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		objectDBs, err := h.objectUsecase.GetListObjectByVideoID(tracingCtx, uint(videoIDNum), fmt.Sprintf("GetListObjectByVideoID-%v", videoIDNum))
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
			h.log.WarnMsg("SearchObject", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    objectDBs,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

func (h *HttpService) GetListObjectByVideoPath() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.SearchObjectHttpRequests.Inc()
		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.SearchObject")
		defer span.Finish()

		path := string(ctx.QueryArgs().Peek("path"))
		objectDBs, err := h.objectUsecase.GetListObjectByVideoPath(tracingCtx, path)
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
			h.log.WarnMsg("SearchObject", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}
		var dataRes = make([]*models.ObjectInteractiveResponse, 0, len(objectDBs))
		for _, objectDB := range objectDBs {
			dataRes = append(dataRes, &models.ObjectInteractiveResponse{
				Name:        objectDB.Description,
				MakerName:   objectDB.MarkerName,
				In:          int(objectDB.TimeStart),
				Out:         int(objectDB.TimeEnd),
				TouchVector: objectDB.TouchVector,
			})
		}
		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    dataRes,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
