package handler

import (
	"github.com/labstack/echo/v4"
	"home-test-tiki/log"
	"net/http"
	"strconv"
	"strings"
)

type StringHandler struct {

}

func (ss *StringHandler) HandlerSplitString(c echo.Context) error {
	message := c.QueryParam("message")
	log.Info("message received: "+ message)
	stringSplited := strings.Split(message, "")
	var mergeString string
	count := 1
	log.Info("start compress message: "+ message)
	for i, _ := range stringSplited{
		if i != 0 && i < len(stringSplited)-1 {
			if stringSplited[i] == stringSplited[i-1] {
				count = count + 1
				if stringSplited[i] != stringSplited[i+1] {
					mergeString = mergeString + strconv.Itoa(count)
				}
			}else {
				mergeString = mergeString + stringSplited[i]
				count = 1
			}
		}else if i == 0{
			mergeString = mergeString + stringSplited[i]
			count = 1
		}else {
			if stringSplited[i] == stringSplited[i-1] {
				count = count + 1
				mergeString = mergeString + strconv.Itoa(count)
			}else {
				mergeString = mergeString + stringSplited[i]
			}
		}
	}
	log.Info("return result: "+ mergeString)
	return c.String(http.StatusOK, "return: " + mergeString)
}