package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryCommentReader interface {
	GetCommentByID() fasthttp.RequestHandler
	GetListCommentByVideoID() fasthttp.RequestHandler
}
