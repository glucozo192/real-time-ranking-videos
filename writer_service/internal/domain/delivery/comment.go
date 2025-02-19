package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryComment interface {
	CreateComment() fasthttp.RequestHandler
	UpdateComment() fasthttp.RequestHandler
	DeleteComment() fasthttp.RequestHandler
	ImportComment() fasthttp.RequestHandler
}
