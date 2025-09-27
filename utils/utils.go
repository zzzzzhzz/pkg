package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	uuid2 "github.com/google/uuid"
	"github.com/zzzzzhzz/pkg/constant"
)

// NewUuid 创建uuid
func NewUuid() string {
	uuid := uuid2.New()
	return uuid.String()
}

// NowUnix 当前时间戳
func NowUnix() int {
	var sysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	return int(time.Now().In(sysTimeLocation).Unix())
}

// FormatFromUnixTime 将时间戳转为 yyyy-mm-dd H:i:s 格式
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(constant.SysTimeFormat)
	} else {
		return time.Now().Format(constant.SysTimeFormat)
	}
}

// FormatFromUnixTimeShort 将时间戳转为 yyyy-mm-dd 格式
func FormatFromUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(constant.SysTimeFormatShort)
	} else {
		return time.Now().Format(constant.SysTimeFormatShort)
	}
}

// ParseTime 将字符串转成时间
func ParseTime(str string) (time.Time, error) {
	var sysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(constant.SysTimeFormat, str, sysTimeLocation)
}

// Random 得到一个随机数
func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}

// Ip4toInt 将字符串的 IP 转化为数字
func Ip4toInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	var sum int64 = 0
	if len(bits) == 4 {
		b0, _ := strconv.Atoi(bits[0])
		b1, _ := strconv.Atoi(bits[1])
		b2, _ := strconv.Atoi(bits[2])
		b3, _ := strconv.Atoi(bits[3])
		sum += int64(b0) << 24
		sum += int64(b1) << 16
		sum += int64(b2) << 8
		sum += int64(b3)
	}
	return sum
}

// GetJsonFmtStr func
func GetJsonFmtStr(data interface{}) string {
	resp, _ := json.Marshal(data)
	respStr := string(resp)
	if respStr == "" {
		respStr = fmt.Sprintf("%+v", data)
	}
	return respStr
}
