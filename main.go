/*
创建时间: 2020/10/28
作者: Administrator
功能介绍:

*/
package main

import (
	"ResServer/global"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

//资源文件
var resp  gin.H



func main() {
	loadVersion()
	global.G_ginEngine.GET("/reload",RLoadVersion)
	global.G_ginEngine.GET("/getVersion",GetVserion)
	global.G_ginEngine.Run(":9090")
}


func loadVersion()  {
	jsonfile,erro := os.Open("version.json")
	if erro != nil {
		fmt.Println(erro)
		return
	}
	jsond :=  json.NewDecoder(jsonfile)
	if jsond == nil {
		fmt.Printf("json decoder 失败")
		return
	}
	jsond.Decode(&resp)
}
func   RLoadVersion(c *gin.Context){
	loadVersion()
	c.JSON(http.StatusOK,resp)
}

func   GetVserion(c *gin.Context){
	c.JSON(http.StatusOK,resp)
}