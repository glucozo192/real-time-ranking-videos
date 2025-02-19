package v1

import (
	"encoding/json"
	"strconv"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/delivery"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/writer_service/internal/dto"

	"github.com/valyala/fasthttp"
)

var _ delivery.HttpDeliveryObject = (*HttpService)(nil)

// UpdateObject
// @Tags Objects
// @Summary Update Object
// @Description Update Object item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Objects
// @Router /objects [post]
func (h *HttpService) UpdateObject() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateObjectHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.UpdateObject")
		defer span.Finish()

		objectId := ctx.UserValue(constants.ID).(string)
		objectIdNum, err := strconv.Atoi(objectId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var createDto models.UpdateObjectRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}
		objectDB, err := h.objectUsecase.Update(tracingCtx, &ent.Objects{
			ID:          uint(objectIdNum),
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			CoordinateX: createDto.CoordinateX,
			CoordinateY: createDto.CoordinateY,
			Length:      createDto.Length,
			Width:       createDto.Width,
			Order:       createDto.Order,
			TimeStart:   createDto.TimeStart,
			TimeEnd:     createDto.TimeEnd,
			MarkerName:  createDto.MarkerName,
			TouchVector: createDto.TouchVector,
		})
		if err != nil {
			h.log.WarnMsg("UpdateObject", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    objectDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// DeleteObject
// @Tags Objects
// @Summary Delete Object
// @Description Delete Object item
// @Accept json
// @Produce json
// @Success 201 {object} success
// @Router /objects [post]
func (h *HttpService) DeleteObject() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteObjectHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.DeleteObject")
		defer span.Finish()

		objectId := ctx.UserValue(constants.ID).(string)
		objectIdNum, err := strconv.Atoi(objectId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.objectUsecase.Delete(tracingCtx, uint(objectIdNum))
		if err != nil {
			h.log.WarnMsg("DeleteObject", err)
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

// CreateObject
// @Tags Objects
// @Summary Create Object
// @Description Create new Object item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Objects
// @Router /objects [post]
func (h *HttpService) CreateObject() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateObjectHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "objectsHandlers.CreateObject")
		defer span.Finish()

		var createDto models.CreateObjectRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		objectDB, err := h.objectUsecase.Create(tracingCtx, &ent.Objects{
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			CoordinateX: createDto.CoordinateX,
			CoordinateY: createDto.CoordinateY,
			Length:      createDto.Length,
			Width:       createDto.Width,
			Order:       createDto.Order,
			TimeStart:   createDto.TimeStart,
			TimeEnd:     createDto.TimeEnd,
		})
		if err != nil {
			h.log.WarnMsg("CreateObject", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    objectDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
