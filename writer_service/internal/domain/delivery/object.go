package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryObject interface {
	CreateObject() fasthttp.RequestHandler
	UpdateObject() fasthttp.RequestHandler
	DeleteObject() fasthttp.RequestHandler
}
