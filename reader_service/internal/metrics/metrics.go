package metrics

import (
	"fmt"

	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ReaderServiceMetrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter

	CreateProductGrpcRequests  prometheus.Counter
	UpdateProductGrpcRequests  prometheus.Counter
	DeleteProductGrpcRequests  prometheus.Counter
	GetProductByIdGrpcRequests prometheus.Counter
	SearchProductGrpcRequests  prometheus.Counter

	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter

	CreateProductKafkaMessages prometheus.Counter
	UpdateProductKafkaMessages prometheus.Counter
	DeleteProductKafkaMessages prometheus.Counter

	SuccessHttpRequests        prometheus.Counter
	ErrorHttpRequests          prometheus.Counter
	CreateProductHttpRequests  prometheus.Counter
	UpdateProductHttpRequests  prometheus.Counter
	DeleteProductHttpRequests  prometheus.Counter
	GetProductByIdHttpRequests prometheus.Counter
	SearchProductHttpRequests  prometheus.Counter

	// Video http
	GetVideoByIDHttpRequests prometheus.Counter
	SearchVideoHttpRequests  prometheus.Counter

	// Object http
	GetObjectByIDHttpRequests prometheus.Counter
	SearchObjectHttpRequests  prometheus.Counter

	// Comment http
	GetCommentByIDHttpRequests prometheus.Counter
	SearchCommentHttpRequests  prometheus.Counter

	// Reaction http
	GetReactionByIDHttpRequests prometheus.Counter
	SearchReactionHttpRequests  prometheus.Counter

	// Viewer http
	GetViewerByIDHttpRequests prometheus.Counter
	SearchViewerHttpRequests  prometheus.Counter
}

func NewReaderServiceMetrics(cfg *config.Config) *ReaderServiceMetrics {
	return &ReaderServiceMetrics{
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
		// Http product
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

		// Http video
		GetVideoByIDHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_video_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get video by id http requests",
		}),
		SearchVideoHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_video_http_requests_total", cfg.ServiceName),
			Help: "The total number of search video http requests",
		}),

		// Http object
		GetObjectByIDHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_object_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get object by id http requests",
		}),
		SearchObjectHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_object_http_requests_total", cfg.ServiceName),
			Help: "The total number of search object http requests",
		}),

		// Http comment
		GetCommentByIDHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_comment_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get comment by id http requests",
		}),
		SearchCommentHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_comment_http_requests_total", cfg.ServiceName),
			Help: "The total number of search comment http requests",
		}),

		// Http reaction
		GetReactionByIDHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_reaction_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get reaction by id http requests",
		}),
		SearchReactionHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_reaction_http_requests_total", cfg.ServiceName),
			Help: "The total number of search reaction http requests",
		}),

		// Http viewer
		GetViewerByIDHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_viewer_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get viewer by id http requests",
		}),
		SearchViewerHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_viewer_http_requests_total", cfg.ServiceName),
			Help: "The total number of search viewer http requests",
		}),
	}
}
