package handler_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	httpHandler "github.com/fahmyabida/eDot/pkg/http/handler"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	echoOpenAPI "github.com/alexferl/echo-openapi"
)

var e *echo.Echo

func TestMain(m *testing.M) {
	e = echo.New()

	e.Use(echoOpenAPI.OpenAPI("../../../docs/openapi.yaml"))

	os.Exit(m.Run())
}

func TestCheckLiveness(t *testing.T) {
	// Given
	givenMethod := http.MethodGet
	givenURL := "/healthcheck/liveness"
	expectedStatusCode := http.StatusOK
	expectedBody := httpHandler.Healthcheck{
		Status: http.StatusText(http.StatusOK),
	}

	// Mock dependency
	g := e.Group("/healthcheck")
	httpHandler.InitHealthcheckHandler(g)

	// When
	req := httptest.NewRequest(givenMethod, givenURL, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	gotBodyInBytes, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %v", err)
	}
	gotBody := httpHandler.Healthcheck{}
	err = json.Unmarshal(gotBodyInBytes, &gotBody)
	if err != nil {
		t.Fatalf("unable to unmarshal response body: %v", err)
	}

	fmt.Println(rec.Code)
	fmt.Println(gotBody)
	// Then
	require.Equal(t, expectedStatusCode, rec.Code)
	require.Equal(t, expectedBody, gotBody)
}

func TestCheckReadiness(t *testing.T) {
	// Given
	givenMethod := http.MethodGet
	givenURL := "/healthcheck/readiness"
	expectedStatusCode := http.StatusOK
	expectedBody := httpHandler.Healthcheck{
		Status: http.StatusText(http.StatusOK),
	}

	// Mock dependency
	g := e.Group("/healthcheck")
	httpHandler.InitHealthcheckHandler(g)

	// When
	req := httptest.NewRequest(givenMethod, givenURL, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	gotBodyInBytes, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %v", err)
	}
	gotBody := httpHandler.Healthcheck{}
	err = json.Unmarshal(gotBodyInBytes, &gotBody)
	if err != nil {
		t.Fatalf("unable to unmarshal response body: %v", err)
	}

	// Then
	assert.Equal(t, expectedStatusCode, rec.Code)
	assert.Equal(t, expectedBody, gotBody)
}
