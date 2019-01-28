package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"morefun"
)

func RegisterHandlers()  *httprouter.Router{
	router := httprouter.New()

	router.GET("/", morefun.Register)

	return router
}

func main()  {
	r:= RegisterHandlers()

	http.ListenAndServe(":8081", r)

}