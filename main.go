package main

import (
	BidController "crud-project/controllers/BidController"
	ProductController "crud-project/controllers/ProductController"
	"crud-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", ProductController.Index)
	r.GET("/api/product/:id", ProductController.Show)
	r.POST("/api/product", ProductController.Create)
	r.PUT("/api/product/:id", ProductController.Update)
	r.DELETE("/api/product/:id", ProductController.Delete)

	r.GET("/api/bids", BidController.AllBid)
	r.GET("/api/bids/:user_id", BidController.ShowBidByUser)
	r.POST("/api/bid", BidController.Create)

	r.Run()
}
