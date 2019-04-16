package handlers

import (
	"log"
	"net/http"

	"github.com/ardanlabs/garagesale/internal/mid"
	"github.com/ardanlabs/garagesale/internal/platform/web"
	"github.com/jmoiron/sqlx"
)

// API constructs an http.Handler with all application routes defined.
func API(db *sqlx.DB, log *log.Logger) http.Handler {

	// Create the variable that contains all Middleware functions.
	mw := mid.Middleware{Log: log}

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.New(log, mw.Errors, mw.Metrics)

	{
		c := Checks{db: db}
		app.Handle(http.MethodGet, "/v1/health", c.Health)
	}

	{
		p := Products{db: db, log: log}

		app.Handle(http.MethodGet, "/v1/products", p.List)
		app.Handle(http.MethodGet, "/v1/products/{id}", p.Get)
		app.Handle(http.MethodPost, "/v1/products", p.Create)
		app.Handle(http.MethodPut, "/v1/products/{id}", p.Update)
		app.Handle(http.MethodDelete, "/v1/products/{id}", p.Delete)

		app.Handle(http.MethodPost, "/v1/products/{id}/sales", p.AddSale)
		app.Handle(http.MethodGet, "/v1/products/{id}/sales", p.ListSales)
	}

	return app
}