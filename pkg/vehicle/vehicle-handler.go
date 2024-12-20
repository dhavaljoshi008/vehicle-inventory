package vehicle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllVehicles(c *gin.Context) {
	c.JSON(http.StatusOK, Vehicles)
}
