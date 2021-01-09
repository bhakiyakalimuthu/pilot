package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Controller struct {
	logger  *zap.Logger
	service *Service
}

func NewController(logger *zap.Logger, service *Service) *Controller {
	return &Controller{
		logger:  logger,
		service: service,
	}
}

func (c *Controller) SetupRouter(router chi.Router) chi.Router {
	router.Route("/profile", func(r chi.Router) {
		r.Post("/create", c.Create)
		r.Get("/{id}", c.Get)
		r.Post("/med-data/create", c.CreateMedicalData)
		r.Post("/diagno-prof/create", c.CreateDiagnosisProfile)
	})
	return router
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Successfully received create request")
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println("invalid request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request data"))
		return
	}
	if err := c.service.CreateUser(r.Context(), &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to create profile"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Successfully received get request")
	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("invalid request param")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request param"))
		return
	}
	d, err := c.service.GetUser(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to process"))
		return
	}
	if err := json.NewEncoder(w).Encode(d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to encode the user data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) CreateMedicalData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Successfully received create request")
	var d MedicalData
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println("invalid request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request data"))
		return
	}
	if err := c.service.CreateMedicalData(r.Context(), &d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to create profile"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) CreateDiagnosisProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Successfully received create request")
	var d MedicalData
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println("invalid request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request data"))
		return
	}
	if err := c.service.CreateMedicalData(r.Context(), &d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to create profile"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
