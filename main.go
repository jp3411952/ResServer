/*
创建时间: 2020/10/28
作者: Administrator
功能介绍:

*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


//func GetName(c *gin.Context){
//	name := c.("name","xxx")
//	action := c.Query("action")
//	message := name + " is " + action
//	c.String(http.StatusOK, message)
//	c.Cookie()
//	//TestRedis()
//}
//
//
//func   PostName(c *gin.Context){
//	name:= c.DefaultPostForm("name","xxx")
//	lastname:= c.DefaultPostForm("lastname","xxx")
//	qqq := c.PostForm("qqq")
//	c.JSON(http.StatusOK,gin.H{
//		"name":name,
//		"lastname":lastname,
//		"qqq":qqq,
//	})
//}
//
//func   uploadfile(c *gin.Context){
//	fh,erro := c.FormFile("file")
//	fmt.Printf(fh.Filename)
//	if erro != nil {
//		c.JSON(http.StatusForbidden,gin.H{
//			"erro" : erro,
//		})
//		return
//	}
//	c.SaveUploadedFile(fh,"./html/c.txt")
//}
//
//func   DowloadFile(c *gin.Context){
//	c.Writer.Header().Add("Content-Disposition",`attachment;filename="c.rar"`)//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
//	//c.Writer.Header().Add("Content-Type", "application/octet-stream")
//	c.Writer.Header().Add("Content-Type", "application/zip")
//	c.File("./html/c.rar")
//}

func Page(){

}

func main() {
	r := gin.Default()
	//r.GET("/getname",GetName)
	//r.POST("/postget",PostName)
	//r.POST("/uploadfile",uploadfile)
	//r.GET("/dfile",DowloadFile)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
	fmt.Println("运行后面")
}