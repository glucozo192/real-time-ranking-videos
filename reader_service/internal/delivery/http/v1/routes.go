package v1

import (
	"github.com/valyala/fasthttp"
)

func (h *HttpService) MapRoutes() {
	prefixReader := "/reader"
	h.routes.GET(prefixReader+"/", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("Default path")
	})

	pathVideo := prefixReader + "/api/v1/videos"
	h.routes.POST(pathVideo+"/search", h.SearchVideo())
	h.routes.GET(pathVideo+"/:id", h.GetVideoByID())
	h.routes.OPTIONS(pathVideo+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	pathObject := prefixReader + "/api/v1/objects"
	h.routes.GET(pathObject+"/list-by-video", h.GetListObjectByVideoID())
	h.routes.GET(pathObject+"/list-by-video-v2", h.GetListObjectByVideoPath())
	h.routes.GET(pathObject+"/get/:id", h.GetObjectByID())
	h.routes.OPTIONS(pathObject+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	pathReaction := prefixReader + "/api/v1/reactions"
	h.routes.GET(pathReaction+"/list-by-video", h.GetListReactionByVideoID())
	h.routes.GET(pathReaction+"/get/:id", h.GetReactionByID())
	h.routes.OPTIONS(pathReaction+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	pathComment := prefixReader + "/api/v1/comments"
	h.routes.GET(pathComment+"/list-by-video", h.GetListCommentByVideoID())
	h.routes.GET(pathComment+"/get/:id", h.GetCommentByID())
	h.routes.OPTIONS(pathComment+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	pathViewer := prefixReader + "/api/v1/viewers"
	h.routes.GET(pathViewer+"/list-by-video", h.GetListViewerByVideoID())
	h.routes.GET(pathViewer+"/get/:id", h.GetViewerByID())
	h.routes.OPTIONS(pathViewer+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})
}
