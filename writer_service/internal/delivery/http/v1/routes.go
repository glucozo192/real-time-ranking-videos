package v1

import (
	"github.com/valyala/fasthttp"
)

func (h *HttpService) MapRoutes() {
	prefixWriter := "/writer"

	h.routes.GET(prefixWriter+"/", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("Default path")
	})

	// video
	pathVideos := prefixWriter + "/api/v1/videos"
	h.routes.POST(pathVideos+"", h.CreateVideo())
	h.routes.POST(pathVideos+"/import", h.ImportVideo())
	h.routes.POST(pathVideos+"/zip/:id", h.ZipVideo())
	h.routes.POST(pathVideos+"/update", h.UpdateVideo())
	h.routes.POST(pathVideos+"/delete/:id", h.DeleteVideo())
	h.routes.OPTIONS(pathVideos+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})
	h.routes.POST(pathVideos+"/import-by-excel", h.ImportDataVideo())

	// comment
	pathComments := prefixWriter + "/api/v1/comments"
	h.routes.POST(pathComments+"", h.CreateComment())
	h.routes.POST(pathComments+"/import", h.ImportComment())
	h.routes.POST(pathComments+"/update/:id", h.UpdateComment())
	h.routes.POST(pathComments+"/delete/:id", h.DeleteComment())
	h.routes.OPTIONS(pathComments+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	// object
	pathObjects := prefixWriter + "/api/v1/objects"
	h.routes.POST(pathObjects+"", h.CreateObject())
	h.routes.POST(pathObjects+"/update/:id", h.UpdateObject())
	h.routes.POST(pathObjects+"/delete/:id", h.DeleteObject())
	h.routes.OPTIONS(pathObjects+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	// reaction
	pathReactions := prefixWriter + "/api/v1/reactions"
	h.routes.POST(pathReactions+"", h.CreateReaction())
	h.routes.POST(pathReactions+"/import", h.ImportReaction())
	h.routes.POST(pathReactions+"/update/:id", h.UpdateReaction())
	h.routes.POST(pathReactions+"/delete/:id", h.DeleteReaction())
	h.routes.OPTIONS(pathReactions+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})

	// viewer
	pathViewers := prefixWriter + "/api/v1/viewers"
	h.routes.POST(pathViewers+"", h.CreateViewer())
	h.routes.POST(pathViewers+"/import", h.ImportViewer())
	h.routes.POST(pathViewers+"/update/:id", h.UpdateViewer())
	h.routes.POST(pathViewers+"/delete/:id", h.DeleteViewer())
	h.routes.OPTIONS(pathViewers+"/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")
	})
}
