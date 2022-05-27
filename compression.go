package main

import (
	"io"
	"strings"

	"github.com/klauspost/compress/zstd"
	"github.com/uptrace/bunrouter"
)

func compressionReader(req bunrouter.Request) (io.Reader, error) {
	switch strings.ToLower(req.Header.Get("content-encoding")) {
	case "zstd":
		return zstd.NewReader(req.Body)

	default:
		return req.Body, nil
	}
}
