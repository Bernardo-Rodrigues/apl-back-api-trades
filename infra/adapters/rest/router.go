package rest

import "github.com/valyala/fasthttp"

func (s *restServer) router() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/report":
			s.reportService.GenerateReport(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
}
