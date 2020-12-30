package main

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"

	"github.com/bhakiyakalimuthu/pilot/internal/app"
)

func main()  {

	r := chi.NewRouter()

	store:= app.NewMemoryStore(zap.L())
	svc := app.NewService(zap.L(),store)
	_ = app.NewController(zap.L(),svc).SetupRouter(r)
	http.ListenAndServe(":8080",r)
}


// main - > controller
// GetProfile (/GetProfile)
// CreateProfile
// UpdateProfile
// DeleteProfile