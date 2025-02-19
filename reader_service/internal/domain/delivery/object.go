package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryObjectReader interface {
	GetObjectByID() fasthttp.RequestHandler
	GetListObjectByVideoID() fasthttp.RequestHandler
}
