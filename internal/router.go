package internal

import (
	"github.com/fadil-tao/yokai-go-test/internal/handler"
	"github.com/ankorstore/yokai/fxhttpserver"
	"go.uber.org/fx"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "", handler.NewExampleHandler),
	)
}
