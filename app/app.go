package app

import (
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"github.com/vitalvas/rspamd-statistics/internal/middleware"
	"github.com/vitalvas/rspamd-statistics/internal/storage"
	"github.com/vitalvas/rspamd-statistics/internal/storage/memstore"
)

type App struct {
	storage storage.Storage
}

func NewApp() *App {
	return &App{
		storage: memstore.NewStorage(),
	}
}

func (app *App) newRouter() http.Handler {
	router := bunrouter.New(
		bunrouter.Use(
			reqlog.NewMiddleware(
				reqlog.WithVerbose(true),
			),
		),
	)

	router.GET("/healthz", func(w http.ResponseWriter, req bunrouter.Request) error {
		_, err := w.Write([]byte("ok"))
		return err
	})

	group := router.Use(middleware.ContentType)

	group.POST("/storage", app.requestFindHandler)
	group.PUT("/storage", app.requestAddHandler)
	group.DELETE("/storage", app.requestDeleteHandler)

	return router
}
