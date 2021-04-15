package services

import (
	"net/http"
)

// LoadIndexPage loads ./template/index.html page
func (s *HelloWorldServiceConfig) NotFound(w http.ResponseWriter, r *http.Request) {
	s.Log.Info().Msg("Url not found.  Redirecting to cma.ca")
	http.Redirect(w, r, "http://cma.ca", http.StatusNotFound)
}
