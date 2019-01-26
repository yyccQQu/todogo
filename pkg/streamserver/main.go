package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todogo/pkg/streamserver/src"
	"log"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *src.ConnLimiter
}

//自己申明一个结构体，通过方法调用返回一个全新的结构体
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = src.NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		src.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func init() {

	log.SetFlags(log.Flags() | log.Llongfile)
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", src.StreamHandler)
	router.POST("/upload/:vid-id", src.UploadHandler)
	router.GET("/testpage", src.TestPageHandler)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
