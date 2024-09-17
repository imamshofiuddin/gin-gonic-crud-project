package main

import (
	BidController "crud-project/controllers/BidController"
	ProductController "crud-project/controllers/ProductController"
	TimeSeriesDataController "crud-project/controllers/TimeSeriesDataController"
	"crud-project/controllers/TokenController"
	"crud-project/controllers/UserController"
	"crud-project/middlewares"
	"crud-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	go TimeSeriesDataController.WriteToInfluxDb()

	api := r.Group("/api")
	{
		api.POST("/token", TokenController.GenerateToken)
		api.POST("/user/register", UserController.CreateUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			//Product API
			secured.GET("/products", ProductController.Index)
			secured.GET("/product/:id", ProductController.Show)
			secured.POST("/product", ProductController.Create)
			secured.PUT("/product/:id", ProductController.Update)
			secured.DELETE("/product/:id", ProductController.Delete)

			//Bid API
			secured.GET("/bids", BidController.AllBid)
			secured.GET("/bids/:user_id", BidController.ShowBidByUser)
			secured.POST("/bid", BidController.Create)

			//Data Series API
			secured.GET("/read", TimeSeriesDataController.ReadFromInfluxDB)
		}
	}

	r.Run()
}
