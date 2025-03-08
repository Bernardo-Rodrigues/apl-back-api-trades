package csv

import "io"

type cvsHandler struct {
	tradesFile  io.Reader
	assetsFiles map[string]io.Reader
}

func NewHandler(tradesFile io.Reader, assetsFiles map[string]io.Reader) *cvsHandler {
	return &cvsHandler{
		tradesFile:  tradesFile,
		assetsFiles: assetsFiles,
	}
}
