package grf

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ErrCheck(err error, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "err", "data": "", "msg": err.Error()})
}
func OkCheck(data interface{}, c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": data, "msg": ""})
	return
}
func GetId(c *gin.Context) (idInt int, err error) {
	id := c.Param("id")
	idInt, err = strconv.Atoi(id)
	return
}
