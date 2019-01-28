package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"morefun"
	"morefun/dbops"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.GET("/", morefun.Register)
	router.GET("/pa", morefun.Hasagain)

	router.POST("/pos/:fun", morefun.Resign)

	router.DELETE("/delete/:fun", morefun.Delfun)

	router.PUT("/put/:fun", dbops.Putfun)

	return router
}

func main()  {
	r:= RegisterHandlers()

	http.ListenAndServe(":8081", r)

}