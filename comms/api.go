package comms

import (
	"id-generator/generator"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Router the HTTP router
var Router *gin.Engine

//CreateRouter - Establish a new HTTP router using the gin framework for our api.
func CreateRouter() {
	Router = gin.Default()

	setupRoutesV1()
	log.Println("[Router] Starting the HTTP Router.")
	Router.Run(":7033")

}

func setupRoutesV1() {
	Router.GET("/id/:length/:prefix", getNewID)
}

func getNewID(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("length"))
	newID, err1 := generator.GenerateID(number, c.Param("prefix"))

	if len(c.Param("prefix")) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Prefix not 3 characters",
		})
		return
	}

	if err != nil || err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failure to generate ID, make sure the length is between 1 and 30 and the prefix is only 3 characters in length.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      newID,
	})
	return
}
