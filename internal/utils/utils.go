package utils

import (
	"fmt"
	"github.com/labstack/echo"
	"time"
)

//GetEST coment
func getEst() time.Time {
	result := time.Now().UTC()
	location, err := time.LoadLocation("EST")
	if err == nil {
		result = result.In(location)
	}
	return result
}

//TimeToString comment
func timeToString(tm time.Time) string {
	var result string
	result = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	return result
}

//ConsoleLog comment
func ConsoleLog(c echo.Context) {
	now := getEst() //time.Now()
	timeMessage := timeToString(now)
	fmt.Println(timeMessage, "-", c.Request().Method, "-", c.Request().URL)
}