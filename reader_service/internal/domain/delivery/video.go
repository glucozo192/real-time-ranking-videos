package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryVideoReader interface {
	GetVideoByID() fasthttp.RequestHandler
	SearchVideo() fasthttp.RequestHandler
}
