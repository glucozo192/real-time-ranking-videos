package v1

import (
	"os"
	"path"
	"strings"

	httpErrors "github.com/glu/video-real-time-ranking/core/pkg/http_errors"
	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/services"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/usecase"
	product2 "github.com/glu/video-real-time-ranking/writer_service/internal/metrics"

	"github.com/georgecookeIW/fasthttprouter"
	"github.com/go-playground/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
	"github.com/xuri/excelize/v2"
)

type HttpService struct {
	routes          *fasthttprouter.Router
	log             logger.Logger
	cfg             *config.Config
	v               *validator.Validate
	videoUsecase    usecase.IVideoUsecase
	viewerUsecase   usecase.IViewerUsecase
	objectUsecase   usecase.IObjectUsecase
	reactionUsecase usecase.IReactionUsecase
	commentUsecase  usecase.ICommentUsecase
	uploadServices  services.IUploadServices
	metrics         *product2.WriterServiceMetrics
}

func NewWriterHttpService(
	routes *fasthttprouter.Router,
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	videoUsecase usecase.IVideoUsecase,
	viewerUsecase usecase.IViewerUsecase,
	objectUsecase usecase.IObjectUsecase,
	reactionUsecase usecase.IReactionUsecase,
	commentUsecase usecase.ICommentUsecase,
	uploadServices services.IUploadServices,
	metrics *product2.WriterServiceMetrics,
) *HttpService {
	return &HttpService{
		routes:          routes,
		log:             log,
		cfg:             cfg,
		v:               v,
		videoUsecase:    videoUsecase,
		viewerUsecase:   viewerUsecase,
		objectUsecase:   objectUsecase,
		reactionUsecase: reactionUsecase,
		commentUsecase:  commentUsecase,
		uploadServices:  uploadServices,
		metrics:         metrics}
}

func (h *HttpService) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}

func (h *HttpService) CloneAndPushResourceByURL(ctx *fasthttp.RequestCtx, url, localPath, destinationPath string) (string, error) {
	resourceName := path.Base(url)

	err := utils.DownloadFile(url, localPath+resourceName)

	_, err = h.uploadServices.UploadFileMedia(ctx,
		localPath+resourceName,
		resourceName,
		destinationPath,
		"",
		true)
	if err != nil {
		h.log.WarnMsg("UploadFileMedia", err)
		httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
		return "", err
	}

	if err = os.Remove(localPath + resourceName); err != nil {
		h.log.WarnMsg("os.Remove", err)
		httpErrors.ErrorCtxResponse(ctx, err, h.cfg.Http.DebugErrorsResponse)
		return "", err
	}

	return resourceName, nil
}

func GetRowsFromImportFile(ctx *fasthttp.RequestCtx, keyFile string, sheetName string) ([][]string, error) {
	file, err := ctx.FormFile(keyFile)
	if err != nil {
		return nil, err
	}

	// Convert uploaded file to io.Reader
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}

	// Now pass the reader
	xl, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}

	// Parse headers
	rows, err := xl.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func ValidateHeadersImportAPI(requiredVideoHeaders, headers []string) bool {
	for _, header := range requiredVideoHeaders {
		if !strings.Contains(strings.Join(headers, ","), header) {
			return false
		}
	}
	return true
}
