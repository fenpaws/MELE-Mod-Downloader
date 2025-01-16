package consume

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NexusRequest struct {
	NexusURL string `json:"nexus_url"`
}

func Consume(modChannel chan string) func(c *gin.Context) {
	return func(c *gin.Context) {
		nexusRequest := NexusRequest{}
		err := c.ShouldBindBodyWithJSON(&nexusRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		modChannel <- nexusRequest.NexusURL
	}
}
