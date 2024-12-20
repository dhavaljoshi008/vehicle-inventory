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
	c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Vehicle with ID = " + strconv.Itoa(id) + " not found"})
}

func CreateVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		fmt.Println("CreateVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "The request body is invalid. Please ensure all required fields are provided and correctly formatted."})
		return
	}
	lastVehicleId := Vehicles[len(Vehicles)-1].ID
	vehicle.ID = lastVehicleId + 1
	Vehicles = append(Vehicles, vehicle)
	c.JSON(http.StatusCreated, vehicle)
}
