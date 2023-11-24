package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertToString(arg interface{}) string {

	if arg == nil {
		return ""
	}

	var strValue string

	switch arg := arg.(type) {
	case string:
		value, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			strValue = fmt.Sprint(value)
		} else {
			strValue = arg
		}

	case float32, float64:
		strValue = fmt.Sprint(arg)

	case time.Time:
		loc, _ := time.LoadLocation("Asia/Kolkata")
		strValue = arg.In(loc).Format("02/01/2006 15:04:05")
	default:
		strValue = fmt.Sprint(arg)
	}
	return strValue
}
