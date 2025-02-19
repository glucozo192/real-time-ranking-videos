package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryViewer interface {
	CreateViewer() fasthttp.RequestHandler
	UpdateViewer() fasthttp.RequestHandler
	DeleteViewer() fasthttp.RequestHandler
	ImportViewer() fasthttp.RequestHandler
}
