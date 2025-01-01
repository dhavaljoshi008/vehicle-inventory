package vehicle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllVehicles(t *testing.T) {
	router := gin.Default()
	router.GET("/vehicles", GetAllVehicles)
	req, _ := http.NewRequest("GET", "/vehicles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var vehicles []Vehicle
	if err := json.Unmarshal(w.Body.Bytes(), &vehicles); err != nil {
		return
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, Vehicles)
}

func TestGetVehicleByID(t *testing.T) {
	router := gin.Default()
	router.GET("/vehicles/:id", GetVehicleById)

	// 200 - OK

	req, _ := http.NewRequest("GET", "/vehicles/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var vehicle Vehicle
	if err := json.Unmarshal(w.Body.Bytes(), &vehicle); err != nil {
		return
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, vehicle)

	// 400 - Bad Request

	reqBadRequest, _ := http.NewRequest("GET", "/vehicles/a", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqBadRequest)
	var responseBadRequest map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &responseBadRequest); err != nil {
		return
	}
	bodyReqBadRequest := gin.H{"error": "Bad Request", "message": "Please provide a valid Vehicle ID"}
	errorValBadRequest, errorValExists := responseBadRequest["error"]
	messageValBadRequest, messageValExists := responseBadRequest["message"]
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, errorValExists)
	assert.Equal(t, bodyReqBadRequest["error"], errorValBadRequest)
	assert.True(t, messageValExists)
	assert.Equal(t, bodyReqBadRequest["message"], messageValBadRequest)

	// 404 - Not Found

	reqNotFound, _ := http.NewRequest("GET", "/vehicles/0", nil)
	w = httptest.NewRecorder()
	var responseReqNotFound map[string]string
	router.ServeHTTP(w, reqNotFound)
	if err := json.Unmarshal(w.Body.Bytes(), &responseReqNotFound); err != nil {
		return
	}
	errorValReqNotFound, errorValExists := responseReqNotFound["error"]
	messageValReqNotFound, messageValExists := responseReqNotFound["message"]
	bodyReqNotFound := gin.H{"error": "Not Found", "message": "Vehicle with ID = 0 not found"}
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, errorValExists)
	assert.Equal(t, bodyReqNotFound["error"], errorValReqNotFound)
	assert.True(t, messageValExists)
	assert.Equal(t, bodyReqNotFound["message"], messageValReqNotFound)
}
