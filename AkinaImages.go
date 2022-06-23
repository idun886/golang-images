package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)
//随机数函数
func isNum(a int) int {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(a)
	return number
}

func main() {
	r := gin.Default()
	//设置图片静态资源路径
	r.Static("images", "./images/")
	//设置 get 的url路径
	r.GET("/Akina", func(c *gin.Context) {
		//将读取的文件传值到 切片中 再遍历
		de, _ := ioutil.ReadDir("images")
		var imgSlice []string
		for _, v := range de {
			imgSlice = append(imgSlice, v.Name())
		}
		//设置url的参数值 type=json,type=img
		query, ok := c.GetQuery("type")
		//检测是否有此参数
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": "没有这个参数啊！",
			})
			//检测参数是否为json(不区分大消息)
		} else if strings.ToLower(query) == strings.ToLower("json") {
			c.JSON(http.StatusOK, gin.H{
				"code":  1,
				"image": "http://127.0.0.1:5220/images/" + imgSlice[isNum(len(de))],
			})

		}else if strings.ToLower(query) == strings.ToLower("img"){
			c.Redirect(302,"images/"+imgSlice[isNum(len(de))])
		} else{
			c.JSON(http.StatusOK,gin.H{
				"code":"这个参数没有这个值啊!",
			})
		}
	})
	r.Run(":5220")

}
