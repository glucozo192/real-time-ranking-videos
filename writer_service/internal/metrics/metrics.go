package metrics

import (
	"fmt"

	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type WriterServiceMetrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter

	CreateProductGrpcRequests  prometheus.Counter
	UpdateProductGrpcRequests  prometheus.Counter
	DeleteProductGrpcRequests  prometheus.Counter
	GetProductByIdGrpcRequests prometheus.Counter
	SearchProductGrpcRequests  prometheus.Counter

	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter

	// Product
	CreateProductKafkaMessages prometheus.Counter
	UpdateProductKafkaMessages prometheus.Counter
	DeleteProductKafkaMessages prometheus.Counter

	// Video
	CreateVideoKafkaMessages prometheus.Counter
	UpdateVideoKafkaMessages prometheus.Counter
	DeleteVideoKafkaMessages prometheus.Counter

	SuccessHttpRequests prometheus.Counter
	ErrorHttpRequests   prometheus.Counter

	// product http
	CreateProductHttpRequests  prometheus.Counter
	UpdateProductHttpRequests  prometheus.Counter
	DeleteProductHttpRequests  prometheus.Counter
	GetProductByIdHttpRequests prometheus.Counter
	SearchProductHttpRequests  prometheus.Counter

	// video http
	CreateVideoHttpRequests prometheus.Counter
	UpdateVideoHttpRequests prometheus.Counter
	DeleteVideoHttpRequests prometheus.Counter
	ImportVideoHttpRequests prometheus.Counter

	// Viewer http
	CreateViewerHttpRequests prometheus.Counter
	UpdateViewerHttpRequests prometheus.Counter
	DeleteViewerHttpRequests prometheus.Counter
	ImportViewerHttpRequests prometheus.Counter

	// Object http
	CreateObjectHttpRequests prometheus.Counter
	UpdateObjectHttpRequests prometheus.Counter
	DeleteObjectHttpRequests prometheus.Counter

	// Comment http
	CreateCommentHttpRequests prometheus.Counter
	UpdateCommentHttpRequests prometheus.Counter
	DeleteCommentHttpRequests prometheus.Counter
	ImportCommentHttpRequests prometheus.Counter

	// c http
	CreateReactionHttpRequests prometheus.Counter
	UpdateReactionHttpRequests prometheus.Counter
	DeleteReactionHttpRequests prometheus.Counter
	ImportReactionHttpRequests prometheus.Counter
}

func NewWriterServiceMetrics(cfg *config.Config) *WriterServiceMetrics {
	return &WriterServiceMetrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of error grpc requests",
		}),
		CreateProductGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_product_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of create product grpc requests",
		}),
		UpdateProductGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_product_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of update product grpc requests",
		}),
		DeleteProductGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_product_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of delete product grpc requests",
		}),
		GetProductByIdGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_product_by_id_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of get product by id grpc requests",
		}),
		SearchProductGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_product_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of search product grpc requests",
		}),

		CreateProductKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_product_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of create product kafka messages",
		}),
		UpdateProductKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_product_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of update product kafka messages",
		}),
		DeleteProductKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_product_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of delete product kafka messages",
		}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of error kafka processed messages",
		}),
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requests_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requests_total", cfg.ServiceName),
			Help: "The total number of error http requests",
		}),

		// Product
		CreateProductHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_product_http_requests_total", cfg.ServiceName),
			Help: "The total number of create product http requests",
		}),
		UpdateProductHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_product_http_requests_total", cfg.ServiceName),
			Help: "The total number of update product http requests",
		}),
		DeleteProductHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_product_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete product http requests",
		}),
		GetProductByIdHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_product_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get product by id http requests",
		}),
		SearchProductHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_product_http_requests_total", cfg.ServiceName),
			Help: "The total number of search product http requests",
		}),

		// Video request
		CreateVideoHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_video_http_requests_total", cfg.ServiceName),
			Help: "The total number of create video http requests",
		}),
		UpdateVideoHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_video_http_requests_total", cfg.ServiceName),
			Help: "The total number of update video http requests",
		}),
		DeleteVideoHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_video_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete video http requests",
		}),
		ImportVideoHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_import_video_http_requests_total", cfg.ServiceName),
			Help: "The total number of Import video http requests",
		}),

		CreateVideoKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_video_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of create video kafka messages",
		}),
		UpdateVideoKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_video_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of update video kafka messages",
		}),
		DeleteVideoKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_video_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of delete video kafka messages",
		}),

		// Comment request
		CreateCommentHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_comment_http_requests_total", cfg.ServiceName),
			Help: "The total number of create comment http requests",
		}),
		UpdateCommentHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_comment_http_requests_total", cfg.ServiceName),
			Help: "The total number of update comment http requests",
		}),
		DeleteCommentHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_comment_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete comment http requests",
		}),
		ImportCommentHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_import_comment_http_requests_total", cfg.ServiceName),
			Help: "The total number of Import comment http requests",
		}),

		// Viewer request
		CreateViewerHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_viewer_http_requests_total", cfg.ServiceName),
			Help: "The total number of create viewer http requests",
		}),
		UpdateViewerHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_viewer_http_requests_total", cfg.ServiceName),
			Help: "The total number of update viewer http requests",
		}),
		DeleteViewerHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_viewer_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete viewer http requests",
		}),
		ImportViewerHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_import_viewer_http_requests_total", cfg.ServiceName),
			Help: "The total number of Import viewer http requests",
		}),

		// Object request
		CreateObjectHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_object_http_requests_total", cfg.ServiceName),
			Help: "The total number of create object http requests",
		}),
		UpdateObjectHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_object_http_requests_total", cfg.ServiceName),
			Help: "The total number of update object http requests",
		}),
		DeleteObjectHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_object_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete object http requests",
		}),

		// Reaction request
		CreateReactionHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_reaction_http_requests_total", cfg.ServiceName),
			Help: "The total number of create reaction http requests",
		}),
		UpdateReactionHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_reaction_http_requests_total", cfg.ServiceName),
			Help: "The total number of update reaction http requests",
		}),
		DeleteReactionHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_reaction_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete reaction http requests",
		}),
		ImportReactionHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_import_reaction_http_requests_total", cfg.ServiceName),
			Help: "The total number of Import reaction http requests",
		}),
	}
}
