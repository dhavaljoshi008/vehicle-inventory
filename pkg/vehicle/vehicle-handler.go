package vehicle

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// Function to check if all the required vehicle fields are present
func checkForRequiredVehicleFields(vehicle Vehicle) error {
	if vehicle.Make == "" || len(strings.TrimSpace(vehicle.Model)) == 0 {
		return errors.New("vehicle make is missing")
	} else if vehicle.Model == "" || len(strings.TrimSpace(vehicle.Model)) == 0 {
		return errors.New("vehicle model is missing")
	} else if vehicle.Year < 1885 {
		return errors.New("vehicle year is missing or invalid")
	} else {
		return nil
	}
}

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
	// Check for required vehicle fields
	if err := checkForRequiredVehicleFields(vehicle); err != nil {
		fmt.Println("CreateVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "The request body is invalid. Please ensure all required fields are provided and correctly formatted."})
		return
	}
	lastVehicleId := Vehicles[len(Vehicles)-1].ID
	vehicle.ID = lastVehicleId + 1
	Vehicles = append(Vehicles, vehicle)
	c.JSON(http.StatusCreated, vehicle)
}

func ReplaceVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ReplaceVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Please provide a valid Vehicle ID"})
		return
	}
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		fmt.Println("ReplaceVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "The request body is invalid. Please ensure all required fields are provided and correctly formatted."})
		return
	}
	// Check for required vehicle fields
	if err := checkForRequiredVehicleFields(vehicle); err != nil {
		fmt.Println("ReplaceVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "The request body is invalid. Please ensure all required fields are provided and correctly formatted."})
		return
	}

	flag := 0 // Default for vehicle not found

	// Find the vehicle by id and replace it
	for i := 0; i < len(Vehicles); i++ {
		if Vehicles[i].ID == id {
			// Update all the properties except ID
			Vehicles[i].Make = vehicle.Make
			Vehicles[i].Model = vehicle.Model
			Vehicles[i].Year = vehicle.Year
			// Return the replaced vehicle from the Vehicle slice
			c.JSON(http.StatusOK, Vehicles[i])
			// Vehicle found, set flag = 1
			flag = 1
			return
		}
	}
	if flag == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Vehicle with ID = " + strconv.Itoa(id) + " not found"})
	}
}

func UpdateVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("UpdateVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Please provide a valid vehicle ID"})
		return
	}
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		fmt.Println("UpdateVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Please ensure that the fields are valid and correctly formatted."})
		return
	}
	for i := 0; i < len(Vehicles); i++ {
		if Vehicles[i].ID == id {
			if vehicle.Make != "" && len(strings.TrimSpace(vehicle.Make)) > 0 {
				Vehicles[i].Make = vehicle.Make
			}
			if vehicle.Model != "" && len(strings.TrimSpace(vehicle.Model)) > 0 {
				Vehicles[i].Model = vehicle.Model
			}
			if vehicle.Year > 1885 {
				Vehicles[i].Year = vehicle.Year
			}
			c.JSON(http.StatusOK, Vehicles[i])
			return
		}
	}
}

func DeleteVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("DeleteVehicle() Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Please provide a valid Vehicle ID"})
	}

	for i := 0; i < len(Vehicles); i++ {
		if Vehicles[i].ID == id {
			fmt.Println("DeleteVehicle() Deleting Vehicle: ", Vehicles[i])
			Vehicles = append(Vehicles[:i], Vehicles[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
