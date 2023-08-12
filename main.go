package main

import (
	"context"
	
	"log"
	"os"
	

	"github.com/bright2704/jwt-api/database"
	routes "github.com/bright2704/jwt-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_"go.mongodb.org/mongo-driver/mongo"
	
)


func main() {


	
	err  := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}
	
	
	router := gin.New()
	router.Use(gin.Logger())

	client := database.DBinstance()
	// collection := OpenCollection(client, os.Getenv("COLLECTION_NAME"))

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"seccess":"Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-2"})
	})
	// Close the MongoDB client when the application exits
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	router.Run(":" + "9000")
}

func DBinstance() {
	panic("unimplemented")
}
