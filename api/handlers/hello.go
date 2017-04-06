package handlers

import (
	//"api/models"
	//"api/services"
	//pb "base/protos/helloworld"
	//"context"
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"
)

// handlers  ==>  models  ==>  services
// 所有的 db 调用在model层,model层 调用 services层的服务,
// db 提供接口
func RegisterHello(router *gin.RouterGroup) {
	r := router.Group("/hello")
	r.GET("/hello", Hello)
}

func Hello(c *gin.Context) {
	//p := fmt.Println

	//jsonTime := models.JSONTime(time.Now())
	//p(jsonTime)

	//jsonTime := models.JSONTime{Time: time.Now()}
	//account := models.Account{Tel: "Lbxxxc44450", Password: "sssdfss", StartAt: jsonTime}

	//db := services.InitDB()
	//db.NewRecord(account)
	//db.Create(&account)
	//
	//cli := services.InitGrpc()
	//cli.SayHello(context.Background(), &pb.HelloRequest{Name: "hello", Num: "2"})

	//p(account)
	//p(time.Now().UTC())

	// var account models.Account

	// db.First(&account, 5) // find product with id 1

	// time := account.StartAt.Format(time.RFC3339Nano)

	// utils.RenderJson(&account)

	// c.JSON(200, )
	//
	c.JSON(http.StatusOK, gin.H{"user": "hello"})
}
