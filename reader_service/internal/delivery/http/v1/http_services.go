package v1

import (
	"github.com/georgecookeIW/fasthttprouter"
	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/usecase"
	product2 "github.com/glu/video-real-time-ranking/reader_service/internal/metrics"
	"github.com/go-playground/validator"
	"github.com/opentracing/opentracing-go"
)

type HttpService struct {
	routes          *fasthttprouter.Router
	log             logger.Logger
	cfg             *config.Config
	v               *validator.Validate
	videoUsecase    usecase.IVideoUsecase
	commentUsecase  usecase.ICommentUsecase
	objectUsecase   usecase.IObjectUsecase
	reactionUsecase usecase.IReactionUsecase
	viewerUsecase   usecase.IViewerUsecase
	metrics         *product2.ReaderServiceMetrics
}

func NewReaderHttpService(
	routes *fasthttprouter.Router,
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	videoUsecase usecase.IVideoUsecase,
	commentUsecase usecase.ICommentUsecase,
	objectUsecase usecase.IObjectUsecase,
	reactionUsecase usecase.IReactionUsecase,
	viewerUsecase usecase.IViewerUsecase,
	metrics *product2.ReaderServiceMetrics,
) *HttpService {
	return &HttpService{
		routes:          routes,
		log:             log,
		cfg:             cfg,
		v:               v,
		videoUsecase:    videoUsecase,
		commentUsecase:  commentUsecase,
		objectUsecase:   objectUsecase,
		reactionUsecase: reactionUsecase,
		viewerUsecase:   viewerUsecase,
		metrics:         metrics,
	}
}

func (h *HttpService) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
