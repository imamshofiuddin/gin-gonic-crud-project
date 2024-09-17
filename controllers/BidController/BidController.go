package BidController

import (
	"crud-project/controllers/TimeSeriesDataController"
	"crud-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllBid(c *gin.Context) {

	var bid []models.Bid

	models.DB.Find(&bid)
	c.JSON(http.StatusOK, gin.H{"bids": bid})

}

func ShowBidByUser(c *gin.Context) {
	var bids []models.Bid

	user_id := c.Param("user_id")

	if err := models.DB.Where("user_id = ?", user_id).Find(&bids).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"bids": bids})
}

func Create(c *gin.Context) {

	var bid models.Bid

	if err := c.ShouldBindJSON(&bid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&bid)
	TimeSeriesDataController.AddBid(int(bid.ItemId))
	c.JSON(http.StatusOK, gin.H{"Bid": bid})
}
