// ! importante
// go get github.com/gin-gonic/gin

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/experiments", getExperiments)
	r.GET("/experiments/:id", getExperiment)
	r.POST("/create-experiment", createExperiment)
	r.POST("/delete-experiment", deleteExperiment)

	fmt.Println("Server avviato sulla porta 8080")
	r.Run(":8080")
}

// Funzioni per ottenere gli esperimenti esistenti
func getExperiments(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all experiments",
	})
}

func getExperiment(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "Get experiment " + id,
	})

}

func createExperiment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create experiment",
	})
}
func deleteExperiment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete experiment",
	})
}
