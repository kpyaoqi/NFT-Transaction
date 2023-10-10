package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"yqnft/NFTManage/routes"
)

// 全局中间件
func initMiddleware(c *gin.Context) {
	//session := sessions.Default(c)
	//userID := session.Get("userId")
	//fmt.Println(userID)
	//if userID != nil {
	//	c.Set("userId", userID)
	//} else {
	//	c.Set("userId", "20230908095458693536608783")
	//}
	//c.Next()
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("root"))
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mySession", store))
	r.Use(initMiddleware)
	r.Static("/static", "./static")
	routes.LoginRoutesInit(r)
	routes.ApiRoutesInit(r)
	r.Run(":8080")
}
