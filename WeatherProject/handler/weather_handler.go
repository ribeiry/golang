package handler

import (
	"WeatherProject/config"
	service "WeatherProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {

	config.Init()
	config.InitCache()
	router := gin.Default()

	router.GET("/temperatura", IndentedJSON)

	router.Run("localhost:8080")
	defer config.DB.Close()

}

func IndentedJSON(c *gin.Context) {

	response, err := service.GetTodayWeather()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar temperatura"})
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}
