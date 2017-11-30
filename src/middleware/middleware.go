package middleware

import (
	"github.com/rpoletaev/mwtest/src/middleware/autch"
	"github.com/rpoletaev/mwtest/src/middleware/compress"
	// "middleware/convertor"
)

var (
	Autch    = autch.New()
	Compress = compress.New()
	// Convertor = convertor.New()
)
