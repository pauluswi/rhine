package trxhistory

import (
	"net/http"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/pauluswi/rhine/internal/entity"
	"github.com/pauluswi/rhine/internal/errors"
	"github.com/pauluswi/rhine/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)

	// endpoint get
	r.Get("/get/<id>", res.get)

	// the following endpoints require a valid JWT
	r.Post("/save", res.save)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	intID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	trx, err := r.service.Get(c.Request.Context(), intID)
	if err != nil {
		return err
	}

	return c.Write(trx)
}

func (r resource) save(c *routing.Context) error {
	var input entity.InputSave
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("Bad Request")
	}
	trxhistory, err := r.service.Save(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(trxhistory, http.StatusCreated)
}
