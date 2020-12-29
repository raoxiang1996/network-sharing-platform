package routes

import (
	"University-Information-Website/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.Run(utils.HttpPort)
}
