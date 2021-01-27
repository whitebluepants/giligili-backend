package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

// DailyRank  排行榜
func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.DailyRank()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}