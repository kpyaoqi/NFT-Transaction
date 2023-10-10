package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Session sessions.Session

func InitSession(c *gin.Context) {
	Session = sessions.Default(c)
}

// 全局中间件
func SessionSet(key string, value interface{}) {
	Session.Set(key, value)
	Session.Save()
}

func SessionDel(key string) {
	Session.Delete(key)
	Session.Save()
}

func SessionGet(key string) interface{} {
	result := Session.Get(key)
	return result
}
