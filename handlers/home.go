package handlers

import (
	"net/http"
	"stockchart/components"
)

// HomeHandler serves the home page and renders the chart component
// It uses the Chart component to render the chart in the response

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		components.Chart().Render(r.Context(), w)
	}
}
