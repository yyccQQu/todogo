package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"todogo/pkg"
)




func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user", pkg.CreateUser)

	router.POST("/user/:user_name", pkg.Login)

	return router
}

func main() {
	pkg.SayHello()

	r:= RegisterHandlers()
	http.ListenAndServe(":8000", r)
}