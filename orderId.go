package faceid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// CreateOrderID 生成随机订单号，格式：yyyyMMddHHmm ss + 14位随机数
func createOrderID() string {
	// 获取毫秒级时间戳
	now := time.Now()
	timeStamp := now.UnixNano()
	// 生成随机数种子
	seed := rand.NewSource(timeStamp)
	randNumber := rand.New(seed) // 重置随机数种子
	datetime, _ := strconv.Atoi(now.Format("20060102150405"))
	orderID := fmt.Sprintf("%014d%02d", datetime, randNumber.Intn(99)) // 组合生成随机数
	return orderID
}
