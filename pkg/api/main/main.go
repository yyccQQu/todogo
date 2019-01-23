package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todogo/pkg/api"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	//validateUserSession(r)

	// ServeHTTP makes the router implement the http.Handler interface.
	// ServeHTTP使路由器实现http。处理程序接口。
	m.r.ServeHTTP(w, r)
}


func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user",api.CreateUser)

	router.POST("/user/:user_name", api.Login)

	return router
}

func main() {

	r:= RegisterHandlers()
	mh := NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", mh)
}


//过程

// 从 handler 进来，验证用户名及密码，返回response
// 从 main -> handler -> dbops调用，拿到信息，结合defs再做进一步处理（消息定义），组装成response后返回