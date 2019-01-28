package dbops

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"io"
	"morefun/defs"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:33061)/agency_dev?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
}



func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {

	w.WriteHeader(errResp.HttpSC) //写入错误码

	resStr, _ := json.Marshal(&errResp.Error) //格式化 错误信息
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}