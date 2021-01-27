package service

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"strings"
)

type DailyRankService struct {

}

func (d *DailyRankService) DailyRank() serializer.Response{
	var videos []model.Video

	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接出错",
				Error:  err.Error(),
			}
		}
	} else{
		return serializer.Response{
			Status: 50001,
			Msg:    "当前数据不足,排行榜无法显示",
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
