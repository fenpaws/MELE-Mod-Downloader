package v1

import (
	"github.com/fenpaws/MELE-Mod-Downloader/internal/api/v1/consume"
	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	Router     *gin.RouterGroup
	modChannel chan string
}

func NewAPIRouter(baseRouter *gin.RouterGroup, modChannel chan string) *APIRouter {
	return &APIRouter{
		Router:     baseRouter,
		modChannel: modChannel,
	}
}

func (api *APIRouter) InitializeRoutes() {

	gole := api.Router.Group("/")
	gole.POST("consume", consume.Consume(api.modChannel))

}
