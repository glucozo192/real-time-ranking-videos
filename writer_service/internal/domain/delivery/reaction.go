package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryReaction interface {
	CreateReaction() fasthttp.RequestHandler
	UpdateReaction() fasthttp.RequestHandler
	DeleteReaction() fasthttp.RequestHandler
	ImportReaction() fasthttp.RequestHandler
}
