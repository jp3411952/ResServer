/*
创建时间: 2020/11/21
作者: Administrator
功能介绍:

*/
package global

import "github.com/gin-gonic/gin"

var (
	G_ginEngine *gin.Engine
)

func init()  {
	G_ginEngine = gin.Default();
}