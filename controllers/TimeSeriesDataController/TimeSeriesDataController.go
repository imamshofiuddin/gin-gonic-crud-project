package TimeSeriesDataController

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var biddingCounts = make(map[int]int)

func AddBid(productID int) {
	biddingCounts[productID]++
}

func WriteToInfluxDb() {
	for {
		time.Sleep(time.Minute)

		for productID, count := range biddingCounts {
			storeToInfluxDB(productID, count)
		}

		biddingCounts = make(map[int]int)
	}
}

func storeToInfluxDB(productID int, count int) {
	influxDBUrl := "http://localhost:8086/write?db=mydb"
	data := fmt.Sprintf("bidding,product_id=%d value=%d", productID, count)

	req, err := http.NewRequest("POST", influxDBUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}
	resp.Body.Close()

	fmt.Printf("InfluxDB response for product_id %d: %s\n", productID, string(body))
}

func ReadFromInfluxDB(c *gin.Context) {
	influxDBUrl := "http://localhost:8086/query?db=mydb&q=SELECT+*+FROM+bidding"

	resp, err := http.Get(influxDBUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.JSON(http.StatusOK, gin.H{"data": string(body)})
}
