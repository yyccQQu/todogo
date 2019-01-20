package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todogo/pkg/api"
)




func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user",api.CreateUser)

	router.POST("/user/:user_name", api.Login)

	return router
}

func main() {

	r:= RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

//过程

// 从 handler 进来，验证用户名及密码，返回response
// 从 main -> handler -> dbops调用，拿到信息，结合defs再做进一步处理（消息定义），组装成response后返回