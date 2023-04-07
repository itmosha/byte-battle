package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_handleApiList(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api", nil)

	s.handleApiList().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "API root")
}
