package vehicle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllVehicles(c *gin.Context) {
	c.JSON(http.StatusOK, Vehicles)
}

func GetVehicleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("GetVehicleById() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Please provide a valid Vehicle ID"})
		return
	}
	for _, vehicle := range Vehicles {
		if vehicle.ID == id {
			c.JSON(http.StatusOK, vehicle)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Vehicle not found"})
}
