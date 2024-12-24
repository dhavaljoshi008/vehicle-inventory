package main

import (
	"github.com/dhavaljoshi008/vehicle-inventory/pkg/vehicle"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/vehicles", vehicle.GetAllVehicles)
	router.GET("/vehicles/:id", vehicle.GetVehicleById)
	router.POST("/vehicles", vehicle.CreateVehicle)
	router.PUT("/vehicles/:id", vehicle.ReplaceVehicle)
	router.PATCH("/vehicles/:id", vehicle.UpdateVehicle)

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}

	if err := router.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
