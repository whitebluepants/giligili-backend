package model

import (
	"giligili/cache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title string
	Info  string
	Url   string
	Avatar string
	UserID uint
}

func (video *Video) AvatarURL() string{
	if video.Avatar == "" {
		return video.Avatar
	}
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (video *Video) AddView() {
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))

	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
