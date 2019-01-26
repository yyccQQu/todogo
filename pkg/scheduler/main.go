package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todogo/pkg/scheduler/src"
	)

// 什么是sheduler
// 叫做 任务，通过普通resetapi方法 不会马上给他结果的任务，
// 这些任务会分发到 scheduler 里面 ，然后通过一些方法定时或者延时触发，settimeout ，interva 《异步任务》

// 延迟操作，需要scheduler，周期，顺时

//scheduler 包含RESTful 的 http server, Timer, task Runer

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", src.VidDelRecHandler)
	return router
}

func main()  {
	//go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe("9001",r)
}


