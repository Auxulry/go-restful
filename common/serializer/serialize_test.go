package serializer

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeRequest(t *testing.T) {
	t.Run("Test Serialize Request Success", func(t *testing.T) {
		body := struct {
			Body string
		}{}
		requestBody := strings.NewReader(`{ "body": "body" }`)
		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", requestBody)

		SerializeRequest(request, &body)

		assert.NotEqual(t, "", body.Body)
	})

	t.Run("Test Serialize Request Failed", func(t *testing.T) {
		body := struct {
			Body string
		}{}
		requestBody := strings.NewReader(`{ body: "body" }`)
		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", requestBody)

		SerializeRequest(request, &body)

		assert.Equal(t, "", body.Body)
	})
}

func TestSerializeWriter(t *testing.T) {
	t.Run("Test Serialize Writer Success", func(t *testing.T) {
		data := struct {
			Data string
		}{}
		recorder := httptest.NewRecorder()

		SerializeWriter(recorder, http.StatusOK, data)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "", data.Data)
	})
}
