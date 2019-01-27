package web

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"github.com/labstack/gommon/log"
)

type HomePage struct {
	Name string
}


func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil && err2 !=nil {
		p := &HomePage{Name:"yyccQQu"}
		t, e := template.ParseFiles("./template/home.html") //搞清 符合编译性语言的规范 路径

		if e!= nil {
			log.Printf("Parsing template home.html error:%s",e)
			return
		}

		t.Execute(w,p) //将responseWriter返回给前端，进行数据渲染。将变量结合模版形成模版引擎
		return
	}

	if len(cname.Value)!=0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

}
