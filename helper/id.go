package helper

import (
	"strconv"
	"time"
)

func IdData(sPrefix string) string {
	// convert now unix timestamp to string
	sId := strconv.FormatInt(time.Now().Unix(), 10)
	return sPrefix + sId
}
