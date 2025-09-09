package scenario

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ScenarioResource struct {
	service *ScenarioService
}

func New(service *ScenarioService) *ScenarioResource {
	return &ScenarioResource{service}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs *ScenarioResource) List(w http.ResponseWriter, r *http.Request) {
	query := GetScenarioQuery(r.URL.Query())

	resp, err := rs.service.FindMany(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := GetScenarioListResponse(resp)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *ScenarioResource) Create(w http.ResponseWriter, r *http.Request) {
	newScenario, err := rs.service.Create()

	if err != nil {
		fmt.Printf("Caught error while creating entity: \"%s\" \n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(newScenario)

	if err != nil {
		fmt.Printf("Caught error while serializing: \"%s\" \n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
