package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1

	return randomNumber
}

type Data struct {
	Wind   int    `json:"wind"`
	Water  int    `json:"water"`
	Status string `json:"status"`
}

func getStatus(water, wind int) string {
	waterStatus := "safe"
	if water < 5 {
		waterStatus = "safe"
	} else if water >= 5 && water <= 8 {
		waterStatus = "caution"
	} else {
		waterStatus = "danger"
	}

	windStatus := "safe"
	if wind < 6 {
		windStatus = "safe"
	} else if wind >= 6 && wind <= 15 {
		windStatus = "caution"
	} else {
		windStatus = "danger"
	}

	if waterStatus == "danger" || windStatus == "danger" {
		return "danger"
	} else if waterStatus == "caution" || windStatus == "caution" {
		return "caution"
	}
	return "safe"
}

func handleRequest(c *gin.Context) {
	water := getRandomNumber()
	wind := getRandomNumber()

	status := getStatus(water, wind)

	data := Data{
		Wind:   wind,
		Water:  water,
		Status: status,
	}

	c.JSON(http.StatusOK, data)
}

func main() {
	router := gin.Default()
	router.GET("/update", handleRequest)

	router.Run(":8080")
}
