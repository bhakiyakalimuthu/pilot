package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

type Controller struct {
	logger *zap.Logger
	service *Service
}

func NewController(logger *zap.Logger, service *Service) *Controller {
	return &Controller{
		logger:  logger,
		service: service,
	}
}

func (c *Controller) SetupRouter( router chi.Router)  chi.Router {
	router.Route("/profile", func(r chi.Router) {
		r.Post("/user",c.Create)
	})
	return router
}


func (c *Controller) Create(w http.ResponseWriter,r *http.Request)  {

}
