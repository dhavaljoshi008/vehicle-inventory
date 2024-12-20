package main

import (
	"github.com/dhavaljoshi008/vehicle-inventory/pkg/vehicle"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/vehicles", vehicle.GetAllVehicles)
	router.GET("/vehicles/:id", vehicle.GetVehicleById)
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	router.Run("127.0.0.1:8080")
}
