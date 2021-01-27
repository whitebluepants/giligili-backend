package cache

import (
	"fmt"
	"strconv"
)

const (
	DailyRankKey = "rank:daily"
)

// 命名空间:分类:id -> 值
// view:video:1 -> 150

// VideoViewKey 视频点击数的key
func VideoViewKey(id uint) string{
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
