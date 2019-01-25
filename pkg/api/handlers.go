package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"encoding/json"
	"todogo/pkg"
	"todogo/pkg/api/defs"
	"todogo/pkg/api/dbops"
	"todogo/pkg/api/session"
	"fmt"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	//io.WriteString(w, "Create Userhandler")
	res, _ := ioutil.ReadAll(r.Body)
	// 在body里面添加
	//{
	//  "user_name": "xyd",
	//  "pwd": "123456"
	//}
	//
	fmt.Println("---->")
	ubody := &defs.UserCredential{}
	fmt.Println("---->",res)


	if err := json.Unmarshal(res, ubody); err != nil { //如果不能格式化json，那就返回错误
		pkg.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		pkg.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username) //将username 以sessionId标记（key: value） 存起来
	su := &defs.SignedUp{Success:true, SessionId:id} //将ID 存进是否注册状态

	if resp, err := json.Marshal(su); err != nil { //如果 不能格式化注册状态
		pkg.SendErrorResponse(w, defs.ErrorInternalFaults)
	}else{
		pkg.SendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}