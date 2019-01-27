package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userhomeHandler)
	router.POST("/userhome", userhomeHandler)

	router.POST("/api", apiHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}

func main()  {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

//将go编译出来，再把html移到真正的webUI下(加载真正的路径）
