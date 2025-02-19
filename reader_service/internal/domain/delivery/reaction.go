package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryReactionReader interface {
	GetReactionByID() fasthttp.RequestHandler
	GetListReactionByVideoID() fasthttp.RequestHandler
}
