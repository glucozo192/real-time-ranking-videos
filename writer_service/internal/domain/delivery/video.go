package delivery

import "github.com/valyala/fasthttp"

type HttpDeliveryVideo interface {
	CreateVideo() fasthttp.RequestHandler
	UpdateVideo() fasthttp.RequestHandler
	DeleteVideo() fasthttp.RequestHandler

	ImportVideo() fasthttp.RequestHandler
	ZipVideo() fasthttp.RequestHandler
}
