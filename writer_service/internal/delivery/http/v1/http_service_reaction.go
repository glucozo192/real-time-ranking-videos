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

var _ delivery.HttpDeliveryReaction = (*HttpService)(nil)

func (h *HttpService) ImportReaction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.ImportReactionHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.ImportReaction")
		defer span.Finish()

		videoID := string(ctx.QueryArgs().Peek("videoID"))
		videoIDNum, err := strconv.Atoi(videoID)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		rows, err := GetRowsFromImportFile(ctx, "file", "Template File Import Reaction")
		if err != nil {
			h.log.WarnMsg("GetRowsFromImportFile", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		headers := rows[0]

		// Validate headers
		if !ValidateHeadersImportAPI(models.RequiredReactionHeaders, headers) {
			h.log.WarnMsg("Invalid headers", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		// Parse rows
		reactionImports := []*ent.Reactions{}
		for idx, row := range rows {
			if row != nil {
				if idx == 0 {
					continue
				}
				if strings.TrimSpace(strings.Join(row, "")) == "" {
					continue
				}
				// Create video
				v := &ent.Reactions{}
				v.VideoID = uint(videoIDNum)

				// Map row values to struct fields using tags
				for index, header := range headers {
					switch header {
					case models.ReactionImportHeaderReaction:
						v.Name = row[index]
					case models.ReactionImportHeaderNumberReaction:
						numInt, err := strconv.Atoi(row[index])
						if err != nil {
							h.log.WarnMsg(fmt.Sprintf("Error when import row %v at field %v", index, models.ReactionImportHeaderNumberReaction), err)
							h.traceErr(span, err)
							httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
							return
						}
						v.Number = numInt
					case models.ReactionImportHeaderTime:
						timeInt, err := strconv.Atoi(row[index])
						if err != nil {
							h.log.WarnMsg(fmt.Sprintf("Error when import row %v at field %v", index, models.ReactionImportHeaderTime), err)
							h.traceErr(span, err)
							httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
							return
						}
						v.TimePoint = float64(timeInt)
					}
				}

				// Append reaction
				reactionImports = append(reactionImports, v)
			}
		}

		if err = h.reactionUsecase.CreateInBulk(tracingCtx, reactionImports); err != nil {
			h.log.WarnMsg("reactionUsecase.CreateInBulk", err)
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

// UpdateReaction
// @Tags Reactions
// @Summary Update Reaction
// @Description Update Reaction item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Reactions
// @Router /reactions [post]
func (h *HttpService) UpdateReaction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateReactionHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.UpdateReaction")
		defer span.Finish()

		reactionId := ctx.UserValue(constants.ID).(string)
		reactionIdNum, err := strconv.Atoi(reactionId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		var createDto models.UpdateReactionRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		reactionDB, err := h.reactionUsecase.Update(tracingCtx, &ent.Reactions{
			ID:          uint(reactionIdNum),
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			Name:        createDto.Name,
			Number:      createDto.Number,
			TimePoint:   createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("UpdateReaction", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    reactionDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}

// DeleteReaction
// @Tags Reactions
// @Summary Delete Reaction
// @Description Delete Reaction item
// @Accept json
// @Produce json
// @Success 201 {object} success
// @Router /reactions [post]
func (h *HttpService) DeleteReaction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.DeleteReactionHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.DeleteReaction")
		defer span.Finish()

		reactionId := ctx.UserValue(constants.ID).(string)
		reactionIdNum, err := strconv.Atoi(reactionId)
		if err != nil {
			h.log.WarnMsg("strconv.Atoi", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		err = h.reactionUsecase.Delete(tracingCtx, uint(reactionIdNum))
		if err != nil {
			h.log.WarnMsg("DeleteReaction", err)
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

// CreateReaction
// @Tags Reactions
// @Summary Create Reaction
// @Description Create new Reaction item
// @Accept json
// @Produce json
// @Success 201 {object} ent.Reactions
// @Router /reactions [post]
func (h *HttpService) CreateReaction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.metrics.CreateReactionHttpRequests.Inc()

		tracingCtx, span := tracing.StartHttpServerTracerSpan(ctx, "reactionsHandlers.CreateReaction")
		defer span.Finish()

		var createDto models.CreateReactionRequest
		if err := json.Unmarshal(ctx.Request.Body(), &createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		reactionDB, err := h.reactionUsecase.Create(tracingCtx, &ent.Reactions{
			VideoID:     createDto.VideoID,
			Description: createDto.Description,
			Name:        createDto.Name,
			Number:      createDto.Number,
			TimePoint:   createDto.TimePoint,
		})
		if err != nil {
			h.log.WarnMsg("CreateReaction", err)
			h.metrics.ErrorHttpRequests.Inc()
			httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
			return
		}

		h.metrics.SuccessHttpRequests.Inc()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		responseJSON, _ := json.Marshal(dto.HttpResponse{
			Status:  fasthttp.StatusOK,
			Message: "Success",
			Data:    reactionDB,
		})
		ctx.SetContentType("application/json")
		ctx.Write(responseJSON)
	}
}
