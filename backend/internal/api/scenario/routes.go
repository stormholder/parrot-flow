package scenario

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ScenarioResource struct {
	store *ScenarioStore
}

func New(store *ScenarioStore) *ScenarioResource {
	return &ScenarioResource{store}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs *ScenarioResource) List(w http.ResponseWriter, r *http.Request) {
	query := GetScenarioQuery(r.URL.Query())

	resp, err := rs.store.List(query)
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
