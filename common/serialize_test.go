package common

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

		err := SerializeRequest(request, &body)

		assert.Nil(t, err)
		assert.NotEqual(t, "", body.Body)
	})

	t.Run("Test Serialize Request Failed", func(t *testing.T) {
		body := struct {
			Body string
		}{}
		requestBody := strings.NewReader(`{ body: "body" }`)
		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", requestBody)

		err := SerializeRequest(request, &body)

		assert.NotNil(t, err)
		assert.Equal(t, "", body.Body)
	})
}

func TestSerializeWriter(t *testing.T) {
	t.Run("Test Serialize Writer Success", func(t *testing.T) {
		data := struct {
			Data string
		}{}
		recorder := httptest.NewRecorder()

		err := SerializeWriter(recorder, http.StatusOK, data)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "", data.Data)
	})

	t.Run("Test Serialize Writer Failed", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		err := SerializeWriter(recorder, http.StatusOK, make(chan int))

		assert.NotNil(t, err)
	})
}
