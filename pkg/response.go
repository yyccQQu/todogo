package pkg

import (
	"net/http"
	"todogo/pkg/api/defs"
	"encoding/json"
	"io"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {

	//w.WriteHeader(errResp.HttpSC) //写入错误码
	//
	//resStr, _ := json.Marshal(&errResp.Error) //格式化 错误信息
	//
	//io.WriteString(w,string(resStr))

	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
