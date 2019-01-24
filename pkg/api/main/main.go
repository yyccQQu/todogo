package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todogo/pkg/api"
	"fmt"
)

type middleWareHandler struct {
	r *httprouter.Router
}

//---1
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	fmt.Println(76543)
	return m
}

//---2
// middleWareHandler执行的函数 //duck type //劫持 serveHTTP 函数
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	api.ValidateUserSession(r)
	fmt.Println(123456)
	// ServeHTTP makes the router implement the http.Handler interface.
	// ServeHTTP使路由器实现http。处理程序接口。
	m.r.ServeHTTP(w, r)  //duck type 长得像鸭子，那么他就是鸭子
}

//1，2之间的联系


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
// main -> middleware -> defs(message,err) -> handlers -> dbops -> response