package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/src/infra/controllers"
	"github.com/eduns/oncar-challenge/backend/tests/mocks/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupServer() *gin.Engine {
	vehicleRepository := repositories.NewVehicleRepositoryInMemory()
	leadRepository := repositories.NewLeadRepositoryInMemory()

	getVehiclesUseCase := usecases.NewGetVehiclesUseCase(&vehicleRepository)
	createLeadUseCase := usecases.NewCreateLeadUseCase(&leadRepository)

	vehicleController := controllers.NewVehicleController(getVehiclesUseCase)
	leadController := controllers.NewLeadController(createLeadUseCase)

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.POST("/leads", leadController.HandleCreateLead)
	router.GET("/vehicles", vehicleController.HandleGetVehicles)

	return router
}

var server = setupServer()

func TestVehicleRoutes(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/vehicles", nil)

	server.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
}

func TestLeadRoutes(t *testing.T) {
	payload := `{"name":"Test","email":"mail@mail.com","phone":"+5512881437331","chosenVehicleId":"1da55e15-f253-488d-8590-e0c8a955898b"}`

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/leads", strings.NewReader(payload))
	request.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(w, request)

	assert.Equal(t, 201, w.Code)
}

func TestLeadRoutesWithInvalidParams(t *testing.T) {
	payload := `{"email":"mail@mail.","phone":"+5512881437331","chosenVehicleId":"1da55e15-f253-488d-8590-e0c8a955898b"}`

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/leads", strings.NewReader(payload))
	request.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(w, request)

	var result map[string]string
	json.Unmarshal([]byte(w.Body.String()), &result)

	errors := strings.Split(string(result["error"]), "\n")

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, errors[0], "Key: 'CreateLeadPayload.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		"Should return an error for invalid 'name' param")
	assert.Equal(t, errors[1], "Key: 'CreateLeadPayload.Email' Error:Field validation for 'Email' failed on the 'email' tag",
		"Should return an error for invalid 'email' param")
}
