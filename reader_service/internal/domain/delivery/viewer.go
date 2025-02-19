package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryViewerReader interface {
	GetViewerByID() fasthttp.RequestHandler
	GetListViewerByVideoID() fasthttp.RequestHandler
}
